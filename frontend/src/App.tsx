import React, { useRef } from 'react'

import AppContext, { AppContextI } from 'AppContext'
import Header from 'components/Header/Header'
import AppRouter from 'components/Router/AppRouter'
import { BrowserRouter, Link } from 'react-router-dom'
import { RestaurantHttpClient } from 'services/RestaurantHttpClient'
import RestaurantsStore from 'stores/Restaurants/RestaurantsStore'
import './App.css'

function App() {
  const appContext = useRef<AppContextI>({
    _restaurantsStore: new RestaurantsStore(new RestaurantHttpClient()),
  })

  return (
    <BrowserRouter>
      <div className="App">
        <Header>
          <div>
            <Link to="/">Restaurants</Link> | <Link to="/catalog">Catalog</Link>
          </div>
          Restaurants Browser
        </Header>
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
