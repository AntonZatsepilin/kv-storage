# tarantool-key-value-storage

## Description

This project is a REST API for key value json storage. It is written in Go and uses Tarantool and Docker. The project includes creation, reading, updating, and deletion of json.


## How to Run


### Step 1 
Сreate a .env file in the root of the project with the following content:
``` .env
TARANTOOL_USER_PASSWORD= ...
TARANTOOL_USER_NAME= ...
```

### Step 2
Make sure you have Docker and Docker Compose installed. Then run the following command:
```
docker-compose up --build
```

### Step 3
tarantool-key-value-storage is on http://localhost:8080  
Swagger UI is on http://localhost:8080/swagger/index.html  