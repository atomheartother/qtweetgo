GO=go build

OUT=qtweet

ENTRY= cmd/qtweet.go

DEPS = $(wildcard *.go) ${wildcard **/*.go}

all: $(OUT)

test:
	go test ./...

run: all
	./$(OUT)

$(OUT): $(DEPS)
	$(GO) -o $(OUT) $(ENTRY)

clean:
	rm -rf $(OUT)

re: clean all