FROM golang:1.14 AS builder

ADD . /go/src/github.com/replicatedhq/troubleshoot-preview
WORKDIR /go/src/github.com/replicatedhq/troubleshoot-preview

RUN make build


FROM debian:stretch-slim

RUN apt-get update -y && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/github.com/replicatedhq/troubleshoot-preview/bin /app

EXPOSE 3000

CMD ["/app/troubleshoot-preview"]
