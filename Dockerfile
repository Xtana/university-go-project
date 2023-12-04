FROM ubuntu:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -0 struct home/main.go

CMD ["./struct"]