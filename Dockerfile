# Start from a minimal image that contains Go compiler
FROM golang:1.22 AS build

# Set work directory inside the container
WORKDIR /app

# Copy Go module files first (go.mod and go.sum)
COPY go.mod go.sum ./

# Download Go module
RUN go mod download

# Copy the entire project
COPY . ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -a -o probe .

# Start a new stage with a busybox image
FROM busybox:1.34.1

# Copy the binary file from the previous stage
COPY --from=build /app/probe /probe

# Expose port 3000
EXPOSE 3000

# Command to run the binary
ENTRYPOINT ["/probe"]