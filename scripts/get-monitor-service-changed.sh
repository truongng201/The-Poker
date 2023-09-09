#!/bin/bash

CURRENT_FOLDER_NAME=$1

# Get the current folder name
SERVICES_CHANGED=()

# Check for changes in the folder name
CHANGED=$(git diff --name-only HEAD~ HEAD | grep "$CURRENT_FOLDER_NAME")
if [[ -n "$CHANGED" ]]; then
  # read line separated by space into bash array 
  while read -r line; do
    # remove the folder name from the 
    line=${line/$CURRENT_FOLDER_NAME\//}
    # remove all the files from the path
    line=${line/\/*/}
    SERVICES_CHANGED+=("$line")
  done <<< "$CHANGED"
  echo ${SERVICES_CHANGED[@]}
  exit 0
else
  exit 0
fi