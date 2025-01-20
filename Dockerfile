FROM golang:1.23 as builder

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy the rest of the application code
COPY . ./

# Step 6: Build the Go binary
RUN go build -o server ./

# Step 10: Expose the application port
EXPOSE 8080

# Step 11: Command to run the binary
CMD ["./server"]