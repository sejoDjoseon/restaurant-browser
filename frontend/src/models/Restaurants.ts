import { Coordinates } from './Coordinates'

export interface RestaurantResponse {
  id: string
  name: string
  image: string
  location: Coordinates
  open: boolean
}

export interface Restaurant {
  id: string
  name: string
  image: string
  location: Coordinates
  open: boolean
}
