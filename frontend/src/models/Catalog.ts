import { MoneyValue } from './Money'

export interface ProductResponse {
  id: string
  name: string
  description: string
  image: string
  price: MoneyValue
}

export interface Product {
  id: string
  name: string
  description: string
  image: string
  price: MoneyValue
}

export interface CategoryResponse {
  category: string
  products: Product[]
}

export interface Category {
  categoryName: string
  products: Product[]
}

export type Catalog = Category[]
