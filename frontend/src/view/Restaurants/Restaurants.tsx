import React from 'react'

import { useAppContext } from 'AppContext'

export default () => {
  const { _restaurantsStore } = useAppContext()!

  return (
    <React.Fragment>
      <div>
        <h2>Restaurants</h2>
      </div>
      <button onClick={() => _restaurantsStore.getRestaurants()}>Holaaa</button>
    </React.Fragment>
  )
}
