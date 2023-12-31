FROM golang:1.21.5
WORKDIR /search-primes
COPY *.go .
COPY sieve/ sieve
RUN go mod init homecredit.vn/prime-go
RUN go mod tidy
RUN go build -o main
CMD ["./main"]