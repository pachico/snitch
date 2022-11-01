# Define argument for builder image with default value
ARG BUILDER_IMAGE=golang:1.21-alpine

# Use scratch as the final image to minimize size
ARG TARGET_IMAGE=scratch

# Builder stage
FROM ${BUILDER_IMAGE} as builder

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Use go vet to examine source code and report suspicious constructs
RUN CGO_ENABLED=0 go vet ./...

# Run tests, remove wildcard to ensure 'go test' command works correctly
RUN CGO_ENABLED=0 go test -v ./...

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -o /snitch \
    -ldflags='-w -s -extldflags "-static"' -a \
    ./cmd/snitch.go

# Final stage using a minimal base image
FROM ${TARGET_IMAGE} as final

# Copy the binary from the builder stage
COPY --from=builder /snitch /snitch

# Command to run the application
CMD ["/snitch"]
