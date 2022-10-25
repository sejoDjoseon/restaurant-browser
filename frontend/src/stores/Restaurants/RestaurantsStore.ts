import { makeObservable, observable, action } from 'mobx'
import { Restaurant } from 'models/Restaurants'
import { RestaurantHttpClient } from 'services/RestaurantHttpClient'

export default class RestaurantsStore {
  @observable
  restaurants?: Restaurant[]

  constructor(private readonly restaurantHttpClient: RestaurantHttpClient) {
    makeObservable(this)
  }

  getRestaurants() {
    this.restaurantHttpClient.getRestaurants().then((restaurants) => {
      this.setRestaurants(restaurants)
    })
  }

  @action
  private setRestaurants(rest: Restaurant[]) {
    this.restaurants = rest
  }
}
