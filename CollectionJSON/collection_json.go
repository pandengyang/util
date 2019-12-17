package CollectionJSON

import (
	"bytes"
	"github.com/pandengyang/utils/StringUtils"
	"text/template"
)

type Data struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Prompt string `json:"prompt"`
}

type Datas struct {
	Data []Data `json:"data"`
}

func Items(items []interface{}, pTemplateStr *string) (jsonStr string, err error) {
	buf := bytes.NewBufferString("")

	t := template.New("").Funcs(template.FuncMap{"embedded_json": StringUtils.EmbeddedJson})
	t, err = t.Parse(*pTemplateStr)
	if err != nil {
		return
	}

	err = t.ExecuteTemplate(buf, "items", items)
	if err != nil {
		return
	}

	jsonStr = buf.String()

	return
}

/* Item json document */
func Item(item interface{}, pTemplateStr *string) (jsonStr string, err error) {
	buf := bytes.NewBufferString("")

	t := template.New("")
	t, err = t.Parse(*pTemplateStr)
	if err != nil {
		return jsonStr, err
	}

	t.ExecuteTemplate(buf, "item", item)
	jsonStr = buf.String()

	return jsonStr, err
}

/* Queries template document */
func Queries(pTemplateStr *string) (jsonStr string, err error) {
	buf := bytes.NewBufferString("")

	t := template.New("")
	t, err = t.Parse(*pTemplateStr)
	t.ExecuteTemplate(buf, "queries", nil)

	jsonStr = buf.String()

	return
}

/* Create template document */
func Template(pTemplateStr *string) (jsonStr string, err error) {
	buf := bytes.NewBufferString("")

	t := template.New("")
	t, err = t.Parse(*pTemplateStr)
	t.ExecuteTemplate(buf, "template", nil)

	jsonStr = buf.String()

	return
}
