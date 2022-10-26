import React, { useState } from 'react'

import { GoogleMap, Marker, useJsApiLoader } from '@react-google-maps/api'
import { Coordinates } from 'models/Coordinates'

const BARCELONA_COORDINATES: Coordinates = { lat: 41.390205, lng: 2.154007 }

interface UseLoadScriptOptions {
  id: string
  googleMapsApiKey: string
}

interface BrowserMapProps {
  onNewPosition?: (point: Coordinates) => void
}

export default ({ onNewPosition }: BrowserMapProps) => {
  const { isLoaded } = useJsApiLoader({
    id: 'google-map-script',
  } as UseLoadScriptOptions)

  const [markerPosition, setMarkerPos] = useState<Coordinates | undefined>(
    undefined,
  )

  const handleOnClick = (e: google.maps.MapMouseEvent) => {
    if (!!e.latLng) {
      const point: Coordinates = {
        lat: e.latLng.lat(),
        lng: e.latLng.lng(),
      }
      setMarkerPos(point)
      onNewPosition && onNewPosition(point)
    }
  }

  return isLoaded ? (
    <GoogleMap
      mapContainerStyle={{ height: '100%', width: '100%' }}
      center={BARCELONA_COORDINATES}
      zoom={14}
      onClick={handleOnClick}
    >
      {/* Child components, such as markers, info windows, etc. */}
      {!!markerPosition && <Marker position={markerPosition}></Marker>}
    </GoogleMap>
  ) : (
    <></>
  )
}
