#!/bin/sh

PYTHON_VERSION=3.12
NEXUS_SDK_WHL=SentinelDFI-2-py2.py3-none-any.whl

# SentinelDFI package is not installed - attempt to install it
if [ ! -d "/usr/local/lib/python${PYTHON_VERSION}/site-packages/SentinelDFI" ]; then
    if [ ! -f "/opt/s1scanner/nexus-sdk/SDK/${NEXUS_SDK_WHL}" ]; then
        echo ""
        echo "ERROR: The SentinelOne Nexus SDK was not found in this container image."
        echo "       Please make sure you have mounted the Nexus SDK distribution files at: /opt/s1scanner/nexus-sdk"
        echo ""
        exit 100
    fi
    pip3 install /opt/s1scanner/nexus-sdk/SDK/${NEXUS_SDK_WHL} >/dev/null 2>&1
fi

# run the scanner
exec /opt/s1scanner/bin/scanner.py "$@"
