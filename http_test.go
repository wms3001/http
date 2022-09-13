package http

import "testing"

func TestHttp_Get(t *testing.T) {
	m := map[string]string{
		"name":    "backy",
		"species": "dog",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	hp := Http{
		"https://www.wishpost.cn/api/v3/oauth/access_token",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := hp.Get()
	println(resp.Code)
	println(resp.HttpCode)
	println(resp.Message)
	println(resp.Data)
}

func TestHttp_Post(t *testing.T) {
	m := map[string]string{
		"name":    "backy",
		"species": "dog",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	hp := Http{
		"https://www.wishpost.cn/api/v3/oauth/access_token",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := hp.Post()
	println(resp.Code)
	println(resp.HttpCode)
	println(resp.Message)
	println(resp.Data)
}

func TestHttp_Delete(t *testing.T) {
	m := map[string]string{
		"name":    "backy",
		"species": "dog",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	hp := Http{
		"https://www.wishpost.cn/api/v3/oauth/access_token",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := hp.Delete()
	println(resp.Code)
	println(resp.HttpCode)
	println(resp.Message)
	println(resp.Data)
}

func TestHttp_Head(t *testing.T) {
	m := map[string]string{
		"name":    "backy",
		"species": "dog",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	hp := Http{
		"https://www.wishpost.cn/api/v3/oauth/access_token",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := hp.Head()
	println(resp.Code)
	println(resp.HttpCode)
	println(resp.Message)
	println(resp.Data)
}

func TestHttp_Options(t *testing.T) {
	m := map[string]string{
		"name":    "backy",
		"species": "dog",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	hp := Http{
		"https://www.wishpost.cn/api/v3/oauth/access_token",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := hp.Options()
	println(resp.Code)
	println(resp.HttpCode)
	println(resp.Message)
	println(resp.Data)
}

func TestHttp_Put(t *testing.T) {
	m := map[string]string{
		"name":    "backy",
		"species": "dog",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	hp := Http{
		"https://www.wishpost.cn/api/v3/oauth/access_token",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := hp.Put()
	println(resp.Code)
	println(resp.HttpCode)
	println(resp.Message)
	println(resp.Data)
}

func TestHttp_Trace(t *testing.T) {
	m := map[string]string{
		"name":    "backy",
		"species": "dog",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	hp := Http{
		"https://www.wishpost.cn/api/v3/oauth/access_token",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := hp.Trace()
	println(resp.Code)
	println(resp.HttpCode)
	println(resp.Message)
	println(resp.Data)
}
