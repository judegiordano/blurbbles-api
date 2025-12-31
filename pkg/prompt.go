package pkg

import (
	"blurbbles/internal/generated"
	"blurbbles/types"

	"github.com/dromara/carbon/v2"
)

var EPOCH *carbon.Carbon

func init() {
	// date of app release
	EPOCH = carbon.Parse("2025-12-30 00:00:00", carbon.UTC)
}

func GetDailyPrompt() types.Prompt {
	now := carbon.Now(carbon.UTC).SetTime(0, 0, 0)
	diff := now.DiffAbsInDays(EPOCH)
	idx := int(diff) % len(generated.STORE.Prompts)
	return generated.STORE.Prompts[idx]
}

func GetPromptById(id string) types.Prompt {
	prompt := generated.STORE.GetById(id)
	return prompt
}
