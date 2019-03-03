FROM golang:1.12-alpine

WORKDIR /go/src/github.com/munenari/eco-report

RUN apk update && \
    apk --no-cache add git
RUN go get -v github.com/oxequa/realize
RUN go get -v github.com/golang/dep/cmd/dep

ENV PATH /go/bin:$PATH

CMD ["sh", "-c", "dep ensure && realize start"]
