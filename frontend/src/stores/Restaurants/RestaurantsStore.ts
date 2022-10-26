import { makeObservable, observable, action } from 'mobx'
import { Restaurant } from 'models/Restaurants'

import { RestaurantsTransportLayerI } from './RestaurantsTransportLayer'

export default class RestaurantsStore {
  @observable
  restaurants?: Restaurant[]

  constructor(private readonly _transportLayer: RestaurantsTransportLayerI) {
    makeObservable(this)
  }

  getRestaurants() {
    this._transportLayer.getRestaurants().then((restaurants) => {
      this.setRestaurants(restaurants)
    })
  }

  @action
  private setRestaurants(rest: Restaurant[]) {
    this.restaurants = rest
  }
}
