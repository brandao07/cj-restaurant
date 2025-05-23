.PHONY: all swagger build run clean

# Generate Swagger docs
swagger:
	swag init --generalInfo main.go --output docs --dir ./cmd,./internal

# Build the Go app
build: swagger
	mkdir -p build
	go build -o ./build/app ./cmd

# Run the app (builds first, cleans up after exit)
run: build
	./build/app

# Remove built binary and docs
clean:
	rm -rf build
