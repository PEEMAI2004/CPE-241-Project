# Use the official Golang image to build the app
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy everything from your project into the container's /app folder
COPY . .

# Download Go dependencies (go.mod, go.sum)
RUN go mod tidy

# Build the Go application, output binary called "app"
RUN go build -o app

# Expose port 6000 (so Docker knows)
EXPOSE 6000

# Start the built app
CMD ["./app"]
