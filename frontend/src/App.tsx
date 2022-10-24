import React from 'react'

import Header from 'components/Header'
import AppRouter from 'components/Router/AppRouter'
import { BrowserRouter, Link } from 'react-router-dom'

import './App.css'

function App() {
  return (
    <BrowserRouter>
      <div className="App">
        <Header>
          <div>
            <Link to="/">Restaurants</Link> | <Link to="/catalog">Catalog</Link>
          </div>
        </Header>

        <div className="App-header">
          <AppRouter />
        </div>
      </div>
    </BrowserRouter>
  )
}

export default App
