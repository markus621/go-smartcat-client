package smartcatclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//==================================================================================

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

//==================================================================================

type (
	//Form fields map for upload form
	Form struct {
		data     *bytes.Buffer
		boundary string
		fileID   int
	}
)

//NewForm init new form fields
func NewForm() *Form {
	return &Form{
		data:     &bytes.Buffer{},
		boundary: randomString(23),
		fileID:   0,
	}
}

//GetContentType getting content type for form
func (f *Form) GetContentType() string {
	return fmt.Sprintf(`multipart/form-data; boundary="%s"`, f.boundary)
}

func (f *Form) writeHeader(name string, length int, isfile bool) {
	f.data.WriteString("--" + f.boundary + "\r\n")
	if isfile {
		f.fileID++
		f.data.WriteString("Content-Type: application/octet-stream\r\n")
		f.data.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"file_%d\"; filename=\"%s\"\r\n", f.fileID, escapeQuotes(name)))
	} else {
		f.data.WriteString("Content-Type: application/json\r\n")
		f.data.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"\r\n", escapeQuotes(name)))
	}
	f.data.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n", length))
}

//Add field to form
func (f *Form) Add(name string, value []byte) {
	f.writeHeader(name, len(value), false)
	f.data.Write(value)
	f.data.WriteString("\r\n")
}

//AddJSON convert object to json and add as field
func (f *Form) AddJSON(name string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	f.Add(name, b)
	return nil
}

//AddFile add text as file
func (f *Form) AddFile(name string, value []byte) {
	f.writeHeader(name, len(value), true)
	f.data.Write(value)
	f.data.WriteString("\r\n")
}

//LoadFile read file to form
func (f *Form) LoadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	value, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	f.AddFile(info.Name(), value)
	return nil
}

//Bytes close form, read, and reset
func (f *Form) Bytes() (r []byte) {
	f.data.WriteString(fmt.Sprintf("--%s--\r\n", f.boundary))
	r, f.fileID = f.data.Bytes(), 0
	f.data.Reset()
	return
}
