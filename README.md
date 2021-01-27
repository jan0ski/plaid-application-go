# plaid-application-go
A Golang client for applying to Plaid via their API

## Getting started
* Go binary
```bash
$ go get github.com/jan0ski/plaid-application-go
$ $GOPATH/bin/plaid-application-go -h
```

* Build from source
```bash
$ git clone https://github.com/jan0ski/plaid-application-go.git
$ cd plaid-application-go/
$ go build -o plaid_apply plaid_apply.go
$ ./plaid_apply -h
```

* Docker image
```bash
docker run jan0skii/jan0skii/plaid_apply:latest -h
```
or build locally
```bash
$ git clone https://github.com/jan0ski/plaid-application-go.git
$ cd plaid-application-go/
$ docker build -t plaid_apply:latest .
$ docker run plaid_apply:latest -h
```

## Example
```bash
./plaid_apply \
--jobID "fd0c5629-a495-443c-8b98-1b1a27ce60af" \
--name "Your Name" \
--email "youremail@domain.com" \
--location "Remote" \
--phone "555-867-5309" \
--resume "https://resume-bucket.s3.amazonaws.com/YourResume.pdf" \
--github "https://github.com/jan0ski/" \
--website "https://your-website-here.com"
--candy "Three Musketeers" \
--superpower "Super speed"
"We got your application and we'll get back to you shortly!"
```
