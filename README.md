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

## init go project

go mod init goparkin_service

## package installation

go get github.com/go-sql-driver/mysql
go get go.mongodb.org/mongo-driver/mongo
go get github.com/eclipse/paho.mqtt.golang
go get github.com/gorilla/websocket
go get github.com/go-redis/redis/v8
go get gopkg.in/gomail.v2
go get github.com/gorilla/mux
go get github.com/joho/godotenv

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/hrbliukaiST/goparkin_service.git
   ```
2. Navigate to the project directory:
   ```sh
   cd goparkin_service
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

1. Run the project:
   ```sh
   go run main.go
   ```

## File Descriptions

- **/config**
  - `config.go`: Configuration settings for the project.
- **/db**
  - `mysql.go`: MySQL database connection and operations.
  - `mongodb.go`: MongoDB database connection and operations.
- **/mqtt**
  - `mqtt.go`: MQTT protocol handling.
- **/websocket**
  - `websocket.go`: WebSocket protocol handling.
- **/redis**
  - `redis.go`: Redis database connection and operations.
- **/email**
  - `email.go`: Email sending functionality.
- **/handlers**
  - `user.go`: HTTP handlers for user-related operations.
- **/models**
  - `user.go`: Data models for user-related data.
- **/routes**
  - `routes.go`: Routing configuration for the application.
- **main.go**: Entry point of the application.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
