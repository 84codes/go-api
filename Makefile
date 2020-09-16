## go-api provider version
version = 1.4.0

clean:
	## remove previous installed go-api

build:
	go build -ldflags "-X 'go/src/github.com/84codes/go-api/api.version=$(version)'" -o go-api_v$(version)

install:
	go install -ldflags "-X 'go/src/github.com/84codes/go-api/api.version=$(version)'"
