FROM golang:1.19-alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create app directory
WORKDIR /app

# Copy all other source code to work directory
COPY . .

# Download all the dependencies that are required
RUN go mod tidy

# Build the application
RUN go build -o binary cmd/main.go

ENTRYPOINT ["/app/binary"]