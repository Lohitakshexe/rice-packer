.PHONY: build run-test-env test

# Build the tool natively (requires Go on host)
build:
	go build -o bin/rice main.go

# Build the Docker testing image
build-docker:
	docker build -f Dockerfile.test -t ricepacker-test-env .

# Run the interactive Docker testing environment
# We mount the current directory into the container so we can run the code
run-test-env: build-docker
	docker run -it --rm \
		-v $(CURDIR):/home/tester/workspace \
		-w /home/tester/workspace \
		ricepacker-test-env /bin/bash

# Run go tests inside docker since the host might not have go installed
test: build-docker
	docker run --rm \
		-v $(CURDIR):/home/tester/workspace \
		-w /home/tester/workspace \
		ricepacker-test-env go test -v ./...

# Download dependencies inside docker
mod-tidy: build-docker
	docker run --rm \
		-v $(CURDIR):/home/tester/workspace \
		-w /home/tester/workspace \
		ricepacker-test-env go mod tidy
