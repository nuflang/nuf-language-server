package lsp

// --------------------
// REQUESTS
// --------------------

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeParams
type InitializeRequestParams struct {
	ProcessId  int        `json:"processId"`
	ClientInfo ClientInfo `json:"clientInfo"`
	Locale     string     `json:"locale"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// --------------------
// RESPONSES
// --------------------

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeResult
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
type ServerCapabilities struct {
	TextDocumentSync   int            `json:"textDocumentSync"`
	CompletionProvider map[string]any `json:"completionProvider"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeResult
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			Message: Message{RPC: "2.0"},
			ID:      id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1,
				CompletionProvider: map[string]any{},
			},
			ServerInfo: ServerInfo{
				Name:    "nuf-lsp",
				Version: "0.0.1",
			},
		},
	}
}
