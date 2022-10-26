import { Restaurant } from 'models/Restaurants'
import { RestaurantHttpClient } from 'services/RestaurantHttpClient'

export interface RestaurantsTransportLayerI {
  getRestaurants: () => Promise<Restaurant[]>
}

export const newRestaurantsTransportLayer = (
  _restaurantHttpClient: RestaurantHttpClient,
): RestaurantsTransportLayerI => ({
  getRestaurants: _restaurantHttpClient.getRestaurants,
})
