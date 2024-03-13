# Build stage
FROM library/golang:1.21-alpine AS builder
WORKDIR /src

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o build/netsly cmd/main.go

# Run stage
FROM library/alpine
WORKDIR /src

COPY --from=builder /src/internal/presentation/webApp/static static/
COPY --from=builder /src/build/netsly build/

# Environmental variables
ENV MODE=dev
ENV WEB_APP_PORT=:3000
ENV WEB_APP_STATIC_PATH=static

EXPOSE 3000

CMD [ "/src/build/netsly -webApp" ]
