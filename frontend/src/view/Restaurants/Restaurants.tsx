import React, { useEffect, useState } from 'react'

import { useAppContext } from 'AppContext'
import { useAutorun } from 'hooks/useAutorun'
import { Restaurant } from 'models/Restaurants'

import BrowserMap from './components/BrowserMap/BrowserMap'
import RestaurantList from './components/RestaurantList/RestaurantList'
import ScreenLayout from './components/ScreenLayout/ScreenLayout'

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
    <ScreenLayout>
      <div style={{ width: '40vw' }}>
        <BrowserMap></BrowserMap>
      </div>
      <div style={{ width: '60vw' }}>
        <div>
          <h2>Restaurants</h2>
        </div>
        {loaging && <h1>Loading</h1>}
        {restaurants && <RestaurantList restaurants={restaurants} />}
      </div>
    </ScreenLayout>
  )
}
