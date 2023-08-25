package main

import (
	"fmt"
	"os"
	"path"

	"go-playground/cmd/gofrome-dao-scanlist/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	dir, _ := os.Getwd()
	err := Database().ImportDBTable(ctx, "", path.Join(dir, "cmd/gofrome-dao-scanlist/init.sql"))
	if err != nil {
		fmt.Println(err)
		return
	}

	g.Dump("从证书绑定的多个设备中，查找指定设备绑定的最新的证书信息")
	var entities []*model.Cert
	err = g.DB().Model("test_cert_deviceId").
		Where("identity", "identity11").
		Where("deviceId", "deviceId3").
		ScanList(&entities, "CertDevId")
	if err != nil {
		panic(err)
	}
	err = g.DB().Model("test_cert_info").
		Where("status", "可用").
		Where("identity", gdb.ListItemValuesUnique(entities, "CertDevId", "Identity")).
		ScanList(&entities, "CertInfo")
	if err != nil {
		panic(err)
	}
	g.Dump(entities)
	g.Dump("===================================================================")
	g.Dump("查找登录账号在当前设备上的证书列表")
	var entities2 []*model.Cert
	err = g.DB().Model("test_cert_deviceId").
		Where("customer_id", 1).
		Where("deviceId", "deviceId2").
		ScanList(&entities2, "CertDevId")
	if err != nil {
		panic(err)
	}
	g.Dump(entities2)
	g.Dump("-----------------------")
	err = g.DB().Model("test_cert_info").
		OrderDesc("id").
		Where("identity", gdb.ListItemValuesUnique(entities2, "CertDevId", "Identity")).
		ScanList(&entities2, "CertInfo")
	if err != nil {
		panic(err)
	}
	g.Dump(entities2)

	g.Dump("--------去除重复的记录， 以最新的为准---")
	var entities3 []*model.Cert
	//
	ss, _ := g.DB().Model("test_cert_info").Fields("id").Group("identity").Array()
	err = g.DB().Model("test_cert_info").
		Where("id", ss).
		OrderDesc("id").
		ScanList(&entities3, "CertInfo")
	if err != nil {
		panic(err)
	}
	g.Dump(entities3)
}
