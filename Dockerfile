FROM golang:1.11-alpine

WORKDIR /go/src/github.com/munenari/eco-report

RUN apk update && \
    apk --no-cache add git
RUN go get -v github.com/oxequa/realize
RUN apk del --purge git

ENV PATH /go/bin:$PATH

CMD ["realize", "start", "--server"]
