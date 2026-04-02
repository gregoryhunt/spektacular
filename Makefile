BINARY := spektacular
VERSION := 0.1.0

.PHONY: build test lint clean install install-local cross harbor-test

build:
	go build -ldflags "-X github.com/jumppad-labs/spektacular/cmd.version=$(VERSION)" -o ./bin/$(BINARY) .

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -f ./bin 

install-local: build
	sudo cp $(BINARY) /usr/local/bin/$(BINARY)

cross:
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o ./bin/$(BINARY)-darwin-arm64  .
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o ./bin/$(BINARY)-darwin-amd64  .
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o ./bin/$(BINARY)-linux-amd64   .
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build -o ./bin/$(BINARY)-linux-arm64   .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/$(BINARY)-windows-amd64.exe .

harbor-test:
	GOOS=linux GOARCH=amd64 go build -o tests/harbor/spec-workflow/environment/spektacular .
	ANTHROPIC_AUTH_TOKEN=$$(python3 -c "import json; print(json.load(open('$$HOME/.claude/.credentials.json'))['claudeAiOauth']['accessToken'])") \
	harbor run -p tests/harbor/spec-workflow -a claude-code -m claude-sonnet-4-6 -o tests/harbor/jobs; \
	echo ""; \
	echo "=== Test Results ==="; \
	cat $$(ls -td tests/harbor/jobs/*/spec-workflow__*/verifier/test-stdout.txt 2>/dev/null | head -1)
