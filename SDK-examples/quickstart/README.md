# SDK Quickstart <!-- omit in toc -->

SDK Quickstarts have been developed to allow you see examples of how to use the SentinelOne Nexus SDK with a variety of languages when embedding it into your technology stack. 

Each Quickstart will produce a malware scanner utilizing the SentinelOne Static AI engine via the Nexus SDK to scan files and/or directories on demand.  Each scanner is embedded within a Docker container and can be utilized on any platform that supports Docker.

## Table of Contents <!-- omit in toc -->

- [System Requirements](#system-requirements)
- [Initial Setup Steps](#initial-setup-steps)
- [Quickstarts by Language](#quickstarts-by-language)
- [Additional References](#additional-references)

## System Requirements

- A 64-bit machine is required - it may be Intel/AMD or Apple Silicon but x86 platforms are not supported by the Quickstarts
- `git` must be installed on your machine
  - [git Client](https://git-scm.com/downloads) - Windows, MacOS, Linux
- Docker must be installed and running on your machine
  - [Docker Engine](https://docs.docker.com/engine/install/) - Linux distributions (free)
  - [Docker Desktop](https://docs.docker.com/desktop/) - Windows, MacOS, Linux (may require licensing)
  - [Rancher Desktop](https://rancherdesktop.io/) - Windows, MacOS, Linux (free)
- Windows systems must have WSL2 installed and available
  - [Installing WSL](https://learn.microsoft.com/en-us/windows/wsl/install)
- You will need a copy of the Nexus SDK distribution ZIP file
  - Contact your SentinelOne Sales Team for a download link

_**NOTES:**_

- _These instructions assume you are using either **Docker Engine** for Linux systems or **Rancher Desktop** with the **dockerd (moby) engine** for MacOS and Windows systems. If you are using **Docker Desktop** you may need to make adjustments._

- _For Windows systems using WSL2, the instructions assume you are using the default image that is included with the latest version of Rancher Desktop. The commands should be universal across images, however, you may need to make adjustments._

## Initial Setup Steps

1. From a location within your home directory (or user profile on Windows), clone this repository from GitHub and change into the SDK-examples folder:
  
   ```sh
   git clone https://github.com/Sentinel-One/s1-integration-examples ./s1-integration-examples
   cd ./s1-integration-examples
   cd ./SDK-examples
   ```

1. Create a folder called `nexus-sdk` and unpack the Nexus SDK distribution ZIP file into it. The resulting folder should look something like this (some files/folders have been omitted for brevity):
   
   ```sh
   └── s1-integration-examples/
    └── SDK-examples/
        └── quickstart/
        └── nexus-sdk/
            ├── Code Samples
            ├── DIGEST
            └── SDK/
                ├── include/
                │   └── libdfi.h
                └── lib/
                    ├── linux/
                    │   ├── arm64/
                    │   │   └── libdfi.so
                    │   └── x64/
                    │       └── libdfi.so
                    ├── windows/
                    │   ├── x64/
                    │   │   └── libdfi64.dll
                    │   └── x86/
                    │       └── libdfi32.dll
                    ├── pyscanner.py
                    └── SentinelDFI-2-py2.py3-none-any.whl
   ```

## Quickstarts by Language

Now that you've made sure your system has met all of the requirements and completed the initial setup steps, you're ready to work on a Quickstart.

To continue, simply choose one of the following languages:

- [Go](./go/README.md)
- [NodeJS](./nodejs/README.md)
- [Python](./python/README.md)

## Additional References

- [Nexus SDK Requirements and Installation](https://community.sentinelone.com/s/article/000005295)
- [Nexus SDK Embedded API User Guide](https://community.sentinelone.com/s/article/000005296)
