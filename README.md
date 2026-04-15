# MTN Server

MTN is a dynamic Minecraft Server frontend and backend ecosystem. 

## Backend Setup

The backend is built in Go. It retrieves Minecraft server ping/status and parses world data, serving it through an API.

1. Ensure Go 1.22+ is installed.
2. `cd backend`
3. Edit/Create `config.json` with the required sections (Redis, MySQL, Lucky DDNS).
4. `go run ./cmd/server` or build it using `go build -o mtn-backend ./cmd/server`.

## Frontend Setup

The frontend is a Vue 3 + Vite Application. 

1. Ensure Node.js 18+ is installed.
2. `cd frontend`
3. `npm install`
4. `npm run dev` to start a development server locally. The Vite configuration automatically proxies `/api` routes to `localhost:8080`.
5. `npm run build` creates production-optimized static assets in `dist/`.

## Contributing

- Please use the GitHub issue template for reporting bugs or feature requests.
- PRs should ensure `go test ./...` in `backend` and `vue-tsc --noEmit` in `frontend` complete without errors.
