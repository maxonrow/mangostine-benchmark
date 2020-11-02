export GO111MODULE = on

all: deps build install proto

deps:
	go mod vendor

build:
	go build $(LDFLAGS) $(TAGS) -mod vendor -o ./build/mangostine ./main.go

install:
	go install $(LDFLAGS) $(TAGS) -mod vendor ./main.go

proto:
	go get -u google.golang.org/grpc
