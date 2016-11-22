#! /bin/bash

_host=$(echo $REDIS_SSL_TUNNEL_HOST | tr -d '\n')
_port=$(echo $REDIS_SSL_TUNNEL_PORT | tr -d '\n')
_dst_ip=$(echo $REDIS_SSL_TUNNEL_DST_IP  | tr -d '\n')
_src_port=$(echo $REDIS_SSL_TUNNEL_SRC_PORT  | tr -d '\n')

if ! hash ssh 2>/dev/null; then
  echo "You need SSH to run SSL tunnels. Running without SSL"
  exit 1
fi

echo "Starting SSL tunnel 127.0.0.1:$_src_port->$_dst_ip:5000"
echo 'ssh -i /var/cert/private -N \'
echo "    compose@$_host -p $_port -L 127.0.0.1:$_src_port:$_dst_ip:5000"

ssh -i /var/cert/private -N \
    compose@$_host -p $_port -L 127.0.0.1:$_src_port:$_dst_ip:5000
