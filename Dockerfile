FROM scratch
EXPOSE 8080

COPY server /
COPY config.gcfg /
COPY assets/ /assets/
ENTRYPOINT ["/server"]
