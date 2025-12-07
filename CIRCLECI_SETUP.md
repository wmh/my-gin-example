# CircleCI 項目設置指南

## 📋 項目信息
- **Repository**: wmh/my-gin-example
- **GitHub URL**: https://github.com/wmh/my-gin-example
- **配置文件**: `.circleci/config.yml` ✅ 已存在

## 🚀 CircleCI 設置步驟

### 步驟 1: 訪問 CircleCI 並登入
1. 打開瀏覽器訪問: https://circleci.com/
2. 點擊 "Sign Up" 或 "Log In"
3. 選擇 "Log in with GitHub"
4. 授權 CircleCI 訪問你的 GitHub 帳戶

### 步驟 2: 添加項目
1. 登入後，點擊左側選單的 "Projects"
2. 在項目列表中找到 `my-gin-example`
3. 點擊 "Set Up Project" 按鈕

### 步驟 3: 配置項目
CircleCI 會自動檢測到 `.circleci/config.yml` 文件

**選項 A - 使用現有配置（推薦）**:
1. 選擇 "Use Existing Config"
2. 選擇分支: `master`
3. 點擊 "Set Up Project"

**選項 B - 手動配置**:
1. 選擇 "Fastest: Use the .circleci/config.yml in my repo"
2. 點擊 "Set Up Project"

### 步驟 4: 觸發首次構建
1. CircleCI 會自動觸發首次構建
2. 或者手動觸發: 點擊 "Trigger Pipeline"
3. 選擇 `master` 分支

### 步驟 5: 配置 GitHub App 授權（解決 SSH 問題）

**當前問題**: `git@github.com: Permission denied (publickey)`

**推薦解決方案 - 使用 GitHub App**:

1. 訪問 CircleCI 項目設置:
   ```
   https://app.circleci.com/settings/project/github/wmh/my-gin-example
   ```

2. 點擊左側的 "Advanced"

3. 找到 "GitHub Permissions" 區域

4. 確認 CircleCI GitHub App 已安裝:
   - 如果未安裝，點擊 "Install CircleCI"
   - 選擇要授權的倉庫
   - 授權 CircleCI 訪問

5. 或者使用 GitHub 整合設置:
   - 訪問: https://github.com/settings/installations
   - 找到 CircleCI
   - 點擊 "Configure"
   - 確保 `my-gin-example` 在授權列表中

### 步驟 6: 配置環境變量（如果需要）

如果你的應用需要環境變量:
1. 進入項目設置: Project Settings → Environment Variables
2. 添加需要的變量（當前項目使用默認配置，可能不需要）

## 📊 預期的 CircleCI 工作流

```
build-and-test (主要任務)
  ├── Checkout 代碼
  ├── 恢復緩存 (Go modules)
  ├── 下載依賴
  ├── 保存緩存
  ├── 運行測試
  └── 構建應用

lint (代碼檢查)
  ├── Checkout 代碼
  ├── 安裝 golangci-lint
  └── 運行 Linter

security-scan (安全掃描)
  ├── Checkout 代碼
  ├── 安裝 gosec
  ├── 運行安全掃描
  └── 保存報告
```

## 🔧 當前配置說明

`.circleci/config.yml` 包含:

### 功能特性:
- ✅ Go 1.23 環境
- ✅ 依賴緩存 (提升構建速度)
- ✅ 單元測試 (`go test -v ./...`)
- ✅ 應用構建
- ✅ 代碼 Lint (golangci-lint)
- ✅ 安全掃描 (gosec)
- ✅ 測試結果保存
- ✅ 安全報告存檔

### 工作流順序:
1. **build-and-test**: 必須首先執行且通過
2. **lint**: 依賴 build-and-test 完成
3. **security-scan**: 依賴 build-and-test 完成

## 🔍 驗證設置

設置完成後，檢查以下內容:

### 1. Pipeline 狀態
訪問: https://app.circleci.com/pipelines/github/wmh/my-gin-example

