# Goal
Your goal it to write MTNetwork webpages with backend. Some criterias below:
- Frontend: Vue.js
- Backend: GOlang
- Frontend deploy: cloudflare pages
- Backend Compile: Github actions
- Compile for: linux, arm64 and amd64
- Config files: json
- Temp data storage: Redis
- Run backend using: systemctl, bash script
- Follow the skills frontend-design, golang-pro, vue-best-practices

Programming criterias:
- All parameters are configurable
- If data can process on frontend, dont do it backend
- Preload all the datas to reduce disk usage
- Add animations and reactions to user everywhere on frontend
- Frontend should dynamically adjust web layout based on the size
- The bar should display in different layout for different size of device (e.g. expanded for laptop, partially folded for ipads, use drop down menu for phones)
- Structurize the folder and keep it clean
- Modular the backend
- Large datas should be formatted with K,M,etc. on frontend with view origin button

# Main page
Refer to .\old_website\index.html, add following descriptions:
- Goal: using rubbish hardware and reletively zero-cost homeserver to create the smoothest MC Server
- Positioning: Valina survival server with strong anticheat and small modifications to enhance playment
- No discriminate with VIPS and players, no admin privilage
- Supports Java and Bedrock connection, support Bedrock & Java account binding
- Stable and opens for long time due to low-cost
- It is a non-profit server

# Server intro
Make the page "Work in progress"
# Core members
Make the page "Work in progress"
# Player list & info
## Things you may need
- User NBT display (playerdata)
    - XpLevel
    - Spigot.ticksLived (/20 equals seconds)
    - foodLevel
    - Health
    - Paper.LastSeen (ms)
    - bukkit.firstPlayed (ms)
    - bukkit.lastKnownName
## The page
Default display a list of player on the web who are active recent 7 days (configurable), read "Paper.LastSeen", and display the player count
User could fuzzy search a player based on uuid or name
Add button to random see a player's info
2D Head should be displayed as icon and name with small uuid below should be below the icon

* delete "be_" prefix when getting player skin
## Player info
It should be at https://{domain}/player/{uuid}
Things should be displayed:
- 3D skin model of the user
- Last activity
- First join
- Played hours
- Linked Java or bedrock account (if present, through reading floodgate database)
- Advancements (with category, default disable locked advancements)
- Stats in minecraft:custom
- Mcmmo status (through reading mcmmo database)
- Ores mining status (mined-used), using pie chart

# Leaderboards
Every line should display Player head, player name, their ranking, and the data, user can click to jump to that player's info
- Skills: Ranking of mcmmo_skills-total
- Playtime: Ranking of minecraft:play_time
- Mining: Ranking of sum of all items in minecraft:mined
- Killing: Ranking of minecraft:mob_kills
- Deaths: Ranking of minecraft:deaths
- Walking: Ranking of minecraft:walk_one_cm
- PVP: ranking of minecraft:player_kills

# Server status
Refreshes per 5 seconds(confgurable) at backend, 
# Connecting to the server
Display JAVA IPV6, BEDROCK IPV6, JAVA IPV4 SRV in config
Display Bedrock IPV4, Java IPV4 through requesting Lucky (Display all stuns, no need to change name)
    Display the domain (matching ip with DDNS domain), ip, port
Add instructions about how to use it in Bedrock and Java
Tell user to choose IPV6 > IPV4
# Wiki
Jump to https://wiki.mcmtn.net/

# Gallery
Make the page "Work in progress"

# Animations
- Full page loading animation: Building the letter "MTN" with different colours of block
- Assest loading animation: colourful letters "MTN" jumping through each other
- Dynamic animation: Display the words with entering animation when the user scrolls down 
- Other animations that you may want to add

# Backend configs
- "World" directory (world\advancements, world\playerdata, world\stats should be used)
- Lucky related config
- Server address config (JAVA IPV6, BEDROCK IPV6, JAVA IPV4 SRV, BEDROCK IPV4 should display through lucky)
- Redis config
- Mcmmo, floodgate MySQL database config
- Other necessary configs

# Instructions
- You should add instructions about deploying the frontend and backend
- Instructions about submitting issues and issue form template


