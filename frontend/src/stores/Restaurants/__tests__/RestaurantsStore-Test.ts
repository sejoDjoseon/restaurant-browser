import { AppError } from 'models/AppError'
import { Restaurant } from 'models/Restaurants'

// eslint-disable-next-line jest/no-mocks-import
import { RestaurantsTransportLayerMock } from '../__mocks__/RestaurantsTransportLayer-Mock'
import RestaurantsStore from '../RestaurantsStore'

const rest1: Restaurant = {
  id: 'id1',
  name: 'string',
  image: 'string',
  location: {
    lat: 0,
    lng: 0,
  },
  open: true,
}

describe('RestaurantsStore_getRestaurant', () => {
  let store: RestaurantsStore

  describe('getRestaurant when restaurant list is stored', () => {
    beforeEach(() => {
      store = new RestaurantsStore(new RestaurantsTransportLayerMock())
      store.restaurants = [rest1]
    })
    test("exists then returns correct, doesn't need fetch", async () => {
      await store.getRestaurant('id1').then((rest) => {
        expect(rest).toEqual(rest1)
      })
    })
    test("not exists then returns error, doesn't need fetch", async () => {
      await expect(store.getRestaurant('id2')).rejects.toThrow(
        new AppError('restaurant not found'),
      )
    })
  })
  describe('getRestaurant when restaurants are not stored', () => {
    beforeEach(() => {
      const tlMock = new RestaurantsTransportLayerMock()
      tlMock.getRestaurants.mockImplementation(
        () =>
          new Promise((resolve) => {
            resolve([rest1])
          }),
      )
      store = new RestaurantsStore(tlMock)
      store.restaurants = undefined
    })
    test('exists returns correct, on resolve fetch', async () => {
      await store.getRestaurant('id1').then((rest) => {
        expect(rest).toEqual(rest1)
      })
    })
    test('not exists returns error, on resolve fetch', async () => {
      await expect(store.getRestaurant('id2')).rejects.toThrow(
        new AppError('restaurant not found'),
      )
    })
  })
})
