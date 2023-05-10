.PHONY: all docs
all: docs speakeasy

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s openapi.yaml

docs:
	go generate ./...

