# Build stage
FROM library/golang:1.21-alpine AS builder
WORKDIR /src

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o build/netsly cmd/app/main.go

# Run stage
FROM library/alpine
WORKDIR /src

COPY --from=builder /src/internal/port/json-rest-api/static static/
COPY --from=builder /src/build/netsly build/

# Environmental variables
ENV MODE=dev
ENV JSON_REST_API_PORT=:3010
ENV JSON_REST_API_STATIC_PATH=static

EXPOSE 3010

CMD /src/build/netsly -entrypoint json-rest-api