確認看到:
- ✅ 綠色勾選表示構建成功
- 🔵 藍色圓圈表示正在運行
- ❌ 紅色 X 表示構建失敗

### 2. 查看構建日誌
1. 點擊任一 Pipeline
2. 查看每個 Job 的執行日誌
3. 確認所有步驟成功執行

### 3. 檢查 Artifacts
1. 進入 security-scan job
2. 點擊 "ARTIFACTS" 標籤
3. 下載並查看 `gosec-report.json`

## 🐛 常見問題排查

### 問題 1: Permission denied (publickey)
**症狀**: 
```
git@github.com: Permission denied (publickey)
```

**解決方案**:
1. 確保 CircleCI GitHub App 已正確安裝和授權（見步驟 5）
2. 或在 Project Settings → SSH Keys 中添加 User Key:
   - 點擊 "Add User Key"
   - 選擇你的 GitHub 帳戶
   - 授權訪問

### 問題 2: Go modules 下載失敗
**症狀**:
```
go: module not found
```

**解決方案**:
確認 `go.mod` 和 `go.sum` 文件已提交到倉庫

### 問題 3: 測試失敗
**症狀**: Tests 步驟顯示紅色

**解決方案**:
1. 本地運行測試確認通過:
   ```bash
   go test -v ./...
   ```
2. 檢查是否需要環境變量或配置文件

### 問題 4: 緩存問題
**症狀**: 每次構建都重新下載依賴

**解決方案**:
檢查 `go.sum` 文件是否有變化，CircleCI 會根據 `go.sum` 的 checksum 決定是否使用緩存

## 📱 添加狀態徽章（可選）

設置成功後，可以在 README.md 中添加 CircleCI 狀態徽章:

```markdown
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/wmh/my-gin-example/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/wmh/my-gin-example/tree/master)
```

不同樣式:
```markdown
<!-- Shield 樣式 -->
[![CircleCI](https://circleci.com/gh/wmh/my-gin-example/tree/master.svg?style=shield)](https://circleci.com/gh/wmh/my-gin-example/tree/master)

<!-- SVG 樣式 -->
[![CircleCI](https://circleci.com/gh/wmh/my-gin-example/tree/master.svg?style=svg)](https://circleci.com/gh/wmh/my-gin-example/tree/master)
```

## 🔔 配置通知（可選）

### Email 通知
默認會發送構建失敗通知到你的 GitHub 郵箱

### Slack 通知
1. 進入 Project Settings → Notifications
2. 點擊 "Add Notification"
3. 選擇 "Slack"
4. 配置 Webhook URL

## 🎯 下一步

設置完成後的建議:

1. ✅ 確認首次構建成功
2. ⚪ 添加狀態徽章到 README
3. ⚪ 配置 Slack 或 Email 通知
4. ⚪ 設置分支保護規則（要求 CircleCI 通過）
5. ⚪ 探索自動部署選項

## 🔗 有用的連結

### CircleCI 資源
- **官方文檔**: https://circleci.com/docs/
- **Go 語言指南**: https://circleci.com/docs/language-go/
- **GitHub 整合**: https://circleci.com/docs/github-integration/
- **配置參考**: https://circleci.com/docs/configuration-reference/

### 項目連結
- **項目首頁**: https://app.circleci.com/projects/project-dashboard/github/wmh/
- **項目設置**: https://app.circleci.com/settings/project/github/wmh/my-gin-example
- **Pipeline**: https://app.circleci.com/pipelines/github/wmh/my-gin-example

## 💡 提示

- 每次 `git push` 到任何分支都會觸發 CircleCI 構建
- Pull Request 也會自動運行 CI 檢查
- 可以在 GitHub PR 頁面看到 CircleCI 檢查狀態
- 構建失敗會阻止 PR 合併（如果設置了分支保護）

---

**設置完成後，CircleCI 將自動為每次代碼提交運行測試和檢查！** 🚀
