# management
to manage User

Pre-requisite : 
    Swagger must be install and running-
        Follow the page for swaggo installation.
        https://github.com/swaggo/gin-swagger 
    Docker must be up and running.
    
To run the project, run following commands:
     swag init --parseDependency  --parseInternal --parseDepth 1 -d api -g ../cmd/main.go
     docker-compose up -d
     go run cmd/main.go -log-level info

To check the Swagger Documentation:
    http://localhost:8000/swagger/index.html