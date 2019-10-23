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

## PACKAGE
Activate and deactivate available packages see on app/moduls/package, see the example on main.go
```
main.go

look at this line:

db.Conn = db.Init()
db.MigrateScheme(db.Conn)
redis.Store = redis.Init()
```


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




