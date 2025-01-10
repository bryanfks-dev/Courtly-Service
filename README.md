# Courtly Service

Courtly Service is a backend service of Courtly: Easily Order Court Application. Courtly-Service is designed to facilitate seamless service management for sports court rentals. This server is built to manage user offerings, handle bookings, and connect efficiently.

## Features

- **API for Booking Management**: Manage bookings and availability for sports fields.
- **Vendor Service Management**: APIs for vendors to add and manage their services.
- **Authentication and Authorization**: Secure access with token-based authentication.
- **Scalable Architecture**: Designed for high performance and scalability.

## Technologies Used

- **Programming Language**: Go (Golang)
- **Framework**: Echo and GORM
- **Database**: MySQL
- **Authentication**: JWT (JSON Web Tokens)

## Getting Started

### Requirement(s):

- **Go**: [Install Go](https://go.dev/doc/install).
- **MySQL**: Ensure MySQL is installed and running.
- **ngrok**: Temporary deploy the service server for midtrans payment callback, [Install ngrok](https://download.ngrok.com).

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
- **PATCH** `/api/v1/users/me/profile-picture` - Update user profile picture with a new profile image

##### Vendors endpoints

- **GET** `/api/v1/vendors/me` - Get current vendor information from database
- **PATCH** `/api/v1/vendors/me/password` - Update vendor password with a new password

##### Orders endpoints

- **GET** `/api/v1/users/me/orders` - Get current user orders overview from database
- **GET** `/api/v1/users/me/orders/reorder` - Create a new order from existing order from database
- **POST** `/api/v1/users/me/orders` - Create a new current user order and payment
- **GET** `/api/v1/users/me/orders/:id` - Get current user order details from database
- **GET** `/api/v1/vendors/me/orders` - Get current vendor orders overview from database
- **GET** `/api/v1/vendors/me/orders/stats` - Get current vendor orders stats from database
- **GET** `/api/v1/vendors/me/orders/:id` - Get current vendor order details from database

##### Courts endpoints

- **GET** `/api/v1/courts` - Get all available courts from database
- **GET** `/api/v1/vendors/:id/courts/:type` - Get vendor courts using vendor id and court type from database
- **GET** `/api/v1/vendors/:id/courts/:type/bookings` - Get vendor court booking datas using vendor id and court type from database
- **GET** `/api/v1/vendors/me/courts/:type` - Get current vendor courts using court type from database
- **PUT** `/api/v1/vendors/me/courts/:type` - Update current vendor courts using court type from database
- **GET** `/api/v1/vendors/courts/:type/bookings` - Get current vendor court type based on court type from database
- **POST** `/api/v1/vendors/me/courts/:type/new` - Create a new court for a court type
- **POST** `/api/v1/vendors/me/courts/:type` - Create a new court for a court type from the existing court
- **GET** `/api/v1/vendors/me/courts/stats` - Get current vendor courts stats from database

##### Reviews endpoints

- **GET** `/api/v1/vendors/:id/courts/:type/reviews` - Get vendor courts type reviews from database
- **GET** `/api/v1/vendors/me/reviews` - Get court reviews related to the vendor from database
- **POST** `/api/v1/vendors/:id/courts/:type/reviews` - Create a new review for current vendor and court type

##### Fees endpoint

- **GET** `/api/v1/fees` - Get all service fees related to bussiness transactions

##### Advertisements endpoint

- **GET** `/api/v1/advertisements` - Get all available advertisements related to the vendor and vendor's court

##### Payment gateway endpoints

- **GET** `/midtrans/payment-callback` - A payment callback endpoint to mark an order status as success

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

## Installation

To start the service, make sure to follow the steps below.

1. Clone the repository:

```bash
git clone https://github.com/bryanfks-dev/Courtly-Service.git
cd Courtly-Service
```

2. Install depedencies:

```bash
go mod tidy
```

2. Clone the environment file

```bash
cp .env.example .env
```

3. Go to [midtrans website](https://dashboard.midtrans.com/), find your midtrans server key, and copy it to clipboard.

4. Set up the environment variables:

```env
# Database Configuration
DB_HOST=<your-db-host>
DB_PORT=<your-db-port>
DB_USERNAME=<your-dh-username>
DB_PASSWORD=<your-dh-password>
DB_DATABASE=courtly_db

# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=3000

# JWT Configuration
JWT_SECRET=<your-jwt-secret>

# Payment Gateway Configuration
MIDTRANS_API_KEY=<your-midtrans-server-key>
```

4. Run the server:

```bash
go run main.go
```

The server will start on port that you've set based on the environtment variable (e.g. :3000).

5. Start ngrok:

```bash
ngrok http <port-to-listen>
```

6. Copy the generated ngrok link, then paste the link into the **Payment Notification URL** in the midtrans website (in dashboard/integration/configurations) with this format:

```txt
<ngrok-url>/midtrans/payment-callback
```

### License

This project is open-source and available under the [MIT License](https://github.com/bryanfks-dev/Courtly-Service/blob/main/LICENSE).
