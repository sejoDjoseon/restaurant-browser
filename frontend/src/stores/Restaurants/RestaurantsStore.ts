import { makeObservable, observable, action } from 'mobx'
import { Coordinates } from 'models/Coordinates'
import { Restaurant } from 'models/Restaurants'

import { RestaurantsTransportLayerI } from './RestaurantsTransportLayer'

export default class RestaurantsStore {
  @observable
  restaurants?: Restaurant[]

  constructor(private readonly _transportLayer: RestaurantsTransportLayerI) {
    makeObservable(this)
  }

  getRestaurants(location?: Coordinates) {
    this._transportLayer.getRestaurants(location).then((restaurants) => {
      this.setRestaurants(restaurants)
    })
  }

  cleanRestaurants() {
    this.setRestaurants()
  }

  @action
  private setRestaurants(rest?: Restaurant[]) {
    this.restaurants = rest
  }
}
