[問題文](https://atcoder.jp/contests/abc307/tasks/abc307_c)

#文字列操作

## 問題文の把握

- 黒いマスと透明なマスからなるシートA, Bがある
- 透明なマスのみのシートCもある
- 黒いマスと透明なマスからなる理想のシートXがある
  - シートAの大きさはそれぞれ縦HA 横WA
  - シートBの大きさはそれぞれ縦HB 横WB
  - シートXの大きさはそれぞれ縦HX 横WX
- マスには `#` `.` が入っている。文字朝WAの文字列が、HAだけあるということ

- 目標として、シートABCから、ABに存在する黒いマス全てを使って、シートXを作成することが目標
  - シートA,Bをマス目に沿って貼り付け、Cに貼り付け。この際、A Bは好きな場所にはっていいが、切り分けたり回転させてはいけない
  - シートCから任意の位置でHX　WXマス切り出す。この際、元となるシートABどちらかで `#` になっていたマスは黒いマスになる
- 上記作業を行い、シートXを作成することができるかを応えよ
- シートXと比較の際は、シートCを回転させてはいけない 

制約
- 1 <= HW <= 10
- A,B,X それぞれマスには `#` `.`のみ
- A,B,X それぞれ裁定１マス黒いマスを含む

```
case1
3 5
#.#..
.....
.#...
2 2
#.
.#
5 3
...
#.#
.#.
.#.
...
=> Yes
↓のようにして、作成可能
...
#.#
.#.
.#.
...

case2
2 2
#.
.#
2 2
#.
.#
2 2
##
##
=> No

case3
1 1
#
1 2
##
1 1
#
=> No

case4
3 3
###
...
...
3 3
#..
#..
#..
3 3
..#
..#
###
=> Yes

```

## ロジック

### caseから考える


case3 から、シートABを元に作成されるシートCの黒いマスの数は、A,Bどちらか黒いマスの数の多い方と同じ数になる  
よって、シートXの黒いマスの数が、A,Bどちらか黒いマスの数の多い方より少ない時は絶対に　No になる

### pseudo code


```
```

### tips

