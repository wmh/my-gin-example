# CLAUDE.md — my-gin-example

Gin + GORM + SQLite 的 Go 範例服務。教學／參考性質的專案。

## 指令

| Command | 用途 |
|---|---|
| `make run` | `go run main.go`（預設 target） |
| `make build` | 編到 `bin/app` |
| `make test` | `go test -v ./...` |
| `make test-cover` | 帶覆蓋率 |
| `make install` | `go mod download && go mod tidy` |
| `make dev` | `air` 熱重載（需先裝 air） |
| `make clean` | 清 `bin/` 與 `data/*.db` |

改動後跑 `make test` 與 `go vet ./...`。

## 結構

```
main.go                  進入點
app/                     handler / service / model
config/                  設定
tests/                   測試
data/                    SQLite db（gitignored 的執行期產物）
examples/                使用範例
scripts/
docker-compose.yml, Dockerfile
```

文件：`README.md`、`QUICKSTART.md`、`FEATURES.md`、`CHANGELOG.md`。

## 注意事項

- 這是**範例專案**，可讀性優先於效能或架構複雜度。加功能時保持每個檔案能被單獨讀懂，不要為了「更正確」而引入抽象層
- `make clean` 會刪掉 `data/` 底下的 SQLite 檔，跑之前確認裡面沒有要留的資料
- 行為有變時同步更新 `README.md` / `FEATURES.md` 與 `CHANGELOG.md`
