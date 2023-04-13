// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestCertDeviceIdDao is the data access object for table test_cert_deviceId.
type TestCertDeviceIdDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns TestCertDeviceIdColumns // columns contains all the column names of Table for convenient usage.
}

// TestCertDeviceIdColumns defines and stores column names for table test_cert_deviceId.
type TestCertDeviceIdColumns struct {
	Id        string //
	Identity  string // 用户定义的证书标识
	DeviceId  string // 设备号
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// testCertDeviceIdColumns holds the columns for table test_cert_deviceId.
var testCertDeviceIdColumns = TestCertDeviceIdColumns{
	Id:        "id",
	Identity:  "identity",
	DeviceId:  "deviceId",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTestCertDeviceIdDao creates and returns a new DAO object for table data access.
func NewTestCertDeviceIdDao() *TestCertDeviceIdDao {
	return &TestCertDeviceIdDao{
		group:   "default",
		table:   "test_cert_deviceId",
		columns: testCertDeviceIdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TestCertDeviceIdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TestCertDeviceIdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TestCertDeviceIdDao) Columns() TestCertDeviceIdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TestCertDeviceIdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TestCertDeviceIdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TestCertDeviceIdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}