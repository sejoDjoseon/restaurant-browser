import { Restaurant } from 'models/Restaurants'

import { RestaurantsTransportLayerI } from '../RestaurantsTransportLayer'

export class RestaurantsTransportLayerMock
  implements RestaurantsTransportLayerI
{
  getRestaurants: jest.Mock<Promise<Restaurant[]>> = jest.fn()
}
