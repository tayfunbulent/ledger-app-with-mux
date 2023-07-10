# http://localhost:8080/app/services/users/create
```
curl --location 'http://localhost:8080/app/services/users/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@test.com",
    "password": "123456789test",
    "username": "Test Test"
}'
```

# http://localhost:8080/app/services/wallets/update
```
curl --location 'http://localhost:8080/app/services/wallets/update' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "amount": 300
}'
```

# http://localhost:8080/app/services/transactions/create
```
curl --location 'http://localhost:8080/app/services/transactions/create' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data-raw '{
    "recipient_email": "testv2@test.com",
    "amount": 100
}'
```

# http://localhost:8080/app/services/users/update-user
```
curl --location 'http://localhost:8080/app/services/users/update-user' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "2",
    "user": {
        "username": "testv3 test",
        "email": "testv3@test.com",
        "password": "123456789test"
    }
}'
```

# http://localhost:8080/app/services/users/delete-user
```
curl --location 'http://localhost:8080/app/services/users/delete-user' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "id": "2"
}'
```

# http://localhost:8080/app/services/users/update-user-role
```
curl --location 'http://localhost:8080/app/services/users/update-user-role' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "id": "2",
    "role": "admin"
}'
```

# http://localhost:8080/app/services/wallets/update-wallet-balance-by-user-id/
```
curl --location 'http://localhost:8080/app/services/wallets/update-wallet-balance-by-user-id' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "id": "2",
    "amount": -250
}'
```

# http://localhost:8080/app/services/transactions/get-wallet-balance-at-time
```
curl --location 'http://localhost:8080/app/services/transactions/get-wallet-balance-at-time' \
--header 'email: test@test.com' \
--header 'password: 123456789test' \
--header 'Content-Type: application/json' \
--data '{
    "userID": "2",
    "time": "2023-06-25T03:00:00Z"
}'
```

# http://localhost:8080/app/services/wallets/get
```
curl --location 'http://localhost:8080/app/services/wallets/get' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

# http://localhost:8080/app/services/transactions/get
```
curl --location 'http://localhost:8080/app/services/transactions/get' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

# http://localhost:8080/app/services/users/get-all-users
```
curl --location 'http://localhost:8080/app/services/users/get-all-users' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

# http://localhost:8080/app/services/users/get-all-users-with-wallet
```
curl --location 'http://localhost:8080/app/services/users/get-all-users-with-wallet' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

# http://localhost:8080/app/services/wallets/get-all-wallets
```
curl --location 'http://localhost:8080/app/services/wallets/get-all-wallets' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

# http://localhost:8080/app/services/users/get-user-by-id/{id}
```
curl --location 'http://localhost:8080/app/services/users/get-user-by-id/5' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```

# http://localhost:8080/app/services/wallets/get-wallet-by-id/{id}
```
curl --location 'http://localhost:8080/app/services/wallets/get-wallet-by-id/3' \
--header 'email: test@test.com' \
--header 'password: 123456789test'
```
