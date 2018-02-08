FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download -t   # "go get -d -v -t ./..."
RUN go-wrapper install    # "go install -v ./..."

RUN go get github.com/tebeka/go2xunit

asd
sssss