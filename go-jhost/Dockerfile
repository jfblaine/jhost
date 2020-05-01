############################
# STEP 1 build executable binary
############################
FROM golang:1.13 AS builder

WORKDIR /go/src/go-jhost
COPY go-jhost.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /go/bin/go-jhost *.go

#########################################
FROM alpine

WORKDIR /go/bin/

RUN apk add --no-cache curl curl-doc

COPY --from=builder /go/bin/go-jhost .

EXPOSE 8080 8443

USER 1001

CMD ["./go-jhost"]

