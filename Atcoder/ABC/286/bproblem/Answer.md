[問題文](https://atcoder.jp/contests/abc286/tasks/abc286_b)

## 問題文の把握

- 文字数 N の文字 S が与えらえる
- 文字列中にある全ての "na" を "nya" に変えよ 

```
case1
4
naan
=> nyaan

case2
4
near
=> near

case3
8
national
=> nyational
```
## ロジック

どんな文字がくるか？case を考えてみる  

```
nnan => nnyan
naan => nyaan
anan => anyan
```

この辺りに配慮しつつ、ロジックを組んでいく  

- 回答用の文字列変数を作成
- 前の文字が n だったよフラグ変数を作成
- 1 文字づつ比較 & 解答用文字列に 1 文字づつ追加
  - A: フラグ ON の時, 文字が "a" だったら "y" を付け足し、フラグをOFF にする
  - B: "n" があったらフラグを ON 
  - C: 比較に使った文字を解答用文字列末尾に追加
- 解答を出力

### pseudo code

```
var n int
var s string

var nFlag bool
var answer string
for i to n
    char = s[i+1]
    if nFlag & char == "a"
        answer += "y"
        nFlag = false
    if char == "n"
        nFlag = true
    answer += char

print(answer)
```

### tips

- time complexity は O(n)