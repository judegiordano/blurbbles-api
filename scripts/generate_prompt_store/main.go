package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"

	"blurble.com/types"
)

func main() {
	var buf bytes.Buffer

	bytes, err := os.ReadFile("./.prompts/PROMPTS.json")
	if err != nil {
		panic(err)
	}

	var prompts []types.Prompt
	if err := json.Unmarshal(bytes, &prompts); err != nil {
		panic(err)
	}

	idx := make(map[string]int)
	for i, v := range prompts {
		idx[v.Id] = i
	}

	buf.WriteString(`
		package generated

		import "blurble.com/types"

		var STORE = types.PromptsStore{
			Prompts: []types.Prompt{
	`)

	for _, p := range prompts {
		buf.WriteString(fmt.Sprintf(`{
				Id:     %q,
				Prompt: %q,
				Genre:  %q,
			},
		`, p.Id, p.Prompt, p.Genre))
	}

	buf.WriteString(`
		},
		Index: map[string]int{
	`)

	for id, index := range idx {
		buf.WriteString(fmt.Sprintf(`			%q: %d,
`, id, index))
	}

	buf.WriteString(`		},
	}`)

	src, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("internal/generated/prompts.go", src, 0644); err != nil {
		panic(err)
	}
}
