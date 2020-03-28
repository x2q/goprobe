#FROM golang

#RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends libpcap-dev libpcap0.8 libpcap0.8-dev \
#    && rm -rf /var/lib/apt/lists/* \
#    && apt-get -y autoremove


FROM golang:alpine

RUN apk update && apk add --no-cache libpcap libpcap-dev git build-base gcc wget

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o /go/bin/goprobe
RUN cp config_sample.tml config.tml


#RUN /go/bin/goprobe install

#ENTRYPOINT ["/go/bin/goprobe"]
