import HttpClient from 'libs/HttpClient'
import { Catalog, Category, CategoryResponse } from 'models/Catalog'

interface CatalogHttpClientI {
  getCatalog: (restaurantID: string, filter?: string) => Promise<Catalog>
}

export default class CatalogHttpClient
  extends HttpClient
  implements CatalogHttpClientI
{
  public constructor() {
    super()
  }

  getCatalog = (restaurantID: string, filter?: string): Promise<Catalog> => {
    let url = `/restaurants/${restaurantID}/catalog`
    if (!!filter) {
      url = `${url}?query=${filter}`
    }

    return this.axiosInstance
      .get<unknown, CategoryResponse[]>(url)
      .then((categories) => {
        return categories.map<Category>((category) => ({
          categoryName: category.category,
          products: category.products,
        }))
      })
      .catch((error) => error)
  }
}
