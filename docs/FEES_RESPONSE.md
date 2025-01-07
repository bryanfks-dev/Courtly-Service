# FEES RESPONSE

This doc will explain fees endpoints in details.

### **GET** `/api/v1/fees`

Endpoint uses to get all service fees related to bussiness transactions.

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
    "app_fee": ...,
  }
}
```

#### Possible HTTP status codes

- `200 OK`: when response is success
