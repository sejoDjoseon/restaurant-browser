import { Coordinates } from 'models/Coordinates'
import { Restaurant } from 'models/Restaurants'
import { RestaurantHttpClient } from 'services/RestaurantHttpClient'

export interface RestaurantsTransportLayerI {
  getRestaurants: (location?: Coordinates) => Promise<Restaurant[]>
}

export const newRestaurantsTransportLayer = (
  _restaurantHttpClient: RestaurantHttpClient,
): RestaurantsTransportLayerI => ({
  getRestaurants: _restaurantHttpClient.getRestaurants,
})
