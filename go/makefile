CC=go
RM=rm
MV=mv
WSDLGEN=wsdl2go


SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
GOOS=linux
GOARCH=amd64

VERSION=1.0
APP=EmutationClient

BUILD_TIME=`date +%FT%T%z`
PACKAGES := github.com/fiorix/wsdl2go

LIBS=

LDFLAGS=

.DEFAULT_GOAL:=test

help:
		@echo "Options possible :"
		@echo "*** test (test all project's units tests)"
		@echo "*** DEFAULT (create binary)"
		@echo "*** client (generate wsdl client go source)"
		@echo "*** fmt (execute gofmt to format code)"
		@echo "*** audit (execute static code check)"
		@echo "*** get (gets all dependencies)"
		@echo "*** init (clean and initialise the project)"
		@echo "*** execute (execute the binary)"
		@echo "*** clean (clean binaries and project)"

copy:
		@cp ../emutation.wsdl .

test: $(APP)
		@GOOS=${GOOS} GOARCH=${GOARCH} go test ./...
		@echo " Tests OK."

$(APP): get $(SOURCES)
		@echo "    Sources compilation ${BUILD_TIME}"
		@GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS} -o ${APP}-${VERSION} $(SOURCEDIR)/main.go
		@echo "    ${APP}-${VERSION} generated."

fmt: audit
		@echo "    Go FMT"
		@$(foreach element,$(SOURCES),go fmt $(element);)

audit: get
		@go tool vet -all $(SOURCES) > audit.log
		@echo "    Audit done"
		@echo "    check the results audit in audit.log file."

get: init copy
		@echo "    Download packages"
		@$(foreach element,$(PACKAGES),go get $(element);)

init: clean
		@echo "    Init of the project"
		@echo "    Version :: ${VERSION}"

execute:
		./${APP}-${VERSION}  

clean:
		@if [ -f "${APP}-${VERSION}" ] ; then rm ${APP}-${VERSION} ; fi
		@echo "    Clean Done"

client: get
		@echo "    Generate client code source."
		@${WSDLGEN} < emutation.wsdl > emutation/client.go


