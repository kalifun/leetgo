package lang

import (
	"path/filepath"

	"github.com/j178/leetgo/config"
	"github.com/j178/leetgo/leetcode"
)

var golangGen = golang{
	baseLang{
		Name:              "Go",
		ShortName:         "go",
		Suffix:            ".go",
		LineComment:       "//",
		BlockCommentStart: "/*",
		BlockCommentEnd:   "*/",
	},
}

type golang struct {
	baseLang
}

func (g golang) ShortName() string {
	return g.baseLang.ShortName
}

func (g golang) Name() string {
	return g.baseLang.Name
}

func (g golang) Generate(q leetcode.QuestionData) ([]FileOutput, error) {
	cfg := config.Get()
	return []FileOutput{
		{
			Filename: filepath.Join(cfg.Go.OutDir, q.TitleSlug, "solution.go"),
			Content:  "package main\n",
		},
		{
			Filename: filepath.Join(cfg.Go.OutDir, q.TitleSlug, "solution_test.go"),
			Content:  "package main\n",
		},
	}, nil
}

func (g golang) GenerateTest(leetcode.QuestionData) ([]FileOutput, error) {
	return nil, NotSupported
}

func (golang) GenerateContest(leetcode.Contest) ([]FileOutput, error) {
	return nil, NotSupported
}

func (golang) GenerateContestTest(leetcode.Contest) ([]FileOutput, error) {
	return nil, NotSupported
}
