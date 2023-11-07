# 📕 go-cli
go-cli 用 Go 打造的第一个命令行工具

# 📕 安裝
首先下載安裝檔 https://golang.org/dl/
安裝 go 1.21.3 或以上的版本

如果版本太低會報錯
```
go: errors parsing go.mod:
C:\github\go-cli\go.mod:3: invalid go version '1.21.3': must match format 1.23
```
安裝指定版本後就可以編譯了

安裝完輸入(基本上不會設錯)
```bash=
go env
```
看看 GOROOT="C:\Go" 與 GOPATH="C:\users\youName\go" 是否正確(基本上不會設錯)

程式碼在 main.go 檔案

# 📕 初始化專案

```shell=
go mod init "專案名稱"
```

# 安裝 Cobra 模組
```shell=
go get -u github.com/spf13/cobra@latest
```

# 安裝 Cobra-cli
```shell=
go install github.com/spf13/cobra-cli@latest
```
可以直接產生範例程式碼 hi.go 到專案的 cmd 目錄下
```shell=
cobra-cli add hi
```

# 📕 編譯專案
```shell=
go build
```

# 📕 執行 CLI
```shell=
./go-cli
```

# 📕 CLI 顯示版本號
```shell=
./go-cli -v
```

# 📕 CLI 帶入參數
```shell=
./go-cli --name Sam
```





