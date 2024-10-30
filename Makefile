GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
BINARY_NAME=app_name

all: watch

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	$(GORUN) main.go

watch:
	reflex -r '\.go$$' -- sh -c '$(GORUN) main.go'

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build run watch clean