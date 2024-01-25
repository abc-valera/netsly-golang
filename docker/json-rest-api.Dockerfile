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
RUN go build -o build/json-rest-api cmd/json-rest-api/main.go

# Run stage
FROM library/alpine
WORKDIR /src

COPY --from=builder /src/internal/port/json-rest-api/static static/
COPY --from=builder /src/build/json-rest-api build/

# Environmental variables
ENV MODE=dev
ENV JSON_REST_API_PORT=:3010
ENV JSON_REST_API_STATIC_PATH=static

EXPOSE 3010

CMD [ "/src/build/json-rest-api" ]
