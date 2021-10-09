.DEFAULT_GOAL = build

COMMIT ?= `git rev-parse --short HEAD 2>/dev/null`
DATE ?= `git log -1 --format=%cd --date=format:"%Y%m%d"`
VERSION ?= `git describe --abbrev=0 --tags $(git rev-list --tags --max-count=1) 2>/dev/null | sed 's/v\(.*\)/\1/'`

ifeq ($(VERSION),)
	VERSION := $(COMMIT)-$(DATE)
endif

COMMIT_FLAG := -X version/version.GitCommit=$(COMMIT)
VERSION_FLAG := -X version/version.Version=$(VERSION)
LDFLAGS := -ldflags "$(COMMIT_FLAG) $(VERSION_FLAG)"

mod:
	go mod download

build: mod
	go build -o garrus $(LDFLAGS)