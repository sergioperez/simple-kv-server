FROM registry.access.redhat.com/ubi8/ubi
COPY simple-kv-server.go /app/simple-kv-server.go
RUN dnf install -y go --setopt=install_weak_deps=False && \
	dnf clean all && \
	rm -rf /var/cache/yum && \
	mkdir -p /app/.cache && \
	chmod 770 /app/.cache
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/simple-kv-server /app/simple-kv-server.go


FROM registry.access.redhat.com/ubi8/ubi-minimal
ENV HOME=/app
COPY --from=0 /app/simple-kv-server /app/simple-kv-server

CMD ["/app/simple-kv-server"]
