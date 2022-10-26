import React, { useEffect, useRef, useState } from 'react'

import { Catalog } from 'models/Catalog'
import { useParams } from 'react-router-dom'

import {
  CatalogTransportLayerI,
  newCatalogTransportLayer,
} from './CatalogTransportLayer'
import CatalogHttpClient from './services/CatalogHttpClient'

export default () => {
  const { id } = useParams()

  const transportLayer = useRef<CatalogTransportLayerI>(
    newCatalogTransportLayer(new CatalogHttpClient()),
  ).current

  const [loading, setLoading] = useState(true)
  const [catalog, setCatalog] = useState<Catalog | undefined>(undefined)

  useEffect(() => {
    !!id &&
      transportLayer.getCatalog(id).then((response) => {
        setLoading(false)
        setCatalog(response)
      })
  }, [id, transportLayer])

  return (
    <div>
      <h2>Catalog</h2>
      {loading && <h3>Loading</h3>}
      {!!catalog && <p>{JSON.stringify(catalog)}</p>}
    </div>
  )
}
