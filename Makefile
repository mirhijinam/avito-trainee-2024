# Run compose-up
compose-up: 
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up

# Run test
test: 
	go test -v ./...
.PHONY: test

