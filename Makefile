.PHONY: build run dev docker-build docker-run test clean

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

.PHONY: api, docs, nxddns

api:
	go run ./cmd/api/main.go

docs:
	swag init --output ./cmd/api/docs --dir ./cmd/api,./internal

nxddns:
	go run ./cmd/nxddns/main.go