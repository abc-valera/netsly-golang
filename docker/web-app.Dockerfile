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
RUN go build -o build/web-app cmd/web-app/main.go

# Run stage
FROM library/alpine
WORKDIR /src

COPY --from=builder /src/internal/port/web-app/static static/
COPY --from=builder /src/build/web-app build/

# Environmental variables
ENV MODE=dev
ENV WEB_APP_PORT=:3000
ENV WEB_APP_STATIC_PATH=static

EXPOSE 3000

CMD [ "/src/build/web-app" ]
