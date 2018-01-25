package muxuser

import (
	"github.com/gorilla/mux"
	"github.com/sourcegraph/go-langserver/pkg/lsp"
)

func init() {
	var p lsp.InitializeParams
	mux.NewRouter()
}
