FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

# Install make
RUN apt-get update && apt-get install -y make

# Ensure Go bin is in PATH for swag
ENV PATH="$PATH:/go/bin"

# Use Makefile to generate Swagger docs, build, and run the app
CMD ["make", "run"]