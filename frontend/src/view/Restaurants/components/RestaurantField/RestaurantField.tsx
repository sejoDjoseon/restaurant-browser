import { Restaurant } from 'models/Restaurants'

import {
  RestaurantName,
  RestaurantImg,
  OpenStatusContainer,
  ImageConainer,
} from './FieldLayout'

interface RestaurantFieldProps {
  restaurant: Restaurant
}

export default ({ restaurant }: RestaurantFieldProps) => (
  <>
    <ImageConainer>
      <RestaurantImg src={`https://${restaurant.image}`} alt="restaurant-pic" />
    </ImageConainer>
    <RestaurantName>
      <h3>{restaurant.name}</h3>
    </RestaurantName>
    <OpenStatusContainer>
      {restaurant.open ? 'Open' : 'Closed'}
    </OpenStatusContainer>
  </>
)
