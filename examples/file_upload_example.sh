#!/bin/bash

# File Upload API Examples for my-gin-example
# Make sure the server is running on http://localhost:8089

BASE_URL="http://localhost:8089"

echo "=== Step 1: Register and Login ==="
curl -s -X POST "$BASE_URL/v2/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "fileuser",
    "email": "file@example.com",
    "password": "file123456",
    "full_name": "File User"
  }' | jq .
echo ""

LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/v2/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "fileuser",
    "password": "file123456"
  }')
echo "$LOGIN_RESPONSE" | jq .

TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')
echo "Token: $TOKEN"
echo ""

# Create test files
echo "Creating test files..."
echo "This is a test document" > /tmp/test_document.txt
echo "Sample data for testing" > /tmp/sample_data.csv
echo ""

echo "=== Step 2: Upload Single File ==="
UPLOAD_RESPONSE=$(curl -s -X POST "$BASE_URL/v2/files/upload" \
  -H "Authorization: Bearer $TOKEN" \
  -F "file=@/tmp/test_document.txt" \
  -F "category=document" \
  -F "description=Test document for demonstration" \
  -F "is_public=false")
echo "$UPLOAD_RESPONSE" | jq .

FILE_ID=$(echo "$UPLOAD_RESPONSE" | jq -r '.id')
echo "Uploaded File ID: $FILE_ID"
echo ""

echo "=== Step 3: Upload Multiple Files ==="
curl -s -X POST "$BASE_URL/v2/files/upload/multiple" \
  -H "Authorization: Bearer $TOKEN" \
  -F "files=@/tmp/test_document.txt" \
  -F "files=@/tmp/sample_data.csv" | jq .
echo ""

echo "=== Step 4: List All Files ==="
curl -s -X GET "$BASE_URL/v2/files?page=1&page_size=10" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "=== Step 5: Get Single File Info ==="
curl -s -X GET "$BASE_URL/v2/files/$FILE_ID" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "=== Step 6: Download File ==="
curl -s -X GET "$BASE_URL/v2/files/$FILE_ID/download" \
  -H "Authorization: Bearer $TOKEN" \
  -o "/tmp/downloaded_file.txt"
echo "File downloaded to /tmp/downloaded_file.txt"
echo "Content:"
cat /tmp/downloaded_file.txt
echo ""
echo ""

echo "=== Step 7: Delete File ==="
curl -s -X DELETE "$BASE_URL/v2/files/$FILE_ID" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

# Cleanup
echo "Cleaning up test files..."
rm -f /tmp/test_document.txt /tmp/sample_data.csv /tmp/downloaded_file.txt

echo "File upload examples completed!"
