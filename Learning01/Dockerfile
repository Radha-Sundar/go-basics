FROM golang:alpine
RUN mkdir /go-basic
COPY . /go-basic
WORKDIR /go-basic
RUN go build -o main .
CMD ["/go-basic/main"]