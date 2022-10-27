import React, { useEffect, useState } from 'react'

import { useAppContext } from 'AppContext'
import NoContent from 'components/NoContent/NoContent'
import Title from 'components/Title/Title'
import { useAutorun } from 'hooks/useAutorun'
import { Coordinates } from 'models/Coordinates'
import { Restaurant } from 'models/Restaurants'

import BrowserMap from './components/BrowserMap/BrowserMap'
import RestaurantList from './components/RestaurantList/RestaurantList'
import {
  ScreenContainer,
  SectionContainer,
} from './components/ScreenLayout/ScreenLayout'

const LEFT_SECTION_VW = 40
const RIGHT_SECTION_VW = 60

export default () => {
  const { _restaurantsStore } = useAppContext()!
  const [restaurants, setRestaurants] = useState<Restaurant[] | undefined>()
  const [loaging, setLoading] = useState(true)

  useEffect(() => {
    !_restaurantsStore.restaurants && _restaurantsStore.getRestaurants()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useAutorun(() => {
    if (!!_restaurantsStore.restaurants) {
      setLoading(false)
      setRestaurants(_restaurantsStore.restaurants)
    }
  })

  const handleNewPosition = (position: Coordinates) => {
    _restaurantsStore.cleanRestaurants()
    setLoading(false)
    _restaurantsStore.getRestaurants(position)
  }

  return (
    <ScreenContainer>
      <SectionContainer widthVW={LEFT_SECTION_VW}>
        <BrowserMap onNewPosition={handleNewPosition}></BrowserMap>
      </SectionContainer>
      <SectionContainer widthVW={RIGHT_SECTION_VW} style={{ overflow: 'auto' }}>
        <Title>Restaurants</Title>
        {loaging && (
          <NoContent>
            <h1>Loading</h1>
          </NoContent>
        )}
        {restaurants && <RestaurantList restaurants={restaurants} />}
      </SectionContainer>
    </ScreenContainer>
  )
}
