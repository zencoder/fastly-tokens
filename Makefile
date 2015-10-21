GO ?= go
COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

all: test cover
fmt:
	$(GO) fmt ./...
test:
	if [ ! -d coverage ]; then mkdir coverage; fi
	$(GO) test -v ./... -race -cover -coverprofile=$(COVERAGEDIR)/ft.coverprofile
cover:
	$(GO) tool cover -html=$(COVERAGEDIR)/ft.coverprofile -o $(COVERAGEDIR)/ft.html
tc: test cover
coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)
clean:
	$(GO) clean
	rm -rf coverage/
godep-save:
	godep save ./...
