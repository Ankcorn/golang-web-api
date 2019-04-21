FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GIN_MODE release
RUN go build -o main .
CMD ["/app/main"]