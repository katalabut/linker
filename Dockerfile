FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/katalabut/linker

# Build the golang-docker command inside the container.
RUN go install github.com/katalabut/linker

# Run the golang-docker command when the container starts.
ENTRYPOINT /go/bin/katalabut/linker

# http server listens on port 8080.
EXPOSE 8080