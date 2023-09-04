#!/bin/bash

set -euo pipefail

i="${1:-10}" # default to 10
while [ "$i" -gt 0 ]; do
   echo $((RANDOM%10))
   i=$((i-1))
done