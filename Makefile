PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.0.1
NAME := tenkiGetter
DIST := $(NAME)-$(VERSION)


tenkiGetter:coverage.out
	go build -o tenkiGetter $(PACKAGE_LIST)

coverage.out:
	go test -covermode=count \
		-coverprofile=coverage.out $(PACKAGE_LIST)

# refer from https://pod.hatenablog.com/entry/2017/06/13/150342
define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/$(NAME)/main.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
#	cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: urleap
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,arm64,.exe)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,linux,arm64,)

distclean: clean
	rm -rf dist

clean:
	rm -f tenkiGetter
