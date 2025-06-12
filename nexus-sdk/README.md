# Nexus SDK Examples <!-- omit in toc -->

## Table of Contents <!-- omit in toc -->

- [Overview](#overview)
- [System Requirements](#system-requirements)
- [Initial Setup Steps](#initial-setup-steps)
- [Examples by Language](#examples-by-language)
- [Additional References](#additional-references)

## Overview

SDK examples have been developed to allow you visualize through code how to create or update an application and embed the SentinelOne Nexus SDK in it. 

Each example will produce a malware scanning application utilizing the SentinelOne Static AI engine via the Nexus SDK to scan files and/or directories on demand.  Each scanner is embedded within a Docker container and can be utilized on any platform that supports Docker.

## System Requirements

- A 64-bit machine is required - it may be Intel/AMD or Apple Silicon but 32-bit platforms are not supported for these examples
- `git` must be installed on your machine
  - [git Client](https://git-scm.com/downloads) - Windows, MacOS, Linux
- `make` must be installed on your machine
  - [GNU Make](https://www.gnu.org/software/make/) - Windows, MacOS, Linux
- Docker must be installed and running on your machine
  - [Docker Engine](https://docs.docker.com/engine/install/) - Linux distributions (free)
  - [Docker Desktop](https://docs.docker.com/desktop/) - Windows, MacOS, Linux (may require licensing)
  - [Rancher Desktop](https://rancherdesktop.io/) - Windows, MacOS, Linux (free)
- Windows systems must have WSL2 installed and available
  - [Installing WSL](https://learn.microsoft.com/en-us/windows/wsl/install)
- You will need a copy of the Nexus SDK distribution files - Please contact your SentinelOne Sales Team to obtain these files.

_**NOTES:**_

- _These instructions assume you are using either **Docker Engine** for Linux systems or **Rancher Desktop** with the **dockerd (moby) engine** for MacOS and Windows systems. If you are using **Docker Desktop** you may need to make adjustments._

- _For Windows systems using WSL2, the instructions assume you are using the default image that is included with the latest version of Rancher Desktop. The commands should be universal across images, however, you may need to make adjustments._

## Initial Setup Steps

1. From a location within your home directory (or user profile on Windows), clone this repository from GitHub and change into the `nexus-sdk` folder:
  
   ```sh
   git clone https://github.com/sentinel-one/s1-integration-examples ./s1-integration-examples
   cd ./s1-integration-examples/nexus-sdk
   ```

2. Create a folder called `_distfiles_` and unpack the Nexus SDK distribution ZIP file into it. The resulting folder should look something like this (some files/folders have been omitted for brevity):
   
   ```sh
   └── s1-integration-examples/
    └── nexus-sdk/
        └── _distfiles_/
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
        └── go/
        └── nodejs/
        └── python/
        └── sample-files/
   ```

## Examples by Language

Now that you've made sure your system has met all of the requirements and completed the initial setup steps, you're ready to work on an example for your language of choice.

To continue, simply choose one of the following languages:

- [Go](./go/README.md)
- [NodeJS](./nodejs/README.md)
- [Python](./python/README.md)

## Additional References

- [Nexus SDK Requirements and Installation](https://community.sentinelone.com/s/article/000005295)
- [Nexus SDK Embedded API User Guide](https://community.sentinelone.com/s/article/000005296)
