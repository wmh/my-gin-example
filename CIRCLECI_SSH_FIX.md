# 🔧 CircleCI SSH Permission Denied 快速修復指南

## ⚠️ 問題描述

如果你看到以下錯誤：

```
git@github.com: Permission denied (publickey).
fatal: Could not read from remote repository.
```

這表示 CircleCI 沒有權限克隆你的 GitHub 倉庫。

---

## ✅ 快速修復（3 分鐘）

### 方法 1: 使用 Deploy Key（CircleCI 推薦）⭐⭐⭐

這是 CircleCI 官方推薦的最佳方式：

#### 什麼是 Deploy Key？

Deploy Key 是**倉庫專用的 SSH 密鑰**：
- CircleCI 已經為你生成了密鑰對
- 公鑰需要添加到 GitHub
- 私鑰已安全存儲在 CircleCI
- **更安全** - 只能訪問單個倉庫

#### 為什麼選擇 Deploy Key？

✅ CircleCI 官方推薦
✅ 更安全 - 限於單個倉庫
✅ 符合最佳實踐
✅ 適合生產環境

#### 步驟：

**步驟 1: 從 CircleCI 複製公鑰**

1. 訪問 SSH Keys 設置頁面：
   
   🔗 https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh

2. 找到 **"Deploy Key"** 區域（通常在頁面頂部）

3. 你會看到：
   ```
   Deploy Key
   ───────────────────────────────────────────
   A deploy key is a repo-specific SSH key.
   
   Public key:
   ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAA...
   ```

4. **複製完整的公鑰**：
   - 點擊公鑰旁邊的 📋 複製按鈕
   - 或手動選擇並複製（從 `ssh-ed25519` 到結尾）

**步驟 2: 添加到 GitHub**

1. 訪問 GitHub 倉庫的 Deploy keys 設置：
   
   🔗 https://github.com/wmh/my-gin-example/settings/keys

2. 點擊綠色按鈕 **"Add deploy key"**

3. 填寫表單：
   - **Title**: `CircleCI Deploy Key`
   - **Key**: 粘貼剛才複製的完整公鑰
   - **Allow write access**: 
     - ⬜ 只需要 pull（運行測試）：不勾選
     - ✅ 需要 push（自動部署）：勾選

4. 點擊 **"Add key"** 保存

**步驟 3: 驗證設置**

1. 返回 CircleCI Pipeline：
   
   🔗 https://app.circleci.com/pipelines/github/wmh/my-gin-example

2. 點擊 **"Rerun workflow from failed"**

3. 查看構建日誌，應該看到：
   ```
   ✅ Cloning git repository
   ✅ Cloning into '.'...
   ✅ Receiving objects: 100%
   ```

---

### 方法 2: 添加 User Key（快速但不推薦）

⚠️ 注意：User Key 可以訪問你的所有倉庫，安全性較低。

如果你想快速測試（不推薦用於生產）：

#### 步驟：

1. 訪問：https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh

2. 滾動到 **"User Key"** 區域

3. 點擊 **"Add User Key"**

4. 授權 GitHub

5. 重新運行 Pipeline

---

## 🔗 直接連結

| 操作 | 連結 |
|------|------|
| SSH Keys 設置 | https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh |
| Pipeline 列表 | https://app.circleci.com/pipelines/github/wmh/my-gin-example |
| 項目設置 | https://app.circleci.com/settings/project/github/wmh/my-gin-example |

---

## 🎥 步驟截圖說明

### 1️⃣ 找到 SSH Keys 頁面

```
Project Settings (右上角)
    ↓
左側選單: SSH Keys
    ↓
User Key 區域
```

### 2️⃣ User Key 區域應該顯示

```
┌─────────────────────────────────────┐
│ User Key                            │
│                                     │
│ No user key                         │
│                                     │
│ [Add User Key]  ← 點擊這個按鈕        │
└─────────────────────────────────────┘
```

### 3️⃣ 授權後應該顯示

```
┌─────────────────────────────────────┐
│ User Key                            │
│                                     │
│ ✓ User key added                    │
│ Fingerprint: aa:bb:cc:dd:...        │
└─────────────────────────────────────┘
```

---

## 🔄 替代方法

### 方法 2: 檢查 GitHub App 授權

如果方法 1 不起作用，檢查 GitHub App：

1. 訪問 GitHub 設置：https://github.com/settings/installations

2. 找到 "CircleCI" 應用

3. 點擊 "Configure"

4. 確保 "wmh/my-gin-example" 在授權列表中

5. 保存後返回 CircleCI 重新運行

---

### 方法 3: 手動 Deploy Key

如果仍然失敗，使用手動 Deploy Key：

#### 在本地生成 SSH Key：

```bash
ssh-keygen -t ed25519 -C "circleci-my-gin-example" -f ~/.ssh/circleci_key
# 不要設置密碼，直接按 Enter
```

