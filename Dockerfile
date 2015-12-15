FROM golang:1.5.1

ENV USER root
WORKDIR /go/src/github.com/schneidexe/tpl

ADD . /go/src/github.com/schneidexe/tpl

RUN mkdir bin

CMD for OS in darwin linux windows; do \
        export GOOS=$OS; \
        for ARCH in 386 amd64; do \
            export GOARCH=$ARCH; \
            echo building $GOOS-$GOARCH; \
            if [ $GOOS = "windows" ]; then \
                go build -o bin/tpl-$GOOS-$GOARCH.exe; \
            else \
                go build -o bin/tpl-$GOOS-$GOARCH; \
            fi \
        done; \
    done
