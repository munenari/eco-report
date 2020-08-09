FROM golang:1.14-alpine

WORKDIR /code

RUN apk update && \
    apk --no-cache add git
RUN go get -v github.com/pilu/fresh

ENV PATH /go/bin:$PATH

CMD ["fresh"]
