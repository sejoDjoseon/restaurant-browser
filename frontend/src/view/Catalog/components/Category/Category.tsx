import React from 'react'

import { Category } from 'models/Catalog'

import Product from '../Product/Product'

interface CategoryProps {
  category: Category
}

export default ({ category }: CategoryProps) => {
  return (
    <>
      <h4>{category.categoryName}</h4>
      {category.products.map((product) => (
        <Product key={product.id} product={product} />
      ))}
    </>
  )
}
