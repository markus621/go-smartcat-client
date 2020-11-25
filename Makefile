lint:
	golangci-lint -v run ./...

generate:
	go generate -v ./...

demo_run:
	cd ./example && go run main.go