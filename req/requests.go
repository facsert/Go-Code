package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
    "encoding/json"
	"net/http"
)


type ReqOpt struct {
    Timeout time.Duration
    Headers map[string]string
    Cookie map[string]string

    Param map[string]string
    Body any
    FormData map[string]string
    File string
}

type Resp[T any] struct {
	Result  bool   `json:"result"`
	Msg     string `json:"msg"`
	Content T      `json:"content"`
}

func HttpGet(url string, opt ReqOpt) (*http.Response, error) {
    client := http.Client{
        Timeout: opt.Timeout,
    }
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("create request err: %w", err)
    }
    req.Header.Set("Content-Type", "application/json")
    if opt.Headers != nil {
        for key, value := range opt.Headers {
            req.Header.Set(key, value)
        }
    }
    if opt.Param != nil { 
        q := req.URL.Query()
        for key, value := range opt.Param {
            q.Add(key, value)
        }
        req.URL.RawQuery = q.Encode()
    }
    if opt.Cookie != nil {
        for key, value := range opt.Cookie {
            req.AddCookie(&http.Cookie{Name: key, Value: value})
        }
    }
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("send request err: %w", err)
    }
    return resp, nil
}

func HttpPost(url string, opt ReqOpt) (*http.Response, error) {
    client := http.Client{
        Timeout: opt.Timeout,
    }
    
    var body io.Reader
    if opt.Body != nil {
        switch v := opt.Body.(type) {
        case string:
            body = strings.NewReader(v)
        case []byte:
            body = bytes.NewReader(v)
        default:
            jsonData, err := json.Marshal(v)
            if err != nil {
                return nil, fmt.Errorf("parse body err: %w", err)
            }
            body = bytes.NewReader(jsonData)
        }
    }
    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return nil, fmt.Errorf("create request err: %w", err)
    }
    req.Header.Set("Content-Type", "application/json")
    if opt.Headers != nil {
        for key, value := range opt.Headers {
            req.Header.Set(key, value)
        }
    }
    if opt.Param != nil { 
        q := req.URL.Query()
        for key, value := range opt.Param {
            q.Add(key, value)
        }
        req.URL.RawQuery = q.Encode()
    }
    if opt.Cookie != nil {
        for key, value := range opt.Cookie {
            req.AddCookie(&http.Cookie{Name: key, Value: value})
        }
    }
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("send request err: %w", err)
    }
    return resp, nil
}

// 解析响应
func Parse[T any](r *http.Response) (*T, error) {
	defer r.Body.Close()

	s, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("code: %d, response: %s", r.StatusCode, string(s))
	}
	var v T
	err = json.Unmarshal(s, &v)
	if err != nil {
		return nil, fmt.Errorf("parse err: %w", err)
	}
	return &v, nil
}


