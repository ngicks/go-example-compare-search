package goexamplecomparesearch_test

import (
	"cmp"
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"fmt"
	"iter"
	"maps"
	"math"
	mathRand "math/rand/v2"
	"slices"
	"testing"

	"github.com/gammazero/deque"
	"github.com/ngicks/go-iterator-helper/collection"
	"github.com/ngicks/go-iterator-helper/hiter"
	"github.com/ngicks/go-iterator-helper/x/exp/xiter"
	"rsc.io/omap"
)

func linearSearch[S ~[]E, E, T any](s S, tgt T, cmp func(e E, t T) int) int {
	for i, e := range s {
		if cmp(e, tgt) == 0 {
			return i
		}
	}
	return -1
}

func sizes() iter.Seq[int] {
	return xiter.Merge(
		xiter.Map(func(i int) int { return i * 5 }, hiter.Range(1, 5)),
		xiter.Map(func(i int) int { return 1 << i }, hiter.Range(1, 12)),
	)
}

func randStr() iter.Seq[string] {
	return hiter.RepeatFunc(func() string {
		var buf [4]byte
		_, err := rand.Read(buf[:])
		if err != nil {
			panic(err)
		}
		return hex.EncodeToString(buf[:])
	}, -1)
}

func Benchmark_linear_search(b *testing.B) {
	for i := range sizes() {
		targetSlice := slices.Sorted(xiter.Limit(randStr(), i))
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			b.Run("average", func(b *testing.B) {
				for j := range b.N {
					_ = linearSearch(targetSlice, targetSlice[j%len(targetSlice)], cmp.Compare)
				}
			})
			b.Run("worst", func(b *testing.B) {
				for range b.N {
					_ = linearSearch(targetSlice, targetSlice[len(targetSlice)-1], cmp.Compare)
				}
			})
		})
	}
}

func Benchmark_binary_search(b *testing.B) {
	for i := range sizes() {
		targetSlice := slices.Sorted(xiter.Limit(randStr(), i))
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			for j := range b.N {
				_, _ = slices.BinarySearchFunc(targetSlice, targetSlice[j%len(targetSlice)], cmp.Compare)
			}
		})
	}
}

func Benchmark_map_lookup(b *testing.B) {
	for i := range sizes() {
		kv := hiter.Collect2(hiter.Pairs(xiter.Limit(randStr(), i), xiter.Limit(randStr(), i)))
		targetMap := maps.Collect(hiter.KeyValues[string, string](kv).Iter2())
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			for j := range b.N {
				_, ok := targetMap[kv[j%len(targetMap)].K]
				if !ok {
					panic("eh")
				}
			}
		})
	}
}

func Benchmark_omap_lookup(b *testing.B) {
	for i := range sizes() {
		kv := hiter.Collect2(hiter.Pairs(xiter.Limit(randStr(), i), xiter.Limit(randStr(), i)))
		targetMap := omap.Map[string, string]{}
		for _, kv := range kv {
			targetMap.Set(kv.K, kv.V)
		}
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			for j := range b.N {
				_, ok := targetMap.Get(kv[j%len(kv)].K)
				if !ok {
					panic("eh")
				}
			}
		})
	}
}

type Range struct {
	Start, End int
}

func ranges(i int) iter.Seq[Range] {
	return xiter.Map(
		func(i int) Range {
			return Range{Start: i * 500, End: i*500 + 500}
		},
		hiter.Range(0, i),
	)
}

func compareRange(r Range, t int) int {
	switch {
	case t < r.Start:
		return 1
	case r.Start <= t && t < r.End:
		return 0
	default: // r.End <= off:
		return -1
	}
}

func Benchmark_find_range_linear_search(b *testing.B) {
	for i := range sizes() {
		bottom := mathRand.N(i)
		targetSlice := slices.Collect(ranges(i))
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			for j := range b.N {
				tgt := (500 * ((j + bottom) % len(targetSlice))) + 1
				idx := linearSearch(targetSlice, tgt, compareRange)
				if idx < 0 {
					panic("eh")
				}
			}
		})
	}
}

func Benchmark_find_range_binary_search(b *testing.B) {
	for i := range sizes() {
		bottom := mathRand.N(i)
		targetSlice := slices.Collect(ranges(i))
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			for j := range b.N {
				tgt := (500 * ((j + bottom) % len(targetSlice))) + 1
				_, ok := slices.BinarySearchFunc(targetSlice, tgt, compareRange)
				if !ok {
					panic("eh")
				}
			}
		})
	}
}

