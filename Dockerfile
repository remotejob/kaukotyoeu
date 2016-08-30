FROM scratch
EXPOSE 8080

COPY godocker /
COPY config.gcfg /
COPY assets/ /assets/
ENTRYPOINT ["/godocker"]
