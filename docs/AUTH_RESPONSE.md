# AUTH RESPONSE

This doc will explain auth endpoints in details.

### **POST** `/api/v1/auth/user/login`

Endpoint uses for sign user with an existing account.

#### Request body needed:

```json
{
  "username": "...",
  "password": "..."
}
```

#### Response body:

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "user": {
      "id": ...,
      "username": "...",
      "phone_number": "...",
      "profile_picture_url": "..."
    },
    "token": "..."
  }
}
```

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when fails validating request body
- `401 UNAUTHORIZE`: when either the password is not valid or username not exists
- `500 INTERNAL SERVER ERROR`: when either fails checking if username is exists or fails to generate token

### **POST** `/api/v1/auth/user/register`

Endpoint uses to register a new user account.

#### Request body needed

```json
{
  "username": "...",
  "phone_number": "...",
  "password": "...",
  "confirm_password": "..."
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "user": {
      "id": ...,
      "username": "...",
      "phone_number": "...",
      "profile_picture_url": "..."
    }
  }
}
```

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when fails validating request body
- `403 FORBIDDEN`: when either username or phone number already exists
- `500 INTERNAL SERVER ERROR`: when either fails checking if username or phone number is exists or fails to hash password or fails to create new user

### **POST** `/api/v1/auth/user/logout`

Endpoints uses to remove user from authenticated status.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": null
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `500 INTERNAL SERVER ERROR`: when fails to blacklist user token

### **POST** `/api/v1/auth/user/logout`

Endpoints uses to remove user from authenticated status.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": null
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `500 INTERNAL SERVER ERROR`: when fails to blacklist user token

### **POST** `/api/v1/auth/user/verify-password`

Endpoint uses to verify current user password.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Request body needed

```json
{
  "password": "..."
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "user": {
      "id": ...,
      "username": "...",
      "phone_number": "...",
      "profile_picture_url": "..."
    }
  }
}
```

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when fails to validate request body
- `401 UNAUTHORIZE`: when the password is invalid
- `500 INTERNAL SERVER ERROR`: when fails to get user

### **POST** `/api/v1/auth/user/logout`

Endpoints uses to remove user from authenticated status.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": null
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `500 INTERNAL SERVER ERROR`: when fails to blacklist user token

### **POST** `/api/v1/auth/vendor/verify-password`

Endpoint uses to verify current vendor password.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Request body needed

```json
{
  "password": "..."
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "vendor": {
      "id": ...,
      "name": "...",
      "email": "...",
      "address": "...",
      "open_time": "...",
      "close_time": "..."
    }
  }
}
```

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when fails to validate request body
- `401 UNAUTHORIZE`: when the password is invalid
- `500 INTERNAL SERVER ERROR`: when fails to get vendor

### **POST** `/api/v1/auth/vendor/login`

Endpoint uses to sign vendor with an existing account.

#### Request body needed:

```json
{
  "email": "...",
  "password": "..."
}
```

#### Response body:

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "vendor": {
      "id": ...,
      "name": "...",
      "email": "...",
      "address": "...",
      "open_time": "...",
      "close_time": "..."
    },
    "token": "..."
  }
}
```

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when fails validating request body
- `401 UNAUTHORIZE`: when either the password is not valid or email not exists
- `500 INTERNAL SERVER ERROR`: when either fails checking if email is exists or fails to generate token

### **POST** `/api/v1/auth/vendor/logout`

Endpoint uses to remove vendor from authenticated status.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": null
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `500 INTERNAL SERVER ERROR`: when fails to blacklist vendor token
