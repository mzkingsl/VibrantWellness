/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_GRAPHQL_ENDPOINT: string
    // add more env variables here if needed
  }
  
  interface ImportMeta {
    readonly env: ImportMetaEnv
  }
  