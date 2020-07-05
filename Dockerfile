FROM golang:1.14.4 as gobuilder

WORKDIR /go/src/atomheartother/qtweet

COPY go.mod .
COPY go.sum .

COPY cmd/ cmd/
COPY pkg/ pkg/

RUN go build cmd/qtweet.go

CMD ["qtweet"]