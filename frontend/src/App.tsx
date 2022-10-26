import React, { useRef } from 'react'

import AppContext, { AppContextI } from 'AppContext'
import Header from 'components/Header/Header'
import AppRouter from 'components/Router/AppRouter'
import { BrowserRouter } from 'react-router-dom'
import { RestaurantHttpClient } from 'services/RestaurantHttpClient'
import RestaurantsStore from 'stores/Restaurants/RestaurantsStore'
import { newRestaurantsTransportLayer } from 'stores/Restaurants/RestaurantsTransportLayer'
import './App.css'

function App() {
  const appContext = useRef<AppContextI>({
    _restaurantsStore: new RestaurantsStore(
      newRestaurantsTransportLayer(new RestaurantHttpClient()),
    ),
  })

  return (
    <BrowserRouter>
      <div className="App">
        <Header>Restaurants Browser</Header>
        <AppContext.Provider value={appContext.current}>
          <div className="container">
            <AppRouter />
          </div>
        </AppContext.Provider>
      </div>
    </BrowserRouter>
  )
}

export default App
