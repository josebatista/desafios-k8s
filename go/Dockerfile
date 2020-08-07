FROM golang:1.14.6-alpine3.12 as builder

COPY ./src/greeting_web_server .
RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM scratch

COPY --from=builder /go/main .

EXPOSE 8000

CMD [ "./main" ]