package models

import "encoding/json"

type KeyValue struct {
	Key   string `json:"key" binding:"required"`
	Value json.RawMessage `json:"value"`
}

type KeyValueResp struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value"`
}

type KeyValueUpdateReq struct {
	Value string `json:"value"`
}