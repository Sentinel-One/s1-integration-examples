#!/usr/bin/env python3

import argparse
import os
import sys
import time
from SentinelDFI import scanner

def scan_file(engine, file):
    """
    Scans the actual file and prints the verdict.

    Args:
        engine (object): The instantiated S1 scan engine
        file (str): The path to the file to scan
    """
    try:
        start_time = time.perf_counter()
        verdict = engine.scanFile(file)
        end_time = time.perf_counter()
        print(f"| INFO | Scanned file: {file}")
        print(f"         Verdict: {verdict}")
        print(f"         Scan time: {(end_time - start_time):.6f} second(s)")
    except scanner.ScannerUnknownFileTypeException as e:
        print(f"| NOTICE | Skipped file: {file}")
        print(f"           Reason: The file type is not supported by the scan engine")
    except scanner.ScannerFileTooSmallException as e:
        print(f"| NOTICE | Skipped file: {file}")
        print(f"           Reason: The file is too small to scan")
    except scanner.ScannerFileTooLargeException as e:
        print(f"| NOTICE | Skipped file: {file}")
        print(f"           Reason: The file is too large to scan")
    except Exception as e:
        print(f"| WARN | Skipped file: {file}")
        print(f"         Reason: {e}")

def scan_path(root_volume_mount_path, target_path, recurse):
    """
    Scans the specified path, printing information about files found.

    Args:
        root_volume_mount_path (str): The path inside the container in which the host's root volume is mounted.
        target_path (str): The file or directory path to scan.
        recurse (bool): If True, recursively scans directories.
    """
    if target_path[0] == '/':
        full_path = os.path.join(root_volume_mount_path, target_path[1:])
    else:
        full_path = os.path.join(root_volume_mount_path, target_path)

    # ensure the path exists
    if not os.path.exists(full_path):
        print(f"| ERROR | Path '{target_path}' does not exist.")
        return

    # create the scanner
    try:
        s = scanner.Scanner(False)
    except Exception as e:
        print(f"| ERROR | Unable to create a new scanner: {e}")
        return

    # scan the file
    if os.path.isfile(full_path):
        scan_file(s, full_path)

    # scan the directory
    elif os.path.isdir(full_path):
        if recurse:
            print(f"| INFO | Recursively scanning directory: {target_path}")
            # os.walk generates the file names in a directory tree
            # by walking the tree top-down or bottom-up.
            for root, dirs, files in os.walk(full_path):
                for file_name in files:
                    file_path = os.path.join(root, file_name)
                    scan_file(s, file_path)
        else:
            print(f"| INFO | Scanning directory: {target_path}")
            try:
                # os.listdir returns a list containing the names of the entries
                # in the directory given by path.
                for entry_name in os.listdir(full_path):
                    entry_path = os.path.join(full_path, entry_name)
                    if os.path.isfile(entry_path):
                        scan_file(s, entry_path)
            except PermissionError:
                print(f"| ERROR | Permission denied: Could not access '{target_path}'")
    else:
        print(f"| WARNING | '{target_path}' is neither a file nor a directory (e.g., a symbolic link or special file). Skipping.")

def main():
    """
    Main function to parse arguments and initiate the scanning process.
    """
    # create the parser
    parser = argparse.ArgumentParser(
        description="Scan a file or directory for files.",
        formatter_class=argparse.RawTextHelpFormatter # for better help message formatting
    )

    # add the path argument
    parser.add_argument(
        "path",
        type=str,
        nargs="?",
        help="The path to the file or directory to scan."
    )

    # add the --base-scan-folder flag
    parser.add_argument(
        "--base-scan-folder",
        type=str,
        default="/mnt", # Default value set to /mnt
        help="Override the root volume mount path (defaults to /mnt)."
    )

    # add the --demo flag
    parser.add_argument(
        "--demo",
        action="store_true",
        help="Enable demo mode."
    )

    # add the --recurse flag
    parser.add_argument(
        "--recurse",
        action="store_true", 
        help="Recursively scan the directory if the path is a directory."
    )

    # scan the path
    args = parser.parse_args()
    if args.demo:
        args.base_scan_folder="/"
        args.path="/opt/s1scanner/sample-files"
        args.recurse=False
    elif args.path == None or args.path == "":
        print(f"| ERROR | Please specify a path to scan!")
        return
    scan_path(args.base_scan_folder, args.path, args.recurse)

if __name__ == "__main__":
    main()
