#!/bin/bash

# API Examples for my-gin-example
# Make sure the server is running on http://localhost:8089

BASE_URL="http://localhost:8089"

echo "=== Health Check ==="
curl -X GET "$BASE_URL/ok"
echo -e "\n"

echo "=== User Registration ==="
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/v2/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demouser",
    "email": "demo@example.com",
    "password": "demo123456",
    "full_name": "Demo User"
  }')
echo "$REGISTER_RESPONSE" | jq .
echo ""

echo "=== User Login ==="
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/v2/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demouser",
    "password": "demo123456"
  }')
echo "$LOGIN_RESPONSE" | jq .

# Extract token
TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')
echo "Token: $TOKEN"
echo ""

echo "=== Get User Profile (Protected) ==="
curl -s -X GET "$BASE_URL/v2/users/profile" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "=== Update User Profile ==="
curl -s -X PUT "$BASE_URL/v2/users/profile" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "updated@example.com",
    "full_name": "Updated Demo User"
  }' | jq .
echo ""

echo "=== Create Product (Protected) ==="
PRODUCT_RESPONSE=$(curl -s -X POST "$BASE_URL/v2/products" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "MacBook Pro 16",
    "description": "High-performance laptop for professionals",
    "price": 2499.99,
    "stock": 15,
    "category": "Laptops",
    "sku": "MBP-16-2024"
  }')
echo "$PRODUCT_RESPONSE" | jq .

# Extract product ID
PRODUCT_ID=$(echo "$PRODUCT_RESPONSE" | jq -r '.id')
echo ""

echo "=== Create Another Product ==="
curl -s -X POST "$BASE_URL/v2/products" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Magic Mouse",
    "description": "Wireless rechargeable mouse",
    "price": 79.99,
    "stock": 50,
    "category": "Accessories",
    "sku": "MM-2024"
  }' | jq .
echo ""

echo "=== List All Products (Public) ==="
curl -s -X GET "$BASE_URL/v2/products?page=1&page_size=10" | jq .
echo ""

echo "=== Get Single Product ==="
curl -s -X GET "$BASE_URL/v2/products/$PRODUCT_ID" | jq .
echo ""

echo "=== Search Products ==="
curl -s -X GET "$BASE_URL/v2/products?search=Mac&page=1&page_size=10" | jq .
echo ""

echo "=== Update Product ==="
curl -s -X PUT "$BASE_URL/v2/products/$PRODUCT_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 2299.99,
    "stock": 20
  }' | jq .
echo ""

echo "=== Legacy API - Hello World ==="
curl -s -X GET "$BASE_URL/v1/example/hello" | jq .
echo ""

echo "=== Legacy API - Custom Auth ==="
curl -s -X GET "$BASE_URL/v1/example/auth/example" \
  -H "X-Auth: PASS" | jq .
echo ""

echo "=== Test Rate Limiting (send 5 requests) ==="
for i in {1..5}; do
  echo "Request $i:"
  curl -s -X GET "$BASE_URL/ok"
  echo ""
done

echo "All examples completed!"
