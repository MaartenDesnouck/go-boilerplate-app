FROM golang:1.14.2-alpine as build

RUN apk --no-cache add bash alpine-sdk
RUN apk --no-cache add ca-certificates

RUN go get github.com/codegangsta/gin
RUN go get github.com/sirupsen/logrus

WORKDIR /var/src

COPY . .

RUN go build


FROM alpine as release

COPY --from=build /var/src/go-boilerplate-app /go-boilerplate-app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD /go-boilerplate-app
