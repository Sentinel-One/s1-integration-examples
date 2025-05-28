#!/bin/bash
# Error Handling
handle_error() {
    echo -e "*** No Verdict. $1   "
}

# Function to process input
process_input() {
    local input_path="$1"
    local recurse="$2"

    if [ -d "$input_path" ]; then
        # If the input is a directory, loop through each file in the directory
        echo "Input is a directory: $input_path"

        if [ "$recurse" == "true" ]; then
            # If recursion is enabled, loop through the directory and its subdirectories
            find "$input_path" -type f | while read file; do
                measure_time python scanFile.py "$file"
            done
        else
            for file in "$input_path"/*; do
                if [ -f "$file" ]; then
                    measure_time python scanFile.py "$file"
                else
                    echo "Skipping non-file item: $file"
                fi
            done
        fi
    elif [ -f "$input_path" ]; then
        # If the input is a file, process that file
        echo "Input is a file: $input_path"
        measure_time python scanFile.py "$input_path"
    else
        echo "Invalid path: $input_path"
    fi
}

# Function to compute elapsed time for a command
measure_time() {
    # print the name of the file being scanned
    echo "Scanning file:  ${@: -1}"

    # Capture the start time (in nanoseconds)
    local start=$(date '+%Y-%m-%d %H:%M:%S.%N')

    # Execute the command passed as arguments to the function   (and suppress any error messages - we'll handle that in handle_error)
    "$@" 2>/dev/null || handle_error " SKIPPED ${@: -1}"

    # Capture the end time (in nanoseconds)
    echo -e "started at ${start}"
    echo -e "  ended at $(date '+%Y-%m-%d %H:%M:%S.%N')"
    echo -e "-------------------------------------------------"
    echo
}

# Function to demo how this works
demo_mode() {
    echo
    echo -e "=========================================================="
    echo -e "   Scanning the NexusSDK.pdf file using python command"
    echo -e "=========================================================="
    measure_time python scanFile.py NexusSDK.pdf 
    echo -e "\n=========================================================="
    echo -e "   Now using the python wheel wrapper with multiple files \n     [python -m SentinelDFI.scanner -i PATH_TO_FILE]"
    echo -e "==========================================================\n"
    measure_time python -m SentinelDFI.scanner -i /app/NexusSDK.pdf
    measure_time python -m SentinelDFI.scanner -i /app/TestFileSuspicious.txt
    measure_time python -m SentinelDFI.scanner -i /app/TestFileMalicious.txt
    echo
}

#  Main Logic
# Frist install the python wheel module from the SDK  -- volume must be mounted to /nexus-sdk
python -m pip install /nexus-sdk/SDK/SentinelDFI-2-py2.py3-none-any.whl 2>/dev/null

if [ "$FILES" = "false" ]; then
    # nothing specified, run the demo files
    demo_mode
else
    # files specified, run the scanner on those
    #measure_time python scanFile.py $FILES
    process_input "$FILES" "$RECURSE"
fi
