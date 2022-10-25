import React, { useEffect, useState } from 'react'

import { useAppContext } from 'AppContext'
import { useAutorun } from 'hooks/useAutorun'
import { Restaurant } from 'models/Restaurants'

import RestaurantList from './components/RestaurantList/RestaurantList'

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
    <React.Fragment>
      <div>
        <h2>Restaurants</h2>
      </div>
      {loaging && <h1>Loading</h1>}
      {restaurants && <RestaurantList restaurants={restaurants} />}
    </React.Fragment>
  )
}
