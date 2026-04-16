

**System Role:** Act as an Expert Full-Stack Developer (Vue.js, Golang) and Minecraft Server Architect. 

**Project Goal:** Build the official website and backend infrastructure for "MTNetwork", a Vanilla-based Minecraft survival server. The project requires reading local server files (NBT data), querying MySQL databases (McMMO, Floodgate), and integrating with APIs, presented on a highly responsive, animated frontend.

---

### 1. Architecture & Tech Stack
* **Frontend:** Vue.js (Follow Vue best practices, modular components).
* **Backend:** Golang (Modularized, pro-level structure).
* **Temporary Data/Caching:** Redis (Cache heavy calculations like NBT parsing).
* **Databases:** MySQL (Read-only access to McMMO and Floodgate tables).
* **Configuration:** `config.json` for all environment variables and parameters.
* **Deployment (Frontend):** Cloudflare Pages.
* **CI/CD (Backend):** GitHub Actions (Compile for Linux `amd64` and `arm64`).
* **Backend Runner:** `systemctl` with a Bash script wrapper.

### 2. Core Engineering Principles
* **Offload Processing:** If data can be processed or formatted on the frontend (e.g., formatting large numbers with K, M, B), do it on the frontend.
* **Performance:** Preload necessary data to reduce disk I/O on the server.
* **Responsiveness:** The layout must adapt dynamically to device sizes. The Navigation Bar must be expanded on desktops, partially folded on tablets, and a hamburger/dropdown menu on mobile.
* **Immersive UX:** Add micro-interactions and reactions to user inputs.

### 3. Backend Requirements (Golang)
* **Modular Structure:** Separate routes, controllers, services, and utilities. 
* **File Reading:** Read from the "World" directory (`world/advancements`, `world/playerdata`, `world/stats`). **Must include a robust NBT parser** to read `.dat` files.
* **Polling API:** A `/api/status` endpoint that refreshes server status every 5 seconds (configurable).
* **Configuration File:** Create a strict JSON config that includes:
    * World directory path.
    * "Lucky" integration configurations.
    * Server addresses (Java IPv6, Bedrock IPv6, Java IPv4 SRV, Bedrock IPv4).
    * Redis URI/credentials.
    * MySQL credentials for McMMO and Floodgate.

### 4. Frontend Views (Vue.js)

**A. Global Animations**
* **Full-page Load:** Building the letters "MTN" with different colored blocks.
* **Asset Loaders:** Colorful letters "MTN" jumping through each other.
* **Scroll Animations:** Text and elements should slide/fade in dynamically as the user scrolls.

**B. Main Page (`/`)**
* **Theme:** Describe the goal of using low-end hardware and a zero-cost homeserver to create the smoothest MC server.
* **Positioning:** Vanilla survival server with strong anti-cheat and small QoL modifications.
* **Values:** Non-profit, open long-term, no VIP discrimination, no admin abuse.
* **Connectivity:** Supports Java and Bedrock, including account binding.

**C. Player List & Search (`/players`)**
* **Default View:** List players active in the last 7 days (read from `Paper.LastSeen`). Display total active player count.
* **Search:** Fuzzy search by UUID or Name.
* **Features:** "Random Player" button.
* **UI:** 2D Head icon (delete "be_" prefix for Bedrock skins), player name, and small UUID below it.

**D. Player Profile (`/player/{uuid}`)**
* **UI:** Display an interactive 3D skin model of the user.
* **Data to extract from NBT/Databases:**
    * Basic info: Last activity, First join (`bukkit.firstPlayed`), Played hours (`Spigot.ticksLived` / 20 / 3600).
    * Linked Java/Bedrock account (from Floodgate MySQL).
    * Advancements (categorized, hide locked by default).
    * Custom stats (`minecraft:custom`).
    * McMMO status (from McMMO MySQL).
    * Ores Mining Status (mined minus used) displayed as a **Pie Chart**.

**E. Leaderboards (`/leaderboards`)**
* Each row needs: Player Head, Name, Rank, and Stat. Clicking a row jumps to their profile.
* **Categories:**
    * Skills: `mcmmo_skills-total`
    * Playtime: `minecraft:play_time`
    * Mining: Sum of all items in `minecraft:mined`
    * Killing: `minecraft:mob_kills`
    * Deaths: `minecraft:deaths`
    * Walking: `minecraft:walk_one_cm`
    * PVP: `minecraft:player_kills`

**F. Connection Guide (`/connect`)**
* Display Java IPv6, Bedrock IPv6, and Java IPv4 SRV from config.
* Fetch Bedrock IPv4 and Java IPv4 dynamically via the "Lucky" API (display all STUNs, keeping original names). Display the matching DDNS domain, IP, and port.
* Provide clear instructions for Bedrock and Java clients.
* Recommend IPv6 over IPv4 to the user.

**G. Work-in-Progress / Redirect Pages**
* `/server-intro`, `/core-members`, `/gallery` -> Display "Work in Progress" UI.
* `/wiki` -> External redirect to `https://wiki.mcmtn.net/`.

### 5. Deliverables & Execution Instructions
Because this is a large project, please output the response in the following steps. Wait for my confirmation between each step.
1.  **Step 1:** Provide the complete directory structure for both the Golang Backend and the Vue.js Frontend.
2.  **Step 2:** Provide the backend configuration structure (`config.json`), the GitHub Actions YAML, and the bash/systemctl deployment scripts.
3.  **Step 3:** Provide the Golang core logic (NBT parsing, Database models, and API Handlers).
4.  **Step 4:** Provide the Vue.js frontend routing, main components, and the 3D skin/animation logic.
5.  **Step 5:** Provide Markdown templates for frontend/backend deployment instructions and an Issue Submission form template.

