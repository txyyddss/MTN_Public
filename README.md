# MTN Server

MTN is a dynamic Minecraft Server frontend and backend ecosystem. 

## Backend Setup

The backend is built in Go. It retrieves Minecraft server ping/status and parses world data, serving it through an API.

1. Ensure Go 1.22+ is installed.
2. `cd backend`
3. Edit/Create `config.json` with the required sections (Redis, MySQL, Lucky DDNS).
4. `go run ./cmd/server` or build it using `go build -o mtn-backend ./cmd/server`.

### Whitelist Management

The backend includes optional whitelist management through RCON, a bearer-token API, and a NapCat OneBot 11 WebSocket bot. Keep committed config values as placeholders. For local testing, replace these keys in your private `backend/config.json`:

- `rcon.host`, `rcon.port`, `rcon.password`
- `onebot.enabled`, `onebot.ws_url`, `onebot.token`, `onebot.qq_group_id`
- `whitelist.api_token`, `whitelist.max_per_qq`

For NapCat forward WebSocket server mode, `onebot.ws_url` should point at the WebSocket root, for example `ws://127.0.0.1:3001/`, so the backend can receive group message events and send `send_group_msg` actions on the same connection.

The frontend admin page is available at `/whitelist` and sends the whitelist token only as a browser-side bearer token.

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
