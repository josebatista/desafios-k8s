FROM golang:1.14.6-alpine3.12 as builder

COPY ./src/greeting_web_server .
ENV CGO_ENABLED 0
RUN GOOS=linux go build main.go

EXPOSE 8000

CMD [ "./main" ]