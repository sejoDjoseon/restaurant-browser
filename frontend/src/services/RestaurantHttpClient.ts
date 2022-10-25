import HttpClient from 'libs/HttpClient/HttpClient'
import { Restaurant, RestaurantResponse } from 'models/Restaurants'

interface RestaurantHttpClientI {
  getRestaurants: () => Promise<Restaurant[]>
}

export class RestaurantHttpClient
  extends HttpClient
  implements RestaurantHttpClientI
{
  public constructor() {
    super()
  }

  getRestaurants = (): Promise<Restaurant[]> => {
    return this.axiosInstance
      .get<unknown, RestaurantResponse[]>('/restaurants')
      .then((restaurants): Restaurant[] => restaurants)
      .catch((error) => error)
  }
}
