BINARY_NAME=golang-job-frame
DOCKER_REPO=gcr.io/typed-app-dev/typed-job-v2

wire:
	echo "wire build..."
	$(shell go env GOPATH)/bin/wire ./...

build: wire clean
build:
	GOARCH=amd64 GOOS=linux go build -tags=nomsgpack -o ./build/${BINARY_NAME}-linux ./main.go

build_dev: wire clean
build_dev:
	go build -tags=nomsgpack -o ./build/${BINARY_NAME}-darwin ./main.go

clean:
	go clean
	-rm -r ./build 

run: export GO_ENV=local
run: export KO_DATA_PATH=kodata
run:
	go run ./build/${BINARY_NAME}-linux

run_dev: export GO_ENV=local
run_dev: export KO_DATA_PATH=kodata
run_dev:
	./build/${BINARY_NAME}-darwin

dev: build_dev run_dev

# ko: build
# ko: export KO_DATA_PATH=kodata
# ko: export KO_DOCKER_REPO=${DOCKER_REPO}
# ko:
# 	ko build --local ./
# # ko publish ./cmd
