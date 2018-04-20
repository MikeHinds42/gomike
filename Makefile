



all: gomike

gomike:
	go build

test1:
	go build


get:
	go get github.com/gorilla/sessions


.PHONE: gofmt2
gofmt2: FORCE
	$(shell find -name *.go -exec gofmt -w '{}' \; )
	echo "hello"

FORCE: ;


clean:
	rm -f gomike
