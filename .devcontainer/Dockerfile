FROM golang:1.23-bullseye
RUN apt-get update && apt-get install -y sudo curl vim
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest