# XmfodcGK5xkD
Library Management System application

## API calls
`curl -X POST -vk -d '{"title": "NoCoincidence"}' localhost:8080/addBook`
`curl -X POST -vk -d '{"username": "archie"}' localhost:8080/createBorrower`
`curl -X POST -vk -d '{"bookId": "1", "borrowerId": "1"}' localhost:8080/borrowBook`