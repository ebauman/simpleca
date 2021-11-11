BINARY=simpleca
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

.DEFAULT_GOAL := build

build:
	go build -o ${BINARY}

release: dist $(PLATFORMS)

dist:
	mkdir dist

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o dist/'${BINARY}-$(os)-$(arch)'

test: fmt vet
	go test ./... -coverprofile=cover.out

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	go clean
	rm cover.out || true
	rm -r dist/ || true