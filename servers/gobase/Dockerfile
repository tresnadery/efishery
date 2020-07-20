FROM golang:1.13-alpine

ARG APP_NAME
ENV APP_NAME=$APP_NAME

RUN mkdir -p /app/$APP_NAME

# Install curl 
RUN apk --no-cache add bash \
	curl \
	git \
	gcc \
	g++ \
	inotify-tools

WORKDIR /app/$APP_NAME

RUN mkdir -p /go/_tools

ADD config/build.sh /go/_tools
ADD config/reflex.conf /go/_tools

RUN chmod +x /go/_tools/build.sh

# Add file watcher
RUN go get -v github.com/cespare/reflex

CMD ["reflex","-c","/go/_tools/reflex.conf"]