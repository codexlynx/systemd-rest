FROM golang:1.15.5-buster AS builder
WORKDIR /go/src/github.com/codexlynx/systemd-rest

RUN apt-get update -y \
    && apt-get install libsystemd-dev gcc-arm-linux-gnueabi -y \
    && mkdir -p /build/dist/

COPY . .
RUN go build -o /build/dist/systemd-rest_amd64 ./cmd/systemd-rest
RUN GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc go build -o /build/dist/systemd-rest_arm ./cmd/systemd-rest

FROM scratch AS binary
COPY --from=builder /build/dist /
