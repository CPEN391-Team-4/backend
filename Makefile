all: install

install:
	go install -v ./src/server/...

.PHONY: install