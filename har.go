// Package har provides types that represent HAR spec version 1.2 and it can parse and dump HAR files according to this version of the spec.
// More details can be found at http://www.softwareishard.com/blog/har-12-spec/
package har

import (
	"bytes"
	"encoding/json"
	"io"
	"time"
)

type Log struct {
	Version string   `json:"version"`
	Creator Creator  `json:"creator"`
	Browser *Browser `json:"browser,omitempty"`
	Pages   []Page   `json:"pages,omitempty"`
	Entries []Entry  `json:"entries"`
	Comment string   `json:"comment,omitempty"`
}

type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

type Page struct {
	StartedDateTime *time.Time  `json:"startedDateTime"`
	Id              string      `json:"id"`
	Title           string      `json:"title"`
	PageTimings     PageTimings `json:"pageTimings,omitempty"`
	Comment         string      `json:"comment,omitempty"`
}

type PageTimings struct {
	OnContentLoad time.Duration `json:"onContentLoad,omitempty"`
	OnLoad        time.Duration `json:"onLoad,omitempty"`
	Comment       string        `json:"comment,omitempty"`
}

type Entry struct {
	Pageref         string        `json:"pageref,omitempty"`
	StartedDateTime *time.Time    `json:"startedDateTime"`
	Time            time.Duration `json:"time"`
	Request         Request       `json:"request"`
	Response        Response      `json:"response"`
	Cache           *Cache        `json:"cache"`
	Timings         Timings       `json:"timings"`
	ServerIPAddress string        `json:"serverIPAddress,omitempty"`
	Connection      string        `json:"connection,omitempty"`
	Comment         string        `json:"comment,omitempty"`
}

type Request struct {
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	HTTPVersion string        `json:"httpVersion"`
	Cookies     []Cookie      `json:"cookies"`
	Headers     []Header      `json:"headers"`
	QueryString []QueryString `json:"queryString"`
	PostData    *PostData     `json:"postData,omitempty"`
	HeadersSize int64         `json:"headersSize"`
	BodySize    int64         `json:"bodySize"`
	Comment     string        `json:"comment,omitempty"`
}

type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText"`
	HTTPVersion string   `json:"httpVersion"`
	Cookies     []Cookie `json:"cookies"`
	Headers     []Header `json:"headers"`
	Content     Content  `json:"content"`
	RedirectURL string   `json:"redirectURL"`
	HeadersSize int64    `json:"headersSize"`
	BodySize    int64    `json:"bodySize"`
	Comment     string   `json:"comment,omitempty"`
}

type Cookie struct {
	Name     string     `json:"name"`
	Value    string     `json:"value"`
	Path     string     `json:"path,omitempty"`
	Domain   string     `json:"domain,omitempty"`
	Expires  *time.Time `json:"expires"`
	HTTPOnly bool       `json:"httpOnly"`
	Secure   bool       `json:"secure"`
	Comment  string     `json:"comment,omitempty"`
}

type Header struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

type QueryString struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

type PostData struct {
	MimeType string  `json:"mimeType,omitempty"`
	Params   []Param `json:"params"`
	Text     string  `json:"text"`
	Comment  string  `json:"comment,omitempty"`
}

type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type Content struct {
	Size        int64  `json:"size"`
	Compression int64  `json:"compression"`
	MimeType    string `json:"mimeType,omitempty"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type Cache struct {
	BeforeRequest *CacheDetails `json:"beforeRequest,omitempty"`
	AfterRequest  *CacheDetails `json:"afterRequest,omitempty"`
	Comment       string        `json:"comment,omitempty"`
}

type CacheDetails struct {
	Expires    *time.Time `json:"expires,omitempty"`
	LastAccess *time.Time `json:"lastAccess,omitempty"`
	ETag       string     `json:"eTag"`
	HitCount   uint64     `json:"hitCount"`
	Comment    string     `json:"comment,omitempty"`
}

type Timings struct {
	Blocked *int64 `json:"blocked,omitempty"`
	DNS     *int64 `json:"dns,omitempty"`
	Connect *int64 `json:"connect,omitempty"`
	Send    int64  `json:"send"`
	Wait    int64  `json:"wait"`
	Receive int64  `json:"receive"`
	SSL     *int64 `json:"ssl,omitempty"`
	Comment string `json:"comment,omitempty"`
}

// Dump unmarshalles the provided instance of Log type to JSON and writes to provided object that implements io.Writer.
func Dump(w io.Writer, log *Log) error {
	root := make(map[string]*Log)
	root["log"] = log
	data, err := json.MarshalIndent(&root, "", "  ")
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	_, err = io.Copy(w, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	return nil
}

// NewLog returns an instance of Log type from io.Reader passed in or an error if content cannot be parsed.
func NewLog(r io.Reader) (*Log, error) {
	var root map[string]json.RawMessage
	dec := json.NewDecoder(r)
	err := dec.Decode(&root)
	if err != nil {
		return nil, err
	}
	var log Log
	err = json.Unmarshal(root["log"], &log)
	if err != nil {
		return nil, err
	}
	return &log, nil
}
