# API Examples

This directory contains example scripts and files demonstrating how to use the API.

## Running the Examples

### Shell Script (api_examples.sh)

This script demonstrates all the API endpoints in sequence.

### File Upload Example (file_upload_example.sh)

This script demonstrates the file upload and management functionality.

**Requirements:**
- Server running on `http://localhost:8089`
- `curl` installed
- `jq` installed (for JSON formatting)

**Usage:**
```bash
# Start the server first
go run main.go

# In another terminal, run the examples
./examples/api_examples.sh

# Or run the file upload examples
./examples/file_upload_example.sh
```

## What the Scripts Demonstrate

### api_examples.sh
1. **Health Check** - Basic connectivity test
2. **User Registration** - Create a new user account
3. **User Login** - Authenticate and get JWT token
4. **Get Profile** - Retrieve user profile (authenticated)
5. **Update Profile** - Modify user information
6. **Create Products** - Add new products (authenticated)
7. **List Products** - Browse products with pagination
8. **Get Product** - Retrieve single product details
9. **Search Products** - Filter products by search term
10. **Update Product** - Modify product information
11. **Legacy API** - Examples using v1 endpoints
12. **Rate Limiting** - Test rate limiter behavior

### file_upload_example.sh
1. **Single File Upload** - Upload a file with metadata
2. **Multiple Files Upload** - Upload multiple files at once
3. **List Files** - Browse uploaded files with pagination
4. **Get File Info** - Retrieve file metadata
5. **Download File** - Download uploaded file
6. **Delete File** - Remove uploaded file

## Manual Testing Examples

### Using Postman

Import the following as a collection or test individually:

**Collection Variables:**
- `base_url`: `http://localhost:8089`
- `token`: (will be set after login)

### Using HTTPie

```bash
# Health check
http GET :8089/ok

# Register
http POST :8089/v2/auth/register username=testuser email=test@example.com password=pass123 full_name="Test User"

# Login
http POST :8089/v2/auth/login username=testuser password=pass123

# Use token (replace YOUR_TOKEN)
http GET :8089/v2/users/profile "Authorization: Bearer YOUR_TOKEN"

# Create product
http POST :8089/v2/products \
  "Authorization: Bearer YOUR_TOKEN" \
  name="Test Product" \
  price:=99.99 \
  stock:=10 \
  sku="TEST-001"
```

## Response Examples

### Successful Login
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "demouser",
    "email": "demo@example.com",
    "full_name": "Demo User",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-03T10:00:00Z"
  }
}
```

### Product List (Paginated)
```json
{
  "data": [
    {
      "id": 1,
      "name": "MacBook Pro 16",
      "description": "High-performance laptop",
      "price": 2499.99,
      "stock": 15,
      "category": "Laptops",
      "sku": "MBP-16-2024",
      "is_active": true,
      "created_at": "2025-12-03T10:00:00Z"
    }
  ],
  "page": 1,
  "page_size": 10,
  "total_items": 1,
  "total_pages": 1
}
```

## Error Handling Examples

### Unauthorized (401)
```json
{
  "error": "Authorization header required"
}
```

### Validation Error (400)
```json
{
  "error": "Key: 'UserRegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

### Not Found (404)
```json
{
  "error": "Product not found"
}
```

### Rate Limited (429)
```json
{
  "error": "Rate limit exceeded"
}
```
