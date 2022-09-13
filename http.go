package http

import (
	"github.com/parnurzeal/gorequest"
	"github.com/wms3001/goCommon"
	"time"
)

type Http struct {
	Url           string            `json:"url"`
	BasicAuthUser string            `json:"basicAuthUser"`
	BasicAuthPass string            `json:"basicAuthPass"`
	Header        map[string]string `json:"header"`
	AppendHeader  map[string]string `json:"appendHeader"`
	Body          string            `json:"body"`
	TimeOut       time.Duration     `json:"timeOut"`
	Type          string            `json:"type"`
	SFile         []byte            `json:"sFile"`
}

func (http *Http) Get() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	request := gorequest.New()
	request.Get(http.Url)
	if http.BasicAuthUser != "" {
		request.SetBasicAuth(http.BasicAuthUser, http.BasicAuthPass)
	}
	if http.TimeOut > 0 {
		request.Timeout(http.TimeOut * time.Second)
	} else {
		request.Timeout(60 * time.Second)
	}
	for k, v := range http.AppendHeader {
		request.AppendHeader(k, v)
	}
	for k, v := range http.Header {
		request.Set(k, v)
	}
	res, body, errs := request.End()
	println(res.StatusCode)
	if res == nil {
		resp.Code = -1
		resp.Message = errs[0].Error()
		resp.HttpCode = 0
	} else {
		if res.StatusCode == 200 {
			resp.Code = 1
			resp.Message = "success"
			resp.HttpCode = res.StatusCode
			resp.Data = body
		} else {
			resp.Code = 1
			if len(errs) > 0 {
				resp.Message = errs[0].Error()
			} else {
				resp.Message = "error"
			}
			resp.HttpCode = res.StatusCode
			resp.Data = body
		}
	}
	return resp
}

func (http *Http) Post() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	request := gorequest.New()
	request.Post(http.Url)

	for k, v := range http.Header {
		request.Set(k, v)
	}
	if http.BasicAuthUser != "" {
		request.SetBasicAuth(http.BasicAuthUser, http.BasicAuthPass)
	}
	if http.TimeOut > 0 {
		request.Timeout(http.TimeOut * time.Second)
	} else {
		request.Timeout(60 * time.Second)
	}
	if http.Type != "" {
		request.Type(http.Type)
	}
	if http.Body != "" {
		request.Send(http.Body)
	}
	if len(http.SFile) > 0 {
		request.SendFile(http.SFile)
	}
	res, body, errs := request.End()
	if res == nil {
		resp.Code = -1
		resp.Message = errs[0].Error()
		resp.HttpCode = 0
	} else {
		if res.StatusCode == 200 {
			resp.Code = 1
			resp.Message = "success"
			resp.HttpCode = res.StatusCode
			resp.Data = body
		} else {
			resp.Code = 1
			if len(errs) > 0 {
				resp.Message = errs[0].Error()
			} else {
				resp.Message = "error"
			}
			resp.HttpCode = res.StatusCode
			resp.Data = body
		}
	}
	return resp
}

func (http *Http) Put() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	request := gorequest.New()
	request.Put(http.Url)
	for k, v := range http.Header {
		request.Set(k, v)
	}
	if http.BasicAuthUser != "" {
		request.SetBasicAuth(http.BasicAuthUser, http.BasicAuthPass)
	}
	if http.TimeOut > 0 {
		request.Timeout(http.TimeOut * time.Second)
	} else {
		request.Timeout(60 * time.Second)
	}
	if http.Type != "" {
		request.Type(http.Type)
	}
	if http.Body != "" {
		request.Send(http.Body)
	}
	if len(http.SFile) > 0 {
		request.SendFile(http.SFile)
	}
	res, body, errs := request.End()
	if res == nil {
		resp.Code = -1
		resp.Message = errs[0].Error()
		resp.HttpCode = 0
	} else {
		if res.StatusCode == 200 {
			resp.Code = 1
			resp.Message = "success"
			resp.HttpCode = res.StatusCode
			resp.Data = body
		} else {
			resp.Code = 1
			if len(errs) > 0 {
				resp.Message = errs[0].Error()
			} else {
				resp.Message = "error"
			}
			resp.HttpCode = res.StatusCode
		}
	}
	return resp
}

