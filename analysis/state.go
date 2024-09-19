package analysis

import "github.com/nuflang/nuf-language-server/lsp"

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
	items := []lsp.CompletionItem{
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
			Documentation: lsp.MarkupContent{
				Kind:  "markdown",
				Value: "```html\n<h1></h1>\n```",
			},
		},
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			Message: lsp.Message{RPC: "2.0"},
			ID:      id,
		},
		Result: items,
	}
}
