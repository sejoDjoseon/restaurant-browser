import React, { useEffect, useRef, useState } from 'react'

import { useAppContext } from 'AppContext'
import { Catalog } from 'models/Catalog'
import { useParams } from 'react-router-dom'

import {
  CatalogTransportLayerI,
  newCatalogTransportLayer,
} from './CatalogTransportLayer'
import Category from './components/Category/Category'
import CategoryContainer from './components/CategoryContainer/CategoryContainer'
import SearchInput from './components/SearchInput/SearchInput'
import CatalogHttpClient from './services/CatalogHttpClient'

export default () => {
  const { id } = useParams()
  const { _restaurantsStore } = useAppContext()!

  const transportLayer = useRef<CatalogTransportLayerI>(
    newCatalogTransportLayer(new CatalogHttpClient()),
  ).current

  const [loading, setLoading] = useState(true)
  const [catalog, setCatalog] = useState<Catalog | undefined>(undefined)

  const handleCatalogResponse = (newCatalog: Catalog) => {
    setLoading(false)
    setCatalog(newCatalog)
  }

  useEffect(() => {
    !!id && transportLayer.getCatalog(id).then(handleCatalogResponse)
  }, [id, transportLayer])

  const handleSearch = (value: string) => {
    setCatalog(undefined)
    setLoading(true)
    !!id && transportLayer.getCatalog(id, value).then(handleCatalogResponse)
  }

  return (
    <div>
      <h2>Catalog</h2>
      <SearchInput onSubmit={handleSearch} />
      {loading && <h3>Loading</h3>}
      {!!catalog &&
        catalog.map((category, index) => (
          <CategoryContainer key={index}>
            <Category category={category} />
          </CategoryContainer>
        ))}
      {catalog?.length === 0 && <h3>We can't find reasults for your search</h3>}
    </div>
  )
}
