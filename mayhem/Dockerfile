FROM golang:1.24 AS builder

COPY . /pgxmock
WORKDIR /pgxmock

RUN go test -c -fuzz=FuzzPgxmock -o /fuzzPgxmock ./mayhem/

FROM debian:bookworm-slim
COPY --from=builder /fuzzPgxmock /fuzzPgxmock
