package httpcli

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var (
	stdCli = NewClient()
)

// 通过HTTP上传文件
// 模拟<form ...><input name="file" type="file" />...</form>
// 参数说明
// url: 上传服务器URL
// form_name: 对应<input>标签中的name
// file_name: 为form表单中的文件名
// path: 为实际要上传的本地文件路径
func PostFile(url, form_name, file_name, path string) (res *http.Response, err error) {
	buf := new(bytes.Buffer) // caveat IMO dont use this for large files, \
	// create a tmpfile and assemble your multipart from there (not tested)
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile(form_name, file_name) //这里的file必须和服务器端的FormFile一致
	if err != nil {
		return
	}
	fd, err := os.Open(path)
	if err != nil {
		return
	}
	defer fd.Close()
	// Write file field from file to upload
	_, err = io.Copy(fw, fd)
	if err != nil {
		return
	}
	// Important if you do not close the multipart writer you will not have a
	// terminating boundry
	w.Close()
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	var client http.Client
	res, err = client.Do(req)
	return
}
