#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN mkdir /go/bin/app
RUN chmod +x /go/bin/app
RUN apk update && apk add openssh

# Set necessary environmet variables needed for the image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

# Copy the code into the container
COPY src/. .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./main.go -o ./docs

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 8080

# For local run: In windows from %APPDATA%/gcloud/ take and copy to the src folder the credentials file. Before, use 'gcloud auth login' from the command line to refresh the tokens. Enable the next 2 lines
#RUN mkdir -p ../gcp/config
#COPY ["application_default_credentials.json", "../gcp/config/"]

# Command to run when starting the container
CMD ["/dist/main"]
