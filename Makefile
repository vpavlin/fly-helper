VERSION := "v0.0.4"
TARGET=_build/flyhelper
PKG=$(shell go list)
LDFLAGS=-ldflags "-X '$(PKG)/cmd/flyhelper.version=$(VERSION)'"

clean:
	rm -f $(TARGET)

build:
	go build -o $(TARGET) $(LDFLAGS) main.go

install:
	cp $(TARGET) ${GOPATH}/bin/