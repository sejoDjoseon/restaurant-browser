import HttpClient from 'libs/HttpClient'
import { Catalog } from 'models/Catalog'

interface CatalogHttpClientI {
  getCatalog: (restaurantID: string) => Promise<Catalog>
}

export default class CatalogHttpClient
  extends HttpClient
  implements CatalogHttpClientI
{
  public constructor() {
    super()
  }

  getCatalog = (restaurantID: string): Promise<Catalog> => {
    return this.axiosInstance
      .get<unknown, Catalog>(`/restaurants/${restaurantID}/catalog`)
      .catch((error) => error)
  }
}
