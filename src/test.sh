curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Alice",
        "email": "bob@example.com",
        "password": "password123"
      }'
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Alice",
        "email": "alice@example.com",
        "password": "password123"
      }'
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
        "email": "alice@example.com",
        "password": "password123"
      }' \
  -c cookies.txt
curl -X POST http://localhost:8080/api/chatReqest \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
        "userS": "alice@example.com",
        "userR": "bob@example.com",
        "status": "pending"
      }'
curl -X POST http://localhost:8080/api/logout \
  -H "Content-Type: application/json" \
  -b cookies.txt

