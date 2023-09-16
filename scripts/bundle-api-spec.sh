#!/usr/bin/env bash

echo ⚙️ Starting to bundle the OpenAPI spec, validate it then generate the server boilerplate...
npx --yes @redocly/cli bundle --config redocly.yaml --output api/hrapp_openapi.yaml
npx @redocly/cli lint --config redocly.yaml
npx @redocly/cli build-docs --config redocly.yaml --output backend/hrapp_openapi.html
npx @redocly/cli stats --config redocly.yaml