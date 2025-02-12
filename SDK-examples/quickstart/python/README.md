# SDK Quickstart for Python <!-- omit in toc -->

Use this directory to see how quickly and easily it is to get started with the SentinelOne Nexus SDK in Python.  You'll build a Docker container and run it.   

By default, it will default to "Demo mode" where the container will scan the included [sample_files](../sample_files/) for malware and return a verdict for each file.

## Table of Contents <!-- omit in toc -->

- [Prerequisite Steps](#prerequisite-steps)
- [Build the Container](#build-the-container)
- [Run the Container](#run-the-container)

## Prerequisite Steps

- Ensure you have met all of the [System Requirements](../README.md#system-requirements)
- Ensure you have completed the [Initial Setup Steps](../README.md#initial-setup-steps)

## Build the Container
The resulting container is "SDK-less" meaning it does **not** contain the Nexus SDK library within it.  The user will need to mount the Nexus SDK library when running it.  
* Build the Docker container
  ```
  docker build -f ./quickstart/python/Dockerfile -t s1scanner-python .
  ```

## Run the Container
As mentioned above, the container does **not** contain the Nexus SDK library within it, so you will need to mount the shared library file as a read-only volume when running it.  The examples below all include the `-v $(pwd)/nexus-sdk:/nexus-sdk ` switch to include the SDK at run time.

Additionally, you can specify flags/settings to either specify a file or directory to scan or mount your local filesystem to scan. For example:
| Flag | Function |
| ---- | -------- |
| -e FILES=<YOUR_FILE> | specify a file or directory to scan |
| -e RECURSE="true" | recurse directory and subdirectories |
| -v $(pwd)/nexus-sdk:/nexus-sdk | mount your copy of the Nexus SDK |
| -v ./:/mnt:ro  | mount your local directory to /mnt in the container |

Here are some example command lines with their usage:
* Demo mode
  ```
  docker run -it --rm  -v $(pwd)/nexus-sdk:/nexus-sdk s1scanner-python
  ```
* Specify a single file to scan
  ```
  docker run -it --rm  -v $(pwd)/nexus-sdk:/nexus-sdk -e FILES="NexusSDK.pdf" s1scanner-python
  ```
* Specify a directory to scan
  ```
  docker run -it --rm  -v $(pwd)/nexus-sdk:/nexus-sdk -e FILES="/app" s1scanner-python
  ```
* Specify a directory to scan recursively
  ```
  docker run -it --rm  -v $(pwd)/nexus-sdk:/nexus-sdk -e FILES="/app" -e RECURSE="true" s1scanner-python
  ```
* Specify a file from the HOST system's current directory to scan
  ```
  docker run --it --rm  -v $(pwd)/nexus-sdk:/nexus-sdk -v ./:/mnt:ro -e FILES="/mnt/quickstart/sample_files" s1scanner-python
  ```

### Sample Output
Here's what to expect when you run in Demo mode:
```
docker run -it --rm  -v $(pwd)/nexus-sdk:/nexus-sdk s1scanner-python

==========================================================
   Scanning the NexusSDK.pdf file using python command
==========================================================
Scanning /app/NexusSDK.pdf
Using SentinelDFI version b'24.1.1.2' (b'7ef1df615d30b87b0b41a26c72b643d54e229bce-Release.arm64')
File hash: b'fda06704bec27d852fc3401e899fcbdd472377a0'
Verdict: benign
Indicators: ``
started at 2025-01-06 18:22:12.867615383
  ended at 2025-01-06 18:22:13.125601509
-------------------------------------------------


==========================================================
   Now using the python wheel wrapper with multiple files
     [python -m SentinelDFI.scanner -i PATH_TO_FILE]
==========================================================

Scanning file:  /app/NexusSDK.pdf
Using SentinelDFI version b'24.1.1.2' (b'7ef1df615d30b87b0b41a26c72b643d54e229bce-Release.arm64')
File hash: b'fda06704bec27d852fc3401e899fcbdd472377a0'
Verdict: benign
Indicators: ``
started at 2025-01-06 18:22:13.126191925
  ended at 2025-01-06 18:22:13.236438717
-------------------------------------------------

Scanning file:  /app/TestFileSuspicious.txt
Using SentinelDFI version b'24.1.1.2' (b'7ef1df615d30b87b0b41a26c72b643d54e229bce-Release.arm64')
File hash: b'7a493e44c7d736a2e4d87bdf04db2ec4aba872d3'
Verdict: suspicious
Indicators: `EICAR-SENTINEL-ANTIVIRUS-SUSP-FILE`
started at 2025-01-06 18:22:13.237138800
  ended at 2025-01-06 18:22:13.284439592
-------------------------------------------------

Scanning file:  /app/TestFileMalicious.txt
Using SentinelDFI version b'24.1.1.2' (b'7ef1df615d30b87b0b41a26c72b643d54e229bce-Release.arm64')
File hash: b'88355b46c777ca86e7045788805c37ebfcc65de2'
Verdict: malware
Indicators: `EICAR-SENTINEL-ANTIVIRUS-TEST-FILE`
started at 2025-01-06 18:22:13.285275175
  ended at 2025-01-06 18:22:13.332424550
-------------------------------------------------
```

More detailed instructions and usage examples can be found on the [SentinelOne Community Site](https://community.sentinelone.com/s/article/000005296).
