# GoMoudle

---

**初始化**

```bash
cd myproject
go mod init myproject
```





`go.mod`文件的组成部分

```go
module example.com/myproject    // 模块名称

go 1.16    // Go 版本声明

require (    // 依赖声明
    github.com/pkg/errors v0.9.1
    golang.org/x/text v0.3.0
)

replace (    // 替换依赖源
    old-module => new-module
)

exclude (    // 排除特定版本的依赖
    bad-module v1.4.0
)
```

依赖管理示例

```go
# 添加新依赖
go get example.com/some-module

# 添加特定版本的依赖
go get example.com/some-module@v1.3.4

# 更新依赖到最新版本
go get -u example.com/some-module

# 更新所有依赖
go get -u all
```





`go mod tidy`作用

1. **移除未使用的依赖**：如果 `go.mod` 或 `go.sum` 文件中包含了一些已经不再使用的依赖，`go mod tidy` 会将它们移除，保持 `go.mod` 文件整洁。
2. **添加遗漏的依赖**：如果 `go.mod` 中缺少某些在代码中 `import` 了但未显式声明的依赖，`go mod tidy` 会自动添加它们。
3. **更新 `go.sum` 文件**：`go.sum` 记录了模块的完整性校验信息，`go mod tidy` 会确保 `go.sum` 文件的内容是最新的，以防止不一致的问题。