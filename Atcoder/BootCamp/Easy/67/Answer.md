[問題文](https://atcoder.jp/contests/abc155/tasks/abc155_c)

#sort
#map

## 問題文の把握

- n 枚の投票用紙がある
- 書かれた投票用紙を、投票数DESC, 辞書順ASCで出力せよ


- 制約
- 1 <= n <= 2 * 10^5
- 1 <= |Si| <= 10
- Si は英小文字のみからなる

```
case1
7
beat
vet
beet
bed
vet
bet
beet

=> beet vet
```

## ロジック

やることは至って単純
入力を受け取り、それを投票数順にソートし、数が同値の場合は辞書順にソートするだけ

ただし、データ構造を考慮する場合にいくつか手順をふむ必要がある

名前を受け取りつつ投票数をカウントする場合、map<string, int>のデータ構造が望ましい
keyに名前、valueに投票数を格納することで、名前を受け取りつつ投票数をカウントすることができる

しかし、mapは順序を保持しない(iteration order is undefined)ので、
投票数順にソートするためには順序を保持できるsliceに変換する

slice に変換した後に投票数順、辞書順にソートするために
名前と投票数を保持する構造体を定義し、それをsliceに格納する

- type kv (key, value) を作成
- kv を要素に持つslice のカスタム型を作成
  - [Less などのメソッドを実装](https://chat.openai.com/share/bff6ee04-3303-4386-8c40-e788e6f0aab9)しsort.Sortできるようにしておく
- map をループで抽出し、kv を作成し、slice に格納
- sort.Sort でINT ASCでソート
- sort.String で辞書順 ASCでソート

### tips

time complexity は O(n log n) 

