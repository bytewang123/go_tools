package main

import (
	"context"
	"github.com/tsuna/gohbase"
	//"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
)

type InsertData struct {
	ColumnFamily string
	Qualifier    string
	Data         []byte
}

type HbaseClient struct {
	Client gohbase.Client
}

func (c *HbaseClient) Init(zkHost string) {
	c.Client = gohbase.NewClient(zkHost)
}

func (c *HbaseClient) Get(tableName string, rowKey string) (*hrpc.Result, error) {
	hbaseReq, err := hrpc.NewGetStr(context.Background(), tableName, rowKey)
	if err != nil {
		return nil, err
	}
	hbaseResp, err := c.Client.Get(hbaseReq)
	if err != nil {
		return nil, err
	}
	return hbaseResp, nil
}

func (c *HbaseClient) Insert(tableName string, rowKey string,
	insertData InsertData) (*hrpc.Result, error) {

	value := map[string]map[string]([]byte){
		insertData.ColumnFamily: {
			insertData.Qualifier: insertData.Data,
		},
	}
	hbaseReq, err := hrpc.NewPutStr(context.Background(),
		tableName, rowKey, value)
	if err != nil {
		return nil, err
	}
	hbaseResp, err := c.Client.Put(hbaseReq)
	if err != nil {
		return nil, err
	}
	return hbaseResp, nil
}
