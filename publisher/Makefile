.PHONY: 

coverage:
	go test -cover -count=1 ./...

test:
	go test -count=1 ./...

run:
	go run ./cmd/publisher

dockerize:
	docker build \
	--tag publisher \
	--build-arg package=./cmd/publisher \
	--build-arg config=config.yaml \
	.

run-docker:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	publisher

run-docker-with-mount:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	--mount type=bind,source=$(PWD)/config/config.dev.yaml,target=/home/deploy/config/config.yaml,readonly \
	publisher

run-docker-with-config:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	--env APP_CONFIG_FILE=config.dev.yaml \
	publisher
