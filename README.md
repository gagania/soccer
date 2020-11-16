1. create database name `soccer` then import soccer.sql
2. change your db setting username, password and database name in config/config.go

2. API access
- show all teams : localhost:8081/api/teams (GET)
- create teams : localhost:8081/api/teams (POST)
- update teams : localhost:8081/api/teams (PUT)
- delete teams : localhost:8081/api/teams/{id} (DELETE)

- show all player : localhost:8081/api/player (GET)
- create player : localhost:8081/api/player (POST)
- update player : localhost:8081/api/player (PUT)
- delete player : localhost:8081/api/player/{id} (DELETE)

3. execute go run main.go


