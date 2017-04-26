package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type Client struct {
	client *http.Client
	Base   *url.URL
	AppID  string
	AppKey string
	Common Service
}

type Service struct {
	Client *Client
}

func NewClient(appid, appkey, host string) *Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Trust self-signed certificates
		},
	}
	url, _ := url.Parse(host)
	c := &Client{client: httpClient, Base: url, AppID: appid, AppKey: appkey}
	c.Common.Client = c
	return c
}

func (c *Client) NewRequest(method, urlstr string, params map[string]string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	u.RawQuery += "api_key=" + c.AppKey
	if len(params) > 0 {
		u.RawQuery += parseParams(params)
	}
	u = c.Base.ResolveReference(u)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, obj interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		// Drain up to 512 bytes and close body to let Transport reuse connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
	}()
	if status := resp.StatusCode; status != 200 && status != 201 {
		return fmt.Errorf("Received status code %d", status)
	}
	if obj != nil {
		if w, ok := obj.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(obj)
			if err == io.EOF {
				err = nil // ignore EOF errors
			}
		}
	}
	return err
}

func parseParams(params map[string]string) string {
	var keys, arr []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		arr = append(arr, url.QueryEscape(k)+"="+url.QueryEscape(params[k]))
	}
	return strings.Join(arr, "&")
}
