# USER RESPONSE

This doc will explain user endpoints in details.

### **GET** `/api/v1/users/me`

Endpoint uses to get current user information from the database.

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

#### Possible HTTP status codes

- `200 OK`: when response is success
- `500 INTERNAL SERVER ERROR`: when fails to get user

### **PATCH** `/api/v1/users/me/password`

Endpoint uses to update user password with a new password.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Request body needed

```json
{
  "old_password": "...",
  "new_password": "...",
  "confirm_password": "..."
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

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response success
- `400 BAD REQUEST`: when either fails to validate request body or old password is invalid or new password and cofirm password not match
- `500 INTERNAL SERVER ERROR`: when either fails getting user or fails hashing password or fails updating user password

### **PATCH** `/api/v1/users/me/username`

Endpoint uses to update user username with a new available username.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Request body needed

```json
{
  "new_username": "..."
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

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response success
- `400 BAD REQUEST`: when either fails to validate request body or username is taken
- `500 INTERNAL SERVER ERROR`: when either fails getting user or fails updating user username
