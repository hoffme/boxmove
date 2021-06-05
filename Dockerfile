FROM golang:latest AS builder
LABEL stage=builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/github.com/hoffme/boxmove
COPY . .
RUN go mod download
RUN make protobuf
RUN go build -o /go/bin/github.com/hoffme/boxmove /go/src/github.com/hoffme/boxmove/cmd/boxmove/main.go

FROM scratch
COPY --from=builder /go/bin/github.com/hoffme/boxmove .
ENTRYPOINT ["./boxmove"]

# docker build -t boxmove .
# docker image prune --filter label=stage=builder