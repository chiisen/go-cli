# 📕 go-cli
go-cli 用 Go 打造的第一个命令行工具

# 📕 安裝
首先下載安裝檔 https://golang.org/dl/
安裝 go 1.23.0 或以上的版本

如果版本太低會報錯
```bash
go: errors parsing go.mod:
C:\github\go-cli\go.mod:3: invalid go version '1.21.3': must match format 1.23
```
安裝指定版本後就可以編譯了

安裝完輸入(基本上不會設錯)
```bash
go env
```
看看 GOROOT="C:\Go" 與 GOPATH="C:\users\youName\go" 是否正確(基本上不會設錯)

程式碼在 main.go 檔案

# 📕 初始化專案

```shell
go mod init "專案名稱"
```

# 安裝 Cobra 模組
```shell
go get -u github.com/spf13/cobra@latest
```

# 安裝 Cobra-cli
```shell
go install github.com/spf13/cobra-cli@latest
```
可以直接產生範例程式碼 hi.go 到專案的 cmd 目錄下
```shell
cobra-cli add hi
```

# 📕 編譯專案
```shell
go build
```
或
```shell
go build -o go-cli.exe
```
-o：這個選項用於指定輸出檔案的名稱和位置。

# 全域執行
```shell
move go-cli.exe C:\Windows\System32
```
遇到【存取被拒】可以在以系統管理員身份運行的命令提示字元中，執行以下命令。
或可以選擇其他目錄，然後將該目錄添加到 PATH 環境變數中：
例如: 將檔案移動到其他目錄，例如 C:\MyGoTools：
```shell
move go-cli.exe C:\MyGoTools
```

# 📕 執行 CLI
```shell
./go-cli

# 全域呼叫
go-cli
```

# 📕 CLI 顯示版本號
```shell
./go-cli -v

# 全域呼叫
go-cli -v
```

# 📕 CLI 帶入參數
```shell
./go-cli --name Sam

# 全域呼叫
go-cli --name Sam
```





