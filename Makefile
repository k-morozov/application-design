OLD_APP=cmd/old/main.go
BOOKING_APP=cmd/booking/main.go
BOOKING_BIN=bin/booking

build:
	go build -o ${BOOKING_BIN} $(BOOKING_APP)

run: build
	./${BOOKING_BIN} -p 8080 -l "debug"

run_old:
	go run $(OLD_APP)

test:
	go test -v ./...

test_race:
	go test -race -v ./...

test_integration:
	pytest ./test/test_integration.py

all: build