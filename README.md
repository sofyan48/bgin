# GO GIN BOILERPLATE
Boilerplate for Golang, Gin

## HOW TO USE

Move project to $GOPATH/src
```
mv boilerplate $GOPATH/src
```

if you change folder name, Customize the import line with the name of your project folder


Install Dep
```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```
Mapping Package
```
dep ensure
```

Set environment, move .env.example to .env

Run
```
go run main.go
```