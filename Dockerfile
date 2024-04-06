FROM golang:1.22.1 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# this is to reduce size of env needed to execute binary that is made above
# alpine image is a lot lighter than golang image used above 
# This creates a smaller container size.
FROM alpine:latest AS production
COPY --from=builder /app .
CMD [ "./app" ]


LABEL maintainer="Daniel Ariza"