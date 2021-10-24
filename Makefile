all:
	make build
	make run

build:
	go build

test:
	go test -cover ./...

run:
	docker-compose up -d
	./terrarium serve module --database-host localhost:27017