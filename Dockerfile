FROM golang:1.21-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o university-go-project ./main.go

CMD ["./university-go-project"]