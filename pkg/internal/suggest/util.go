package suggest

import (
	"github.com/agext/levenshtein"
	"go.minekube.com/brigodier"
	"sort"
	"strings"
)

type suggestion struct {
	text  string
	score float64
}

func Build(builder *brigodier.SuggestionsBuilder, input string, candidates []string) *brigodier.Suggestions {
	if input == "" {
		return builder.Build()
	}
	given := input[strings.LastIndex(input, " ")+1:]
	var result []suggestion
	for _, text := range candidates {
		score := Score(given, text)
		if score < 0.2 {
			continue
		}
		result = append(result, suggestion{
			text:  text,
			score: score,
		})
	}
	sortSuggestions(result)
	for _, s := range result {
		builder.Suggest(s.text)
	}
	return builder.Build()
}

func sortSuggestions(s []suggestion) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].score > s[j].score
	})
}

func Score(given, suggestion string) float64 {
	i := len(given)
	if len(suggestion) < i {
		i = len(suggestion)
	}
	return levenshtein.Similarity(given, suggestion[:i], nil)
}
