FROM golang:latest

COPY . /go/src/burn-resource
WORKDIR /go/src/burn-resource

COPY go.mod ./
RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o burn-resource main.go
RUN chmod +x burn-resource
CMD ["./burn-resource"]
