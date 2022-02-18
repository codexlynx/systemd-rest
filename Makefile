build: dist/systemd-rest_amd64

test:
	@MODE=debug PORT=7777 ./test/run.sh $(ARGS)

release:
	@echo release

gofmt:
	@gofmt -s -w .

clean:
	@rm dist/*

dist/systemd-rest_amd64:
	@DOCKER_BUILDKIT=1 docker build -f build/build.dockerfile --target binary --output dist/ .

all: build test release
.PHONY: test release clean gofmt
