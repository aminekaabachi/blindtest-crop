
FROM golang:1.14-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/blindtest-crop

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./blindtest-crop .

# Run the binary program produced by `go install`
ENTRYPOINT ["./blindtest-crop"]