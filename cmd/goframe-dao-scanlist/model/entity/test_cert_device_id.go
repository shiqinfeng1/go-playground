// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TestCertDeviceId is the golang structure for table test_cert_deviceId.
type TestCertDeviceId struct {
	Id        uint64      `json:"id"        ` //
	Identity  string      `json:"identity"  ` // 用户定义的证书标识
	DeviceId  string      `json:"deviceId"  ` // 设备号
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}