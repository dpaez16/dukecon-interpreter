CC=go build
EXE_NAME=dukecon

.PHONY: build tests

build:
	$(CC) -o $(EXE_NAME)

tests:
	go test ./tests

clean:
	rm -f $(EXE_NAME)
	go clean -testcache