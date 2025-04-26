<template>
  <div>
    <h2>Search for a US State</h2>
    <Typeahead ref="typeaheadRef" @hover-state="highlightState" @clear-state="resetMap" />
    <GoogleMap ref="mapRef"/>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Typeahead from './components/Typeahead.vue'
import GoogleMap from './components/GoogleMap.vue'

const mapRef = ref()
const typeaheadRef = ref()

function highlightState(stateName: string | null) {
  if (mapRef.value?.highlightState) {
    mapRef.value.highlightState(stateName)
  }
}

function resetMap() {
  // call to reset function
  mapRef.value?.resetHighlight()
  typeaheadRef.value?.resetSelectedState()
}
</script>
