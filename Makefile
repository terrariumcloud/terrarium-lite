all:
	make build
	make run

build:
	go build

run:
	docker-compose up -d
	./terrarium serve module --postgres-host localhost:5432 --postgres-password terrarium --postgres-sslmode disable