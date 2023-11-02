FROM golang:1.19-alpine as builder
RUN apk add --no-cache git
RUN go install go.uber.org/mock/mockgen@latest

FROM golang:1.19-alpine
WORKDIR /bin
RUN export GO111MODULE=on
COPY --from=builder /go/bin/mockgen mockgen