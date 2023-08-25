package main

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type DatabaseIntf interface {
	ImportDBTable(ctx context.Context, dbName, sqlFilePath string) (err error)
}

// database 上下文管理服务
var databaseService = databaseServiceImpl{}

type databaseServiceImpl struct{}

func Database() DatabaseIntf {
	return &databaseService
}

// 创建数据库表
func (d *databaseServiceImpl) ImportDBTable(ctx context.Context, dbName, sqlFilePath string) (err error) {
	sqlArr, err := readSqlFile(sqlFilePath)
	if err != nil {
		return err
	}

	if len(sqlArr) == 0 {
		return gerror.Newf("sqlArr=%v", sqlArr)
	}
	var db gdb.DB
	if dbName != "" {
		db = g.DB(dbName) //指定数据库
	} else {
		db = g.DB() //默认数据库
	}

	for _, item := range sqlArr {
		_, err = db.Exec(ctx, item)
		if err != nil {
			return err
		}
	}

	return
}

// 读取sql文件
func readSqlFile(path string) (sqlArr []string, err error) {

	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return nil, inputError
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	var (
		flag   bool = true
		buffer bytes.Buffer
	)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		inputString = strings.TrimSpace(inputString)
		if strings.HasPrefix(inputString, "--") {
			flag = false
		}
		if strings.HasPrefix(inputString, "/*") {
			flag = false
		}

		if flag && inputString != "" {
			if strings.HasSuffix(inputString, ";") {
				if buffer.Len() == 0 {
					sqlArr = append(sqlArr, inputString)
				} else {
					buffer.WriteString(inputString)
					sqlArr = append(sqlArr, buffer.String())
					buffer.Reset()
				}
			} else {
				buffer.WriteString(inputString + " ")
			}
		}

		if !flag && strings.HasPrefix(inputString, "*/") {
			flag = true
		}

		if !flag && strings.HasPrefix(inputString, "--") {
			flag = true
		}

		if readerError == io.EOF {
			break
		}
	}
	return
}

type SqlArr []string

func (arr SqlArr) Split(num int) (result SqlArr) {
	if len(arr) == 0 {
		return
	}
	var (
		buffer bytes.Buffer
		i      int
	)

	for _, item := range arr {
		buffer.WriteString(item)
		i++
		if i == num {
			result = append(result, buffer.String())
			buffer.Reset()
			i = 0
		}
	}

	if buffer.Len() > 0 {
		result = append(result, buffer.String())
	}

	return

}
