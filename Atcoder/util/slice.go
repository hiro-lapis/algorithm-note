package util

import "sort"

/* make slice,grid */

// makeSliceI デフォルト値つきのINT sliceを作成
func makeSliceI(l int, def int) []int {
	slice := make([]int, l)
	for i := 0; i < l; i++ {
		slice[i] = def
	}
	return slice
}

// InsertValue スライスの特定のインデックスに新しい値を挿入する関数
func InsertValue(slice []int, index, value int) []int {
	// スライスを拡張する
	slice = append(slice, 0)

	// 指定した位置の要素以降をひとつずつ後ろにずらす
	copy(slice[index+1:], slice[index:])

	// 新しい値を指定した位置に挿入する
	slice[index] = value

	return slice
}

// makeGrid h*wマスのint多次元配列作成
func makeGrid(h, w int) [][]int {
	index := make([][]int, h)
	data := make([]int, h*w)
	for i := 0; i < h; i++ {
		index[i] = data[i*w : (i+1)*w]
	}
	return index
}

// CreateAlphabetMap a~zまでの英小文字をキーとし、int型の値を持つマップを作成する関数
func CreateAlphabetMap() map[string]int {
	alphabetMap := make(map[string]int)
	// a~zまでの英小文字を走査してマップに追加
	for i := 'a'; i <= 'z'; i++ {
		alphabetMap[string(i)] = 0
	}
	return alphabetMap
}

/* END make slice,grid */

/* Sort */

/* Sort */

func Sort(l []int, o string) []int {
	if o == "desc" {
		sort.Slice(l, func(i, j int) bool {
			return l[i] > l[j]
		})
	} else {
		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
	}
	return l
}

// Sorts sort string
func Sorts(sl []string) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
}

/* END Sort */

// Delete index指定のDelete
func Delete(a []int, i int) []int {
	a = append(a[:i], a[i+1:]...)
	return a
}

/* START Map */

// mapSortedKeys map INT値の昇順に参照するkey配列を取得 asc/desc指定必須
func mapSortedKeys(list map[int]int, asc bool) []int {
	keys := make([]int, 0, len(list))
	for i := range list {
		keys = append(keys, i)
	}

	if asc {
		sort.SliceStable(keys, func(i, j int) bool {
			return list[keys[i]] < list[keys[j]]
		})
	} else {
		sort.SliceStable(keys, func(i, j int) bool {
			return list[keys[i]] > list[keys[j]]
		})
	}
	return keys
}

/* END Map */
