FROM alpine

RUN apk add curl

RUN echo "*/5 * * * * curl feed:8080" > /etc/crontabs/root

CMD [ "crond", "-f" ]