FROM registry.access.redhat.com/ubi8/ubi

COPY simple-kv-rest.go /app/simple-kv-rest.go
RUN dnf install -y go &&\
	mkdir /app/.cache &&\
	chmod 770 /app/.cache

ENV HOME=/app

ENTRYPOINT go run /app/simple-kv-rest.go
