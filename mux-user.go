package muxuser

import (
	"github.com/sourcegraph/go-langserver.git/pkg/lsp"
	"github.com/gorilla/mux.git"
)

func init() {
	var p lsp.InitializeParams
	mux.NewRouter()
}
