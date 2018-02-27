FROM golang:1.8
COPY . .
RUN go get -d -v ./...
RUN apt-get -y update && apt-get install -y libxrender1 libxext6 libfontconfig
RUN go build
