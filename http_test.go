package http

import (
	"encoding/json"
	"testing"
)

type testCase struct {
	url    string
	params map[string]string
}

var c *Client

func init() {
	c = &Client{}
	c.Init()
}

func Test_Get(t *testing.T) {
	testCases := map[string]*testCase{
		"normal_post": &testCase{
			url: "http://localhost:80/data",
		},
	}
	for k, v := range testCases {
		t.Logf("start to test [%+v]", k)
		result, err := c.Get(v.url)
		t.Logf("[%+v]\tresult:%+v\terr:%+v", k, result, err)
	}
}

func Test_PostParams(t *testing.T) {
	testCases := map[string]*testCase{
		"normal_post": &testCase{
			url: "http://localhost:80/info.php",
			params: map[string]string{
				"md5s":    "01243208fb4e83165538242d85f25a91\t0\t\t\n",
				"product": "internal",
				"combo":   "internal",
				"v":       "2",
			},
		},
	}
	for k, v := range testCases {
		t.Logf("start to test [%+v]", k)
		result, err := c.PostParams(v.url, v.params)
		t.Logf("[%+v]\tresult:%+v\terr:%+v", k, result, err)
	}
}

func Test_PostJson(t *testing.T) {
	testCases := map[string]*testCase{
		"normal_post": &testCase{
			url: "http://localhost:80/query.php",
			params: map[string]string{
				"md5": "01243208fb4e83165538242d85f25a91",
			},
		},
	}
	for k, v := range testCases {
		t.Logf("start to test [%+v]", k)
		d := map[string][]string{}
		d["data"] = []string{v.params["md5"]}
		jbyte, err := json.Marshal(d)
		if err != nil {
			t.Errorf("%+v", err)
		}
		result, err := c.PostJson(v.url, jbyte)
		t.Logf("[%+v]\tresult:%+v\terr:%+v", k, result, err)
	}
}
