FROM golang:1.13-alpine

ENV APP_NAME=auth-service

RUN mkdir /app

# Install curl 
RUN apk --no-cache add bash \
	curl \
	git \
	gcc \
	g++ \
    inotify-tools \
	tzdata

RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime

RUN echo "Asia/Jakarta" > /etc/timezone

WORKDIR /app

COPY . .

CMD ["go","run","main.go"]
