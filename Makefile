build:
	podman build -t backend-tech-test:0.1.0 -f ./container/server/containerfile .

deploy:
	podman-compose -f compose.yaml up

undeploy:
	podman-compose -f compose.yaml down

run:
	go run cmd/server/main.go

test:
	go test -v -race ./...