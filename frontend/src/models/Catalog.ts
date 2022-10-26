import { MoneyValue } from './Money'

export interface Product {
  id: string
  name: string
  description: string
  image: string
  price: MoneyValue
}

export interface Category {
  category: string
  products: Product[]
}

export type Catalog = Category[]
