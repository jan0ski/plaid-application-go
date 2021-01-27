FROM golang:1.15-alpine as build

WORKDIR /build
COPY go.* ./
RUN go mod download

COPY ./ ./
RUN go build -o plaid_apply plaid_apply.go 
ENTRYPOINT ["./plaid_apply"]