package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"utils/function"
)

const MAX_PS = 1000

type HttpRequest struct {
	UrlEncodedBody map[string][]string
	JsonBody       map[string]interface{} //可存储body解密的json解析后的结构体
	BodyRaw        []byte                 //可存储body解密的data
	request        *http.Request
	Uid            uint32 //用户的UID
}

func (hr *HttpRequest) GetParam(name string) string {
	return hr.request.URL.Query().Get(name)
}

func (hr *HttpRequest) PostParam(name string) string {
	return hr.request.FormValue(name)
}

func (hr *HttpRequest) IP() string {
	ips := strings.Split(hr.request.Header.Get("X-Forwarded-For"), ",")
	if len(ips[0]) > 3 {
		return ips[0]
	} else {
		addr := strings.Split(hr.request.RemoteAddr, ":")
		return addr[0]
	}
}

func (hr *HttpRequest) Cookie(name string) (*http.Cookie, error) {
	return hr.request.Cookie(name)
}

func (hr *HttpRequest) Host() string {
	return hr.request.Host
}

func (hr *HttpRequest) UserAgent() string {
	return hr.request.UserAgent()
}

func (hr *HttpRequest) Via() string {
	return hr.request.Header.Get("Via")
}

func (hr *HttpRequest) X_Wap_profile() string {
	var s string
	s = hr.request.Header.Get("x-wap-profile")
	if s == "" {
		s = hr.request.Header.Get("X-Wap-Profile")
	}
	if s == "" {
		s = hr.request.Header.Get("X-WAP-PROFILE")
	}
	return s
}

//检查Body中的字段是否齐全
func (hr *HttpRequest) EnsureBody(keys ...string) (string, bool) {
	for _, key := range keys {
		if _, ok := hr.JsonBody[key]; !ok {
			return key, false
		}
	}
	return "", true
}

//带默认值的解析
func (hr *HttpRequest) ParseOpt(params ...interface{}) error {
	if len(params)%3 != 0 {
		return errors.New("params count invalid")
	}
	for i := 0; i < len(params); i += 3 {
		key := function.ToString(params[i])
		v, ok := hr.JsonBody[key]
		var e error
		switch ref := params[i+1].(type) {
		case *string:
			if ok {
				*ref = function.ToString(v)
			} else {
				*ref = function.ToString(params[i+2])
			}

		case *float64:
			if ok {
				*ref, e = function.ToFloat64(v)
			} else {
				*ref, e = function.ToFloat64(params[i+2])
			}
		case *int:
			if ok {
				*ref, e = function.ToInt(v)
			} else {
				*ref, e = function.ToInt(params[i+2])
			}
		case *uint32:
			if ok {
				*ref, e = function.ToUint32(v)
			} else {
				*ref, e = function.ToUint32(params[i+2])
			}
		case *uint64:
			if ok {
				*ref, e = function.ToUint64(v)
			} else {
				*ref, e = function.ToUint64(params[i+2])
			}
		case *int64:
			if ok {
				*ref, e = function.ToInt64(v)
			} else {
				*ref, e = function.ToInt64(params[i+2])
			}
		case *int8:
			if ok {
				*ref, e = function.ToInt8(v)
			} else {
				*ref, e = function.ToInt8(params[i+2])
			}
		case *uint:
			if ok {
				*ref, e = function.ToUint(v)
			} else {
				*ref, e = function.ToUint(params[i+2])
			}
		case *bool:
			if ok {
				*ref, e = function.ToBool(v)
			} else {
				*ref, e = function.ToBool(params[i+2])
			}
		case *[]string:
			if ok {
				*ref, e = function.ToStringSlice(v)
			} else {
				*ref = params[i+2].([]string)
			}
		case *[]uint32:
			if ok {
				*ref, e = function.ToUint32Slice(v)
			} else {
				*ref = params[i+2].([]uint32)
			}
		case *map[string]interface{}:
			if ok {
				switch m := v.(type) {
				case map[string]interface{}:
					*ref = m
				default:
					e = errors.New(fmt.Sprintf("%v is not map[string]iterface{}", key, reflect.TypeOf(v)))
				}
			} else {
				*ref = params[i+2].(map[string]interface{})
			}
		case *interface{}:
			if ok {
				*ref = v
			} else {
				*ref = params[i+2]
			}
		default:
			return errors.New(fmt.Sprintf("unknown type %v ", key, reflect.TypeOf(ref)))
		}
		if e != nil {
			return errors.New(fmt.Sprintf("parse [%v] error:%v", key, e.Error()))
		}
	}
	return nil
}

