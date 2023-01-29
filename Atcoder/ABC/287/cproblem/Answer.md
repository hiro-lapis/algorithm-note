[問題文](https://atcoder.jp/contests/abc287/tasks/abc287_c)

## 問題文の把握

- N vertex M edge の単純方向グラフがある
- vertexには、1,2,...M の番号が振られている
- edge i は vertex ui, vi を結んでいる
- この graph が path graph であるか判定せよ

- path graph: 
  - cycle (循環)になってない
  - vertex 全てが edge で繋がっている


## ロジック



### pseudo code


```
var n, m

if n-1 > m
    print("No")
    return
```

### tips

