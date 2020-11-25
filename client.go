package smartcatclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	userAgent = "go-smartcat-client"
)

type (
	Client struct {
		conf Config
		cli  *http.Client

		debug  bool
		writer io.Writer
	}
)

//NewClient init client
func NewClient(c Config) *Client {
	cli := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
		},
		Timeout: 5 * time.Second,
	}
	return NewCustomClient(c, cli)
}

//NewClient init client
func NewCustomClient(c Config, cli *http.Client) *Client {
	return &Client{
		conf: c,
		cli:  cli,
	}
}

func (v *Client) Debug(is bool, w io.Writer) {
	v.debug, v.writer = is, w
}

func (v *Client) call(method, path string, req json.Marshaler, resp json.Unmarshaler) (error, int) {
	var (
		body []byte
		err  error
	)

	switch method {
	case http.MethodGet, http.MethodHead, http.MethodConnect, http.MethodOptions:
		body = nil
	default:
		body, err = req.MarshalJSON()
		if err != nil {
			return errors.Wrap(err, "marshal request"), 0
		}
	}

	creq, err := http.NewRequest(method, v.conf.URL+path, bytes.NewReader(body))
	if err != nil {
		return errors.Wrap(err, "create request"), 0
	}

	creq.Header.Set("User-Agent", userAgent)
	creq.Header.Set("Connection", "keep-alive")
	creq.Header.Set("Accept", "*/*")
	creq.Header.Set("Content-Type", "application/json")
	creq.Header.Set("Authorization", v.conf.AuthToken())

	cresp, err := v.cli.Do(creq)
	if err != nil {
		return errors.Wrap(err, "make request"), 0
	}

	code := cresp.StatusCode
	if code == 200 {
		body, err = v.readBody(cresp.Body, resp)
		v.writeDebug(code, path, body, err)
		switch err {
		case nil:
			return nil, code
		case io.EOF:
			return errors.New(cresp.Status), code
		default:
			return errors.Wrap(err, "unmarshal response"), code
		}
	}
	if code >= 400 && code < 500 {
		msg := ErrorResponse{}
		body, err = v.readBody(cresp.Body, &msg)
		v.writeDebug(code, path, body, err)
		switch err {
		case nil:
			return msg, code
		case io.EOF:
			return errors.New(cresp.Status), code
		default:
			return errors.Wrap(err, "unmarshal error response"), code
		}
	}

	var raw json.RawMessage
	body, err = v.readBody(cresp.Body, &raw)
	v.writeDebug(code, path, body, err)
	return ErrUnknown, code
}

func (v *Client) readBody(rc io.ReadCloser, resp json.Unmarshaler) (b []byte, err error) {
	b, err = ioutil.ReadAll(rc)
	if err != nil {
		return
	}
	err = resp.UnmarshalJSON(b)
	return
}

func (v *Client) writeDebug(code int, url string, body []byte, err error) {
	if v.debug {
		fmt.Fprintf(v.writer, "[%d] %s err: %+v raw:%s \n", code, url, err, body)
	}
}
