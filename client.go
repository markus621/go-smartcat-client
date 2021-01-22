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
	userAgent = `go-smartcat-client`
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
func (c *Client) Debug(is bool, w io.Writer) {
	c.debug, c.writer = is, w
}

func (c *Client) raw(method, path string, req []byte) (code int, body []byte, err error) {
	code, body, err = c.call(method, path, body, `application/json`)
	return
}

func (c *Client) json(method, path string, req json.Marshaler, resp json.Unmarshaler) (code int, err error) {
	var body []byte
	if req != nil {
		body, err = req.MarshalJSON()
		if err != nil {
			return 0, errors.Wrap(err, "marshal request")
		}
	}
	code, body, err = c.call(method, path, body, `application/json`)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

func (c *Client) form(method, path string, fields *Form, resp json.Unmarshaler) (code int, err error) {
	if fields == nil {
		err = ErrEmptyRequest
		return
	}
	var body []byte
	code, body, err = c.call(method, path, fields.Bytes(), fields.GetContentType())
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

func (c *Client) call(method, path string, body []byte, contentType string) (int, []byte, error) {
	creq, err := http.NewRequest(method, c.conf.URL+path, bytes.NewReader(body))
	c.requestDebug(method, path, body, err)
	if err != nil {
		return 0, nil, errors.Wrap(err, "create request")
	}

	creq.Header.Set("User-Agent", userAgent)
	creq.Header.Set("Connection", "keep-alive")
	creq.Header.Set("Accept", "*/*")
	creq.Header.Set("Content-Type", contentType)
	creq.Header.Set("Authorization", c.conf.AuthToken())

	cresp, err := c.cli.Do(creq)
	if err != nil {
		return 0, nil, errors.Wrap(err, "make request")
	}

	code := cresp.StatusCode
	switch code {
	case 200:
		body, err = c.read(cresp.Body)
	case 204:
	case 404, 415:
		body, err = nil, errors.New(cresp.Status)
	default:
		msg := ErrorResponse{}
		body, err = c.read(cresp.Body)
		if err == nil {
			if err = json.Unmarshal(body, &msg); err != nil {
				err = ErrorResponse{Message: string(body)}
			} else {
				err = msg
			}
		}
	}

	c.responseDebug(code, err, body)

	switch err {
	case nil:
		return code, body, nil
	case io.EOF:
		return code, body, errors.New(cresp.Status)
	default:
		return code, body, err
	}
}

func (c *Client) read(rc io.ReadCloser) (b []byte, err error) {
	defer rc.Close() //nolint: errcheck
	b, err = ioutil.ReadAll(rc)
	return
}

func (c *Client) requestDebug(method, path string, body []byte, err error) {
	if c.debug {
		fmt.Fprintf(c.writer, "REQ: %s:%s err: %v raw:%s \n", method, path, err, body)
	}
}

func (c *Client) responseDebug(code int, err error, body []byte) {
	if c.debug {
		fmt.Fprintf(c.writer, "RES: [%d] err: %v raw:%s \n", code, err, body)
	}
}
