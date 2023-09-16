//go:build generate
// +build generate

package api

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
)

// The below code runs when we run go generate on this folder and is part of the Go toolchain.
//go:generate echo 📦 Generating API code from OpenAPI spec...
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config server.cfg.yaml ../api/hrapp_openapi.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config types.cfg.yaml ../api/hrapp_openapi.yaml
//go:generate echo ✅ Finished generating API code!
