# CircleCI 設置指南

## 快速設置（5 分鐘）

### 1. 添加項目

訪問 https://circleci.com/ 並用 GitHub 登入，然後添加項目：

```
Projects → 找到 my-gin-example → Set Up Project → Use Existing Config → master
```

### 2. 修復 SSH 錯誤

**錯誤**: `git@github.com: Permission denied (publickey)`

**解決方案**：

#### Deploy Key（推薦，自動同步）

CircleCI 的 Deploy Key 會自動同步到 GitHub，無需手動複製貼上。

**步驟**：

1. 訪問 https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh

2. 刪除舊的 Deploy Key（如果有）：
   - 點擊 "Delete Deploy Key"

3. 添加新的 Deploy Key：
   - 點擊 "Add Deploy Key"
   - 系統會自動生成並同步到 GitHub

4. 驗證 GitHub：
   - 訪問 https://github.com/wmh/my-gin-example/settings/keys
   - 應該會自動看到新的 Deploy Key

5. 重新運行 Pipeline

**注意**：Deploy Key 由 CircleCI 自動管理，會自動出現在 GitHub 倉庫設置中。

#### User Key（備選，快速）

1. 訪問 https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
2. 點擊 "Add User Key" → 授權 GitHub
3. 重新運行 Pipeline

### 3. 驗證

訪問 https://app.circleci.com/pipelines/github/wmh/my-gin-example

確認看到：
- ✅ Cloning git repository (成功)
- ✅ build-and-test (運行中)
- ✅ lint (運行中)
- ✅ security-scan (運行中)

## 本地驗證

```bash
./scripts/verify_circleci.sh
```

## 配置說明

`.circleci/config.yml` 包含：

- Go 1.23 環境
- 測試和構建
- golangci-lint 檢查
- gosec 安全掃描
- 依賴緩存

## 常見問題

**Q: Permission denied (publickey)?**
A: 按照上面的方法 A 或 B 添加 SSH key

**Q: 測試失敗?**
A: 本地運行 `go test ./...` 檢查

**Q: Deploy Key vs User Key?**
A: Deploy Key 更安全（僅限單倉庫，自動同步到 GitHub），User Key 更快（訪問所有倉庫）

**Q: Deploy Key 會自動出現在 GitHub?**
A: 是的，CircleCI 添加 Deploy Key 後會自動同步到 GitHub 倉庫的 Deploy keys 列表

**Q: 如何重新生成 Deploy Key?**
A: 在 CircleCI SSH Keys 頁面刪除舊的，點擊 "Add Deploy Key"，新的會自動同步到 GitHub

## 連結

- SSH Keys: https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
- Pipeline: https://app.circleci.com/pipelines/github/wmh/my-gin-example
- GitHub Deploy Keys: https://github.com/wmh/my-gin-example/settings/keys
- 文檔: https://circleci.com/docs/
