## Run with docker :

1. `sudo docker build -t server .`
2. `sudo docker run -d --name go_app --link postgres:postgres -p 3000:3000 server go run server.go`
