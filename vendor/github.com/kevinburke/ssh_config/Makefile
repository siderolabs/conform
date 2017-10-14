STATICCHECK := $(shell command -v staticcheck)

lint:
	go vet ./...
ifndef STATICCHECK
	go get -u honnef.co/go/tools/cmd/staticcheck
endif
	staticcheck ./...

test: lint
	@# the timeout helps guard against infinite recursion
	go test -timeout=30ms ./...
