PROGRAM = quine
SOURCE = main.go

build:
	go build -o $(PROGRAM) $(SOURCE)

clean:
	rm -f $(PROGRAM)

fmt:
	gofmt -w $(SOURCE)

vet:
	go vet $(SOURCE)

run:
	go run $(SOURCE)