//不带默认值的解析
func (hr *HttpRequest) Parse(params ...interface{}) error {
	if len(params)%2 != 0 {
		return errors.New("params count must be odd")
	}
	for i := 0; i < len(params); i += 2 {
		key := function.ToString(params[i])
		if v, ok := hr.JsonBody[key]; ok {
			var e error
			switch ref := params[i+1].(type) {
			case *string:
				*ref = function.ToString(v)
			case *float64:
				*ref, e = function.ToFloat64(v)
			case *int:
				*ref, e = function.ToInt(v)
			case *uint16:
				*ref, e = function.ToUint16(v)
			case *uint32:
				*ref, e = function.ToUint32(v)
			case *uint64:
				*ref, e = function.ToUint64(v)
			case *int64:
				*ref, e = function.ToInt64(v)
			case *int16:
				*ref, e = function.ToInt16(v)
			case *int8:
				*ref, e = function.ToInt8(v)
			case *uint:
				*ref, e = function.ToUint(v)
			case *map[string]interface{}:
				switch m := v.(type) {
				case map[string]interface{}:
					*ref = m
				default:
					e = errors.New("value is not map[string]iterface{}")
				}
			case *[]string:
				*ref, e = function.ToStringSlice(v)
			case *[]uint32:
				*ref, e = function.ToUint32Slice(v)
			case *interface{}:
				*ref = v
			default:
				return errors.New(fmt.Sprintf("unknown type %v ", reflect.TypeOf(ref)))
			}
			if e != nil {
				return errors.New(fmt.Sprintf("parse [%v] error:%v", key, e.Error()))
			}
			if key == "ps" {
				ps, e := function.ToUint64(v)
				if e == nil && ps > MAX_PS {
					return errors.New("ps too large")
				}
			}
		} else {
			return errors.New(fmt.Sprintf("%v not provided", key))
		}
	}
	return nil
}

func (hr *HttpRequest) PostFile(param string, filename string) (e error) {
	// hr.request.ParseMultipartForm(1024 * 1024 * 10)
	file, _, e := hr.request.FormFile(param)
	if e != nil {
		return errors.New(fmt.Sprintf("FormFile: %v", e.Error()))
	}

	defer file.Close()

	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		return errors.New(fmt.Sprintf("ReadAll: %v", e.Error()))
	}
	e = ioutil.WriteFile(filename, bytes, os.ModePerm)
	return
}

//获取HTTP post的详细信息
func (hr *HttpRequest) PostFileInfo(param string) (bytes []byte, filename string, e error) {
	// hr.request.ParseMultipartForm(1024 * 1024 * 10)
	file, handler, e := hr.request.FormFile(param)
	if e != nil {
		e = errors.New(fmt.Sprintf("FormFile: %v", e.Error()))
		return
	}

	defer file.Close()
	filename = handler.Filename
	bytes, e = ioutil.ReadAll(file)
	if e != nil {
		e = errors.New(fmt.Sprintf("ReadAll: %v", e.Error()))
		return
	}
	// e = ioutil.WriteFile(filename, bytes, os.ModePerm)
	return
}

func (hr *HttpRequest) GetParams(params ...interface{}) error {
	if len(params)%2 != 0 {
		return errors.New("params count must be odd")
	}
	for i := 0; i < len(params); i += 2 {
		key := function.ToString(params[i])
		if v := hr.request.URL.Query().Get(key); v != "" {
			var e error
			switch ref := params[i+1].(type) {
			case *string:
				*ref = function.ToString(v)
			case *float64:
				*ref, e = function.ToFloat64(v)
			case *int:
				*ref, e = function.ToInt(v)
			case *int8:
				*ref, e = function.ToInt8(v)
			case *int16:
				*ref, e = function.ToInt16(v)
			case *int32:
				*ref, e = function.ToInt32(v)
			case *int64:
				*ref, e = function.ToInt64(v)
			case *uint:
				*ref, e = function.ToUint(v)
			case *uint8:
				*ref, e = function.ToUint8(v)
			case *uint16:
				*ref, e = function.ToUint16(v)
			case *uint32:
				*ref, e = function.ToUint32(v)
			case *uint64:
				*ref, e = function.ToUint64(v)
			case *map[string]interface{}:
				e = errors.New("do not support map[string]iterface{}")
			case *[]interface{}:
				e = errors.New("do not support []iterface{}")
			case *[]string:
				ll := strings.Split(v, ",")
				*ref, e = function.ToStringSlice(ll)
			case *[]uint32:
				ll := strings.Split(v, ",")
				*ref, e = function.ToUint32Slice(ll)
			case *interface{}:
				*ref = v
			default:
				return errors.New(fmt.Sprintf("unknown type %v ", reflect.TypeOf(ref)))
			}
			if e != nil {
				return errors.New(fmt.Sprintf("parse [%v] error:%v", key, e.Error()))
			}
		} else {
			return errors.New(fmt.Sprintf("%v not provided value", key))
		}
	}
	return nil
}

