GO ?= go
STATICCHECK ?= staticcheck
GOCACHE ?= $(CURDIR)/.gocache
GOMODCACHE ?= $(CURDIR)/.modcache
BIN := $(CURDIR)/bin/ito

GOENV := GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE)

.PHONY: build fmt vet staticcheck clean

build: fmt vet staticcheck $(BIN)

$(BIN):
	@mkdir -p $(dir $@)
	$(GOENV) $(GO) build -o $@ .

fmt:
	$(GOENV) $(GO) fmt ./...

vet:
	$(GOENV) $(GO) vet ./...

staticcheck:
	@command -v $(STATICCHECK) >/dev/null || (echo "staticcheck コマンドが見つかりません。STATICCHECK=path/to/staticcheck で指定してください。" >&2; exit 1)
	$(GOENV) $(STATICCHECK) ./...

clean:
	rm -rf $(BIN) $(GOCACHE) $(GOMODCACHE)
