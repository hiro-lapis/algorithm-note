[問題文](https://atcoder.jp/contests/abc122/tasks/abc122_b)

## 問題文の把握

- 英大文字で構成される文字列 S が与えられる
- S の部分文字列であるような最も長い ACGT 文字列を求めよ
- ACGT 文字列: A,C,G,T 以外を含まない文字列

＊ 先頭と末尾から 0 文字以上を取り去って得られる文字列  

```
ATCODER の部分文字列
=> TCO, AT, CODER, ATCODER, (空文字列) が含まれ、AC は含まれない
```

- 制約
- S は長さ 1~10

## ロジック

部分文字列の定義からみるに、文字列をアレコレ改変する必要はなさそう  
`先頭と後ろから 0 文字以上を取り去って` のロジックは、ループで解決できる  

- 文字の先頭から i 文字取り去る
- 文字の後ろから j 文字取り去る

文字数もMAX10文字と少ないため、上記両方で文字を切り出して、ACFT文字になりうる場合は文字数をカウント  
文字数の最大値を更新する、という処理を全探索で処理する

### pseudo code

```
var s 
var n s.length

max = 0 
for i to n
    for j to n - i
        partial = s[i:j] // 先頭からi文字、末尾からj文字目を切り出し
        count = count ACGT
        if count > max
            max = count 
print(max)

function count ACGT
    count = 0
    for each letter of s
        if  letter contained ACGT
            count++
        else  
            count--
    return count
```

### tips

解答時間: 25m

time complexity はfactorial O(n!)  
先頭から1文字切り出すごとに、末尾から切り出せる文字の組み合わせが減るため   

・文字列を for range で切り出すと、ASCII code として出力される    
文字列として出力したい場合は `string(s)` で cast できる   
できたら cast なしでロジック組めるようになりたひ  