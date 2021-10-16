# start from the golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# install git
RUN apk update && apk add --no-cache git

# Set current directory as working directory
WORKDIR /app

# copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# copy the source
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .


# Finally our multi-stage to build a small image
# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .

# Run executable
CMD ["./main"]