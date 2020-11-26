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
	//Client client connecting to the server
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

//NewCustomClient init client
func NewCustomClient(c Config, cli *http.Client) *Client {
	return &Client{
		conf: c,
		cli:  cli,
	}
}

//Debug enable logging of responses
func (v *Client) Debug(is bool, w io.Writer) {
	v.debug, v.writer = is, w
}

func (v *Client) call(method, path string, req json.Marshaler, resp json.Unmarshaler) (int, error) {
	var (
		body []byte
		err  error
	)

	if req != nil {
		body, err = req.MarshalJSON()
		if err != nil {
			return 0, errors.Wrap(err, "marshal request")
		}
	}

	creq, err := http.NewRequest(method, v.conf.URL+path, bytes.NewReader(body))
	if err != nil {
		return 0, errors.Wrap(err, "create request")
	}

	creq.Header.Set("User-Agent", userAgent)
	creq.Header.Set("Connection", "keep-alive")
	creq.Header.Set("Accept", "*/*")
	creq.Header.Set("Content-Type", "application/json")
	creq.Header.Set("Authorization", v.conf.AuthToken())

	cresp, err := v.cli.Do(creq)
	if err != nil {
		return 0, errors.Wrap(err, "make request")
	}

	code := cresp.StatusCode
	switch code {
	case 200:
		body, err = v.readBody(cresp.Body, resp)
	case 204:
	case 404:
		body, err = nil, errors.New(cresp.Status)
	case 400, 401, 403:
		msg := ErrorResponse{}
		body, err = v.readBody(cresp.Body, &msg)
		if err != nil {
			err = ErrorResponse{Message: string(body)}
		}
	default:
		var raw json.RawMessage
		body, err = v.readBody(cresp.Body, &raw)
		if err == nil {
			err = ErrUnknown
		}
	}

	v.writeDebug(code, method, path, body, err)

	switch err {
	case nil:
		return code, nil
	case io.EOF:
		return code, errors.New(cresp.Status)
	default:
		return code, errors.Wrap(err, "unmarshal response")
	}
}

func (v *Client) readBody(rc io.ReadCloser, resp json.Unmarshaler) (b []byte, err error) {
	b, err = ioutil.ReadAll(rc)
	if err != nil || resp == nil {
		return
	}
	err = resp.UnmarshalJSON(b)
	return
}

func (v *Client) writeDebug(code int, method, path string, body []byte, err error) {
	if v.debug {
		fmt.Fprintf(v.writer, "[%d] %s:%s err: %+v raw:%s \n", code, method, path, err, body)
	}
}
