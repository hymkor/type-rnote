ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=set
    DEL=del
    NUL=nul
else
    SET=export
    DEL=rm
    NUL=/dev/null
endif

NAME=$(lastword $(subst /, ,$(abspath .)))
VERSION=$(shell git.exe describe --tags 2>$(NUL) || echo v0.0.0)
GOOPT=-ldflags "-s -w -X main.version=$(VERSION)"
EXE=$(shell go env GOEXE)

all:
	go fmt
	$(SET) "CGO_ENABLED=0" && go build $(GOOPT)

test:
	go test -v

_dist:
	$(SET) "CGO_ENABLED=0" && go build $(GOOPT)
	zip -9 $(NAME)-$(VERSION)-$(GOOS)-$(GOARCH).zip $(NAME)$(EXE)

dist:
	$(SET) "GOOS=linux" && $(SET) "GOARCH=386"   && $(MAKE) _dist
	$(SET) "GOOS=linux" && $(SET) "GOARCH=amd64" && $(MAKE) _dist
	$(SET) "GOOS=windows" && $(SET) "GOARCH=386"   && $(MAKE) _dist
	$(SET) "GOOS=windows" && $(SET) "GOARCH=amd64" && $(MAKE) _dist

clean:
	$(DEL) *.zip $(NAME)$(EXE)

manifest:
	make-scoop-manifest *-windows-*.zip > $(NAME).json

release:
	gh release create -d -t $(VERSION) $(VERSION) $(wildcard $(NAME)-$(VERSION)-*.zip)

.PHONY: release clean dist manifest _dist test all
