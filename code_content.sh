#!/bin/bash

# Output file where the result will be saved
OUTPUT_FILE="golang_app_files.txt"

# Name of this script to exclude it from the output
SCRIPT_NAME=$(basename "$0")

# Remove the previous output file if it exists
if [ -f "$OUTPUT_FILE" ]; then
    rm "$OUTPUT_FILE"
fi

# Loop through all specified files, excluding .pb.go files and the script itself
find . -type f \( -name "*.go" -o -name "*.proto" -o -name "*.mod" -o -name "*.sum" -o -name "*.yml" -o -name "Makefile" \) \
    ! -name "*.pb.go" ! -name "$SCRIPT_NAME" | while read file; do
    # Print the relative file path and its content
    echo "File Path: $file" >> "$OUTPUT_FILE"
    echo "-------------------" >> "$OUTPUT_FILE"
    cat "$file" >> "$OUTPUT_FILE"
    echo -e "\n\n" >> "$OUTPUT_FILE"
done

# Print the directory tree to the output file
echo "Directory Tree:" >> "$OUTPUT_FILE"
echo "-------------------" >> "$OUTPUT_FILE"
tree -L 3 >> "$OUTPUT_FILE"  # Modify the depth (e.g., 3) if you want more or fewer levels of directories in the tree

echo "File listing and content saved to $OUTPUT_FILE."