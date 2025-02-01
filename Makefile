.PHONY: build run dev docker-build docker-run test clean docs cli

build:
	go build -o nxddns ./cmd/nxddns

run:
	go run ./cmd/nxddns/main.go

dev:
	docker compose -f .devcontainer/docker-compose.yml up -ds

docker-build:
	docker build -t nxddns .

docker-run:
	docker run -it -p 8080:8080 nxddns

# Run tests
test:
	go test ./...

# Clean the project (remove build artifacts)
clean:
	rm -f nxddns

docs:
	swag init --output ./cmd/nxddns/docs --dir ./cmd/nxddns,./internal

cli:
	go run ./cmd/cli/main.go