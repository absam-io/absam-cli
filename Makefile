# absam-cli Makefile

TARGET = absam-cli
PREFIX = /usr
BINDIR = ${PREFIX}/bin

build:
	@echo Building ${TARGET}
	@go build -o ${TARGET} main.go

run:
	@go run main.go

clean:
	@echo Cleaning
	@rm -rf bin
	@rm ${TARGET}

install:
	@echo Installing ${TARGET} at ${BINDIR}
	@mv ${TARGET} ${BINDIR}

uninstall:
	@echo Removing ${TARGET} from ${BINDIR}
	@rm -f ${BINDIR}/${TARGET}

compile:
	@echo Compiling ${TARGET} for 32-Bit systems
	@GOOS=freebsd GOARCH=386 go build -o bin/${TARGET}-freebsd-386 main.go
	@GOOS=linux GOARCH=386 go build -o bin/${TARGET}-linux-386 main.go
	@echo Compiling ${TARGET} for 64-Bit systems
	@GOOS=freebsd GOARCH=amd64 go build -o bin/${TARGET}-freebsd-amd64 main.go
	@GOOS=linux GOARCH=amd64 go build -o bin/${TARGET}-linux-amd64 main.go

all: build install
