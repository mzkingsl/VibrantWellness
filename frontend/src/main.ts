// import { createApp } from 'vue'
// import App from './App.vue'
// import {
//   ApolloClient,
//   InMemoryCache,
//   createHttpLink
// } from '@apollo/client/core'
// import { DefaultApolloClient } from '@vue/apollo-composable'

// // Make sure your environment variable is set correctly
// // It should be something like http://localhost:8080/query
// const httpLink = createHttpLink({
//   uri: import.meta.env.VITE_GRAPHQL_ENDPOINT || 'http://localhost:8080/query'
// })

// const apolloClient = new ApolloClient({
//   link: httpLink,
//   cache: new InMemoryCache(),
//   defaultOptions: {
//     query: {
//       fetchPolicy: 'network-only' // Ensures fresh data
//     }
//   }
// })

// const app = createApp(App)

// // Provide the Apollo client to the entire app
// app.provide(DefaultApolloClient, apolloClient)

// app.mount('#app')
import { createApp, h, provide } from 'vue'
import App from './App.vue'
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink
} from '@apollo/client/core'
import { DefaultApolloClient } from '@vue/apollo-composable'
import { gql } from 'graphql-tag' // if you want manual query testing

const httpLink = createHttpLink({
  uri: import.meta.env.VITE_GRAPHQL_ENDPOINT
})

const apolloClient = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache()
})


apolloClient.query({
  query: gql`
    query {
      states(filter: "a") {
        name
        abbreviation
      }
    }
  `
})
.then(result => {
  console.log('Apollo manual query SUCCESS:', result)
})
.catch(error => {
  console.error('Apollo manual query FAILURE:', error)
})


createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient)
  },
  render: () => h(App)
  
}).mount('#app')



