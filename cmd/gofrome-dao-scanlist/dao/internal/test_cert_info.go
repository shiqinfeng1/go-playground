// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestCertInfoDao is the data access object for table test_cert_info.
type TestCertInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TestCertInfoColumns // columns contains all the column names of Table for convenient usage.
}

// TestCertInfoColumns defines and stores column names for table test_cert_info.
type TestCertInfoColumns struct {
	Id                string //
	Version           string // 版本号
	Identity          string // 用户定义的证书标识
	CertType          string // 证书类型
	IssuerIdentity    string // 颁发机构标识
	Subject           string // 证书主体信息
	ValidityNotbefore string // 证书有效期开始
	ValidityNotafter  string // 证书有效期结束
	CustomerName      string // 用户名称
	CustomerIdcard    string // 用户身份证号
	CustomerId        string // 客户id
	SerialNumber      string // 证书序列号
	Status            string // 证书状态
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// testCertInfoColumns holds the columns for table test_cert_info.
var testCertInfoColumns = TestCertInfoColumns{
	Id:                "id",
	Version:           "version",
	Identity:          "identity",
	CertType:          "cert_type",
	IssuerIdentity:    "issuer_identity",
	Subject:           "subject",
	ValidityNotbefore: "validity_notbefore",
	ValidityNotafter:  "validity_notafter",
	CustomerName:      "customer_name",
	CustomerIdcard:    "customer_idcard",
	CustomerId:        "customer_id",
	SerialNumber:      "serial_number",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewTestCertInfoDao creates and returns a new DAO object for table data access.
func NewTestCertInfoDao() *TestCertInfoDao {
	return &TestCertInfoDao{
		group:   "default",
		table:   "test_cert_info",
		columns: testCertInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TestCertInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TestCertInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TestCertInfoDao) Columns() TestCertInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TestCertInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TestCertInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TestCertInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}