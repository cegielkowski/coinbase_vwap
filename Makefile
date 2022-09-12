run-server:
	go run ./cmd/main.go

build-api:
	go build ./cmd/main.go

create-docker-image:
	docker build -t cegielkowski/coinbase_vwap .

run-docker:
	docker run cegielkowski/coinbase_vwap

test:
	go test -v ./...

mock:
	mockery --all