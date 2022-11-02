import React, { createContext, useContext } from 'react'

import { RestaurantsStoreI } from 'stores/Restaurants/RestaurantsStore'

export interface AppContextI {
  _restaurantsStore: RestaurantsStoreI
}

const AppContext = createContext<AppContextI | null>(null)

export const useAppContext = () => {
  return useContext(AppContext)
}

/* HOC to inject store to any functional or class component */
export const withAppContext =
  <P extends object>(Component: React.ComponentType<P>) =>
  (props: P) => {
    return <Component {...props} _appContext={useAppContext()} />
  }

export default AppContext
