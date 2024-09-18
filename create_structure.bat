@echo off

md config
cd config
echo > config.go
cd ..

md db
cd db
echo > mysql.go
echo > mongodb.go
cd ..

md mqtt
cd mqtt
echo > mqtt.go
cd ..

md websocket
cd websocket
echo > websocket.go
cd ..

md redis
cd redis
echo > redis.go
cd ..

md email
cd email
echo > email.go
cd ..

md handlers
cd handlers
echo > user.go
cd ..

md models
cd models
echo > user.go
cd ..

md routes
cd routes
echo > routes.go
cd ..

echo main.go > main.go