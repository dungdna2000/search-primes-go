FROM golang:latest 
WORKDIR /search-primes
COPY go.mod go.sum /search-primes/
RUN go mod download
COPY *.go /search-primes/
RUN go build -o main
CMD ["./main"]