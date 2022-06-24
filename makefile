pre-commit: mod test

mod:
	go mod tidy
	go mod verify

test:
	go vet
	go test ./... -cover -timeout 60s -race -cpu 4

demo:
	go run ./examples/demo