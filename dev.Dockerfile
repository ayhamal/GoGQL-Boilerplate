# Specifies a parent image
FROM golang:1.19-alpine
# Creates work directory
WORKDIR /app
# Copy application data into image
COPY . .
# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./
# Installs Go dependencies
RUN go mod download
# Copy only `.go` files, if you want all files to be copied then replace `with `COPY . .` for the code below.
COPY *.go .

# Install our third-party application for hot-reloading capability.
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build -o ./bin/" -command="./bin/gogql-boilerplate" -directory="./"
