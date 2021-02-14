package utils

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

func DeepCopy(v1, v2 interface{}) error {
	if bs, err := jsoniter.Marshal(v1); err != nil {
		return err
	} else {
		return jsoniter.Unmarshal(bs, v2)
	}
}

func CompareObj(v1, v2 interface{}) (result bool, err error) {
	bs1, err := jsoniter.Marshal(v1)
	if err != nil {
		return result, err
	}
	bs2, err := jsoniter.Marshal(v2)
	if err != nil {
		return result, err
	}
	if bytes.Compare(bs1, bs2) != 0 {
		return false, nil
	}
	return true, nil
}
