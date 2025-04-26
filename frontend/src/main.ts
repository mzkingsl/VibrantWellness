import { createApp, h, provide } from 'vue'
import App from './App.vue'
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink
} from '@apollo/client/core'
import { DefaultApolloClient } from '@vue/apollo-composable'
import { gql } from 'graphql-tag' 

const httpLink = createHttpLink({
  uri: import.meta.env.VITE_GRAPHQL_ENDPOINT
})

const apolloClient = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache()
})

// was used to manually test when connecting frontend
// apolloClient.query({
//   query: gql`
//     query {
//       states(filter: "a") {
//         name
//         abbreviation
//       }
//     }
//   `
// })
// .then(result => {
//   console.log('manual query SUCCESS:', result)
// })
// .catch(error => {
//   console.error('manual query FAILURE:', error)
// })


createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient)
  },
  render: () => h(App)
  
}).mount('#app')



