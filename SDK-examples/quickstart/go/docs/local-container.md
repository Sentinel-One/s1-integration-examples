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

Because the underlying container utilizes the Linux shared object library version of the Nexus SDK, you will need to build a platform-specific container for your OS unless it is capable of running other architectures (eg: Apple Silicon CPU running Rosetta).

To build the container, run the following command:

**Intel/AMD-based systems**

```sh
docker build -f ./quickstart/go/build/Dockerfile.amd64 -t s1scanner-go:latest .
```

**ARM-based systems (eg: Apple Silicon)**

```sh
docker build -f ./quickstart/go/build/Dockerfile.arm64 -t s1scanner-go:latest .
```

## Running the Container

You are now ready to run your local version of the container with the Nexus SDK embedded within it.

Making sure you're in the root of the repository, create an alias for the Docker command to simplify typing:

```sh
alias s1scanner="docker run --rm -v /:/mnt:ro s1scanner-go"
```

### Sample Scanner Commands

- To print a simple help screen use the `-h` or `--help` option:
  
  ```sh
  s1scanner -h
  ```

- To show the version of the Nexus SDK use the `-v` or `--version` option:
  
  ```sh
  s1scanner -v
  ```

- To run a simple demo without scanning local files use the `--demo` option:
  
  ```sh
  s1scanner --demo
  ```
  
- To scan a single file or directory: 

  ```sh
  s1scanner FULL_PATH_TO_FILE_OR_DIR
  ```

  where `FULL_PATH_TO_FILE_OR_DIR` is the **absolute path** to the file or directory to scan

  _**NOTE:** Relative file or directory paths will not work as the container has no context as to what your current working directory is._
  
  _Examples:_

  ```sh
  s1scanner /Users/joshhogle/s1-sdk-examples/sample_files/NexusSDK.pdf

  s1scanner "/mnt/c/Users/Josh Hogle/s1-sdk-examples/sample_files/NexusSDK.pdf"
 
  s1scanner /home/josh/s1-sdk-examples/sample_files/NexusSDK.pdf
  ```

- You can also scan using wildcards (ie: file globs) and even use local environment variables or command interpolation:

  ```sh
  s1scanner -r FULL_PATH_TO_FILE_OR_DIR_WITH_WILDCARD
  ```

  _Example:_

  ```sh
  s1scanner $(pwd)/../sample_files/*.txt
  ```

- To scan a directory recursively, use the `-r` or `--recurse` option:

  ```sh
  s1scanner -r FULL_PATH_TO_DIR
  ```

  _Example:_

  ```sh
  s1scanner -r $(pwd)/../sample_files
  ```

- To output the messages in JSON and use `jq` to prettify the results (note that [`jq`](https://jqlang.github.io/jq/download/) must be installed on your system for this to work):
  
  ```sh
  s1scanner --json -r FULL_PATH_TO_DIR | jq .
  ```

  _Example:_

  ```sh
  s1scanner --json -r $(pwd)/../sample_files | jq .
  ```
  
- Some additional command-line options that may be of interest:
  
  | Option                | Description                                                        |
  | --------------------- | ------------------------------------------------------------------ |
  | `--max-depth`         | maximum depth to of an archive file (eg: `.zip` or `.tar`) to scan |
  | `--max-scan-duration` | maximum amount of time to allow the scanner to scan a file         |

  Use the `-h` option for a full list of options and help.

## Additional References

- [Developing Your Own Application using Go](./app-developer.md)
- [Back to the SDK Quickstart for Go](../README.md)
- [Back to All SDK Quickstarts](../../README.md)
  
