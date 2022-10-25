import React from 'react'

import Cell from 'components/Cell/Cell'
import { Restaurant } from 'models/Restaurants'

import RestaurantField from '../RestaurantField/RestaurantField'

interface RestaurantListProps {
  restaurants: Restaurant[]
}

export default ({ restaurants }: RestaurantListProps) => (
  <>
    {restaurants.map((r) => (
      <Cell key={r.id} style={{ marginBottom: 20 }}>
        <RestaurantField restaurant={r} />
      </Cell>
    ))}
  </>
)