func (hr *HttpRequest) GetParamOpt(params ...interface{}) error {
	if len(params)%3 != 0 {
		return errors.New("params count invalid")
	}
	for i := 0; i < len(params); i += 3 {
		key := function.ToString(params[i])
		v := hr.request.URL.Query().Get(key)
		var e error
		switch ref := params[i+1].(type) {
		case *string:
			if v != "" {
				*ref = function.ToString(v)
			} else {
				*ref = function.ToString(params[i+2])
			}
		case *float64:
			if v != "" {
				*ref, e = function.ToFloat64(v)
			} else {
				*ref, e = function.ToFloat64(params[i+2])
			}
		case *int:
			if v != "" {
				*ref, e = function.ToInt(v)
			} else {
				*ref, e = function.ToInt(params[i+2])
			}
		case *int8:
			if v != "" {
				*ref, e = function.ToInt8(v)
			} else {
				*ref, e = function.ToInt8(params[i+2])
			}
		case *int16:
			if v != "" {
				*ref, e = function.ToInt16(v)
			} else {
				*ref, e = function.ToInt16(params[i+2])
			}
		case *int32:
			if v != "" {
				*ref, e = function.ToInt32(v)
			} else {
				*ref, e = function.ToInt32(params[i+2])
			}
		case *int64:
			if v != "" {
				*ref, e = function.ToInt64(v)
			} else {
				*ref, e = function.ToInt64(params[i+2])
			}
		case *uint:
			if v != "" {
				*ref, e = function.ToUint(v)
			} else {
				*ref, e = function.ToUint(params[i+2])
			}
		case *uint8:
			if v != "" {
				*ref, e = function.ToUint8(v)
			} else {
				*ref, e = function.ToUint8(params[i+2])
			}
		case *uint16:
			if v != "" {
				*ref, e = function.ToUint16(v)
			} else {
				*ref, e = function.ToUint16(params[i+2])
			}
		case *uint32:
			if v != "" {
				*ref, e = function.ToUint32(v)
			} else {
				*ref, e = function.ToUint32(params[i+2])
			}
		case *uint64:
			if v != "" {
				*ref, e = function.ToUint64(v)
			} else {
				*ref, e = function.ToUint64(params[i+2])
			}
		case *bool:
			if v != "" {
				*ref, e = function.ToBool(v)
			} else {
				*ref, e = function.ToBool(params[i+2])
			}
		case *[]string:
			if v != "" {
				*ref, e = function.ToStringSlice(v)
			} else {
				*ref = params[i+2].([]string)
			}
		case *[]uint32:
			if v != "" {
				*ref, e = function.ToUint32Slice(v)
			} else {
				*ref = params[i+2].([]uint32)
			}
		case *map[string]interface{}:
			e = errors.New("do not support map[string]iterface{}")
		case *[]interface{}:
			e = errors.New("do not support []iterface{}")
		default:
			return errors.New(fmt.Sprintf("unknown type %v ", key, reflect.TypeOf(ref)))
		}
		if e != nil {
			return errors.New(fmt.Sprintf("parse [%v] error:%v", key, e.Error()))
		}
	}
	return nil
}

//不带默认值的EncodeUrl解析
func (hr *HttpRequest) ParseEncodeUrl(params ...interface{}) error {
	if len(params)%2 != 0 {
		return errors.New("params count must be odd")
	}
	for i := 0; i < len(params); i += 2 {
		key := function.ToString(params[i])
		if vs := hr.UrlEncodedBody[key]; len(vs)>0 {
			v:=vs[0]
			var e error
			switch ref := params[i+1].(type) {
			case *string:
				*ref = function.ToString(v)
			case *float64:
				*ref, e = function.ToFloat64(v)
			case *int:
				*ref, e = function.ToInt(v)
			case *int8:
				*ref, e = function.ToInt8(v)
			case *int16:
				*ref, e = function.ToInt16(v)
			case *int32:
				*ref, e = function.ToInt32(v)
			case *int64:
				*ref, e = function.ToInt64(v)
			case *uint:
				*ref, e = function.ToUint(v)
			case *uint8:
				*ref, e = function.ToUint8(v)
			case *uint16:
				*ref, e = function.ToUint16(v)
			case *uint32:
				*ref, e = function.ToUint32(v)
			case *uint64:
				*ref, e = function.ToUint64(v)
			case *bool:
				*ref, e = function.ToBool(v)
			case *map[string]interface{}:
				e = errors.New("do not support map[string]iterface{}")
			case *[]interface{}:
				e = errors.New("do not support []iterface{}")
			case *[]string:
				ll := strings.Split(v, ",")
				*ref, e = function.ToStringSlice(ll)
			case *[]uint32:
				ll := strings.Split(v, ",")
				*ref, e = function.ToUint32Slice(ll)
			case *interface{}:
				*ref = v
			default:
				return errors.New(fmt.Sprintf("unknown type %v ", reflect.TypeOf(ref)))
			}
			if e != nil {
				return errors.New(fmt.Sprintf("parse [%v] error:%v", key, e.Error()))
			}
		} else {
			return errors.New(fmt.Sprintf("%v not provided value", key))
		}
	}
	return nil
}
