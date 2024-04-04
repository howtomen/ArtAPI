FROM golang:1.22.1

LABEL maintainer="Daniel Ariza"

WORKDIR /app

COPY . . 

RUN go mod download

RUN go build -o api 

EXPOSE 8000

CMD ["./api"]
