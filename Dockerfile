FROM golang:1.8
COPY . .
RUN go get -d -v ./...
RUN apt-get -y update && apt-get install -y libxrender1 libxext6 libfontconfig

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main .
EXPOSE 5000
CMD ["/app/main"]
