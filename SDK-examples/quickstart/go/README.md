# SDK Quickstart for Go <!-- omit in toc -->

This folder contains a sample file scanner written in Go that utilizes the SentinelOne Nexus SDK Linux shared object library.

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

There are two ways utilize the SDK for this Quickstart - the fastest way is to use the pre-built "SDK-less" container and simply mount the shared object library from the SDK distribution files as a read-only volume within the container.  If you wish to use this method, continue reading below.

The other choice is to clone this repository and build the container locally on your own machine.  If you choose this method, please read the [Building and Running a Local Container](./docs/local-container.md) documentation.

## Running the SDK-less Container

The "SDK-less" container does **not** contain the Nexus SDK library within it, so you will need to mount the shared library file as a read-only volume when running it.

You will need to mount the appropriate version of the shared library into your container depending on the architecture of your machine.  If you are running on an Intel/AMD-based system, then you'll need to mount the `SDK/lib/linux/x64/libdfi.so` library.  ARM-base systems such as Apple Silicon Macs must mount the `SDK/lib/linux/arm64/libdfi.so` library.

_**NOTE:** Even though the commands below refer to Linux, it is safe to run them on MacOS and Windows as the container itself is running Linux, so we use the Linux version of the shared library._

Making sure you're in the root of the repository, create an alias for the Docker command to simplify typing:

**Intel/AMD-based systems**

```sh
alias s1scanner="docker run --rm -v $(pwd)/nexus-sdk/SDK/lib/linux/x64/libdfi.so:/lib/libnexus.unstripped.so:ro -v /:/mnt:ro --platform linux/amd64 ghcr.io/sentinel-one/s1-sdk-examples/s1scanner-go"
```

**ARM-based systems (eg: Apple Silicon)**

```sh
alias s1scanner="docker run --rm -v $(pwd)/nexus-sdk/SDK/lib/linux/arm64/libdfi.so:/lib/libnexus.unstripped.so:ro -v /:/mnt:ro --platform linux/arm64 ghcr.io/sentinel-one/s1-sdk-examples/s1scanner-go"
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

## Developing Your Own Application

For information on using the code within the quickstart to develop your own Go application, please [click here](./docs/app-developer.md).

## Additional References

- [Developing Your Own Application using Go](./docs/app-developer.md)
- [Building and Running a Local Container](./docs/local-container.md)
- [Quickstart Developers Guide](./docs/qs-developer.md) (for making changes to this quickstart itself)
- [Back to All SDK Quickstarts](../README.md)
