
GO ?= go

HTMX_VERSION ?= $(shell cat ./HTMX_VERSION)
NUNJUCKS_VERSION ?= $(shell cat ./NUNJUCKS_VERSION)

LDFLAGS_COMMON = -X main.htmx_version=$(HTMX_VERSION) -X main.nunjucks_version=$(NUNJUCKS_VERSION)
GO_BUILD := $(GO) build $(EXTRA_FLAGS) -ldflags "$(LDFLAGS_COMMON)"

.DEFAULT: all

.PHONY: all
%/main: %/main.go
	$(GO_BUILD) -o $@ $@.go

.PHONY: all
all: \
	active-search/main \
	animations/main \
	bulk-update/main \
	click-to-edit/main \
	click-to-load/main \
	debug/main \
	delete-row/main	\
	dialogs/main \
	edit-row/main \
	infinite-scroll/main \
	inline-validation/main \
	keyboard-shortcuts/main \
	lazy-load/main \
	modal-bootstrap/main \
	multi-swap/main \
	path-params/main \
	progress-bar/main \
	response-targets/main \
	restored/main \
	sortable/main \
	tab/main \
	value-select/main \
	websocket/main \
	websocket-echo/main

# use golines to limit line length
# https://github.com/segmentio/golines
.PHONY: all
format:
	gofmt -w */*.go
	golines -w */*.go

.PHONY: all
clean:
	rm -rf */main
