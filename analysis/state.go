package analysis

import (
	"strings"

	"github.com/nuflang/nuf-language-server/lsp"
)

type State struct {
	// Map of file names to content
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Completion(
	id int,
	uri string,
	context lsp.CompletionContext,
	position lsp.Position,
) lsp.CompletionResponse {
	items := make([]lsp.CompletionItem, 0)

	if position.Character == 0 {
		items = []lsp.CompletionItem{
			{
				Label:            "section_title",
				Kind:             3,
				InsertTextFormat: 2,
				TextEdit: lsp.TextEdit{
					Range: lsp.Range{
						Start: position,
						End:   position,
					},
					NewText: "section_title(\"$0\")",
				},
			},
			{
				Label:            "section",
				Kind:             3,
				InsertTextFormat: 2,
				TextEdit: lsp.TextEdit{
					Range: lsp.Range{
						Start: position,
						End:   position,
					},
					NewText: "section(\"$0\")",
				},
			},
		}
	}

	for row, line := range strings.Split(s.Documents[uri], "\n") {
		if row == position.Line {
			if position.Character > 0 && line[position.Character-1] == '"' && line[position.Character] == '"' && line[:position.Character-1] == "section(" {
				items = []lsp.CompletionItem{
					{
						Label:            "main",
						InsertTextFormat: 2,
						TextEdit: lsp.TextEdit{
							Range: lsp.Range{
								Start: position,
								End:   position,
							},
							NewText: "main",
						},
						Documentation: lsp.MarkupContent{
							Kind:  "markdown",
							Value: "```html\n<main></main>\n```",
						},
					},
					{
						Label:            "site_navigation",
						InsertTextFormat: 2,
						TextEdit: lsp.TextEdit{
							Range: lsp.Range{
								Start: position,
								End:   position,
							},
							NewText: "site_navigation",
						},
						Documentation: lsp.MarkupContent{
							Kind:  "markdown",
							Value: "```html\n<nav></nav>\n```",
						},
					},
					{
						Label:            "region",
						InsertTextFormat: 2,
						TextEdit: lsp.TextEdit{
							Range: lsp.Range{
								Start: position,
								End:   position,
							},
							NewText: "region",
						},
						Documentation: lsp.MarkupContent{
							Kind:  "markdown",
							Value: "```html\n<section></section>\n```",
						},
					},
				}
			}
		}
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			Message: lsp.Message{RPC: "2.0"},
			ID:      id,
		},
		Result: items,
	}
}
