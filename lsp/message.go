package lsp

/*
See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#message
*/
type Message struct {
	RPC string `json:"jsonrpc"`
}

/*
See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#requestMessage
*/
type Request struct {
	Message
	ID     int    `json:"id"`
	Method string `json:"method"`
}

/*
See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage
*/
type Response struct {
	Message
	ID int `json:"id"`
}

/*
See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#notificationMessage
*/
type Notification struct {
	Message
	Method string `json:"method"`
}
