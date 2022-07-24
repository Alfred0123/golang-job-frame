BINARY_NAME=golang-job-frame
KO_DOCKER_REPO=gcr.io/typed-app-dev/typed-job-v2

# build: clean pre_build
build: wire clean
build:
	GOARCH=amd64 GOOS=linux go build -tags=nomsgpack -o ./build/${BINARY_NAME}-linux ./cmd/main.go
	zip ./build/function.zip ./build/${BINARY_NAME}-linux

run: export GO_ENV=local
run: export PORT=8080
run: export KO_DATA_PATH=cmd/kodata
run:
	$(shell go env GOPATH)/bin/air

build_and_run: build run

#* - 를 앞에 붙이면 오류무시
clean:
	go clean
	-rm -r ./build 
# rm ${BINARY_NAME}-darwin
# rm ${BINARY_NAME}-linux
# rm ${BINARY_NAME}-windows

wire:
	echo "wire build..."
	$(shell go env GOPATH)/bin/wire ./...

ko: build
ko: export KO_DATA_PATH=cmd/kodata
ko: export KO_DOCKER_REPO=${KO_DOCKER_REPO}
ko:
	ko build --local ./cmd
# ko publish ./cmd