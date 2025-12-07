# CircleCI 設置指南

## 快速設置（5 分鐘）

### 1. 添加項目

訪問 https://circleci.com/ 並用 GitHub 登入，然後添加項目：

```
Projects → 找到 my-gin-example → Set Up Project → Use Existing Config → master
```

### 2. 修復 SSH 錯誤

**錯誤**: `git@github.com: Permission denied (publickey)`

**解決方案**（選擇一種）：

#### 方法 A: Deploy Key（推薦）

1. 訪問 https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
2. 複製 "Deploy Key" 區域的公鑰（完整的 `ssh-ed25519 AAAA...`）
3. 訪問 https://github.com/wmh/my-gin-example/settings/keys
4. 點擊 "Add deploy key"
   - Title: `CircleCI`
   - Key: 粘貼公鑰
   - 不勾選 "Allow write access"
5. 重新運行 Pipeline

#### 方法 B: User Key（快速）

1. 訪問 https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
2. 點擊 "Add User Key" → 授權 GitHub
3. 重新運行 Pipeline

**推薦使用 Deploy Key**（更安全，只能訪問單個倉庫）

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
A: Deploy Key 更安全（僅限單倉庫），User Key 更快（訪問所有倉庫）

## 連結

- SSH Keys: https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
- Pipeline: https://app.circleci.com/pipelines/github/wmh/my-gin-example
- GitHub Deploy Keys: https://github.com/wmh/my-gin-example/settings/keys
- 文檔: https://circleci.com/docs/
