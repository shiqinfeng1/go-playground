package model

import "go-playground/cmd/gofrome-dao-scanlist/model/entity"

type Cert struct {
	CertDevId *entity.TestCertDeviceId
	CertInfo  *entity.TestCertInfo
}
