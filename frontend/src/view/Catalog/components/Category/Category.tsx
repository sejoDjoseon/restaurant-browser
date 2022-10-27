import React from 'react'

import { Category } from 'models/Catalog'

import Product from '../Product/Product'
import { ProductsContainer, ProductsGrid } from '../ScreenLayout/ScreenLayout'

interface CategoryProps {
  category: Category
}

export default ({ category }: CategoryProps) => {
  return (
    <>
      <h2>{category.categoryName}</h2>
      <ProductsContainer>
        <ProductsGrid>
          {category.products.map((product) => (
            <Product key={product.id} product={product} />
          ))}
        </ProductsGrid>
      </ProductsContainer>
    </>
  )
}
