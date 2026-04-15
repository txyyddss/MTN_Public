# MTN Deployment Guide

## Prerequisites
- Node.js >= 20
- Go >= 1.20
- MySQL Server (for McMMO & Floodgate databases)
- Redis Server (for data caching/temporary persistence)
- Internet access (for DDNS/Lucky sync if configured)

## Backend Deployment (Linux/amd64/arm64)

The backend handles metrics gathering, STUN/DDNS matching through Lucky, and parsing of Minecraft data.

### 1. Build
```bash
cd backend
# For Linux AMD64
GOOS=linux GOARCH=amd64 go build -o mtn_backend ./cmd/server

# For Linux ARM64
GOOS=linux GOARCH=arm64 go build -o mtn_backend_arm64 ./cmd/server
```

### 2. Configuration (`config.json`)
Ensure the `backend/config.json` points to the correct paths:
- `WorldDir`: Your actual Minecraft world path, must contain `/advancements`, `/playerdata`, `/stats`.
- `MySQL`: DSNs to McMMO and Floodgate.
- `Lucky`: API info for connection detection.
- `Redis`: Connection string.

### 3. Running as a Service (systemd)
Create a service file at `/etc/systemd/system/mtn-backend.service`:

```ini
[Unit]
Description=MTNetwork Backend
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/path/to/MTN_Public/backend
ExecStart=/path/to/MTN_Public/backend/mtn_backend
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Reload and Start:
```bash
systemctl daemon-reload
systemctl enable --now mtn-backend
```

## Frontend Deployment (Cloudflare Pages)

The frontend is built using Vue.js and Vite and should be deployed via Cloudflare Pages or Github Actions.

### 1. Build
```bash
cd frontend
npm install
npm run build
```
The output will be inside the `dist/` directory.

### 2. Deploy via Cloudflare Pages
1. Go to your Cloudflare Dashboard and select **Workers & Pages**.
2. Click **Create an Application** -> **Pages** -> **Connect to Git**.
3. Select your repository.
4. Set the build settings:
   - Framework preset: **Vue**
   - Build command: `npm run build`
   - Output directory: `dist`
5. Click **Save and Deploy**.

### 3. API Proxy Setup (Important)
Since standard Vue requires an endpoint to fetch `/api/*`, ensure Cloudflare routes `/api/*` to your backend public server IP/Domain; this is done inside Cloudflare routing or handled via a Reverse Proxy (like NGINX or Lucky) on your main server.

Example NGINX Reverse Proxy for API:
```nginx
location /api/ {
    proxy_pass http://localhost:8080/api/;
    proxy_set_header Host $host;
}
```
