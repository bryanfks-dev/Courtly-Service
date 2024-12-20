# BOOKING RESPONSE

This doc will explain booking endpoints in details.

### **POST** `/api/v1/users/me/bookings`

Endpoint uses to create current user court bookings in database.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Request body needed

```json
{
  "vendor_id": ...,
  "date": "...",
  "bookings": [
    {
      "court_id": ...,
      "book_times": ["...", "...", ...]
    },
    {...},
    {...},
    ...
  ]
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
- `400 BAD REQUEST`: when there is an invalid value in the request body field
- `500 INTERNAL SERVER ERROR`: when either fails to create order or fails to create bookings or fails to commit database transaction
