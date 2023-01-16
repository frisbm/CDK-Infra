#!/bin/sh
# Script to generate diagrams for cdk stacks

set -ef

rm -rf ./.github/images
mkdir -p ./.github/images

STACKS=$(cdk list)
STACKS_ARRAY=($(echo "$STACKS" | tr ' ' '\n'))

for i in "${STACKS_ARRAY[@]}"
do
  echo "Generating diagram for: $i"
  npx cdk-dia --collapse false --stacks "$i" $1> /dev/null
  mv diagram.png "./.github/images/$i.png"
done
rm diagram.dot
