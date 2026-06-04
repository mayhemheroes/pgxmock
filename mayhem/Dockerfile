FROM golang:1.25 AS builder
ENV GOTOOLCHAIN=auto

COPY . /pgxmock
WORKDIR /pgxmock

RUN go test -c -fuzz=FuzzPgxmock -o /fuzzPgxmock ./mayhem/

FROM debian:bookworm-slim
COPY --from=builder /fuzzPgxmock /fuzzPgxmock
