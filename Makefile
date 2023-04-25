PACKAGE_LIST := $(shell go list ./...)


tenkiGetter:
	go build -o tenkiGetter $(PACKAGE_LIST)
test:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
clean:
	rm -f tenkiGetter
