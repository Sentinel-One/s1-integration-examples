# Simple File Scanner in Python <!-- omit in toc -->

This folder contains a sample file scanner written in Python that utilizes the SentinelOne Nexus SDK Python library.  

## Table of Contents <!-- omit in toc -->

- [Prerequisite Steps](#prerequisite-steps)
- [Embedded SDK Container vs SDK-less Container](#embedded-sdk-container-vs-sdk-less-container)
- [Running the SDK-less Container](#running-the-sdk-less-container)
  - [Sample Scanner Commands](#sample-scanner-commands)
- [Developing Your Own Application](#developing-your-own-application)
- [Additional References](#additional-references)

## Prerequisite Steps

- Ensure you have met all of the [System Requirements](../README.md#system-requirements)
- Ensure you have completed the [Initial Setup Steps](../README.md#initial-setup-steps)

## Embedded SDK Container vs SDK-less Container

There are two ways utilize the SDK for this example application - the fastest way is to use the pre-built "SDK-less" container and simply mount the Python library from the SDK distribution files as a read-only volume within the container.  If you wish to use this method, continue reading below.

The other choice is to clone this repository and build the container locally on your own machine.  If you choose this method, please read the [Building and Running a Local Container](./docs/local-container.md) documentation.

## Running the SDK-less Container

The "SDK-less" container does **not** contain the Nexus SDK library within it, so you will need to mount the shared library file as a read-only volume when running it.

Making sure your current working directory is the **_root folder of the repository_**, create an alias for the Docker command to simplify typing:

**Intel/AMD-based systems**

```sh
alias s1scanner="docker run --rm -v $(pwd)/nexus-sdk/_distfiles_/SDK:/opt/s1scanner/nexus-sdk/SDK:ro -v /:/mnt:ro --platform linux/amd64 ghcr.io/s1integrations/nexus-sdk/python/s1scanner"
```

**ARM-based systems (eg: Apple Silicon)**

```sh
alias s1scanner="docker run --rm -v $(pwd)/nexus-sdk/_distfiles_/SDK:/opt/s1scanner/nexus-sdk/SDK:ro -v /:/mnt:ro --platform linux/arm64 ghcr.io/s1integrations/nexus-sdk/python/s1scanner"
```

### Sample Scanner Commands

- To print a simple help screen use the `-h` or `--help` option:
  
  ```sh
  s1scanner -h
  ```

- To run a simple demo without scanning local files use the `--demo` option:
  
  ```sh
  s1scanner --demo
  ```
  
- To scan a file or directory: 

  ```sh
  s1scanner FULL_PATH_TO_FILE_OR_DIR
  ```

  where `FULL_PATH_TO_FILE_OR_DIR` is the **absolute path** to the file or directory to scan

  _**NOTE:**_
  
  _Relative file or directory paths will not work as the container has no context as to what your current working directory is._
  
  _Also be aware that running this within the Dev Containers environment will not work since it is already running inside a container.  Be sure to run the command from an OS terminal window instead of through the Terminal window in VS Code._
 
  
  _Examples:_

  ```sh
  s1scanner /Users/joshhogle/s1-integration-examples/nexus-sdk/sample-files/NexusSDK.pdf

  s1scanner "/mnt/c/Users/Josh Hogle/s1-integration-examples/nexus-sdk/sample-files/NexusSDK.pdf"
 
  s1scanner /home/josh/s1-integration-examples/nexus-sdk/sample-files/NexusSDK.pdf
  ```

- To scan a directory recursively, use the `-r` or `--recurse` option:

  ```sh
  s1scanner -r FULL_PATH_TO_DIR
  ```

  _Example:_

  ```sh
  s1scanner -r $(pwd)/../sample-files
  ```

## Developing Your Own Application

For information on using the code within the example to develop your own Go application, please [click here](./docs/app-developer.md).

## Additional References

- [Developing Your Own Application using Python](./docs/app-developer.md)
- [Building and Running a Local Container](./docs/local-container.md)
- [Developer Guide for Python Example](./docs/developer.md) (for making changes to this example itself)
- [Back to Nexus SDK Examples](../README.md)