func Benchmark_find_range_omap(b *testing.B) {
	for i := range sizes() {
		bottom := mathRand.N(i)
		cont := hiter.KeyValues[int, Range]{{}}
		targetMap := omap.Map[int, Range]{}
		for r := range ranges(i) {
			targetMap.Set(r.End, r)
		}
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			for j := range b.N {
				tgt := (500 * ((j + bottom) % i)) + 1
				kv := hiter.AppendSeq2(cont[:0], xiter.Limit2(targetMap.Scan(tgt, math.MaxInt), 1))
				if len(kv) != 1 {
					panic("not found")
				}
			}
		})
	}
}

func Benchmark_deque_merge_sort_slice_conversion(b *testing.B) {
	rng := hiter.RepeatFunc(func() int { return mathRand.N[int](1000) }, -1)
	randNumArray := slices.Collect(xiter.Limit(rng, 2048))

	b.Run("[]int", func(b *testing.B) {
		b.ResetTimer()
		deque_merge_sort(b, randNumArray[:], cmp.Compare)
	})

	type bigStruct struct {
		Key string
		Mah [2048]byte
	}
	bigStructs := make([]bigStruct, 2048)
	for i := range 2048 {
		bigStructs[i] = bigStruct{
			Key: hiter.StringsCollect(4*8*2, xiter.Limit(randStr(), 8)),
		}
	}

	b.Run("[]bigStruct", func(b *testing.B) {
		b.ResetTimer()
		deque_merge_sort(b, bigStructs, func(i, j bigStruct) int { return cmp.Compare(i.Key, j.Key) })
	})

	starBigStructs := make([]*bigStruct, 2048)
	for i := range 2048 {
		starBigStructs[i] = &bigStruct{
			Key: hiter.StringsCollect(4*8*2, xiter.Limit(randStr(), 8)),
		}
	}

	b.Run("[]*bigStruct", func(b *testing.B) {
		b.ResetTimer()
		deque_merge_sort(
			b,
			starBigStructs,
			func(i, j *bigStruct) int {
				if i == nil {
					return -1
				}
				return cmp.Compare(i.Key, j.Key)
			},
		)
	})
}

func deque_merge_sort[T any](b *testing.B, input []T, cmp func(l T, r T) int) {
	for i := range sizes() {
		d := deque.New[T]()
		for ele := range xiter.Limit(slices.Values(input), i) {
			d.PushBack(ele)
		}
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			b.Run("slice_version", func(b *testing.B) {
				b.ResetTimer()
				for range b.N {
					ok := slices.IsSortedFunc(mergeSortFunc(input[:i], cmp), cmp)
					if !ok {
						panic("eh")
					}
				}
			})
			b.Run("converted_to_slice", func(b *testing.B) {
				b.ResetTimer()
				s := make([]T, d.Len())
				for i := range d.Len() {
					s[i] = d.At(i)
				}
				for range b.N {
					ok := slices.IsSortedFunc(slices.Collect(collection.MergeSortFunc(s, cmp)), cmp)
					if !ok {
						panic("eh")
					}
				}
			})
			b.Run("converted_to_slice_no_collect", func(b *testing.B) {
				b.ResetTimer()
				s := make([]T, d.Len())
				for i := range d.Len() {
					s[i] = d.At(i)
				}
				for range b.N {
					var prev T
					for n := range collection.MergeSortFunc(s, cmp) {
						if cmp(prev, n) > 0 {
							panic("oh?")
						}
						prev = n
					}
				}
			})
			b.Run("no_conversion", func(b *testing.B) {
				b.ResetTimer()
				for range b.N {
					ok := slices.IsSortedFunc(slices.Collect(collection.MergeSortSliceLikeFunc(d, cmp)), cmp)
					if !ok {
						panic("eh")
					}
				}
			})
		})
	}
}

func mergeSortFunc[S ~[]T, T any](m S, cmp func(l, r T) int) S {
	if len(m) <= 1 {
		return m
	}
	left, right := m[:len(m)/2], m[len(m)/2:]
	left = mergeSortFunc(left, cmp)
	right = mergeSortFunc(right, cmp)
	return mergeFunc(left, right, cmp)
}

func mergeFunc[S ~[]T, T any](l, r S, cmp func(l, r T) int) S {
	m := make(S, len(l)+len(r))
	var i int
	for i = 0; len(l) > 0 && len(r) > 0; i++ {
		if cmp(l[0], r[0]) < 0 {
			m[i] = l[0]
			l = l[1:]
		} else {
			m[i] = r[0]
			r = r[1:]
		}
	}
	for _, t := range l {
		m[i] = t
		i++
	}
	for _, t := range r {
		m[i] = t
		i++
	}
	return m
}
