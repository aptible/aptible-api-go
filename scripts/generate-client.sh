#!/bin/sh
# There isn't a way to automatically remove outdated files at this time
# https://github.com/OpenAPITools/openapi-generator/issues/15669
# Diffing the tracked FILES with comm allows us to work around this

set -e

if [ "$#" -lt "2" ]; then
  echo "Usage: $(basename "$0") DOC_ENV SPEC_PATH"
  exit 1
fi

PACKAGE_NAME=aptibleapi
TMP_FILE="$(mktemp)"
trap 'rm -f "$TMP_FILE"' EXIT

doc_url="https://documentation-${1}.s3.amazonaws.com/openapi/v1/${2}"
echo "Generating client from ${doc_url}"

# Save the current openapi-generator managed files
cp "${PACKAGE_NAME}/.openapi-generator/FILES" "$TMP_FILE" &> /dev/null || true

# Generate the client
docker run --rm -it -v "$(pwd):/local" "openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_TAG}" \
  generate -g go -i "$doc_url" \
  -o "/local/${PACKAGE_NAME}" --template-dir /local/templates --name-mappings _type=MetaType \
  --git-user-id aptible --git-repo-id "aptible-api-go/${PACKAGE_NAME}" \
  --additional-properties "packageName=${PACKAGE_NAME},disallowAdditionalPropertiesIfNotPresent=false"

# Remove files that are no longer being managed by the generator
# -2 -3 removes the comm lines that are unique to the second file and shared by
# both files, leaving only the lines unique to the first file
cd "$PACKAGE_NAME"
outdated_files="$(comm -2 -3 "$TMP_FILE" .openapi-generator/FILES)"
if [ -n "$outdated_files" ]; then
  echo "Removing outdated files:"
  echo "${outdated_files}"
  echo "${outdated_files}" | xargs -n 10 rm
fi
