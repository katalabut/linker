FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/katalabut/linker

RUN cd /go/src/github.com/katalabut/linker && go get && go build
# Build the golang-docker command inside the container.
#RUN go build
#RUN go install github.com/katalabut/linker

#ADD ./static /go/bin/static
WORKDIR /go/src/github.com/katalabut/linker
# Run the golang-docker command when the container starts.
ENTRYPOINT /go/src/github.com/katalabut/linker/linker

# http server listens on port 8080.
EXPOSE 8080