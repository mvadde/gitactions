# Define the Go compiler
GO = go

# Define the main target
all: build

# Build the project
build:
	$(GO) build -o myapp main.go

# Run the project
run: build
	./myapp

# Clean up build files
clean:
	rm -f myapp
