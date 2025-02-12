#!/bin/bash

git config --global --add safe.directory /workspaces/s1-integration-examples

# copy Nexus SDK files for development
if [ -d ./SDK-examples/nexus-sdk ]; then
    # Go Quickstart
    cp ./SDK-examples/nexus-sdk/SDK/include/libdfi.h ./quickstart/go/pkg/scanner/nexus/include/
    cp ./SDK-examples/nexus-sdk/SDK/lib/linux/arm64/libdfi.so ./quickstart/go/pkg/scanner/nexus/lib/linux/arm64/
    cp ./SDK-examples/nexus-sdk/SDK/lib/linux/x64/libdfi.so ./quickstart/go/pkg/scanner/nexus/lib/linux/amd64/
fi
