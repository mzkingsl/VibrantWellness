<template>
    <div id="map" style="width: 100%; height: 500px; margin-top: 20px;"></div>
  </template>
  
  <script setup lang="ts">
  import { onMounted, ref, defineExpose } from 'vue'
  
  const map = ref<google.maps.Map | null>(null)
  const dataLayer = ref<google.maps.Data | null>(null)
  const stateBounds = ref<Map<string, google.maps.LatLngBounds>>(new Map())
  
  // default view settings
  const defaultCenter = { lat: 39.8283, lng: -98.5795 }
  const defaultZoom = 4
  
  onMounted(() => {
    try {
      const mapElement = document.getElementById('map')
      if (!mapElement) throw new Error('Map container not found')
      
      //creating map instance with default settings
      map.value = new google.maps.Map(mapElement, {
        center: defaultCenter,
        zoom: defaultZoom,
        mapTypeControl: false
      })
  
      // building a data layer on the map and loading US state geojsons
      dataLayer.value = new google.maps.Data({ map: map.value })
      dataLayer.value.loadGeoJson('/us-states.json', {}, features => {
        features?.forEach(feature => {
          const name = feature.getProperty('name')
          const bounds = new google.maps.LatLngBounds()
          const geometry = feature.getGeometry()
          
          // computing each state's geographic bounds
          geometry?.forEachLatLng(latLng => {
            bounds.extend(latLng)
          })
          
          stateBounds.value.set(name, bounds)
        })
      })
  
      // styling
      dataLayer.value.setStyle({
        fillColor: 'transparent',
        strokeWeight: 1,
        strokeColor: '#666',
      })
  
    } catch (error) {
      console.error('map error:', error)
      const fallback = document.getElementById('map')
      if (fallback) fallback.innerHTML = '<p>map error</p>'
    }
  })
  
  function highlightState(stateName: string | null) {
  if (!dataLayer.value || !map.value) return

  // if no state is selected, clear any current highlights
  if (!stateName) {
    resetHighlight()
    return
  }

  // zoom map to state boundaries 
  const bounds = stateBounds.value.get(stateName)
  if (bounds && !bounds.isEmpty()) {
    map.value.fitBounds(bounds, 20)
  }

  dataLayer.value.setStyle(feature => {
    const name = feature.getProperty('name')
    return {
      fillColor: name === stateName ? '#FF0000' : 'transparent',
      fillOpacity: name === stateName ? 0.3 : 0,
      strokeWeight: name === stateName ? 3 : 1,
      strokeColor: name === stateName ? '#FF0000' : '#666',
    }
  })
}

function resetHighlight() {
  // reset to default view
  if (map.value) {
    map.value.setCenter(defaultCenter)
    map.value.setZoom(defaultZoom)
  }
  
  // reset styles
  dataLayer.value?.revertStyle()
  dataLayer.value?.setStyle({
    fillColor: 'transparent',
    strokeWeight: 1,
    strokeColor: '#666',
  })
}
  
  defineExpose({ 
    map,
    dataLayer,
    highlightState,
    resetHighlight 
  })
  </script>