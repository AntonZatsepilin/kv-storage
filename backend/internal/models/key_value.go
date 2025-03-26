package models

import "encoding/json"

type KeyValue struct {
    Key   string `json:"key" binding:"required" example:"my_key"`
    Value json.RawMessage `json:"value" swaggertype:"string" example:"{\"data\":\"value\"}"`
}

type KeyValueResp struct {
    Key   string `json:"key" example:"my_key"`
    Value string `json:"value" example:"{\"data\":\"value\"}"`
}

type KeyValueUpdateReq struct {
    Value string `json:"value" example:"new_value"`
}