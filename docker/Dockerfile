# build the server binary
FROM golang:1.13.5 AS builder
LABEL stage=server-intermediate
WORKDIR /go/src/github.com/egbot7/eggbot
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/eggbot .

# copy the server binary from builder stage; run the server binary
FROM alpine:latest AS runner
WORKDIR /bin
COPY --from=builder /go/src/github.com/egbot/eggbot/bin/eggbot .

ENTRYPOINT ["eggbot"]
