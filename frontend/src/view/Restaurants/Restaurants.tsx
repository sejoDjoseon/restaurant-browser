import React, { useEffect, useState } from 'react'

import { useAppContext } from 'AppContext'
import { useAutorun } from 'hooks/useAutorun'
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

  return (
    <ScreenContainer>
      <SectionContainer widthVW={LEFT_SECTION_VW}>
        <BrowserMap></BrowserMap>
      </SectionContainer>
      <SectionContainer widthVW={RIGHT_SECTION_VW}>
        <div>
          <h2>Restaurants</h2>
        </div>
        {loaging && <h1>Loading</h1>}
        {restaurants && <RestaurantList restaurants={restaurants} />}
      </SectionContainer>
    </ScreenContainer>
  )
}
