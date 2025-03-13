# GlobalWebIndex Engineering Challenge

# REST API
The implementation consists of a single REST API that for now is not in the docker-compose,
so it will need to be executed with go run main.go
The project structure follows a layered architecture with controller, service & repository

# Security
The only endpoint that requires a user token is /protected/assets/:assetId
which can be obtained from /authenticate with the following request body:
```
{
    "username": "user",
    "password": "pass"
}
```

# Docker
For now, only the database is in the docker container.

# DB Migrations
Database migrations are handled with Flyway.

# TODO 
-Add unit tests for the service layer

-Add pagination

-review int64 data types and convert to unsigned
