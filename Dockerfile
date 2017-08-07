FROM golang:1.8

RUN mkdir -p /go/src/finger
WORKDIR /go/src/finger

COPY . /go/src/finger

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

CMD ["go-wrapper", "run"]
