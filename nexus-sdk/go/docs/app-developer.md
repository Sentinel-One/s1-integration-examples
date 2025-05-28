# Developing Your Own Application using Go <!-- omit in toc -->

The code in this repository can be used as a starting point for developing your own Golang application using the SentinelOne Nexus SDK.  The code within the `pkg/scanner` folder can be used directly in your own application with minimal adjustment.  

## Table of Contents <!-- omit in toc -->

- [Prerequisite Steps](#prerequisite-steps)
- [System Requirements](#system-requirements)
- [Application Setup](#application-setup)
  - [Additional Steps for Windows](#additional-steps-for-windows)
  - [Additional Steps for MacOS](#additional-steps-for-macos)
- [Using the `scanner` Package](#using-the-scanner-package)
- [Compiling Your Application](#compiling-your-application)
- [Distributing Your Application](#distributing-your-application)
  - [Windows](#windows)
  - [Linux](#linux)
  - [MacOS](#macos)
- [Additional References](#additional-references)

## Prerequisite Steps

- Ensure you have met all of the [System Requirements](../README.md#system-requirements)
- Ensure you have completed the [Initial Setup Steps](../README.md#initial-setup-steps)

## System Requirements

- Go must be installed on your machine
  - [go Download](https://go.dev/dl/) - Windows, MacOS, Linux
- You will need to install a supported `C` language compiler for your platform
  - **Linux** - You will need to install the `gcc` package for your distribution
  - **MacOS** - You cannot compile the package directly from MacOS as the Nexus SDK is not supported for MacOS systems. You can use Visual Studio Code's **DevContainers** extension for development, though. ([see below](#additional-steps-for-macos))
  - **Windows** - MSYS2 and MINGW64 must be installed on your machine
    - [MSYS2 Distribution](https://www.msys2.org/) - Windows
    - MINGW64 can be installed from an MSYS2 terminal by running the `pacman -S --needed base-devel mingw-w64-x86_64-toolchain` and installing all packages

## Application Setup

1. Create and configure your new Go application as you would any Go application.
2. Clone this repository from GitHub to a local folder on your machine:
  
   ```sh
   git clone https://github.com/sentinel-one/s1-integration-examples ./s1-integration-examples
   cd ./s1-integration-examples
   ```

3. Copy the contents of the `nexus-sdk/go/pkg/scanner` folder to your application's working folder. For example:

   ```sh
   cp -R ./nexus-sdk/go/pkg/scanner/* $HOME/workspace/src/github.com/mycompany/myapp/pkg/scanner/
   ```

4. Now unpack the Nexus SDK distribution files as follows into the `nexus` folder within the destination folder you just copied the files to:

| SDK Distribution File               | `pkg/scanner/nexus` Destination Path |
| ----------------------------------- | ------------------------------------ |
| `/SDK/include/libdfi.h`             | => `include/libdfi.h`                |
| `/SDK/lib/linux/arm64/libdfi.so`    | => `lib/linux/arm64/libdfi.so`       |
| `/SDK/lib/linux/x64/libdfi.so`      | => `lib/linux/amd64/libdfi.so`       |
| `/SDK/lib/windows/x86/libdfi32.dll` | => `lib/windows/x86/dfi.dll`         |
| `/SDK/lib/windows/x64/libdfi64.dll` | => `lib/windows/amd64/dfi.dll`       |

You only need to include the shared object libraries for the operating systems and architectures you plan to build and distribute your application for. 

### Additional Steps for Windows

Because the Nexus SDK does not include a library to perform linking with CGO on Windows, you will need to perform a couple of extra steps to generate the library.

1. Change into the `nexus/lib/windows/amd64` folder and create a new file called `dfi.def` and add the following contents:
   
   ```text
   LIBRARY dfi.dll

   EXPORTS
   init
   cleanup
   scan_file
   scan_file_with_depth
   scan_file_with_timeout
   get_version
   find_file_type
   ```

2. Now run the following command to generate the linking library for CGO:

   ```cmd
   dlltool --dllname dfi.dll --output-lib libdfi.a --def dfi.def
   ```

3. Repeat the process for the `nexus/lib/windows/x86` folder if you plan to support 32-bit versions of Windows.

### Additional Steps for MacOS

Because the Nexus SDK is not supported for MacOS systems directly, you must develop from within a Linux container or VM.  The most efficient way to develop on a MacOS system is to utilize [Visual Studio Code](https://code.visualstudio.com/) with the [DevContainers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) in order to develop using a Linux container from with VS Code.

For more information about the **DevContainers Extension**, please [see the documentation](https://code.visualstudio.com/docs/devcontainers/containers).

If you wish to use this repository as a starting point, you can refer to the contents of the `.devcontainer` folder at the root of the repository.

## Using the `scanner` Package

The `scanner` package you copied into your application is essentially a self-contained package that utilizes the Nexus SDK to scan individual files.

`Engine` is the base object type to use for implementing scanning within your own code.  Its `ScanBytes` function allows you to scan a raw `bytes.Buffer` for cases where you may not be reading directly from a file.  The `ScanFile` function is a wrapper that reads the file into a buffer and then calls the `ScanBytes` function.

The code inside the `scanner` package is fully documented and can be viewed using a standard tool such as [`pkgsite`](https://github.com/golang/pkgsite).

## Compiling Your Application

Simply build your Go application as you would any other go application.  If you are using a single machine to perform cross-platform development, you may want to use a tool such as [`goreleaser`](https://goreleaser.com/).

Please refer to the numerous articles on the web regarding cross-compiling Go applications as that is out of scope for this document.

## Distributing Your Application

### Windows

You will need to distribute the `dfi.dll` for the appropriate architecture with your application binary.  The DLL should either be in the same folder as your application or a standard location in which Windows can locate shared library files.

### Linux

You will need to distribute the `libdfi.so` for the appropriate architecture with your application binary.  The shared object library must be renamed to `libnexus.unstripped.so` on the destination machine and be placed into a folder where `ld` can find it such as `/usr/lib` or somewhere in the `LD_LIBRARY_PATH` configured for the system.

### MacOS

The Nexus SDK libraries do not support MacOS, so you will not be able to run a native application on MacOS using the Nexus SDK.  Instead, use an alternative solution such as running the application from within a Docker container running Linux.

## Additional References

- [Back to Simple File Scanner in Go](../README.md)
- [Back to Nexus SDK Examples](../../README.md)
  