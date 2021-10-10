.PHONY: test
test:
	go test ./...

.PHONY: dev
dev:
	go build -v -o longform ./cmd/longform

.PHONY: container
container: dev test
	docker build -f Dockerfile . -t longform

.PHONY: clean
clean:
	rm -f longform
	docker image rm -f longform

.PHONY: run
run: container
	docker run -p 8080:8080 longform

