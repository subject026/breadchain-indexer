build:
	go build -o ./build/breadchain-api ./cmd/main/main.go

run:
	./build/breadchain-api

dev:
	air --build.cmd "go build -o ./build/breadchain-api ./cmd/main/main.go" --build.bin "./build/breadchain-api"