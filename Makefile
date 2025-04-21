# builds and runs the docker container
build-run:
	docker-compose up --build

# runs tests
run-test:
	go test ./... -race

# tidy
tidy:
	go mod tidy

# create migrations
migration:
	goose create -dir ./migrations $(f) sql