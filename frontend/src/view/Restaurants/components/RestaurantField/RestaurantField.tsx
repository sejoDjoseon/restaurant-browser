import { Restaurant } from 'models/Restaurants'

import {
  RestaurantName,
  RestaurantImg,
  OpenStatusContainer,
} from './FieldLayout'

interface RestaurantFieldProps {
  restaurant: Restaurant
}

export default ({ restaurant }: RestaurantFieldProps) => (
  <>
    <RestaurantImg src={`https://${restaurant.image}`} alt="restaurant-pic" />
    <RestaurantName>
      <h3>{restaurant.name}</h3>
    </RestaurantName>
    <OpenStatusContainer>
      {restaurant.open ? 'Open' : 'Closed'}
    </OpenStatusContainer>
  </>
)
