.PHONY: 

coverage:
	go test -cover -count=1 ./...

test:
	go test -count=1 ./...

run:
	go run ./cmd/notifier

dockerize:
	docker build \
	--tag notifier \
	--build-arg package=./cmd/notifier \
	--build-arg config=config.yaml \
	.

run-docker:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	notifier

run-docker-with-mount:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	--mount type=bind,source=$(PWD)/config/config.dev.yaml,target=/home/deploy/config/config.yaml,readonly \
	notifier

run-docker-with-config:
	docker run -d --rm -it \
	-p 8080:8080 -p 9090:9090 \
	--env APP_CONFIG_FILE=config.container.yaml \
	notifier
