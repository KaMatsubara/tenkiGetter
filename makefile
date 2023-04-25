PACKAGE_LIST := $(shell go list ./...)
tenkiGetter:
	go build -o tenkiGetter $(PACKAGE_LIST)
test:
	go test -cover $(PACKAGE_LIST) -coverprofile=cover.out
clean:
	rm -f tenkiGetter
