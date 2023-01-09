[問題文](https://atcoder.jp/contests/code-festival-2016-qualb/tasks/codefestival_2016_qualB_b)

## 問題文の把握

- N 人の学生参加者
- 学生の情報は3種類あり文字列 S で表される
- S の i 文字目は、予選での順位を表す
- S の i 文字値は、学生の種類を表す
  - a 国内
  - b 海外
  - c どちらでもない
- 予選通過判定
  - a: 予選通過確定者が a+b 人未満なら通過
  - b: a の条件に加え、海外学生の中での順位がb 位以内なら通過
  - c: 通過できない
- 解答
  - S の文字1つづつ(各順位の学生)に対して、予選通過の有無を出力する

## 予選通過の判定について

「予選通過確定者が a+b 人未満なら通過」  
解答例を見てみると、、、  
A は「国内参加者の予選通過枠数」  
B は「国内・海外参加者の予選通過枠数」  
であるということがわかる  

となると、  
- A + B は予選通過者の上限である
- 順位(何文字めか)が影響するのは海外参加のみ
ということがわかる

## ロジック
- 予選通過枠上限(A+B 以下、limit), 現在の通過者数(以下、count) の変数を定義
- 次に現れる海外学生の順位管理変数(b-rank 初期値 1)を定義
- ループで学生のN回判定を行う()
  - c なら `No`
  - a なら
    - C が limit 未満なら`OK` C をincrement
    - それ以外は `NO`
  - b なら
    - a と同じ条件判定 かつ b-rank が B 以内なら `Yes` し、C をincrement
    - b-rank をincrement

文字数文だけループを行うため、time complexity は O(n) である
## pseudo code

```
int limit = A + B
int C
int bRank = 1

for i=0 to S
    s = S[i]
    if c
        print 'No'
    else if a
        if C < limit
            print 'Yes'
            C++
    else if b
        if C < limit & bn <= B
            print 'Yes'
            C++
        else
            print 'No'
        bn ++
```

## tips

S 文字列数 > a + b の場合、「既にcount が limit と同値になっていて、学生の種別に関係なく No と出力することになる」というケースが起こりうる
この場合、疑似コードの for 文直下に「count が limit と同値なら No と入力して次のループへ行く」という処理を追加することで処理を軽量化できる
