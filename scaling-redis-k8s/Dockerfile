FROM golang:1.7.3-alpine

COPY . /go/src/github.com/gronnbeck/examples/scaling-redis-k8s

RUN go install -v github.com/gronnbeck/examples/scaling-redis-k8s

CMD scaling-redis-k8s -redis-addr=$REDIS_ADDR -redis-pass=$REDIS_PASS