#### 添加公鑰到 GitHub：

1. 複製公鑰：
   ```bash
   cat ~/.ssh/circleci_key.pub
   ```

2. 訪問：https://github.com/wmh/my-gin-example/settings/keys

3. 點擊 "Add deploy key"

4. Title: `CircleCI Deploy Key`

5. 粘貼公鑰內容

6. ✅ 勾選 "Allow write access"

7. 點擊 "Add key"

#### 添加私鑰到 CircleCI：

1. 複製私鑰：
   ```bash
   cat ~/.ssh/circleci_key
   ```

2. 訪問：https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh

3. 在 "Additional SSH Keys" 區域點擊 "Add SSH Key"

4. 填寫：
   - Hostname: `github.com`
   - Private Key: 粘貼私鑰內容

5. 點擊 "Add SSH Key"

6. 重新運行 Pipeline

---

## ✅ 驗證修復成功

修復成功後，在 CircleCI 構建日誌中應該看到：

```
✓ Cloning git repository - git@github.com:wmh/my-gin-example.git
  - Creating a blobless clone for better performance
✓ Cloning into '.'...
✓ Receiving objects: 100% (XXX/XXX), done.
✓ Resolving deltas: 100% (XXX/XXX), done.
```

而不是：

```
✗ git@github.com: Permission denied (publickey).
✗ fatal: Could not read from remote repository.
```

---

## 🎯 我應該使用哪個方法？

| 方法 | 適用場景 | 難度 | 推薦度 |
|------|---------|------|--------|
| User Key | 個人項目、快速設置 | ⭐ | ⭐⭐⭐⭐⭐ |
| GitHub App | 組織項目、多倉庫 | ⭐⭐ | ⭐⭐⭐⭐ |
| Deploy Key | 需要細粒度控制 | ⭐⭐⭐ | ⭐⭐⭐ |

**建議：先試方法 1（User Key），最簡單快速！**

---

## 💡 常見問題

### Q: User Key 和 Deploy Key 有什麼區別？

**User Key**:
- 使用你的個人 GitHub 權限
- 可以訪問你有權限的所有倉庫
- 設置最簡單

**Deploy Key**:
- 只能訪問特定倉庫
- 更安全，權限範圍小
- 設置稍複雜

### Q: 為什麼 CircleCI 需要 SSH Key？

CircleCI 使用 SSH 協議從 GitHub 克隆代碼。就像你在本地使用 `git clone` 需要認證一樣，CircleCI 也需要認證才能訪問你的倉庫。

### Q: 這是安全的嗎？

是的。CircleCI 是 GitHub 官方認證的 CI/CD 工具，添加 User Key 只是授權 CircleCI 代表你訪問倉庫。

### Q: 如果我刪除了 SSH Key 會怎樣？

刪除後，CircleCI 將無法克隆倉庫，構建會失敗。你需要重新添加 Key。

---

## 🔍 檢查清單

設置完成後，確認：

- [ ] 在 SSH Keys 頁面看到 User Key 或 Deploy Key
- [ ] 重新運行了失敗的 Pipeline
- [ ] 構建日誌顯示成功克隆倉庫
- [ ] 沒有 "Permission denied" 錯誤
- [ ] 所有 jobs (build-and-test, lint, security-scan) 都開始運行

---

## 🆘 還是不行？

如果按照上述步驟仍然失敗：

1. **檢查網絡連接**
   - CircleCI 能否訪問 GitHub？
   - 是否有防火牆限制？

2. **檢查倉庫權限**
   - 你的 GitHub 帳戶是否有倉庫訪問權限？
   - 倉庫是否是私有的？

3. **清除並重試**
   - 刪除現有的 SSH Keys
   - 重新添加 User Key
   - 重新運行 Pipeline

4. **查看 CircleCI 支援**
   - 文檔：https://circleci.com/docs/github-integration/
   - 社區論壇：https://discuss.circleci.com/

5. **聯繫我**
   - 提供錯誤日誌
   - 說明嘗試過的步驟

---

## 📌 總結

**推薦方法（Deploy Key）：**

1. 訪問 CircleCI：https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
2. 複製 Deploy Key 公鑰
3. 訪問 GitHub：https://github.com/wmh/my-gin-example/settings/keys
4. 添加 Deploy key
5. 重新運行 Pipeline
6. ✅ 完成！

**預計時間：3 分鐘**

**為什麼不用 User Key？**
- User Key 可以訪問所有倉庫（安全風險）
- CircleCI 官方推薦使用 Deploy Key
- Deploy Key 更符合最佳實踐

---

修復完成後，CircleCI 會自動運行測試、Lint 和安全掃描。🚀

如需更多幫助，請查看 `CIRCLECI_SETUP.md` 獲取完整設置指南。
