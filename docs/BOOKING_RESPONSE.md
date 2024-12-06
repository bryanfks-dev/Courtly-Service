# BOOKING RESPONSE

This doc will explain booking endpoints in details.

### **GET** `/api/v1/users/me/bookings`

Endpoint uses to get current user bookings from database.

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
    "bookings": [
      {
        "id": ...,
        "order": {
          "id": ...,
          "payment_method": "...",
          "price": ...,
          "app_fee": ...,
          "status": ...
        },
        "court": {
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
          "image_url": "..."
        }
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
- `500 INTERNAL SERVER ERROR`: when fails getting user bookings
