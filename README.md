# Palindrom EE

Coding-exercise for interview with Amedia.
This crud-app allows you to create/read/update/delete users, and check if their first/last name are palindromes.


## Build
``` 
go mod tidy && go run ./cmd/webserver
```  

## Run with docker
```
docker build -t amedia-crud-app .
```
```
docker run -p 8080:8080 amedia-crud-app
```

## Run with docker-compose
```
docker-compose up
```

## Example usage:
``` 
# Create new user
curl -X POST 'localhost:8080/api/user' --data 'firstname=abba&lastname=not-palindrome'

# Read users
curl -X GET 'localhost:8080/api/user/all'
curl -X GET 'localhost:8080/api/user/<userid>'

# Update user
curl -X PUT 'localhost:8080/api/user/<userid>' --data 'firstname=abba&lastname=not-a-palindrome'

# Check if the names are palindromes
curl -X GET 'localhost:8080/api/user/<userid>/ispalindrome'

# Delete user
curl -X DELETE 'localhost:8080/api/user/<userid>'
```
