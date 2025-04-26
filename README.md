# US States Map Typeahead

A full-stack demo that lets users search and highlight US states on a Google Map via a live GraphQL/Vue Typeahead interface.

## Tech Stack

- **Backend:** Go, gqlgen (GraphQL server)
- **Frontend:** Vue 3, TypeScript, Vite, Apollo Client
- **Map Integration:** Google Maps JavaScript API (Data layer + GeoJSON)

## Prerequisites

- Go 1.18+  
- Node.js 14+ and npm (or yarn/pnpm)  
- A Google Maps API key (with Maps and Data layer enabled)

## Setup

1. **Clone the repository**  
   ```bash
   git clone <repository_url>
   cd <repository_folder>
   ```

2. **Environment variables**  
   Create a `.env` file in the project root:
   ```
   VITE_GOOGLE_MAPS_API_KEY=YOUR_GOOGLE_MAPS_KEY
   VITE_GRAPHQL_ENDPOINT=http://localhost:8080/graphql
   ```

## Backend

1. **Generate GraphQL code** (if not already committed)  
   ```bash
   go run github.com/99designs/gqlgen@latest generate
   ```

2. **Run the server**  
   ```bash
   go run .
   ```
   The GraphQL endpoint will be available at `http://localhost:8080/graphql`.

## Frontend

1. **Install dependencies**  
   ```bash
   npm install
   # or yarn install, pnpm install
   ```

2. **Run in development mode**  
   ```bash
   npm run dev
   ```
   Visit `http://localhost:5173` (default Vite port) to view the app.

## Production Build

- **Backend:**  
  ```bash
  go build -o server
  ```

- **Frontend:**  
  ```bash
  npm run build
  ```
  The compiled assets will be in `dist/`.

## Usage

1. Type a US state name (e.g., “California”) in the search box.  
2. Hover over suggestions to zoom and highlight that state on the map.  
3. Click a suggestion to select and lock in the highlight.

## Environment & Deployment

- Keep your `.env` out of version control (it’s in `.gitignore`).  
- Use different API keys or endpoints for staging/production by swapping `.env` files.

## Troubleshooting

- If the map doesn’t load, check the browser console for API key or network errors.  
- Ensure your GraphQL server is running before launching the frontend.

## License

MIT
