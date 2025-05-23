# Documentation
## API
Use `docs/swagger.json` to see server methods
## Start
### Prepare `.env` file from `.env.example`
```
# Optional
LOG_LEVEL=4 # 1 - Fatal Error, 2 - Error, 3 Warning, 4 - Info, 5 - Debug, 3 by default
# Required
PORT=8080 # Server port

# Optional
MUSIC_INFO_URL= # Url to get data of new song, example: https://google.com

# Required (all), Postgres
DB_USER=user 
DB_PASS=pass 
DB_NAME=name
DB_PORT=5353
DB_HOST=localhost
```
### Launch
```
go run cmd/server/main.go
```