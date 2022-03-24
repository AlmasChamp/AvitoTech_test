FROM golang:1.17.6
RUN mkdir /app
WORKDIR /app
ADD . .
ADD go.mod .
ADD go.sum .
RUN go mod download
RUN go build -o main main.go
CMD ["./main"]
