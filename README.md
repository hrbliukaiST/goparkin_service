# goparkin_service

Backend service for goparkin

go restful APIs that will support mysql DB, mongo DB, MQTT, WebSocket, redise, Email Function.

/goparkin_service
|-- /config
| |-- config.go
|-- /db
| |-- mysql.go
| |-- mongodb.go
|-- /mqtt
| |-- mqtt.go
|-- /websocket
| |-- websocket.go
|-- /redis
| |-- redis.go
|-- /email
| |-- email.go
|-- /handlers
| |-- user.go
|-- /models
| |-- user.go
|-- /routes
| |-- routes.go
|-- main.go

https://github.com/hrbliukaiST/goparkin_service.git

go mod init goparkin_service

go get github.com/go-sql-driver/mysql
go get go.mongodb.org/mongo-driver/mongo
go get github.com/eclipse/paho.mqtt.golang
go get github.com/gorilla/websocket
go get github.com/go-redis/redis/v8
go get gopkg.in/gomail.v2
go get github.com/gorilla/mux
go get github.com/joho/godotenv

go run main.go
