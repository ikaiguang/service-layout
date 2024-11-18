# new project

```shell
# 克隆项目
git clone https://github.com/ikaiguang/service-layout.git
# 移除git仓库地址
git remote remove origin
# 全局替换go.mod
github.com/ikaiguang/service-layout -> github.com/your-name/your-service
# 运行
# go mod download && go mod tidy
make run-all-in-one
make testing-all-in-one# 删除文件
rm README-NEW-PROJECT.md
```