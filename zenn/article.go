package zenn

import (
	"io"
	"path"
	"strings"
	"text/template"

	"github.com/ikawaha/qiita-zenn/qiita"
)

var ZennTmpl = `---
title: "{{ replace .Title "\"" "\\\"" }}"
emoji: "{{ .Emoji }}"
type: "{{ .Type }}"
topics: [{{ join .Topics "," }}]
published: {{ .Published }}
---
{{ .Body }}
`

type ZennArticle struct {
	Title     string
	Emoji     string
	Type      string
	Topics    []string
	Published bool
	Body      string
	ImageURLs []string
	Slug      string
}

func NewZennArticleFromArticle(a qiita.Article) ZennArticle {
	slug := path.Base(a.URL)
	topics := make([]string, 0, len(a.Tags))
	for _, v := range a.Tags {
		topics = append(topics, v.Name)
	}
	return ZennArticle{
		Title:     a.Title,
		Emoji:     "😀",
		Type:      "tech",
		Topics:    topics,
		Published: false,
		Body:      a.Body,
		ImageURLs: a.ImageURLs,
		Slug:      slug,
	}
}

var tmpl = template.Must(template.New("zenn").Funcs(template.FuncMap{
	"join":    strings.Join,
	"replace": strings.ReplaceAll,
}).Parse(ZennTmpl))

func (a ZennArticle) Write(w io.Writer) error {
	if rs := []rune(a.Title); len(rs) > 60 {
		a.Title = string([]rune(a.Title)[:60])
	}
	return tmpl.Execute(w, a)
}
