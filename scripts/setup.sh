#!/usr/bin/env bash
echo setup server ->> redis server
apt-get update
apt-get install redis-server
echo edit bind_ip 0.0.0.0
    supervised systemd
nano /etc/redis/redis.conf
echo run-> systemctl start redis-server