import React from 'react'

import Cell from 'components/Cell/Cell'
import { Restaurant } from 'models/Restaurants'
import { useNavigate } from 'react-router-dom'

import RestaurantField from '../RestaurantField/RestaurantField'
import { ItemList } from '../ScreenLayout/ScreenLayout'

interface RestaurantListProps {
  restaurants: Restaurant[]
}

export default ({ restaurants }: RestaurantListProps) => {
  const navigate = useNavigate()
  return (
    <ItemList>
      {restaurants.map((r) => (
        <Cell
          key={r.id}
          style={{ marginBottom: 20, maxWidth: 800 }}
          onClick={() => navigate(`restaurants/${r.id}`)}
        >
          <RestaurantField restaurant={r} />
        </Cell>
      ))}
    </ItemList>
  )
}
