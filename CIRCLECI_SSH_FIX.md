# 🔧 CircleCI SSH Permission Denied 快速修復指南

## ⚠️ 問題描述

如果你看到以下錯誤：

```
git@github.com: Permission denied (publickey).
fatal: Could not read from remote repository.
```

這表示 CircleCI 沒有權限克隆你的 GitHub 倉庫。

---

## ✅ 快速修復（5 分鐘）

### 方法 1: 添加 User Key（推薦）⭐

這是最簡單快速的方法：

#### 步驟：

1. **打開 SSH Keys 設置頁面**
   
   🔗 https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
   
   或者：
   - 進入你的項目
   - 點擊右上角 "Project Settings"
   - 左側菜單選擇 "SSH Keys"

2. **添加 User Key**
   
   - 滾動到頁面的 "User Key" 區域
   - 點擊綠色按鈕 **"Add User Key"**

3. **授權 GitHub**
   
   - 會彈出 GitHub 授權視窗
   - 點擊 **"Authorize with GitHub"**
   - 如果需要，輸入你的 GitHub 密碼確認

4. **確認添加成功**
   
   - 應該看到 "User key added" 的成功訊息
   - 頁面會顯示一個 fingerprint

5. **重新運行 Pipeline**
   
   - 返回 Pipeline 頁面
   - 點擊 **"Rerun workflow from failed"**
   - 或點擊 **"Rerun from start"**

6. **驗證成功** ✅
   
   在構建日誌中應該看到：
   ```
   ✓ Cloning git repository
   ✓ Cloning into '.'...
   ✓ Receiving objects: 100%
   ```

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

**最快的修復方法：**

1. 訪問 https://app.circleci.com/settings/project/github/wmh/my-gin-example/ssh
2. 點擊 "Add User Key"
3. 授權 GitHub
4. 重新運行 Pipeline
5. ✅ 完成！

**預計時間：2-3 分鐘**

---

修復完成後，CircleCI 會自動運行測試、Lint 和安全掃描。🚀

如需更多幫助，請查看 `CIRCLECI_SETUP.md` 獲取完整設置指南。
