# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /src

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o build/netsly cmd/main.go

# Run stage
FROM alpine
WORKDIR /src

COPY --from=builder /src/internal/presentation/jsonApi/reststatic static/
COPY --from=builder /src/build/netsly build/

# Environmental variables
ENV MODE=dev
ENV JSON_API_PORT=:3010
ENV JSON_API_STATIC_PATH=static

EXPOSE 3010

CMD /src/build/netsly -entrypoint jsonApi
