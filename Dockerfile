FROM scratch
COPY --from=gcr.io/distroless/base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=gcr.io/distroless/base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=gcr.io/distroless/base /etc/passwd /etc/passwd
COPY --from=gcr.io/distroless/base /etc/group /etc/group
COPY fixerapi /app
USER 1000
CMD ["/app"]