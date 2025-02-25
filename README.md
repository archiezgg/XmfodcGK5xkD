# XmfodcGK5xkD
Library Management System application

## Set up application locally
Requirements: Docker, docker-compose.

In order to have the application spin up with it's database, run:
`docker-compose up --build`

This will ensure that the application is rebuilt everytime, and will also spin up the application container, the database, and establish networking between them. For more information, please see the `docker-compose.yml` configuration.

## API calls
`curl -X POST -vk -d '{"title": "NoCoincidence"}' localhost:8080/addBook`
`curl -X POST -vk -d '{"username": "archie"}' localhost:8080/createBorrower`
`curl -X POST -vk -d '{"bookId": "1", "borrowerId": "1"}' localhost:8080/borrowBook`
`curl -X POST -vk -d '{"borrowerId": "1"}' localhost:8080/getBorrower`
`curl -X POST -vk -d '{"borrowerId": "1"}' localhost:8080/borrowedBooks`

## Missing components
Since I've ran out of time (3 hours), there are components yet to be developed:
- Open Telemetry: I've never actually worked with OT before, so it would have been a couple of hours more to implement. The little dig I did led me to two libraries:
    - https://pkg.go.dev/go.opentelemetry.io/otel/metric#section-readme : I would use this to implement measure metrics
    - https://pkg.go.dev/go.opentelemetry.io/otel/trace: I would use this pkg to implement spans
- Unit tests: I didn't write any unit of code that wouldn't rely on external dependencies. I thought about writing tests for database operations and handlers, but that would require mocking, but again, I decided that the time would be too short for that.
    - On a previous pet project of mine, I've already written meaningful (or so I hope :) ) tests: https://github.com/archiezgg/cashcalc-backend/blob/master/services/pricing_service_test.go
    - But then again, as a DevOps Engineer, I never had the responsibility to write unit tests, so I am far from being decent with that, I'll admit