export function drawPolygon(
    map: google.maps.Map,
    path: google.maps.LatLngLiteral[],
    currentPolygon: google.maps.Polygon | null
  ): google.maps.Polygon {
    if (currentPolygon) {
      currentPolygon.setMap(null)
    }
  
    const newPolygon = new window.google.maps.Polygon({
      paths: path,
      strokeColor: '#FF0000',
      strokeOpacity: 0.8,
      strokeWeight: 2,
      fillColor: '#FF0000',
      fillOpacity: 0.35,
    })
  
    newPolygon.setMap(map)
  
    const bounds = new window.google.maps.LatLngBounds()
    path.forEach(coord => bounds.extend(coord))
    map.fitBounds(bounds)
  
    return newPolygon
  }
  