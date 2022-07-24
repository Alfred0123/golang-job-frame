BINARY_NAME=golang-job-frame
DOCKER_REPO=gcr.io/typed-app-dev/typed-job-v2

wire:
	echo "wire build..."
	$(shell go env GOPATH)/bin/wire ./...

build: wire clean
build:
	GOARCH=amd64 GOOS=linux go build -tags=nomsgpack -o ./build/${BINARY_NAME}-linux ./cmd/main.go

clean:
	go clean
	-rm -r ./build 

run: export GO_ENV=local
run: export KO_DATA_PATH=kodata
run:
	go ./build/${BINARY_NAME}-linux

dev: build run

# ko: build
# ko: export KO_DATA_PATH=kodata
# ko: export KO_DOCKER_REPO=${DOCKER_REPO}
# ko:
# 	ko build --local ./
# # ko publish ./cmd
