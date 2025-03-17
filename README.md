# Slot Tree

## 1. 專案介紹
`slot_tree` 是一個用於計算 slot machine (老虎機) 組合得分的 Golang 專案。該專案透過 `payLine` 建立樹狀結構，並使用 **深度優先搜尋 (DFS, Depth First Search)** 來計算得分。

## 2. 主要功能
- **建立 Slot Tree**: 根據 `payLine` 構建樹狀結構，儲存可能的遊戲組合。
- **深度優先搜尋 (DFS)**: 遍歷 `slot_tree`，計算所有可能組合的得分。
- **計算得分 (Scoring)**: 依據遊戲規則計算符合條件的路徑總分。

## 3. 系統架構
### 3.1 樹狀結構 (Slot Tree)
每個節點 (Node) 代表一個 slot 機轉輪 (Reel) 的可能結果，透過 `payLine` 來決定連接關係。

範例結構 (假設有 3 個轉輪，每個轉輪有 3 種可能的結果)：
```
         Root
        /  |  \
      A1  A2  A3  (第一個轉輪)
     / | \ / | \ / | \
   B1 B2 B3 B1 B2 B3 B1 B2 B3  (第二個轉輪)
  /|\ /|\ /|\ /|\ /|\ /|\ /|\ /|\
 C1 C2 C3 ... C1 C2 C3 ... C1 C2 C3  (第三個轉輪)
```

### 3.2 深度優先搜尋 (DFS)
1. **從 Root 節點開始遍歷**
2. **依據 `payLine` 選擇可行路徑**
3. **抵達葉節點 (Leaf Node) 時計算對應得分**
4. **回溯並累積所有可能組合的得分**

## 4. 安裝與使用方式
### 4.1 環境需求
- Golang 1.20+
- 安裝 `go mod` 來管理依賴

### 4.2 安裝與執行
1. 下載專案
   ```sh
   git clone https://github.com/ryanChen100/slot_tree.git
   cd slot_tree
   ```

2. 安裝依賴
   ```sh
   go mod tidy
   ```

3. 執行 Slot Tree 計算
   ```sh
   go run main.go
   ```

### 4.3 測試
執行單元測試:
```sh
go test ./...
```

## 5. 主要程式架構
```plaintext
slot_tree/
│── main.go          # 入口程式，執行 slot tree 計算
│── tree.go          # 樹狀結構的建構與遍歷 (DFS)
```

## 6. 參數設定 (Config)
- **Slot 組合數據**: 在 `tree.go` 定義 slot machine 可能的組合。

## 7. 貢獻方式
歡迎 Pull Request 或 Issue 討論改進建議。

## 8. 授權
MIT License

