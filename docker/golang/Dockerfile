FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go get github.com/pilu/fresh
CMD ["fresh"]
