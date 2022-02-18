FROM golang:1.15.5-buster AS builder
WORKDIR /go/src/github.com/codexlynx/systemd-rest

RUN apt-get update -y \
    && apt-get install libsystemd-dev -y \
    && mkdir -p /build/dist/

COPY . .
RUN go build -o /build/dist/systemd-rest-x86_64 ./cmd/systemd-rest

FROM scratch AS binary
COPY --from=builder /build/dist /
