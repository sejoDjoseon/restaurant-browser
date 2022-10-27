import React from 'react'

import Box from 'components/Box/Box'
import { Product } from 'models/Catalog'
import { stringifyMoney } from 'models/Money'

import ProductImg from '../ProductImg/ProductImg'
import { BoxInfo, Price } from './BoxLatout'

interface ProductProps {
  product: Product
}

export default ({ product }: ProductProps) => {
  return (
    <Box>
      <ProductImg src={`https://${product.image}`} />
      <BoxInfo>
        <h4>{product.name}</h4>
        <p>{product.description}</p>
        <Price>
          <p>{stringifyMoney(product.price)}</p>
        </Price>
      </BoxInfo>
    </Box>
  )
}
