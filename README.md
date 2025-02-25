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