# Specifies a parent image
FROM golang:1.19-alpine AS builder
# Creates an app directory to hold your app’s source code
WORKDIR /app
# Copies go.mod and to the working directory
COPY go.mod .
# Installs Go dependencies
RUN go mod download
# Copies everything from your root directory into /app
COPY *.go .
# Builds your app with optional configuration
RUN go build -o ./bin/prod_build
# Specifies bin as the final stage
FROM builder AS bin
# Copies the binary file from the builder stage
COPY --from=builder ./bin/prod_build /
# Tells Docker which network port your container listens on
EXPOSE 4000
# Specifies the executable command that runs when the container starts
CMD [ "./bin/prod_build" ]