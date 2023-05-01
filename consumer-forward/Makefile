.PHONY: 

coverage:
	go test -cover -count=1 ./...

test:
	go test -count=1 ./...

run:
	go run ./cmd/consumer

dockerize:
	docker build \
	--tag consumer \
	--build-arg package=./cmd/consumer \
	--build-arg config=config.yaml \
	.


