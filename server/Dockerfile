# Stage 1: Build the Go app
FROM golang:1.23.1-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go app
RUN go build -o myapp .

# Stage 2: Run the app in a smaller image
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .
COPY app.env .

# Expose the port the app will run on
EXPOSE 9090

# Command to run the app
CMD ["./myapp"]
