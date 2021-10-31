all:
	make build
	make run

build:
	docker-compose build

run:
	docker-compose up

mac:
	lima nerdctl compose up --build