import { render } from '@testing-library/react'
import AppContext from 'AppContext'
import { Coordinates } from 'models/Coordinates'
import { Restaurant } from 'models/Restaurants'
import { act } from 'react-dom/test-utils'
// eslint-disable-next-line jest/no-mocks-import
import { RestaurantsStoreMock } from 'stores/Restaurants/__mocks__/RestaurantsStore-Mock'

import Restaurants from '../Restaurants'

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

const mockOnNewPosition = jest.fn()

jest.mock(
  '../components/BrowserMap/BrowserMap',
  () => (props: { onNewPosition?: (point: Coordinates) => void }) => {
    props.onNewPosition &&
      mockOnNewPosition.mockImplementation(props.onNewPosition)
    const MapMock = (_props: any) => <>'browser-map-component-mock'</>
    return <MapMock {...props} />
  },
)

describe('RestaurantsTest', () => {
  test('on new position, call to getRestaurants with location', () => {
    const rStore = new RestaurantsStoreMock()
    rStore.getRestaurants.mockImplementation(
      () =>
        new Promise((resolve) => {
          resolve([rest1])
        }),
    )
    render(
      <AppContext.Provider value={{ _restaurantsStore: rStore }}>
        <Restaurants />
      </AppContext.Provider>,
    )

    expect(rStore.getRestaurants).toHaveBeenCalled()

    const point: Coordinates = { lat: 1, lng: 1 }
    act(() => {
      mockOnNewPosition(point)
    })

    expect(rStore.getRestaurants).toHaveBeenCalledWith(point)
  })
})
