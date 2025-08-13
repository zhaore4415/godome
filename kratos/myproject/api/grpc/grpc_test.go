package grpc

import (
	"io/ioutil"
	"testing"

	"bsi/kratos/micro/auth"
	"bsi/kratos/micro/infras"
)

const packageName = "hello"

func TestParseAsset(t *testing.T) {
	infos, err := auth.ParseAsset(packageName, ".")
	if err != nil {
		panic(err)
	}

	for _, info := range infos {
		t.Log(info)
	}

	json := infras.ToJson(infos)
	if err = ioutil.WriteFile("../../internal/asset/grpc.json", []byte(json), 0644); err != nil {
		t.Fatal(err.Error())
	}
}
