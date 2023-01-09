FROM golang:1.18 as builder

COPY . /pgxmock
WORKDIR /pgxmock/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzPgxmock.a && \
    clang -fsanitize=fuzzer fuzzPgxmock.a -o fuzzPgxmock.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /pgxmock/mayhem/fuzzPgxmock.libfuzzer /