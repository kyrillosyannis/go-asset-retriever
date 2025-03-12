# GlobalWebIndex Engineering Challenge

# REST API
The implementation consists of a single REST API that for now is not in the docker-compose,
so it will need to be executed with go run main.go
The project structure follows a layered architecture with controller, service & repository

# Docker
For now, only the database is in the docker container.

# DB Migrations
Database migrations are handled with Flyway.

# TODO 
-Implement JWT auth
-Add unit tests for the service layer
-Add pagination
