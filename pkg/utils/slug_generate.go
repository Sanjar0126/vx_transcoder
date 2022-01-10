package utils

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/teris-io/shortid"
)

func SlugGenerate(input string) string {
	shortId, err := shortid.Generate()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s-%s", slug.Make(FirstN(input)), shortId)
}
