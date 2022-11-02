import { Restaurant } from 'models/Restaurants'

import { RestaurantsStoreI } from '../RestaurantsStore'

export class RestaurantsStoreMock implements RestaurantsStoreI {
  restaurants?: Restaurant[]

  getRestaurants = jest.fn()

  cleanRestaurants = jest.fn()

  getRestaurant = jest.fn()
}
