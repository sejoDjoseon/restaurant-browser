import { Restaurant } from 'models/Restaurants'

interface RestaurantFieldProps {
  restaurant: Restaurant
}

export default ({ restaurant }: RestaurantFieldProps) => (
  <>
    <img src={`https://${restaurant.image}`} alt="restaurant-pic"></img>
  </>
)
