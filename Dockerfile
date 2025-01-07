# Step 1: Build the Go app
FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and Download Dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with statically linked binary
RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags '-s -w' -o go-app .

# Debug: List contents of /app to confirm the binary exists
RUN ls -l /app

# Step 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Optionally, add dependencies like libc6-compat (if needed)
# RUN apk add --no-cache libc6-compat

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/go-app .

# Debug: List contents of /root/ to confirm binary is copied correctly
RUN ls -l /root/

# Ensure the binary is executable
RUN chmod +x /root/go-app

# Command to run the application
CMD ["./go-app"]

