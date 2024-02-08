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

COPY --from=builder /src/build/netsly build/

# Environmental variables
ENV MODE=dev
ENV GRPC_API_PORT=:3020
ENV GRPC_API_STATIC_PATH=static

EXPOSE 3020

CMD [ "/src/build/netsly -grpc-api" ]
