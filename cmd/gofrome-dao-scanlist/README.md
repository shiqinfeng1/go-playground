# 数据库动态关联查询
- 1:N的查询
- N:1的查询

>使用goframe封装特性函数实现

# 首次使用
- 修改config.yaml中数据库的配置
- 把 main.go中数据操作的代码先注释掉，保留导入数据的代码
- 运行 `go run cmd/gofrome-dao-scanlist/*.go`， 导入测试数据到数据库
- 按照下面步骤生成代码

# 生成dao和do文件的步骤
1. 把config.yaml 复制到go-playground下
1. 在go-playground下执行
    ```
    ~/gf gen dao -p cmd/gofrome-dao-scanlist -t test_cert_deviceId,test_cert_info
    ```
    成功运行后会生成dao和model文件夹及文件