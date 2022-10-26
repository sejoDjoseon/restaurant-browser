import { Catalog } from 'models/Catalog'

import CatalogHttpClient from './services/CatalogHttpClient'

export interface CatalogTransportLayerI {
  getCatalog: (restaurantID: string) => Promise<Catalog>
}

export const newCatalogTransportLayer = (
  _catalogHttpClient: CatalogHttpClient,
): CatalogTransportLayerI => ({
  getCatalog: _catalogHttpClient.getCatalog,
})
