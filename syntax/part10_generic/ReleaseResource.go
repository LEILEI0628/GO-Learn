package main

import "encoding/json"

func ReleaseResource[R json.Marshaler](r R) { // [R json.Marshaler]：可以传入接口或any（其他类型不支持）
	r.MarshalJSON()
}
