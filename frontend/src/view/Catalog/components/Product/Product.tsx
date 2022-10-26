import React from 'react'

import Box from 'components/Box/Box'
import { Product } from 'models/Catalog'
import { stringifyMoney } from 'models/Money'

import ProductImg from '../ProductImg/ProductImg'

interface ProductProps {
  product: Product
}

export default ({ product }: ProductProps) => {
  return (
    <Box>
      <ProductImg src={`https://${product.image}`} />
      <p>{product.name}</p>
      <p>{product.description}</p>
      <p>{stringifyMoney(product.price)}</p>
    </Box>
  )
}
