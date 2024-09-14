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

func (s *State) Completion(id int, uri string, context lsp.CompletionContext) lsp.CompletionResponse {
	items := []lsp.CompletionItem{}

	if context.TriggerKind == 2 {
		if context.TriggerCharacter == ">" {
			items = []lsp.CompletionItem{
				{Label: "banner"},
				{Label: "header"},
				{Label: "complementary"},
				{Label: "aside"},
				{Label: "contentinfo"},
				{Label: "footer"},
				{Label: "form"},
				{Label: "main"},
				{Label: "navigation"},
				{Label: "nav"},
				{Label: "region"},
				{Label: "section"},
				{Label: "search"},
			}
		}
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{Message: lsp.Message{RPC: "2.0"}, ID: id},
		Result:   items,
	}
}
