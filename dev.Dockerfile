# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /src
COPY . .
RUN go build -o flugo cmd/http/main.go

# Run stage
FROM alpine
WORKDIR /src
COPY --from=builder /src/flugo .
COPY .dev.env .

EXPOSE 3000

CMD [ "/src/flugo" ]
