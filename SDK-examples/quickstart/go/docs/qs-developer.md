# Quickstart Developer Guide <!-- omit in toc -->

This document is intended for anyone who may be actually adding features to or working on the Golang SDK Quickstart itself.  It is _not_ targeted at developing your own Go application using the Nexus SDK. [Click here](./app-developer.md) to view that documentation instead.

## Table of Contents <!-- omit in toc -->

- [System Requirements](#system-requirements)
  - [MacOS Development](#macos-development)
- [Visual Studio Code DevContainers](#visual-studio-code-devcontainers)
- [Building the Application](#building-the-application)
- [Built-in Code Documentation](#built-in-code-documentation)
- [Building the SDK-less Containers for GHCR](#building-the-sdk-less-containers-for-ghcr)
- [Additional References](#additional-references)

## System Requirements

- `git` must be installed on your machine
  - [git Client](https://git-scm.com/downloads) - Windows, MacOS, Linux
- Docker must be installed and running on your machine
  - [Docker Engine](https://docs.docker.com/engine/install/) - Linux distributions (free)
  - [Docker Desktop](https://docs.docker.com/desktop/) - Windows, MacOS, Linux (may require licensing)
  - [Rancher Desktop](https://rancherdesktop.io/) - Windows, MacOS, Linux (free)
- You will need a copy of the Nexus SDK distribution ZIP file
- Visual Studio Code is the recommended IDE to use and what these instructions are based on
  - [Visual Studio Code](https://code.visualstudio.com/) - Windows, MacOS, Linux (free)
- The following free extensions for Visual Studio Code are recommended:
  - [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) - **Required** for development on MacOS systems - Allows you to develop from within a Linux container within VS Code
  - [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens) - Useful for seeing what changes have been made
  - [Go](https://marketplace.visualstudio.com/items?itemName=golang.go) - Enables Go code highlighting and tools
  - [Markdown All in One](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one) - Tools for working with Markdown-formatted files
- You must install Go locally if you are not using the **Dev Containers** extension within VS Code
  - [Go](https://go.dev/dl/) - Windows, MacOS, Linux (free)

### MacOS Development

If you are developing on a MacOS system, you cannot directly run the resulting binary file on your machine.  The Nexus SDK currently only supports either Windows or Linux-based systems.  In order to develop on MacOS, you must use **Visual Studio Code** with the **Dev Containers** extension or an equivalent alternative solution such as a virtual machine.

Your other option is also to build the container after every code change and test your changes that way, but it becomes much harder to troubleshoot as features such as debugging that are a part of VS Code are not available to you.

## Visual Studio Code DevContainers

The **DevContainers** extension basically allows you to run your development environment inside of Visual Studio Code from within a container of your choosing.  This repo has already been configured to use this feature.  When you open the repository from within VS Code and have the extension installed, you can either choose to manually open the repository in a container or VS Code should automatically prompt you to do so.

On Windows and Linux this is not necessarily required, but it is recommended in order to have a consistent development environment across platforms.  MacOS does require this if you want to be able to directly compile and run the Go binary as it will be dynamically linked with the Linux version of the Nexus SDK.

For much more information on **DevContainers**, please refer to the [Visual Studio Code website](https://code.visualstudio.com/docs/devcontainers/containers).

The remainder of this document assumes you are using **DevContainers** for development.

## Building the Application

To compile the application for local testing, simply open a new Terminal within VS Code and run the following command from the `quickstart/go` folder within the repository:

```sh
go build -o dist/s1scanner ./cmd/s1scanner
```

You'll need to copy over the shared object library after the first build as well:

```sh
if [ "$(uname -m)" = "x86_64" ]; then
  cp ./pkg/scanner/nexus/lib/linux/amd64/libdfi.so ./dist/libnexus.unstripped.so
else
  cp ./pkg/scanner/nexus/lib/linux/arm64/libdfi.so ./dist/libnexus.unstripped.so
fi
```

You can now run the scanner, but you'll need to tell it where to find the shared object library by setting `LD_LIBRARY_PATH` in the environment or as part of the application environment:

```sh
LD_LIBRARY_PATH=./dist ./dist/s1scanner -v
```

## Built-in Code Documentation

All of the code within the Go Quickstart is fully documented according to Go's documentation standards. 

To view the Go documentation for the project, simply run the following command within a Terminal inside VS Code:

```sh
pkgsite -http :9000
```

You should be prompted to open a browser, but if not, you can open a one and navigate to `http://localhost:9000` to view the documentation.

_**NOTE:** If your machine already has something running on port 9000, choose a different port number._

## Building the SDK-less Containers for GHCR

In order to build and push the SDK-less version of the container images to GitHub Container Registry, you'll need to do the following:

1. You'll need to create and authorize a GitHub personal access token in order to push container images to GHCR.
   1. First login to your GitHub account and navigate to `https://github.com/settings/tokens` which should take you to the **Personal access tokens (classic)** page.
   2. Generate a new **classic** access token giving it a name and whatever expiration you want and then checking **write:packages** for the scope.
   3. Save the newly generated token somewhere safe.
   4. Click the **Configure SSO** button to authorize the token with the **Sentinel-One** organization.
1. Authenticate to `ghcr.io` using your GitHub username and the personal access token you just created:
   
   ```sh
   docker login ghcr.io
   ```

1. If you have not already set up a multi-platform build container for Docker, you'll need to run the following command to do so:
   
   ```sh
   docker buildx create --name container --driver=docker-container
   ```

1. Make sure you are working from the root directory of the repository and then run the following command:

   ```sh
   docker buildx build --tag ghcr.io/sentinel-one/s1-sdk-examples/s1scanner-go:latest --platform linux/arm64,linux/amd64 --builder container --push -f ./quickstart/go/build/Dockerfile.nosdk .
   ```

   This will build and push both a 64-bit Intel/AMD and ARM version of the container to the container registry at once.

## Additional References

- [Back to the SDK Quickstart for Go](../README.md)
- [Back to All SDK Quickstarts](../../README.md)
  