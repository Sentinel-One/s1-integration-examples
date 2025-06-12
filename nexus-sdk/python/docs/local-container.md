# Building and Running a Local Container <!-- omit in toc -->

This document outlines the process of building the container locally on your own machine and running it.  The resulting container will contain the Nexus SDK library embedded within it.

## Table of Contents <!-- omit in toc -->

- [Prerequisite Steps](#prerequisite-steps)
- [Building the Container](#building-the-container)
- [Running the Container](#running-the-container)
  - [Sample Scanner Commands](#sample-scanner-commands)
- [Additional References](#additional-references)

## Prerequisite Steps

- Ensure you have met all of the [System Requirements](../../README.md#system-requirements)
- Ensure you have completed the [Initial Setup Steps](../../README.md#initial-setup-steps)

## Building the Container

To build the container, run the following command:

```sh
make container
```

## Running the Container

You are now ready to run your local version of the container with the Nexus SDK embedded within it.

Create an alias for the Docker command to simplify typing:

```sh
alias s1scanner="docker run --rm -v /:/mnt:ro python/s1scanner"
```

**NOTE:**
_Do not try to run the commands below within a VS Code Terminal if you are using DevContainers and have activated the container environment.  The reason being is that the mounts will not match up properly, so you'll want to alias the `s1scanner` command and run the sample commands below from a terminal outside of VS Code._

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

## Additional References

- [Developing Your Own Application using Python](./docs/app-developer.md)
- [Developer Guide for Python Example](./docs/developer.md) (for making changes to this example itself)
- [Back to Nexus SDK Examples](../README.md)
