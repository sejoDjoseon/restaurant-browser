import React from 'react'

import { Route, Routes } from 'react-router-dom'
import Catalog from 'view/Catalog/Catalog'
import Restaurants from 'view/Restaurants/Restaurants'

export default () => {
  return (
    <Routes>
      <Route path="/" element={<Restaurants />} />

      <Route path="/catalog" element={<Catalog />} />
    </Routes>
  )
}
