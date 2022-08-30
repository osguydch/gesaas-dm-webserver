BINARY=rmserver

MODULE=RmServer
DESC=RunManager Server repo
VERSION=`git describe --tags --long`
BUILD=`date +%F\ %T`

LDFLAGS=-ldflags "-X 'rm/common.Module=${MODULE}' -X 'rm/common.Desc=${DESC}' -X 'rm/common.Version=${VERSION}' -X 'rm/common.Build=${BUILD}'"

build:
	go build ${LDFLAGS} -o ${BINARY}

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY}; fi
