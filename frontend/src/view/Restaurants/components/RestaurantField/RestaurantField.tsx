import { Restaurant } from 'models/Restaurants'

import FieldContent from '../FieldContent/FieldContent'
import RestaurantImg from '../RestaurantImg/RestaurantImg'

interface RestaurantFieldProps {
  restaurant: Restaurant
}

export default ({ restaurant }: RestaurantFieldProps) => (
  <>
    <RestaurantImg src={`https://${restaurant.image}`} alt="restaurant-pic" />
    <FieldContent>
      <h3>{restaurant.name}</h3>
    </FieldContent>
  </>
)
