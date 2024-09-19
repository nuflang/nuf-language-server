package lsp

type DocumentUri = string

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#range
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentItem
type TextDocumentItem struct {
	URI        DocumentUri `json:"uri"`
	LanguageId string      `json:"languageId"`
	Version    int         `json:"version"`
	Text       string      `json:"text"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentIdentifier
type TextDocumentIdentifier struct {
	URI DocumentUri `json:"uri"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#versionedTextDocumentIdentifier
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentPositionParams
type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#position
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// --------------------
// REQUESTS
// --------------------

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionParams
type CompletionParams struct {
	TextDocumentPositionParams
	Context CompletionContext `json:"Context"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionContext
type CompletionContext struct {
	TriggerKind      CompletionTriggerKind `json:"triggerKind"`
	TriggerCharacter string                `json:"triggerCharacter"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionTriggerKind
type CompletionTriggerKind = int

// --------------------
// RESPONSES
// --------------------

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItem
type CompletionItem struct {
	Label            string                     `json:"label"`
	LabelDetails     CompletionItemLabelDetails `json:"labelDetails"`
	Kind             CompletionItemKind         `json:"kind"`
	Tags             []CompletionItemTag        `json:"tags"`
	Detail           string                     `json:"detail"`
	Documentation    any                        `json:"documentation"` // string | MarkupContent
	Preselect        bool                       `json:"preselect"`
	SortText         string                     `json:"sortText"`
	FilterText       string                     `json:"filterText"`
	InsertText       string                     `json:"insertText"`
	InsertTextFormat InsertTextFormat           `json:"insertTextFormat"`
	InsertTextMode   InsertTextMode             `json:"insertTextMode"`
	TextEdit         TextEdit                   `json:"textEdit"`
	// TODO: TextEditText
	// TODO: AdditionalTextEdits
	// TODO: CommitCharacters
	// TODO: Command
	// TODO: Data
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemKind
type CompletionItemKind = int

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemTag
type CompletionItemTag = int

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#markupContentInnerDefinition
type MarkupContent struct {
	Kind  MarkupKind `json:"kind"`
	Value string     `json:"value"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#markupContent
type MarkupKind = string

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItemLabelDetails
type CompletionItemLabelDetails struct {
	Detail      string `json:"detail"`
	Description string `json:"description"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#insertTextFormat
type InsertTextFormat = int

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#insertTextMode
type InsertTextMode = int

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textEdit
type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

// --------------------
// NOTIFICATIONS
// --------------------

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didOpen
type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocumentParams `json:"params"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didOpenTextDocumentParams
type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didChange
type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didChangeTextDocumentParams
type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentContentChangeEvent
type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
