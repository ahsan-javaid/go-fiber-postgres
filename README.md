## Go Postgres Base project

### Requirements
- Postgres
- Gorm

### Built with
- Gorm
- fiber v2

### Setup 
```
Create .env file using env.sample file
Specify database name, username and password
```

How to run
```
 go mod download
 go run main.go
```

### API
Login
```
curl -X POST \
  http://localhost:8080/users/login \
  -H 'content-type: application/json' \
  -d '{
    "email": "aki@gmail.com",
  	"password": "1234567"
}'
```

Signup
```
curl -X POST \
  http://localhost:8080/users/signup \
  -H 'content-type: application/json' \
  -d '{
    "email": "aki@gmail.com",
    "password": "1234567",
    "name":"aki"
}'
```

Get all books
```
curl -X GET \
  http://localhost:8080/api/books \
  -H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG4iLCJleHAiOjE2NTk2NjAzODMsImlkIjoxLCJuYW1lIjoiam9obiBkb2UifQ.Bkus3Uj4ftPJ6taRro5R1P8DA_iil7h0Zd9IZgTGlS4'
```
