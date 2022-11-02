import { makeObservable, observable, action } from 'mobx'
import { AppError } from 'models/AppError'
import { Coordinates } from 'models/Coordinates'
import { Restaurant } from 'models/Restaurants'

import { RestaurantsTransportLayerI } from './RestaurantsTransportLayer'

export interface RestaurantsStoreI {
  restaurants?: Restaurant[]
  getRestaurants: (location?: Coordinates) => void
  cleanRestaurants: () => void
  getRestaurant: (rstID: string) => Promise<Restaurant>
}

export default class RestaurantsStore implements RestaurantsStoreI {
  @observable
  restaurants?: Restaurant[]

  constructor(private readonly _transportLayer: RestaurantsTransportLayerI) {
    makeObservable(this)
  }

  getRestaurants(location?: Coordinates) {
    this.fetchRestaurants(location)
  }

  cleanRestaurants() {
    this.setRestaurants()
  }

  getRestaurant(rstID: string): Promise<Restaurant> {
    return new Promise<Restaurant>(async (resolve, reject) => {
      let restaurant: Restaurant | undefined
      if (!!this.restaurants) {
        restaurant = this.restaurants.find((val) => val.id === rstID)
      } else {
        const restaurants = await this.fetchRestaurants()

        restaurant = restaurants.find((val) => val.id === rstID)
      }
      if (!!restaurant) resolve(restaurant)
      reject(new AppError('restaurant not found'))
    })
  }

  private fetchRestaurants(location?: Coordinates): Promise<Restaurant[]> {
    return this._transportLayer.getRestaurants(location).then((restaurants) => {
      this.setRestaurants(restaurants)
      return restaurants
    })
  }

  @action
  private setRestaurants(rest?: Restaurant[]) {
    this.restaurants = rest
  }
}
