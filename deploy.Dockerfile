# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /src
COPY . .
RUN go build -o flugo cmd/http/main.go

# Run stage
FROM alpine
WORKDIR /src
COPY --from=builder /src/flugo /src/cmd/http/
COPY .deploy.env .
COPY docs docs
ENV CONFIG_PATH /src/.deploy.env
ENV HTTP_DOCS_PATH /src/docs/http

EXPOSE 3000

CMD [ "/src/cmd/http/flugo" ]
