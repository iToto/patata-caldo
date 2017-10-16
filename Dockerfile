# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy binary to container
ADD . /go/src/github.com/iToto/patata-caldo

# Build the authorized command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
WORKDIR /go/src/github.com/iToto/patata-caldo
RUN go get github.com/tools/godep
RUN godep go install ./...

# Run the authorized command by default when the container starts.
ENTRYPOINT /go/bin/patata-caldo

ENV PORT 8080

# Document that the service listens on port 8080.
EXPOSE 8080