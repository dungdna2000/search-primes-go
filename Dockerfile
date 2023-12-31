FROM golang:latest 
WORKDIR /search-primes
RUN go mod init homecredit.vn/primes
COPY *.go /search-primes/
RUN go mod tidy
RUN go build -o main
CMD ["./main"]