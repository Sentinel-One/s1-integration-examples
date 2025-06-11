FROM --platform=${BUILDPLATFORM} python:3.12-slim

WORKDIR /opt/s1scanner

COPY ./_distfiles_/SDK/SentinelDFI-2-py2.py3-none-any.whl /opt/s1scanner/nexus-sdk/SDK/
COPY ./python/docker-entrypoint.sh /
COPY ./python/scanner.py /opt/s1scanner/bin/
COPY ./sample-files /opt/s1scanner/sample-files/

RUN set -e && \
    pip3 install /opt/s1scanner/nexus-sdk/SDK/SentinelDFI-2-py2.py3-none-any.whl && \
    chmod +x /docker-entrypoint.sh /opt/s1scanner/bin/scanner.py

ENTRYPOINT ["/docker-entrypoint.sh"]
