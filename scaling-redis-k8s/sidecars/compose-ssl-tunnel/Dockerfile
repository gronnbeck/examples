FROM b.gcr.io/uc-prox-docker-registry/go:latest

RUN apk --update --no-cache add bash openssh-client

COPY . $WORKDIR

RUN adduser -S -H app

RUN mkdir -p /home/app/.ssh
RUN cat known_hosts >> /home/app/.ssh/known_hosts

USER app

CMD bash start.sh
