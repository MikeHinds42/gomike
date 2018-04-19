



all: main

main:
	go build -o ./build/$@ $@.go

test1:
	go build -o ./build/$@ web/tests/$@.go


.PHONE: gofmt2
gofmt2: FORCE
	$(shell find -name *.go -exec gofmt -w '{}' \; )
	echo "hello"

FORCE: ;


