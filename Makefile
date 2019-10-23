# Build And Development
test:
	@go test -v -cover -covermode=atomic ./...

depend:
	@go get github.com/meongbego/bgin

build:
	@go build -o bgin main.go

run:
	@go run main.go

production:
	@go run main.go -e production

docker-stop:
	@docker-compose down

docker-image: 
	@docker build . -t bgin

docker-run:
	@docker-compose up

.PHONY: test depend build  run stop docker docker-stop docker-image docker-run
