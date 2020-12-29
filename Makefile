build: dist/systemd-rest

test:
	@MODE=debug PORT=7777 ./test/run.sh $(ARGS)

release:
	@echo release

gofmt:
	@find src -type f -name "*.go" -execdir go fmt \;

clean:
	@rm dist/systemd-rest

dist/systemd-rest:
	@DOCKER_BUILDKIT=1 docker build -f src/dockerfile.build --target binary --output dist/ .

all: build test release
.PHONY: test release clean gofmt
