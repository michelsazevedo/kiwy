FROM golang:alpine

RUN apk update && apk upgrade && apk --no-cache add make

ENV APP kiwy
ENV APP_HOME /usr/local/$APP
WORKDIR $APP_HOME

COPY . $APP_HOME

CMD ["make", "run"]
