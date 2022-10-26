import HttpClient from 'libs/HttpClient'
import { Catalog, Category, CategoryResponse } from 'models/Catalog'

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
      .get<unknown, CategoryResponse[]>(`/restaurants/${restaurantID}/catalog`)
      .then((categories) => {
        return categories.map<Category>((category) => ({
          categoryName: category.category,
          products: category.products,
        }))
      })
      .catch((error) => error)
  }
}
