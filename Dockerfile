FROM registry.access.redhat.com/ubi8/ubi

RUN dnf install -y go
COPY simple-kv-rest.go /app/simple-kv-rest.go

ENTRYPOINT go run /app/simple-kv-rest.go
