FROM golang:1.15-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build .

FROM alpine:3.9

RUN mkdir -p /app/kubernetes-validation-beeyond/conf
COPY --from=build_base /tmp/app/kubernetes-validation-beeyond /app/kubernetes-validation-beeyond
COPY ./conf/config.yml /app/kubernetes-validation-beeyond/conf/config.yml

EXPOSE 8180
WORKDIR /app/kubernetes-validation-beeyond
CMD ["./kubernetes-validation-beeyond"]
