import React, { useEffect, useRef, useState } from 'react'

import { useAppContext } from 'AppContext'
import NoContent from 'components/NoContent/NoContent'
import Title from 'components/Title/Title'
import { Catalog } from 'models/Catalog'
import { Restaurant } from 'models/Restaurants'
import { useParams } from 'react-router-dom'

import {
  CatalogTransportLayerI,
  newCatalogTransportLayer,
} from './CatalogTransportLayer'
import Category from './components/Category/Category'
import {
  CategoryContainer,
  ScreenContainer,
  SearcherContainer,
} from './components/ScreenLayout/ScreenLayout'
import SearchInput from './components/SearchInput/SearchInput'
import CatalogHttpClient from './services/CatalogHttpClient'

export default () => {
  const { id } = useParams()
  const { _restaurantsStore } = useAppContext()!

  const transportLayer = useRef<CatalogTransportLayerI>(
    newCatalogTransportLayer(new CatalogHttpClient()),
  ).current

  const [loading, setLoading] = useState(true)
  const [restaurant, setRestaurant] = useState<Restaurant | undefined>(
    undefined,
  )
  const [catalog, setCatalog] = useState<Catalog | undefined>(undefined)

  const handleCatalogResponse = (newCatalog: Catalog) => {
    setLoading(false)
    setCatalog(newCatalog)
  }

  useEffect(() => {
    if (id) {
      _restaurantsStore.getRestaurant(id).then(setRestaurant)
      transportLayer.getCatalog(id).then(handleCatalogResponse)
    }
  }, [_restaurantsStore, id, transportLayer])

  const handleSearch = (value: string) => {
    setCatalog(undefined)
    setLoading(true)
    !!id && transportLayer.getCatalog(id, value).then(handleCatalogResponse)
  }

  return (
    <ScreenContainer>
      {restaurant && <Title>Catalog: {restaurant.name}</Title>}
      {loading && <NoContent>Loading</NoContent>}
      {!!catalog && (
        <>
          <SearcherContainer>
            <SearchInput onSubmit={handleSearch} />
          </SearcherContainer>
          {catalog.map((category, index) => (
            <CategoryContainer key={index}>
              <Category category={category} />
            </CategoryContainer>
          ))}
        </>
      )}
      {catalog?.length === 0 && (
        <NoContent>
          <h2>We can't find results for your search</h2>
        </NoContent>
      )}
    </ScreenContainer>
  )
}
