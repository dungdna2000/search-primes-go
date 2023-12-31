FROM golang:latest 
WORKDIR /search-primes
RUN go mod init homecredit.vn/prime-go
COPY *.go /search-primes/
RUN go mod download
RUN go build -o main
CMD ["./main"]