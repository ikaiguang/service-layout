# service layout

**启动服务；启动请检查配置文件；如果使用配置中心，请跳过执行，先阅读`服务配置`部分**

参考配置文件：[config_all.yaml](https://github.com/ikaiguang/go-srv-kit/blob/main/testdata/configs/configs/config_all.yaml)

```shell

# 查看帮助
make help

# 运行服务
make run-service

# 测试服务
make testing-service
```

## 服务配置

* 配置文件：[app.yaml](./app/service-layout/configs/app.yaml)

**示例如下：**

```yaml
# app 程序
app:
  server_name: xxx-service
  # 配置方式；值：local、consul、etcd
  config_method: consul
```

如果`app.config_method`配置是使用配置中心(consul、etcd、...)，首先把配置写入配置中心：

```shell
# 执行
make store-configuration
# or
#* `conf`： 启动读取配置
#* `source_dir`： 被存储的配置文件所在文件夹
#* `store_dir`： 存储到配置中心位置
go run ./app/service-layout/cmd/store-configuration/... -conf=./app/service-layout/configs \
-source_dir=./app/service-layout/configs \
-store_dir=go-micro-saas/service-layout/develop/v1.0.0

go run ./app/service-layout/cmd/store-configuration/... -conf=./app/service-layout/configs \
-source_dir=./app/uuid-service/configs \
-store_dir=go-micro-saas/uuid-service/develop/v1.0.0
```

## 文档

> 生成文档：`make protoc-api-protobuf`

* api文档swagger [xxx.swagger.json](api/testing-service/v1/services/testdata.service.v1.swagger.json)
