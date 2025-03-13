# GlobalWebIndex Engineering Challenge

# REST API
The implementation consists of a single REST API.

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
Database and go app are in a docker-compose file.

# DB Migrations
Database migrations are handled with Flyway.

# TODO 
-Add unit tests for the service layer

-Add pagination

-review int64 data types and convert to unsigned
