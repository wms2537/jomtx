FROM --platform=linux alpine
ARG BUILDARCH

ENV LOCAL=/usr/local

COPY build/jomtxd-linux-${BUILDARCH} ${LOCAL}/bin/jomtxd

ENTRYPOINT [ "jomtxd" ]