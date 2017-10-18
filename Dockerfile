# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM busybox

# Copy binary to container
COPY ./bin/potato /potato

ENV PORT 9999

CMD ["/potato"]