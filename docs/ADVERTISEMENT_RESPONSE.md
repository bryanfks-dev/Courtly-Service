# FEES RESPONSE

This doc will explain advertisement endpoints in details.

### **GET** `/api/v1/advertisements`

Endpoint uses to get all available advertisements related to the vendor and vendor's court.

#### Response body

```json
{
  "success": ...,
  "message": "...",
  "data": {
    "ads": [
      {
        "id": ...,
        "image_url": "...",
        "vendor": {
          "id": ...,
          "name": "...",
          "address": "...",
          "open_time": "...",
          "close_time": "..."
        },
        "court_type": "..."
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
- `500 INTERNAL SERVER ERROR`: when fails to get advertisements
