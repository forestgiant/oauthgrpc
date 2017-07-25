FROM debian

MAINTAINER Forest Giant Inc <docker@forestgiant.com>

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
ENV ADDRESS=127.0.0.1:9001

RUN apt-get update
RUN apt-get install -y build-essential
RUN apt-get install -y git
RUN apt-get install -y golang
RUN go get -v -d github.com/thales17/oauthgrpc/webserver
RUN cd /go/src/github.com/thales17/oauthgrpc/webserver && go install
RUN rm -rf /go/pkg /go/src /usr/local/go

EXPOSE 9001

CMD webserver -address=$ADDRESS