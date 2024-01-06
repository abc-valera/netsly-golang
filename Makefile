### Code generation ###
generate_single_openapi_file:
	podman run --rm -v ${PWD}:/spec:Z \
	redocly/cli bundle \
	-o /spec/gen/openapi/openapi.yaml \
	/spec/internal/port/http/schema/openapi.yaml 
generate_http_code:
	go generate gen/ogen/generate.go
generate_http_docs:
	podman run --rm -v ${PWD}:/spec:Z \
	redocly/cli build-docs \
	-o /spec/docs/http/index.html \
	/spec/internal/port/http/schema/openapi.yaml
generate_http:
	make generate_single_openapi_file
	make generate_http_code
	make generate_http_docs
generate_tailwindcss:
	npx tailwindcss -i internal/port/htmx/static/style/input.css -o internal/port/htmx/static/style/style.css

generate_grpc_code:
	rm -f gen/pb/*.pb.go
	protoc internal/port/grpc/proto/*.proto \
    --go_out=gen/pb \
    --go_opt=paths=source_relative \
	--go-grpc_out=gen/pb \
    --go-grpc_opt=paths=source_relative \
    --proto_path=internal/port/grpc/proto

generate_db_code:
	go generate gen/ent/generate.go

### Podman commands ###
run_flugo-db:
	podman run \
	--name flugo-db \
	-p 5432:5432 \
	-e POSTGRES_USER=flugo \
	-e POSTGRES_PASSWORD=flugo \
	-e POSTGRES_DB=flugo \
	-d postgres:15-alpine
run_flugo-redis:
	podman run \
	--name flugo-redis \
	-p 6379:6379 \
	-d redis/redis-stack:latest

### Docker commands ###
docker_run_flugo-db:
	docker run \
	--name flugo-db \
	-p 5432:5432 \
	-e POSTGRES_USER=flugo \
	-e POSTGRES_PASSWORD=flugo \
	-e POSTGRES_DB=flugo \
	-d postgres:15-alpine
docker_run_flugo-redis:
	docker run \
	--name flugo-redis \
	-p 6379:6379 \
	-d redis/redis-stack:latest
compose_build_flugo-api:
	docker compose -f dev.docker-compose.yml build flugo-api
compose_run_docker:
	docker compose -f dev.docker-compose.yml up


### Local run commands ###
run_infrastructure_local:
	podman rm -f flugo-db
	podman rm -f flugo-redis
	make run_flugo-db
	make run_flugo-redis
	sleep 3
run_flugo_htmx_local:
	go build -o build/flugo-htmx cmd/htmx/main.go
	./build/flugo-htmx
run_flugo_http_local:
	go build -o build/flugo-http cmd/http/main.go
	./build/flugo-http
run_flugo_ws_local:
	go build -o build/flugo-ws cmd/ws/main.go
	./build/flugo-ws
run_flugo_grpc_local:
	go build -o build/flugo-grpc cmd/grpc/main.go
	./build/flugo-grpc


### Other ###
evans_client:
	evans -p 3030 -r repl
