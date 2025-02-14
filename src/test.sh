#!/bin/bash
# Register Bob (ignore error if email exists)
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Bob",
        "email": "bob@example.com",
        "password": "password123"
      }' || true

# Register Alice (ignore error if email exists)
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Alice",
        "email": "alice@example.com",
        "password": "password123"
      }' || true

# Login as Alice
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
        "email": "alice@example.com",
        "password": "password123"
      }' \
  -c cookiesalice.txt

# Login as Bob
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
        "email": "bob@example.com",
        "password": "password123"
      }' \
  -c cookiesbob.txt

# Alice sends a chat request to Bob (ignore if already exists)
curl -X POST http://localhost:8080/api/chatReqest \
  -H "Content-Type: application/json" \
  -b cookiesalice.txt \
  -d '{
        "userS": "alice@example.com",
        "userR": "bob@example.com",
        "status": "pending"
      }' || true

# Bob checks chat requests
curl -X POST http://localhost:8080/api/seeChatReqests \
  -H "Content-Type: application/json" \
  -b cookiesbob.txt \
  -d '{
        "userEmail": "bob@example.com"
      }'

# Bob accepts Alice's chat request (fixing JSON structure)
curl -X POST http://localhost:8080/api/acceptChatReqest \
  -H "Content-Type: application/json" \
  -b cookiesbob.txt \
  -d '{
        "userEmail": "bob@example.com",
        "userS": "alice@example.com"
      }'

