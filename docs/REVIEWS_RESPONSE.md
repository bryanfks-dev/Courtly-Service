# REVIEWS RESPONSE

This doc will explain reviews endpoints in details.

### **GET** `/api/v1/vendors/:id/courts/:type/reviews`

Endpoint uses to get vendor courts type reviews from database.

#### Query parameter (optional)

```js
?rating=...
```

> **rating** query parameter should contains the rating that contains a number 1 to 5 to filter data

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
    "total_rating": ...,
    "total_reviews": ...,
    "stars": {
      "1": ...,
      "2": ...,
      "3": ...,
      "4": ...,
      "5": ...
    },
    "reviews": [
      {
        "id": ...,
        "user": {
          "id": ...,
          "username": "...",
          "profile_picture_url": "..."
        },
        "court_type": "...",
        "rating": ...,
        "review": "...",
        "date": "..."
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
- `400 BAD REQUEST`: when either invalid vendor id or invalid court type or invalid rating query parameter
- `500 INTERNAL SERVER ERROR`: when either fails to get total rating or fails to get total reviews or fails to get reviews

### **GET** `/api/v1/vendors/me/reviews`

Endpoint uses to get court reviews related to the vendor from database.

#### Query parameter (optional)

```js
?rating=...
```

> **rating** query parameter should contains the rating that contains a number 1 to 5 to filter data

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
    "total_rating": ...,
    "total_reviews": ...,
    "stars": {
      "1": ...,
      "2": ...,
      "3": ...,
      "4": ...,
      "5": ...
    },
    "reviews": [
      {
        "id": ...,
        "user": {
          "id": ...,
          "username": "...",
          "profile_picture_url": "..."
        },
        "court_type": "...",
        "rating": ...,
        "review": "...",
        "date": "..."
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
- `400 BAD REQUEST`: when either invalid court type or invalid rating query parameter
- `500 INTERNAL SERVER ERROR`: when either fails to get total rating or fails to get total reviews or fails to get reviews

### **POST** `/api/v1/vendors/:id/courts/:type/reviews`

Endpoint uses to create a new review for current vendor and court type.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Request body needed

```json
{
  "rating": ...,
  "review": "..."
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "review": {
      "id": ...,
      "user": {
        "id": ...,
        "username": "...",
        "profile_picture_url": "..."
      },
      "court_type": "...",
      "rating": ...,
      "review": "...",
      "date": "..."
    }
  }
}
```

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response is success
- `400 BAD REQUEST`: when either invalid vendor id or invalid court type or fails to validate request body
- `403 FORBIDDEN`: when either user haven't book the court or user has reviewed the court
- `500 INTERNAL SERVER ERROR`: when either fails to check user has book a court or fails getting court data or fails checking court has been reviewed by user or fails to get court type data or fails to create review
