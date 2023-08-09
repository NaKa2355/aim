FROM golang:1.20.7-alpine3.18

RUN apk update && \
    apk add bash make && \
    apk add --upgrade grep

WORKDIR ./aim
COPY ./ ./
RUN make build
RUN make install

WORKDIR ../
RUN rm -rf ./aim
RUN go clean --modcache
RUN mkdir /var/run/aimd
RUN chmod 755 /var/run/aimd

CMD ["aim"]