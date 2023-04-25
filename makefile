PACKAGE_LIST := $(shell go list ./...)
tenkiGetter:
	go build -o tenkiGetter $(PACKAGE_LIST)
test:
	go test -coverpkg=$(PACKAGE_LIST)
clean:
	rm -f tenkiGetter
