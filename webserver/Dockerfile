FROM debian

MAINTAINER Forest Giant Inc <docker@forestgiant.com>

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
ENV TICK1=""
ENV TICK2=""
ENV TICK3=""
ENV TICK4=""
ENV TICK5=""
ENV PORT=9001

RUN apt-get update
RUN apt-get install -y build-essential
RUN apt-get install -y git
RUN apt-get install -y golang
RUN go get -v -d github.com/forestgiant/oauthgrpc/webserver
RUN cd /go/src/github.com/forestgiant/oauthgrpc/webserver && go install
RUN rm -rf /go/pkg /go/src /usr/local/go

EXPOSE 9001

CMD webserver -tick1=$TICK1 -tick2=$TICK2 -tick3=$TICK3 -tick4=$TICK4 -tick5=$TICK5 -port=$PORT 