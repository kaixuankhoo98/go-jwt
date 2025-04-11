# go-jwt

This project is a backend service written in Go that implements JSON Web Token (JWT) authentication for a Gym tracking app. It provides secure user authentication and authorization, and the logic to handle creating workouts. 

## Features

- User registration and login.
- JWT-based authentication and token generation.
- Middleware for protected routes.
- Token validation and expiration handling.
- Workout/Exercise/Sets models, services and controllers, behind protected endpoints.

## Installation

1. Clone the repository and navigate to the file:
    ```bash
    git clone https://github.com/<your-username>/go-jwt.git

    cd go-jwt
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Copy the `.env.dist` file to a `.env` file and adjust as appropriate.

3. Create a database instance and connect it to the app via environment variables `.env`. By default, a postgres server is used, but this can be updated in `go.mod` and `dbConnection.go`. For example, a connection string might look like:
    ```
    "host=<hostname> user=test password=password dbname=database port=5432 sslmode=disable"
    ```

4. Run the application:
    ```bash
    go run main.go
    ```
    
    Or for easier development with rebuilding on save, use:
    ```bash
    compiledaemon --command="./go-jwt"
    ```

## Usage

- Use the `/register` endpoint to create a new user.
- Use the `/login` endpoint to authenticate and receive a JWT.
- Access protected routes by including the JWT in the `Authorization` header.
- Protected routes are setup in the `routes/` folder
