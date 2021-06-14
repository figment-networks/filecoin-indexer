# ------------------------------------------------------------------------------
# Builder Image
# ------------------------------------------------------------------------------
FROM golang:1.15 AS build

WORKDIR /build

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux

RUN make build

# ------------------------------------------------------------------------------
# Target Image
# ------------------------------------------------------------------------------
FROM alpine:3.10 AS runner

RUN adduser --system --uid 1234 figment

USER 1234

WORKDIR /app

COPY --from=build /build/filecoin-indexer /app/bin/filecoin-indexer
COPY --from=build /build/migrations /app/migrations

EXPOSE 8080

CMD ["/app/bin/filecoin-indexer"]
