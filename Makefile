COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
  COVERAGEDIR = $(CIRCLE_ARTIFACTS)
endif

ifdef VERBOSE
V = - v
else
.SILENT:
endif

all: test cover
fmt:
	go fmt ./...
test:
	mkdir -p coverage
	go test $(V) ./... -race -cover -coverprofile=$(COVERAGEDIR)/ft.coverprofile
cover:
	go tool cover -html=$(COVERAGEDIR)/ft.coverprofile -o $(COVERAGEDIR)/ft.html

coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN)

clean:
	go clean
	rm -rf coverage/
