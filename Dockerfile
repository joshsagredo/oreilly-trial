FROM distroless:static-nonroot-amd64

ADD oreilly-trial-amd64 /usr/local/bin/oreilly-trial
USER nonroot
ENTRYPOINT ["oreilly-trial"]
