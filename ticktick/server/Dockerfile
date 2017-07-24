FROM debian

MAINTAINER Forest Giant Inc <docker@forestgiant.com>

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
ENV PORT 9001
ENV NAME TICKDOCK
ENV TPS 2

RUN  \
     apt-get update && \ 
     apt-get install -y build-essential && \
     apt-get install -y git && \
     apt-get install -y golang && \
     go get -v -d gitlab.fg/go/oauthgrpc && \
     cd /go/src/gitlab.fg/go/oauthgrpc/ticktick/server && \
     go install && \
     rm -rf /go/pkg /go/src /usr/local/go

EXPOSE $PORT

CMD server -name=$NAME -port=$PORT -tps=$TPS