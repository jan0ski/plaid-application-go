FROM golang:1.15-alpine

WORKDIR /build
COPY go.* ./
RUN go mod download

COPY plaid_apply.go plaid_apply.go
RUN go build -o /bin/plaid_apply plaid_apply.go 
ENTRYPOINT [ "/bin/plaid_apply" ]