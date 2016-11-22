#! /bin/bash

_shouldExit=false

if [ -z "$REDIS_MASTER_URL" ]; then
  echo "REDIS_MASTER_URL is missing."
  _shouldExit=true
fi

if [ -z "$REDIS_MASTER_PORT" ]; then
  echo "REDIS_MASTER_PORT is missing."
  _shouldExit=true
fi

if [ -z "$REDIS_MASTER_PASS" ]; then
  echo "REDIS_MASTER_PASS is missing."
  _shouldExit=true
fi

if $_shouldExit; then
  echo ""
  echo "One or more variables is missing. Exiting."
  exit 1
fi

_masterUrl=$(echo $REDIS_MASTER_URL | tr -d '\n')
_masterPort=$(echo $REDIS_MASTER_PORT | tr -d '\n')
_masterPass=$(echo $REDIS_MASTER_PASS | tr -d '\n')
_port=$(echo $REDIS_PORT | tr -d '\n')

echo "master:" $REDIS_MASTER_URL":"$REDIS_PORT
echo "port": $REDIS_PORT

cat redis.conf.tmpl | \
sed -e s/'$REDIS_MASTER_URL'/$_masterUrl/g | \
sed -e s/'$REDIS_MASTER_PORT'/$_masterPort/g | \
sed -e s/'$REDIS_MASTER_PASS'/$_masterPass/g | \
sed -e s/'$REDIS_PORT'/${_port:-6379}/g \
> redis.conf

redis-server redis.conf
