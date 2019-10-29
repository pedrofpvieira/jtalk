FROM scylladb/scylla:3.0.10

COPY ./docker/scylladb/scylla-init.sh /scylla-init.sh

ENTRYPOINT ["/scylla-init.sh"]