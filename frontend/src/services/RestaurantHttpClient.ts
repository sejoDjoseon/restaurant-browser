import HttpClient from 'libs/HttpClient'
import { Coordinates } from 'models/Coordinates'
import { Restaurant, RestaurantResponse } from 'models/Restaurants'

interface RestaurantHttpClientI {
  getRestaurants: (location?: Coordinates) => Promise<Restaurant[]>
}

export class RestaurantHttpClient
  extends HttpClient
  implements RestaurantHttpClientI
{
  public constructor() {
    super()
  }

  getRestaurants = (location?: Coordinates): Promise<Restaurant[]> => {
    let url = '/restaurants'
    if (!!location) {
      url = `${url}?latitude=${location.lat}&longitude=${location.lng}`
    }

    return this.axiosInstance
      .get<unknown, RestaurantResponse[]>(url)
      .then((restaurants): Restaurant[] => restaurants)
      .catch((error) => error)
  }
}
