# --- Stage 1: The Build Stage ---
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the final binary.
# CGO_ENABLED=0 ensures the binary is statically linked and portable.
# -o /goapp sets the output path and name.
RUN CGO_ENABLED=0 go build -ldflags='-w -s' -o /goapp ./cmd/main.go


# --- Stage 2: The Final Production Stage ---
# Use a minimal image (alpine) to reduce size and attack surface
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /goapp .

# Set the entry point to run the compiled application
CMD ["./goapp"]