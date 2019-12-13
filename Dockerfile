FROM golang:1.13 as builder

RUN mkdir -p /csv_parser/

WORKDIR /csv_parser

COPY . .

RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
    CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags "-s -w \
    -X github.com/cybervagabond/csv_parser/pkg/version.REVISION=${GIT_COMMIT}" \
    -a -o bin/csv_parser cmd/csv_parser/*

FROM alpine:latest

RUN addgroup -S app \
    && adduser -S -g app app \
    && apk --no-cache add \
    curl openssl netcat-openbsd mc

WORKDIR /home/app

COPY --from=builder /csv_parser/bin/csv_parser .

RUN chown -R app:app ./

USER app


CMD ["./csv_parser"]
