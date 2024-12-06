# Courtly-Service

Courtly-Service is a backend service of Courtly: Easily Order Court Application. Courtly-Service built with Golang, with several library like Echo web framework and GORM.

## Endpoint List

There are several available endpoints, such as:

##### Auth endpoints

- **POST** `/api/v1/auth/user/login` - Sign user with an existing account
- **POST** `/api/v1/auth/user/register` - Register a new user account
- **POST** `/api/v1/auth/user/logout` - Remove user from authenticated status
- **POST** `/api/v1/auth/user/verify-password` - Verify current user password
- **POST** `/api/v1/auth/vendor/verify-password` - Verify current vendor password
- **POST** `/api/v1/auth/vendor/login` - Sign vendor with an existing account
- **POST** `/api/v1/auth/vendor/logout` - Remove vendor from authenticated status

##### Users endpoints

- **GET** `/api/v1/users/me` - Get current user information from the database
- **PATCH** `/api/v1/users/me/password` - Update user password with a new password
- **PATCH** `/api/v1/users/me/username` - Update user username with a new available username

##### Vendors endpoints

- **GET** `/api/v1/vendors/me` - Get current vendor information from database
- **PATCH** `/api/v1/vendors/me/password` - Update vendor password with a new password

##### Booking endpoints

- **GET** `/api/v1/users/me/bookings` - Get current user bookings from database

##### Orders endpoints

- **GET** `/api/v1/vendors/me/orders` - Get current vendor orders from database
- **GET** `/api/v1/vendors/me/orders/stats` - Get current vendor orders stats from database

##### Courts endpoints

- **GET** `/api/v1/courts` - Get all available courts from database
- **GET** `/api/v1/courts/:id` - Get court finromation from the database using court id
- **GET** `/api/v1/courts/types/:type` - Get available courts using type from database
- **GET** `/api/v1/vendors/me/courts/types/:type` - Get current vendor courts using court type from database
- **POST** `/api/v1/vendors/me/courts/types/:type/new` - Create a new court for a court type
- **POST** `/api/v1/vendors/me/courts/types/:type` - Create a new court for a court type from the existing court
- **GET** `/api/v1/vendors/me/courts/stats` - Get current vendor courts stats from database

##### Reviews endpoints

- **GET** `/api/v1/vendors/:id/courts/types/:type/reviews` - Get vendor courts type reviews from database
- **GET** `/api/v1/vendors/me/reviews` - Get court reviews related to the vendor from database
- **POST** `/api/v1/vendors/:vendorID/courts/:courtID/reviews` - Create a new review for current vendor and court type

## Response

Since this service relies on REST API, so json is pretty much needed here. All endpoints have the same json response structure **(in general)**, which shown as an example below.

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

To see responses structure in details, please read this doc.

---

[![visit-docs](https://img.shields.io/badge/visit-response--docs-blue)](https://github.com/bryanfks-dev/Courtly-Service/blob/main/docs/RESPONSES.md)

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
