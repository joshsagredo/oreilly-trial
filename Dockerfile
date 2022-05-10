FROM gcr.io/distroless/static:nonroot

ADD oreilly-trial-amd64 /usr/local/bin/oreilly-trial
USER nonroot
ENTRYPOINT ["oreilly-trial"]
