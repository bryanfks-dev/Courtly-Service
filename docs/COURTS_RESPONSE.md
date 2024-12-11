# COURTS RESPONSE

This doc will explain courts endpoints in details.

### **GET** `/api/v1/courts`

Endpoint uses to get all available courts from database.

#### Query parameter (optional)

```js
?type=...
```

> **type** query parameter should contains the court type to filter data

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
    "courts": [
      {
        "id": ...,
        "name": "...",
        "vendor": {
          "id": ...,
          "name": "...",
          "address": "...",
          "open_time": "...",
          "close_time": "..."
        },
        "type": "...",
        "price": ...,
        "image_url": "...",
      },
      {...},
      {...},
      ...
    ]
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when court type is invalid
- `500 INTERNAL SERVER ERROR`: when fails to get courts

### **PATCH** `/api/v1/vendors/me/password`

Endpoint uses to update vendor password with a new password.

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
- `500 INTERNAL SERVER ERROR`: when either fails getting vendor or fails hashing password or fails updating vendor password
