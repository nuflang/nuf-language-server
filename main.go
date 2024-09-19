package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/nuflang/nuf-language-server/analysis"
	"github.com/nuflang/nuf-language-server/lsp"
	"github.com/nuflang/nuf-language-server/rpc"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			continue
		}

		handleMessage(writer, state, method, content)
	}
}

func handleMessage(writer io.Writer, state analysis.State, method string, content []byte) {
	switch method {
	case "initialize":
		handleInitializeMethod(writer, content)
	case "textDocument/didOpen":
		handleTextDocumentDidOpenMethod(state, content)
	case "textDocument/didChange":
		handleTextDocumentDidChangeMethod(state, content)
	case "textDocument/completion":
		handleTextDocumentCompletionMethod(writer, state, content)
	}
}

func handleInitializeMethod(writer io.Writer, content []byte) {
	var request lsp.InitializeRequest
	if err := json.Unmarshal(content, &request); err != nil {
		return
	}

	msg := lsp.NewInitializeResponse(request.ID)
	writeResponse(writer, msg)
}

func handleTextDocumentDidOpenMethod(state analysis.State, content []byte) {
	var request lsp.DidOpenTextDocumentNotification
	if err := json.Unmarshal(content, &request); err != nil {
		return
	}

	state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
}

func handleTextDocumentDidChangeMethod(state analysis.State, content []byte) {
	var request lsp.DidChangeTextDocumentNotification
	if err := json.Unmarshal(content, &request); err != nil {
		return
	}

	for _, change := range request.Params.ContentChanges {
		state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
	}
}

func handleTextDocumentCompletionMethod(writer io.Writer, state analysis.State, content []byte) {
	var request lsp.CompletionRequest
	if err := json.Unmarshal(content, &request); err != nil {
		return
	}

	response := state.Completion(
		request.ID,
		request.Params.TextDocument.URI,
		request.Params.Context,
		request.Params.Position,
	)
	writeResponse(writer, response)
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
