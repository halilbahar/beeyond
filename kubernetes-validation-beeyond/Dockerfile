FROM golang:1.15-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build .

FROM alpine:3.9

RUN mkdir -p /app/yaml-validation/conf
COPY --from=build_base /tmp/app/yaml-validation /app/yaml-validation
COPY ./conf/config.yml /app/yaml-validation/conf/config.yml

EXPOSE 8180
WORKDIR /app/yaml-validation
CMD ["./yaml-validation"]
