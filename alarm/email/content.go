package email

import (
	"bytes"
	"html/template"
)

var (
	emailTemplate = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
    </head>
    <body style="overflow: hidden;">

<div style="font-family:'Microsoft YaHei';color: #919191;font-size: 13px;line-height:20px">
请勿回复此邮件
</div><br><br>
<div style="font-family:'Microsoft YaHei';font-size: 14px;line-height:20px;text-indent: 2em;">
{{ .Content }}
</div><br><br>
<div style="font-family:'Microsoft YaHei';color: #919191;font-size: 13px;line-height:20px">
<p>
Best Regards<br><br>
Ckeyer/ Won/
<strong>王传健</strong><br><br>
Mobile: +86 18002276350<br>
E-mail: me@ckeyer.com<br>
Website: http://www.ckeyer.com/
</p>
</div>
</body></html>
`
)

func SetContent(content string) []byte {
	rw := bytes.NewBuffer(nil)
	t, _ := template.New("emailTmp").Parse(emailTemplate)
	// template.New("sd").ParseFiles(...)
	data := make(map[string]interface{})
	data["Content"] = content
	t.Execute(rw, data)
	return rw.Bytes()
}
