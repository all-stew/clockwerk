FROM alpine
ADD clockwerk /clockwerk
COPY config.yaml /
ENTRYPOINT ["/clockwerk"]
