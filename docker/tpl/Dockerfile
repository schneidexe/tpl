FROM alpine

ENV VERSION 0.4.5
ENV SRC_PATH  /src
ENV DST_PATH /dest
ENV DEPTH 1

ADD https://github.com/schneidexe/tpl/releases/download/v$VERSION/tpl-linux-amd64 /usr/bin/tpl
ADD run.sh /run.sh
RUN apk add findutils \
    && chmod +x /usr/bin/*

CMD /run.sh
