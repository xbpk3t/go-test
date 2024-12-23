package gm

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func Md2HTML(md string) string {
	if md == "" {
		return ""
	}
	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(
			extension.GFM,
			extension.Strikethrough,
			extension.TaskList,
			extension.Linkify,
			extension.Table,
			extension.DefinitionList,
			extension.Footnote,
			extension.Typographer,
			extension.NewTypographer(
				extension.WithTypographicSubstitutions(extension.TypographicSubstitutions{
					extension.LeftSingleQuote:  []byte("&sbquo;"),
					extension.RightSingleQuote: nil, // nil disables a substitution
				}),
			),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := markdown.Convert([]byte(md), &buf); err != nil {
		return ""
	}

	return buf.String()
}
