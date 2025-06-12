# Developer Guide for Go Example <!-- omit in toc -->

This document is intended for anyone who may be actually adding features to or working on the Golang SDK example itself.  It is _not_ targeted at developing your own Go application using the Nexus SDK. [Click here](./app-developer.md) to view that documentation instead.

## Table of Contents <!-- omit in toc -->

- [Prerequisite Steps](#prerequisite-steps)
- [Additional System Requirements](#additional-system-requirements)
- [Visual Studio Code DevContainers](#visual-studio-code-devcontainers)
- [Building the Application](#building-the-application)
- [Built-in Code Documentation](#built-in-code-documentation)
- [Building the SDK-less Containers for GHCR](#building-the-sdk-less-containers-for-ghcr)
- [Additional References](#additional-references)

## Prerequisite Steps

- Ensure you have met all of the [System Requirements](../README.md#system-requirements)
- Ensure you have completed the [Initial Setup Steps](../README.md#initial-setup-steps)

## Additional System Requirements

- Visual Studio Code is the recommended IDE to use and what these instructions are based on
  - [Visual Studio Code](https://code.visualstudio.com/) - Windows, MacOS, Linux (free)
- The following free extensions for Visual Studio Code are recommended:
  - [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) - **Required** for development on MacOS systems - Allows you to develop from within a Linux container within VS Code
  - [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens) - Useful for seeing what changes have been made
  - [Go](https://marketplace.visualstudio.com/items?itemName=golang.go) - Enables Go code highlighting and tools
  - [Markdown All in One](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one) - Tools for working with Markdown-formatted files
- You must install Go locally if you are not using the **Dev Containers** extension within VS Code
  - [Go](https://go.dev/dl/) - Windows, MacOS, Linux (free)

## Visual Studio Code DevContainers

This example code requires the use of an extension in Visual Studio called **DevContainers**, which
basically allows you to run your development environment inside of Visual Studio Code from within a container of your choosing.

When you open the repository from within VS Code and have the extension installed, you can either choose to manually open the repository in a container or VS Code should automatically prompt you to do so.

On Windows and Linux this is not necessarily required, but it is recommended in order to have a consistent development environment across platforms.  MacOS does require this if you want to be able to directly compile and run the Go binary as it will be dynamically linked with the Linux version of the Nexus SDK.

For much more information on **DevContainers**, please refer to the [Visual Studio Code website](https://code.visualstudio.com/docs/devcontainers/containers).

The remainder of this document assumes you are using **DevContainers** for development.

## Building the Application

To compile the application for local testing, simply open a new Terminal within VS Code and run the following command from the `nexus-sdk/go` folder within the repository:

```sh
make s1scanner
```

You can now run the scanner using the command:

```sh
LD_LIBRARY_PATH=./dist ./dist/s1scanner --base-path-folder=/ [flags] <file_or_dir ...>
```

replacing `[flags]` with any actual arguments to `s1scanner` and `<file_or_dir>` with the absolute path to any file(s) or directory(ies) you wish to scan.  

Be sure to use **absolute** paths, otherwise you will get unexpected behavior when running it outside of a container.

## Built-in Code Documentation

All of the code within the Go Quickstart is fully documented according to Go's documentation standards. 

To view the Go documentation for the project, simply run the following command within a Terminal inside VS Code:

```sh
make doc-server
```

You should be prompted to open a browser, but if not, you can open a one and navigate to `http://localhost:9000` to view the documentation.

_**NOTE:** If your machine already has something running on port 9000, choose a different port number by adding `DOCSERVER_PORT=new_port_number` to the command._

## Building the SDK-less Containers for GHCR

In order to build and push the SDK-less version of the container images to GitHub Container Registry, you'll need to do the following:

1. You'll need to create and authorize a GitHub personal access token in order to push container images to GHCR as follows:
   - Log into your GitHub account and navigate to `https://github.com/settings/tokens` which should take you to the **Personal access tokens (classic)** page.
   - Generate a new **classic** access token giving it a name and whatever expiration you want and then checking **write:packages** for the scope.
   - Save the newly generated token somewhere safe.
   
2. Authenticate to `ghcr.io` using your GitHub username and the personal access token you just created:
   
   ```sh
   docker login ghcr.io
   ```

3. Now just run the following command:

   ```sh
   make nosdk-container
   ```

   This will build and push both a 64-bit Intel/AMD and ARM version of the container to the container registry at once.

## Additional References

- [Back to Simple File Scanner in Go](../README.md)
- [Back to Nexus SDK Examples](../../README.md)
  