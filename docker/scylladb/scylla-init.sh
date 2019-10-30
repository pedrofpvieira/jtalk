#!/bin/bash
CQL="DROP KEYSPACE jtalk; CREATE KEYSPACE jtalk WITH replication ={'class': 'SimpleStrategy', 'replication_factor' : 1};"

until echo $CQL | cqlsh; do
  echo "cqlsh: Scylla is unavailable to initialize - will retry later"
  sleep 2
done &

exec /docker-entrypoint.py "$@"