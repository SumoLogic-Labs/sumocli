FROM scratch
USER nonroot
ENTRYPOINT ["/sumocli"]
COPY sumocli /
