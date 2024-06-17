FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./
COPY *.html ./

RUN go build -o /server-image

EXPOSE 8080

CMD [ "/server-image" ]