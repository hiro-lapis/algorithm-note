# Priority Queue

Priority Queue は配列にアルゴリズムに即した実装を加えたデータ構造のこと  

## About queue
そもそも Queue とは FIFO(first in first out) といって、  
最初に入った要素が最初に処理されるようなデータ構造になっている  
逆にstack は LIFO(最後に入った要素が最初に処理される)である  

Priority Queue は、キューに優先度順の並び替えを付与したデータ構造である

```
(3, 6, 9)
add(4)
(3, 6, 9, 4) // push append

(3, 6, 4, 9) // push append
(3, 4, 6, 9) // fixed position
```
通常のキューであれば末尾に要素を追加するのみだが、
Priority Queue は追加後に一定の基準に従い要素の並び替えが行われる  

これにより最短経路問題のダイクストラ法などで効率の良い探索が行える