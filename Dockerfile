FROM golang:1.14

WORKDIR /go/src/go-transactions

ARG root_dir
COPY ${root_dir} .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-transactions"]

EXPOSE 8080
