REST API using Docker, Go Chi and Postgres - [Link](https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/)


## REST API
- `POST /items` to add a new item to the list
- `GET /items` to fetch all existing items in the list
- `GET /items/{itemId}` to fetch a single item from the list using its ID
- `PUT /items/{itemId}` to update an existing item
- `DELETE /items/{itemId}` to delete an item from the list

## Directory structure
- `db`: The code here is responsible for interacting directly with our database. This way, the database engine is properly separated from the rest of the application
- `handler`: The handler package creates and handles our API routes using chi
- `models`: Contains Go structs that can be bounded to database objects or transformed into their equivalent JSON format
- The Dockerfile defines the base image and commands required to have our API server up and running. The docker-compose.yml defines our app dependencies (the server using the Dockerfile and the database using the official postgres docker image). The Docker website has a detailed reference for both Dockerfiles and docker-compose
- `.env`: This holds our application environment variables (such as database credentials)
- `main.go` is our application entry point. It will be responsible for reading environment variables, setting up the database as well as starting and stopping the API server
