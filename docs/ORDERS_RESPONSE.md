# ORDERS RESPONSE

This doc will explain orders endpoints in details.

### **GET** `/api/v1/users/me/orders`

Endpoint uses to get current user orders overview from database.

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
    "orders": [
      {
        "id": ...,
        "date": "...",
        "court_type": "...",
        "vendor": {
          "id": ...,
          "name": "...",
          "address": "...",
          "open_time": "...",
          "close_time": "..."
        },
        "price": ...,
        "app_fee": ...,
        "status": "...",
        "reviewed": ...
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
- `500 INTERNAL SERVER ERROR`: when fails getting user orders or fails to check if user has reviewed vendor's court type

### **POST** `/api/v1/users/me/orders`

Endpoint uses to create a new current user order and payment.

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
    "payment_token": "..."
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when either fails to parse date or fails to parse time
- `500 INTERNAL SERVER ERROR`: when either fails to create order, fails to begin transaction, fails to get court, fails to create transaction, fails to update payment token

### **GET** `/api/v1/users/me/orders/:id`

Endpoint uses to get current user order details from database.

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
    "order_detail": {
      "id": ...,
      "date": "...",
      "price": ...,
      "app_fee": ...,
      "payment_token": "...",
      "status": "...",
      "bookings": [
        {
          "id": ...,
          "court": {
            "id": ...,
            "name": "...",
            "vendor" {
              "id": ...,
              "name": "...",
              "address": "...",
              "open_time": "...",
              "close_time": "..."
            },
            "type": "...",
            "price": ...,
            "image_url": "..."
          },
          "book_start_time": "...",
          "book_end_time": "..."
        },
        {...},
        {...},
        ...
      ]
    }
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when order is invalid
- `500 INTERNAL SERVER ERROR`: when fails to get order detail

### **GET** `/api/v1/vendors/me/orders`

Endpoint uses to get current vendor orders from database.

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
  "data": [
    {
      "id": ...,
      "date": "...",
      "user": {
        "id": ...,
        "username" : "...",
        "profile_picture_url": "..."
      },
      "court": {
        "id": ...,
        "name": "...",
        "type": "...",
        "price": ...,
        "image_url": "..."
      },
      "book_start_time": "...",
      "book_end_time": "...",
      "price": ...
    },
    {...},
    {...},
    ...
  ]
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when court type is invalid
- `500 INTERNAL SERVER ERROR`: when fails to get vendor orders

### **GET** `/api/v1/vendors/me/orders/stats`

Endpoint uses to get current vendor orders stats from database.

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
    "total_orders": ...,
    "total_orders_today": ...,
    "recent_orders": [
      {
      "id": ...,
      "date": "...",
      "user": {
        "id": ...,
        "username" : "...",
        "profile_picture_url": "..."
      },
      "court": {
        "id": ...,
        "name": "...",
        "type": "...",
        "price": ...,
        "image_url": "..."
      },
      "book_start_time": "...",
      "book_end_time": "...",
      "price": ...
    },
      {...},
      {...}
    ]
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
- `500 INTERNAL SERVER ERROR`: when either fails to get total orders or fails to get total orders today or tails to get recent orders
