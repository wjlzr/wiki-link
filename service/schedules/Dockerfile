# build stage
FROM golang:alpine as build-env
RUN apk add build-base
MAINTAINER Oldfritter <leon.zcf@gmail.com>

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"

WORKDIR /src

COPY . /src

RUN cd /src && \
  go mod download

RUN  go build -o bin/schedule service/schedules/schedule.go
  

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/schedule /app/bin/schedule
COPY --from=build-env /src/config/config.yml.production /app/config/config.yml

ENV TZ Asia/Shanghai

CMD ["/app/bin/schedule"]
