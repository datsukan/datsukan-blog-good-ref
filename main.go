package main

import (
	"github.com/datsukan/datsukan-blog-good-core/usecase"
	"github.com/datsukan/datsukan-blog-good-lambda-frame/frame"
)

func main() {
	frame.Exec(ref)
}

func ref(articleID string) (int, error) {
	return usecase.Ref(articleID)
}