func (http *Http) Delete() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	request := gorequest.New()
	request.Delete(http.Url)
	for k, v := range http.Header {
		request.Set(k, v)
	}
	if http.BasicAuthUser != "" {
		request.SetBasicAuth(http.BasicAuthUser, http.BasicAuthPass)
	}
	if http.TimeOut > 0 {
		request.Timeout(http.TimeOut * time.Second)
	} else {
		request.Timeout(60 * time.Second)
	}
	if http.Type != "" {
		request.Type(http.Type)
	}
	if http.Body != "" {
		request.Send(http.Body)
	}
	if len(http.SFile) > 0 {
		request.SendFile(http.SFile)
	}
	res, body, errs := request.End()
	if res == nil {
		resp.Code = -1
		resp.Message = errs[0].Error()
		resp.HttpCode = 0
	} else {
		if res.StatusCode == 200 {
			resp.Code = 1
			resp.Message = "success"
			resp.HttpCode = res.StatusCode
			resp.Data = body
		} else {
			resp.Code = 1
			if len(errs) > 0 {
				resp.Message = errs[0].Error()
			} else {
				resp.Message = "error"
			}
			resp.HttpCode = res.StatusCode
		}
	}
	return resp
}

func (http *Http) Head() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	request := gorequest.New()
	request.Head(http.Url)
	for k, v := range http.Header {
		request.Set(k, v)
	}
	if http.BasicAuthUser != "" {
		request.SetBasicAuth(http.BasicAuthUser, http.BasicAuthPass)
	}
	if http.TimeOut > 0 {
		request.Timeout(http.TimeOut * time.Second)
	} else {
		request.Timeout(60 * time.Second)
	}
	if http.Type != "" {
		request.Type(http.Type)
	}
	if http.Body != "" {
		request.Send(http.Body)
	}
	if len(http.SFile) > 0 {
		request.SendFile(http.SFile)
	}
	res, body, errs := request.End()
	if res == nil {
		resp.Code = -1
		resp.Message = errs[0].Error()
		resp.HttpCode = 0
	} else {
		if res.StatusCode == 200 {
			resp.Code = 1
			resp.Message = "success"
			resp.HttpCode = res.StatusCode
			resp.Data = body
		} else {
			resp.Code = 1
			if len(errs) > 0 {
				resp.Message = errs[0].Error()
			} else {
				resp.Message = "error"
			}
			resp.HttpCode = res.StatusCode
		}
	}
	return resp
}

func (http *Http) Options() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	request := gorequest.New()
	request.Options(http.Url)
	for k, v := range http.Header {
		request.Set(k, v)
	}
	if http.BasicAuthUser != "" {
		request.SetBasicAuth(http.BasicAuthUser, http.BasicAuthPass)
	}
	if http.TimeOut > 0 {
		request.Timeout(http.TimeOut * time.Second)
	} else {
		request.Timeout(60 * time.Second)
	}
	if http.Type != "" {
		request.Type(http.Type)
	}
	if http.Body != "" {
		request.Send(http.Body)
	}
	if len(http.SFile) > 0 {
		request.SendFile(http.SFile)
	}
	res, body, errs := request.End()
	if res == nil {
		resp.Code = -1
		resp.Message = errs[0].Error()
		resp.HttpCode = 0
	} else {
		if res.StatusCode == 200 {
			resp.Code = 1
			resp.Message = "success"
			resp.HttpCode = res.StatusCode
			resp.Data = body
		} else {
			resp.Code = 1
			if len(errs) > 0 {
				resp.Message = errs[0].Error()
			} else {
				resp.Message = "error"
			}
			resp.HttpCode = res.StatusCode
		}
	}
	return resp
}

func (http *Http) Trace() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	request := gorequest.New()
	request.CustomMethod("TRACE", http.Url)
	for k, v := range http.Header {
		request.Set(k, v)
	}
	if http.BasicAuthUser != "" {
		request.SetBasicAuth(http.BasicAuthUser, http.BasicAuthPass)
	}
	if http.TimeOut > 0 {
		request.Timeout(http.TimeOut * time.Second)
	} else {
		request.Timeout(60 * time.Second)
	}
	if http.Type != "" {
		request.Type(http.Type)
	}
	if http.Body != "" {
		request.Send(http.Body)
	}
	if len(http.SFile) > 0 {
		request.SendFile(http.SFile)
	}
	res, body, errs := request.End()
	if res == nil {
		resp.Code = -1
		resp.Message = errs[0].Error()
		resp.HttpCode = 0
	} else {
		if res.StatusCode == 200 {
			resp.Code = 1
			resp.Message = "success"
			resp.HttpCode = res.StatusCode
			resp.Data = body
		} else {
			resp.Code = 1
			if len(errs) > 0 {
				resp.Message = errs[0].Error()
			} else {
				resp.Message = "error"
			}
			resp.HttpCode = res.StatusCode
		}
	}
	return resp
}
