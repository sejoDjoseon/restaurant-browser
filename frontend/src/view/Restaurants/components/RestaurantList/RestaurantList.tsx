import React from 'react'

import Cell from 'components/Cell/Cell'
import { Restaurant } from 'models/Restaurants'
import { useNavigate } from 'react-router-dom'

import RestaurantField from '../RestaurantField/RestaurantField'

interface RestaurantListProps {
  restaurants: Restaurant[]
}

export default ({ restaurants }: RestaurantListProps) => {
  const navigate = useNavigate()
  return (
    <>
      {restaurants.map((r) => (
        <Cell
          key={r.id}
          style={{ marginBottom: 20 }}
          onClick={() => navigate(`restaurants/${r.id}`)}
        >
          <RestaurantField restaurant={r} />
        </Cell>
      ))}
    </>
  )
}
