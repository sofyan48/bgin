# BGIN
Create Your Micro Service with Gin Golang

## TREE
```
├── Dockerfile
├── Gopkg.lock
├── Gopkg.toml
├── README.md
├── app
│   ├── config
│   │   └── config.go
│   ├── controller
│   │   ├── api
│   │   │   ├── health.go
│   │   │   ├── kafka.go
│   │   │   ├── login.go
│   │   │   └── ping.go
│   │   └── routes.go
│   ├── helper
│   │   └── rest.go
│   ├── libs
│   │   └── utils.go
│   ├── middlewares
│   │   └── auth.go
│   ├── models
│   │   └── loginModels.go
│   ├── moduls
│   │   ├── migration
│   │   │   └── scheme.go
│   │   └── package
│   │       ├── database.go
│   │       ├── etcd.go
│   │       ├── kafka.go
│   │       └── redis.go
│   └── server
│       ├── router.go
│       └── server.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── test
    └── ping_test.go
```

## HOW TO INSTALL 
```
go get github.com/meongbego/bgin
```

Test in Running

```
go run main.go
```

Test Build 

```
go build main.go
```

To run production
```
go run main.go -e production or ./build -e production
```

## ENVIRONMENT SETUP

### LOCAL DEVELOPMENT ENV
if using local development use .env file
```
cp .env.example .env
```
now setup your env

### PRODUCTION
if production mode environtment setup in os env, example in your terminal execute this script
```
export GIN_MODE=release
export APP_PORT=5000
export APP_HOST=0.0.0.0
export DB_HOST=roach1
export DB_PORT=26257
export DB_NAME=bgin
export DB_USER=root
export DB_PASSWORD=
export ACL_ADDR=172.19.0.0/24
export REDIS_HOST=rdcaches
export REDIS_PORT=6379
export KAFKA_HOST_PORT=localhost:9092
```
if your production dockerizing see docker-compose.yml to export environtment


## PACKAGE
Activate and deactivate available packages see on app/moduls/package, see the example on main.go
```
main.go

look at this line:

db.Conn = db.Init()
db.MigrateScheme(db.Conn)
redis.Store = redis.Init()
```

Package:
1. kafka
2. elasticsearch
3. etcd
4. database (cockroachdb)
5. redis

## DATABASE SCHEME AN IMPORTED TABLE
to create a database schema look at the app/moduls/migration/ folder then edit the schem.go file

```
app/moduls/migration/scheme.go

look at this line:

db.AutoMigrate (& LoginScheme {}, & Userdata {})
```
## CONTROLLER
to create a controller schema look at the app/controller/api folder

```
app/controller/api
```

## ROUTES
to create route in look at app/controller folder then edit routes.go
```
app/controller/routes.go

look at this line:

api := router.Group("api")
{
    ping := new(controller.PingController)
    health := new(controller.HealthController)
    login := new(controller.LoginController)
    // create rest api models
    api.GET("/ping", ping.Status)
    api.GET("/health", health.Status)
    api.POST("/login", login.LoginUsers)
    api.GET("/login/list", login.ListLogin)
}
```

## MODELS
to create a models look at the app/models folder

```
app/models/
```
see loginModels.go for examples


## DEVELOPMENT MODE

To activate Live Reload install air 
### on macOS

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air
chmod +x /usr/local/bin/air
```

### on Linux

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
chmod +x /usr/local/bin/air
```

### on Windows

```
curl -fLo ~/air.exe \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/windows/air.exe
```

see watcher.conf setting watcher file for air config now

### Starting
go to your project path
```
air -c watcher.conf
```

## DOCKERIZING
if you dockerizing this project follow this step

```
docker build -t your_tagging .
```

then see docker-compose.yml edit file for your configuration and install docker-compose for execute this script

```
version: '3'
services:
  rdcaches:
    image: redis
    command: ["redis-server"]
    ports:
      - "6379:6379"

  roach1:
    image: cockroachdb/cockroach
    command: start --insecure
    ports:
      - "26257:26257"
      - "8080:8080"
    volumes:
      - ./cockroach-data/roach1:/cockroach/cockroach-data

  roach2:
    image: cockroachdb/cockroach
    command: start --insecure --join=roach1
    volumes:
      - ./cockroach-data/roach2:/cockroach/cockroach-data
    links:
      - roach1
  
  bgin:
    image: your_tagging
    ports:
      - "6968:5000"
    environment:
      - GIN_MODE=release
      - APP_PORT=5000
      - APP_HOST=0.0.0.0
      - DB_HOST=roach1
      - DB_PORT=26257
      - DB_NAME=bgin
      - DB_USER=root
      - DB_PASSWORD=
      - ACL_ADDR=172.19.0.0/24
      - REDIS_HOST=rdcaches
      - REDIS_PORT=6379
    command: ./main -e production
    links:
      - roach1
      - rdcaches
```

then start compose

```
docker-compose up
```

stop container
```
docker-compose stop; docker-compose rm -f
```



