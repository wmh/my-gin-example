#!/bin/bash

# CircleCI 配置驗證腳本

set -e

echo "=========================================="
echo "  CircleCI 配置驗證"
echo "=========================================="
echo

# 檢查配置文件
echo "�� 檢查 CircleCI 配置文件..."
if [ -f ".circleci/config.yml" ]; then
    echo "✅ 配置文件存在: .circleci/config.yml"
else
    echo "❌ 配置文件不存在"
    exit 1
fi
echo

# 檢查 Go 環境
echo "🔧 檢查 Go 環境..."
if command -v go &> /dev/null; then
    GO_VERSION=$(go version)
    echo "✅ Go 已安裝: $GO_VERSION"
else
    echo "❌ Go 未安裝"
    exit 1
fi
echo

# 檢查依賴
echo "📦 檢查 Go modules..."
if [ -f "go.mod" ] && [ -f "go.sum" ]; then
    echo "✅ go.mod 和 go.sum 存在"
else
    echo "❌ Go modules 文件缺失"
    exit 1
fi
echo

# 運行測試（模擬 CircleCI）
echo "🧪 運行測試..."
if go test -v ./... > /dev/null 2>&1 ; then
    echo "✅ 所有測試通過"
else
    echo "❌ 測試失敗"
    exit 1
fi
echo

# 構建應用（模擬 CircleCI）
echo "🔨 構建應用..."
if go build -o bin/server . ; then
    echo "✅ 構建成功"
    rm -f bin/server
else
    echo "❌ 構建失敗"
    exit 1
fi
echo

echo "=========================================="
echo "  配置摘要"
echo "=========================================="
echo "Repository: wmh/my-gin-example"
echo "配置文件: .circleci/config.yml"
echo "Go 版本: $(go version | awk '{print $3}')"
echo

echo "=========================================="
echo "  下一步: CircleCI 網站設置"
echo "=========================================="
echo
echo "1. 訪問 CircleCI:"
echo "   https://circleci.com/"
echo
echo "2. 使用 GitHub 登入"
echo
echo "3. 添加項目:"
echo "   https://app.circleci.com/projects/project-dashboard/github/wmh/"
echo
echo "4. 選擇 'my-gin-example' 並點擊 'Set Up Project'"
echo
echo "5. 選擇 'Use Existing Config' 並選擇 master 分支"
echo
echo "6. 解決 SSH 問題:"
echo "   - 訪問: https://app.circleci.com/settings/project/github/wmh/my-gin-example"
echo "   - 點擊 'Advanced' → 'GitHub Permissions'"
echo "   - 確保 CircleCI GitHub App 已授權"
echo
echo "7. 查看構建狀態:"
echo "   https://app.circleci.com/pipelines/github/wmh/my-gin-example"
echo
echo "=========================================="
echo "  ✅ 本地驗證完成！"
echo "=========================================="
echo
echo "詳細設置指南請查看: CIRCLECI_SETUP.md"
echo
