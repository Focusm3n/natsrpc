package codegen

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

var sTpl *template.Template

func init() {
	var err error
	sTpl, err = template.New("st").Parse(serviceTmpl)
	if err != nil {
		panic(err)
	}
}

// Template 模板
func ServiceTemplate() *template.Template {
	return sTpl
}

// GenFile 生成代码文件
func GenFile(t *template.Template, data FileSpec, file string) error {
	src, err := GenText(t, data)
	if nil != err {
		return err
	}

	if b, err := format.Source(src); nil != err {
		if errFile := ioutil.WriteFile(file, src, os.ModePerm); errFile != nil {
			return errFile
		}
		return err
	} else {
		return ioutil.WriteFile(file, b, os.ModePerm)
	}
}

// GenText 模板执行
func GenText(t *template.Template, data interface{}) ([]byte, error) {
	w := &bytes.Buffer{}
	if err := t.Execute(w, data); nil != err {
		return []byte(""), err
	}
	return w.Bytes(), nil
}
