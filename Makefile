# Code generation
generate_http_code:
	go generate gen/ogen/generate.go

generate_grpc_code:
	protoc internal/infrastructure/port/grpc/proto/*.proto \
    --go_out=gen/proto \
    --go_opt=paths=source_relative \
    --proto_path=.

generate_db_code:
	go generate gen/ent/generate.go

# Docker commands
run_flugo-db_container:
	docker run \
	--name flugo-db \
	-p 5432:5432 \
	-e POSTGRES_USER=flugo \
	-e POSTGRES_PASSWORD=flugo \
	-e POSTGRES_DB=flugo \
	-d postgres:15-alpine
run_flugo-redis_container:
	docker run \
	--name flugo-redis \
	-p 6379:6379 \
	-d redis/redis-stack:latest
run_docker-compose:
	docker compose -f dev.docker-compose.yml up

# Local run commands
run_infrastructure_local:
	docker rm -f flugo-db
	docker rm -f flugo-redis
	make run_flugo-db_container
	make run_flugo-redis_container
	sleep 3
run_flugo_http_local:
	go build -o build/flugo-api cmd/http/main.go
	./build/flugo-api
run_all_local:
	make run_infrastructure_local
	make run_flugo_http_local
