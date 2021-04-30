FROM golang:1.16.3-buster AS builder

ARG servername="server1"

WORKDIR /opt/server/
COPY . /opt/server/
ENV GOPATH=/go \
  GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN go build -o /opt/server/app ${servername}.go

# --------------------------------------------------------

FROM debian:buster-slim
COPY --from=builder /opt/server/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]