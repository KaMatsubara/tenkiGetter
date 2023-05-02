PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.0.1
NAME := tenkiGetter
DIST := $(NAME)-$(VERSION)


tenkiGetter: coverage.out
	go build -o tenkiGetter $(PACKAGE_LIST)

coverage.out:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)

clean:
	rm -f tenkiGetter
