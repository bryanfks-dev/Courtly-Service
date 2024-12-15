# COURTS RESPONSE

This doc will explain courts endpoints in details.

### **GET** `/api/v1/courts`

Endpoint uses to get all available courts from database.

#### Query parameter (optional)

```js
?type=...
```

> **type** query parameter should contains the court type to filter data

```js
?search=...
```

> **search** query parameter should contains the vendor name to search courts based on vendor name

To use multiple query parameter in a place, use this format:

```js
?type=...&search=...
```

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
        "rating": ...,
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

### **GET** `/api/v1/courts/:id`

Endpoint uses to get court finromation from the database using court id.

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": {
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
      "image_url": "...",
      "rating": ...,
    }
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response success
- `400 BAD REQUEST`: when court id is invalid
- `500 INTERNAL SERVER ERROR`: when fails getting court using id

### GET `/api/v1/vendors/me/courts/types/:type`

Endpoint uses to get current vendor courts using court type from database.

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
    "courts": [{
      "id": ...,
      "name": "...",
      "type": "...",
      "price": ...,
      "image_url": "..."
    },
    {...},
    {...},
    ...
    ]
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response success
- `400 BAD REQUEST`: when court type is invalid
- `500 INTERNAL SERVER ERROR`: when fails getting courts using court type

### **POST** `/api/v1/vendors/me/courts/types/:type/new`

Endpoint uses to create a new court for a court type.

#### Request header needed

```json
{
  "Authorization": "Bearer <token here>"
}
```

#### Request body needed

```json
{
  "price_per_hour": ...,
  "court_image": "..."
}
```

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "court": {
      "id": ...,
      "name": "...",
      "type": "...",
      "price": ...,
      "image_url": "..."
    }
  }
}
```

> **message** field possibly return a string or a map of string (likely a form error)

#### Possible HTTP status codes

- `200 OK`: when response success
- `400 BAD REQUEST`: when either court type is invalid or fails to validate request body
- `403 FORBIDDEN`: when a vendor with current court type already exists
- `500 INTERNAL SERVER ERROR`: when either fails to check if court exists in current court type or fails to decode court image or fails to save court image or fails to create new court

### **POST** `/api/v1/vendors/me/courts/types/:type`

Endpoint uses to create a new court for a court type from the existing court.

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
    "court": {
      "id": ...,
      "name": "...",
      "type": "...",
      "price": ...,
      "image_url": "..."
    }
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response success
- `400 BAD REQUEST`: when court type is invalid
- `403 FORBIDDEN`: when a vendor with current court type is not exists
- `500 INTERNAL SERVER ERROR`: when either fails to check if court exists in current court type or fails to create new court

### **GET** `/api/v1/vendors/me/courts/stats`

Endpoint uses to get current vendor courts stats from database.

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
    "football_court_count": ...,
    "basketball_court_count": ...,
    "tennis_court_count": ...,
    "volleyball_court_count": ...,
    "badminton_court_count": ...
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response success
- `500 INTERNAL SERVER ERROR`: when fails to get vendor courts stats
