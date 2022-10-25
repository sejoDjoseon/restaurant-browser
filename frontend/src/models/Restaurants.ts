import { Coordinates } from './Coordinates'

export interface RestaurantResponse {
  id: string
  name: string
  image: string
  location: Coordinates
}

export interface Restaurant {
  id: string
  name: string
  image: string
  location: Coordinates
}
