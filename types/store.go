package types

type PromptsStore struct {
	Prompts []Prompt
	Index   map[string]int // UUID → index in slice
}

func (s *PromptsStore) GetById(id string) Prompt {
	idx, ok := s.Index[id]
	if !ok {
		return Prompt{}
	}
	return s.Prompts[idx]
}
