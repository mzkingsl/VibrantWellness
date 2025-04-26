<template>
    <div class="typeahead">
      <input
        type="text"
        v-model="inputValue"
        @input="searchStates"
        placeholder="Search for a state"
        class="input"
      />
      <ul v-if="inputValue.trim().length > 0 && suggestions.length" class="dropdown" @mouseleave="handleMouseLeave">
        <li
            v-for="state in suggestions"
            :key="state.abbreviation"
            @mouseover="handleMouseOver(state.name)"
            @mouseleave="handleMouseLeave"
            @mousedown.prevent="selectState(state.name)"
        >
            {{ state.name }} ({{ state.abbreviation }})
        </li>
      </ul>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, watch } from 'vue'
  import { useQuery } from '@vue/apollo-composable'
  import { gql } from 'graphql-tag'
  
  const emit = defineEmits(['hover-state', 'clear-state', null])
  
  const inputValue = ref('')
  const suggestions = ref<Array<{name: string, abbreviation: string}>>([])
  const allStates = ref<Array<{name: string, abbreviation: string}>>([])
  const selectedState = ref<string | null>(null)
  
  const STATES_QUERY = gql`
    query AllStates($filter: String!) {
      states(filter: $filter) {
        name
        abbreviation
      }
    }
  `
  
  const { result } = useQuery(STATES_QUERY, () => ({
    filter: inputValue.value
  }))
  
  watch(result, (newValue) => {
    if (newValue?.states) {
      allStates.value = newValue.states
      suggestions.value = newValue.states
    }
  })
  
  function searchStates() {
    const query = inputValue.value.toLowerCase()
    if (!query) {
      // when the input becomes empty, fire hover-state:null
      suggestions.value = []
      emit('hover-state', null)
      return
    }
    suggestions.value = allStates.value.filter(state =>
      state.name.toLowerCase().includes(query)
    )
  }
  
  function handleMouseOver(stateName: string) {
    emit('hover-state', stateName)
  }
  
  function handleMouseLeave() {
    // whenever the mouse leaves the dropdown or an item, reset the map
   emit('hover-state', null)
 }


  
  function selectState(stateName: string) {
    inputValue.value = stateName
    selectedState.value = stateName
    suggestions.value = []
    emit('hover-state', stateName)
  }
  
  function resetSelectedState() {
  inputValue.value = ''
  selectedState.value = null
  // force immediate reset
  emit('hover-state', null)
}
  
  defineExpose({
    resetSelectedState
  })
  </script>
  

  
  
  <style scoped>
  .typeahead {
    width: 300px;
    margin: 1em auto;
    position: relative;
  }
  .input {
    width: 100%;
    padding: 8px;
    font-size: 16px;
  }
  .dropdown {
    list-style: none;
    padding: 0;
    margin: 0;
    border: 1px solid #ccc;
    position: absolute;
    background: white;
    width: 100%;
    z-index: 10;
  }
  .dropdown li {
    padding: 8px;
    cursor: pointer;
  }
  .dropdown li:hover {
    background: #f0f0f0;
  }
  </style>
  