tidy:
	go mod tidy

run:
	go run -race main.go

dcbuild:
	docker-compose build --no-cache

dcup:
	docker-compose up --force-recreate --build

dcclean:
	docker-compose down --volumes --remove-orphans
