# For the image to be ran, the following environmental variables must be set:
# - POSTGRES_URL
#
# - REDIS_PORT
# - REDIS_USER
# - REDIS_PASS
#
# - ACCESS_TOKEN_DURATION
# - REFRESH_TOKEN_DURATION
#
# - EMAIL_SENDER_ADDRESS
# - EMAIL_SENDER_PASSWORD
#
# All other environmental variables are already set to their default values in this Dockerfile.

# Build stage
FROM library/golang:1.21-alpine AS builder
WORKDIR /src

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o build/grpc-api cmd/grpc-api/main.go

# Run stage
FROM library/alpine
WORKDIR /src

COPY --from=builder /src/build/web-grpc-api build/

# Environmental variables
ENV MODE=dev
ENV GRPC_API_PORT=:3020
ENV GRPC_API_STATIC_PATH=static

EXPOSE 3020

CMD [ "/src/build/grpc-api" ]
