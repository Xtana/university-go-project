FROM ubuntu:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -0 struct ./main.go

CMD ["./struct"]