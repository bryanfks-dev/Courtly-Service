# Courtly-Service

Courtly-Service is a backend service of Courtly: Easily Order Court Application. Courtly-Service built with Golang, with several library like Echo web framework and GORM.

## Endpoint List

There are several available endpoints, such as:

- **POST** `/api/v1/auth/login` - Sign user with an existing account
- **POST** `/api/v1/auth/register` - Register a new user account
- **POST** `/api/v1/auth/logout` - Remove user from authenticated user status
- **POST** `/api/v1/auth/verify-password` - Verify current user password
- **GET** `/api/v1/users/me` - Get current user information from the database
- **GET** `/api/v1/users/:id` - Get user information from the database using user id
- **PATCH** `/api/v1/users/me/password` - Update user password with a new password
- **PATCH** `/api/v1/users/me/username` - Update user username with a new available username

## Response

Since this service relies on REST API, so json is pretty much needed here. All endpoints have the same json response structure, which shown as an example below.

```json
{
  "success": true,
  "message": "...",
  "data": {
    "..."
  }
}
```

From the example, we can see there are 3 properties in the response body, which will be explained below.

- `"success"` - The success status of the response, this could be either true or false.
- `"message"` - The message of the response, this could be either a success message, failed message, or error messages for inputs
- `"data"` - The passed data to the frontend server. The data structure could be different for each endpoint response, so make sure to read the response carefully

## Run Proccess

To start the service, make sure to follow the steps below.

1. Make sure you're in the root of the project. If not, simply run the command below.

```bash
cd courtly-service
```

2. Clone the environment file

```bash
cp .env.example .env
```

3. Setup the configuration inside the environment file.<br>
   Open the environment file and start editing the configuration, like database connection configuration, jwt, and server.

4. Start the service.<br>
   To start the service simply run `go run cmd/main.go`, or to make things easier, run the `start.bash` script via terminal with `./start.bash` command.
