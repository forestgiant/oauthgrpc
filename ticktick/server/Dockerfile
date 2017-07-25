FROM debian

MAINTAINER Forest Giant Inc <docker@forestgiant.com>

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
ENV PORT=9001
ENV NAME=TICKDOCK
ENV TPS=2

RUN apt-get update
RUN apt-get install -y build-essential
RUN apt-get install -y git
RUN apt-get install -y golang
RUN go get -v -d github.com/forestgiant/oauthgrpc/ticktick/server
RUN cd /go/src/github.com/forestgiant/oauthgrpc/ticktick/server && go install
RUN rm -rf /go/pkg /go/src /usr/local/go

EXPOSE $PORT

CMD server -name=$NAME -port=$PORT -tps=$TPS