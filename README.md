# go-example-compare-search

To answer my own question, I maintain this bench mark.

- Do trust this result too much
  - I've run this on my own desk top, while many tabs open on web browser.

By this benchmark, the conclusion should be like;

- Always use `map[K]V` to look up some value where possible.
- If `map[K]V` is not possible, like finding the interval from a slice of ranges where some specific value is in, use linear search if `n < 15`, otherwise use binary search.
- The merge sort is faster when `[]T` -> `[]T`, while `[]T` -> `iter.Seq[T]` impose less `B/op` because it does not collect result to a slice?

Oh is that correct?

```
# go version
go version go1.23.0 linux/amd64
# go test -benchmem -run=^$ -bench ^Benchmark_ github.com/ngicks/go-example-compare-search
goos: linux
goarch: amd64
pkg: github.com/ngicks/go-example-compare-search
cpu: AMD Ryzen 9 7900X 12-Core Processor
Benchmark_linear_search/05/average-24                   163089920                7.411 ns/op           0 B/op          0 allocs/op
Benchmark_linear_search/05/worst-24                     100000000               11.64 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/10/average-24                   89264334                13.29 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/10/worst-24                     52114400                23.23 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/15/average-24                   61601652                19.93 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/15/worst-24                     35459283                36.29 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/20/average-24                   45398718                26.10 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/20/worst-24                     26341130                46.24 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/32/average-24                   29225172                41.56 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/32/worst-24                     15365752                77.36 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/64/average-24                   15160077                81.67 ns/op            0 B/op          0 allocs/op
Benchmark_linear_search/64/worst-24                      7458610               150.0 ns/op             0 B/op          0 allocs/op
Benchmark_linear_search/128/average-24                   7580826               158.0 ns/op             0 B/op          0 allocs/op
Benchmark_linear_search/128/worst-24                     3998996               304.7 ns/op             0 B/op          0 allocs/op
Benchmark_linear_search/256/average-24                   3792145               312.5 ns/op             0 B/op          0 allocs/op
Benchmark_linear_search/256/worst-24                     1979226               611.2 ns/op             0 B/op          0 allocs/op
Benchmark_linear_search/512/average-24                   1893045               666.9 ns/op             0 B/op          0 allocs/op
Benchmark_linear_search/512/worst-24                     1000000              1208 ns/op               0 B/op          0 allocs/op
Benchmark_linear_search/1024/average-24                   870090              1275 ns/op               0 B/op          0 allocs/op
Benchmark_linear_search/1024/worst-24                     535987              2398 ns/op               0 B/op          0 allocs/op
Benchmark_linear_search/2048/average-24                   484886              2605 ns/op               0 B/op          0 allocs/op
Benchmark_linear_search/2048/worst-24                     246751              4593 ns/op               0 B/op          0 allocs/op
Benchmark_binary_search/05-24                           91856056                12.91 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/10-24                           77617227                16.11 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/15-24                           71351706                17.64 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/20-24                           61261255                19.61 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/32-24                           56280925                22.50 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/64-24                           39821739                25.82 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/128-24                          43539828                28.02 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/256-24                          39369121                30.82 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/512-24                          35089714                34.27 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/1024-24                         23959916                49.63 ns/op            0 B/op          0 allocs/op
Benchmark_binary_search/2048-24                         22539330                54.86 ns/op            0 B/op          0 allocs/op
Benchmark_map_lookup/05-24                              162232400                7.608 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/10-24                              194106210                6.379 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/15-24                              201200707                5.786 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/20-24                              196718079                6.123 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/32-24                              193016404                6.323 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/64-24                              189487412                6.449 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/128-24                             183538749                6.460 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/256-24                             179702038                6.648 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/512-24                             184442653                6.436 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/1024-24                            193396764                6.352 ns/op           0 B/op          0 allocs/op
Benchmark_map_lookup/2048-24                            178591278                6.895 ns/op           0 B/op          0 allocs/op
Benchmark_omap_lookup/05-24                             139542962                8.808 ns/op           0 B/op          0 allocs/op
Benchmark_omap_lookup/10-24                             82839219                14.93 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/15-24                             78329481                15.89 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/20-24                             63956792                17.52 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/32-24                             53782372                23.76 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/64-24                             46823437                26.83 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/128-24                            41089666                29.60 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/256-24                            28583588                44.57 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/512-24                            20126126                60.10 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/1024-24                           17926822                70.02 ns/op            0 B/op          0 allocs/op
Benchmark_omap_lookup/2048-24                           13843434                87.53 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_linear_search/05-24                402010849                2.811 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_linear_search/10-24                235565398                5.005 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_linear_search/15-24                172058935                7.142 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_linear_search/20-24                131924370                9.065 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_linear_search/32-24                80708805                14.76 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_linear_search/64-24                41623030                28.31 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_linear_search/128-24               22883349                49.23 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_linear_search/256-24               13968463                88.76 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_linear_search/512-24                7452919               160.8 ns/op             0 B/op          0 allocs/op
Benchmark_find_range_linear_search/1024-24               3706216               315.5 ns/op             0 B/op          0 allocs/op
Benchmark_find_range_linear_search/2048-24               1937893               631.0 ns/op             0 B/op          0 allocs/op
Benchmark_find_range_binary_search/05-24                187132701                6.490 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_binary_search/10-24                157734927                7.918 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_binary_search/15-24                142144263                8.398 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_binary_search/20-24                132693452                9.167 ns/op           0 B/op          0 allocs/op
Benchmark_find_range_binary_search/32-24                100000000               10.35 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_binary_search/64-24                96111488                13.52 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_binary_search/128-24               83026279                14.84 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_binary_search/256-24               71119357                15.99 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_binary_search/512-24               69895915                17.47 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_binary_search/1024-24              62711910                18.77 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_binary_search/2048-24              48656415                25.07 ns/op            0 B/op          0 allocs/op
Benchmark_find_range_omap/05-24                         21016641                56.10 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/10-24                         20373734                58.41 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/15-24                         18684675                62.20 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/20-24                         19301973                63.37 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/32-24                         17904736                61.68 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/64-24                         19428325                63.78 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/128-24                        17193960                63.13 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/256-24                        17947966                67.67 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/512-24                        16809176                73.11 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/1024-24                       13354203                78.98 ns/op           64 B/op          3 allocs/op
Benchmark_find_range_omap/2048-24                       14086114                83.86 ns/op           64 B/op          3 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/05/slice_version-24           11021620               105.0 ns/op           104 B/op          4 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/05/converted_to_slice-24        442659              2614 ns/op            2728 B/op         96 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/05/converted_to_slice_no_collect-24             427394              2553 ns/op            2608 B/op         92 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/05/no_conversion-24                             229683              4892 ns/op            3008 B/op        106 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/10/slice_version-24                            5028654               243.4 ns/op           288 B/op          9 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/10/converted_to_slice-24                        178975              6608 ns/op            5976 B/op        207 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/10/converted_to_slice_no_collect-24             187180              6488 ns/op            5728 B/op        202 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/10/no_conversion-24                             104522             11590 ns/op            6536 B/op        227 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/15/slice_version-24                            3059004               394.4 ns/op           488 B/op         14 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/15/converted_to_slice-24                         99855             12008 ns/op            9096 B/op        317 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/15/converted_to_slice_no_collect-24             105771             11617 ns/op            8848 B/op        312 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/15/no_conversion-24                              65438             17883 ns/op            9936 B/op        347 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/20/slice_version-24                            1956776               638.4 ns/op           736 B/op         19 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/20/converted_to_slice-24                         72378             15527 ns/op           12472 B/op        428 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/20/converted_to_slice_no_collect-24              82158             15382 ns/op           11968 B/op        422 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/20/no_conversion-24                              45027             25354 ns/op           13592 B/op        468 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/32/slice_version-24                            1213912              1001 ns/op            1280 B/op         31 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/32/converted_to_slice-24                         43376             28141 ns/op           19960 B/op        692 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/32/converted_to_slice_no_collect-24              42408             27670 ns/op           19456 B/op        686 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/32/no_conversion-24                              30598             40019 ns/op           21752 B/op        756 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/64/slice_version-24                             572827              2151 ns/op            3072 B/op         63 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/64/converted_to_slice-24                         20068             59629 ns/op           40441 B/op       1397 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/64/converted_to_slice_no_collect-24              20294             57783 ns/op           39425 B/op       1390 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/64/no_conversion-24                              14340             83475 ns/op           44025 B/op       1525 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/128/slice_version-24                            225879              5286 ns/op            7168 B/op        127 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/128/converted_to_slice-24                         8260            139420 ns/op           81401 B/op       2806 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/128/converted_to_slice_no_collect-24              8508            140004 ns/op           79362 B/op       2798 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/128/no_conversion-24                              6052            183823 ns/op           88571 B/op       3062 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/256/slice_version-24                            115543             10816 ns/op           16384 B/op        255 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/256/converted_to_slice-24                         4088            297456 ns/op          163329 B/op       5623 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/256/converted_to_slice_no_collect-24              4234            289409 ns/op          159236 B/op       5614 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/256/no_conversion-24                              2990            420488 ns/op          177678 B/op       6135 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/512/slice_version-24                             49156             25545 ns/op           36864 B/op        511 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/512/converted_to_slice-24                         1623            734082 ns/op          327171 B/op      11256 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/512/converted_to_slice_no_collect-24              1514            741425 ns/op          318988 B/op      11246 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/512/no_conversion-24                              1207           1006855 ns/op          355844 B/op      12280 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/1024/slice_version-24                            20854             56242 ns/op           81920 B/op       1023 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/1024/converted_to_slice-24                         782           1508165 ns/op          663737 B/op      22522 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/1024/converted_to_slice_no_collect-24              748           1567200 ns/op          638552 B/op      22510 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/1024/no_conversion-24                              733           1714410 ns/op          721175 B/op      24570 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/2048/slice_version-24                             8354            144641 ns/op          180225 B/op       2047 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/2048/converted_to_slice-24                         271           4159280 ns/op         1337583 B/op      45052 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/2048/converted_to_slice_no_collect-24              279           4342915 ns/op         1277542 B/op      45038 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]int/2048/no_conversion-24                              289           4186707 ns/op         1452215 B/op      49148 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/05/slice_version-24                       194048              5172 ns/op           27136 B/op          4 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/05/converted_to_slice-24                   61423             19231 ns/op           56016 B/op         96 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/05/converted_to_slice_no_collect-24        82056             14463 ns/op           23240 B/op         92 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/05/no_conversion-24                        42658             29224 ns/op           56296 B/op        106 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/10/slice_version-24                        78336             13912 ns/op           76032 B/op          9 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/10/converted_to_slice-24                   25798             49892 ns/op          123018 B/op        207 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/10/converted_to_slice_no_collect-24        33164             37420 ns/op           49281 B/op        202 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/10/no_conversion-24                        17048             69321 ns/op          123577 B/op        227 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/15/slice_version-24                        41626             25693 ns/op          136576 B/op         14 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/15/converted_to_slice-24                   17358             71478 ns/op          149059 B/op        317 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/15/converted_to_slice_no_collect-24        20768             56457 ns/op           75322 B/op        312 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/15/no_conversion-24                        10000            104410 ns/op          149897 B/op        347 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/20/slice_version-24                        35974             35323 ns/op          201216 B/op         19 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/20/converted_to_slice-24                   10000            108994 ns/op          257023 B/op        428 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/20/converted_to_slice_no_collect-24        14270             82457 ns/op          101364 B/op        422 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/20/no_conversion-24                         7146            149032 ns/op          258138 B/op        468 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/32/slice_version-24                        18688             71279 ns/op          382978 B/op         31 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/32/converted_to_slice-24                    6112            164219 ns/op          319526 B/op        692 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/32/converted_to_slice_no_collect-24         8144            142163 ns/op          163866 B/op        686 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/32/no_conversion-24                         5344            219282 ns/op          321306 B/op        756 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/64/slice_version-24                         6956            154992 ns/op          905220 B/op         63 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/64/converted_to_slice-24                    2856            378676 ns/op          650061 B/op       1397 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/64/converted_to_slice_no_collect-24         3180            315300 ns/op          330558 B/op       1390 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/64/no_conversion-24                         2410            534363 ns/op          653597 B/op       1525 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/128/slice_version-24                        2764            367666 ns/op         2080781 B/op        127 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/128/converted_to_slice-24                   1306            813273 ns/op         1311217 B/op       2806 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/128/converted_to_slice_no_collect-24        1695            652438 ns/op          663988 B/op       2798 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/128/no_conversion-24                        1228            943143 ns/op         1318177 B/op       3062 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/256/slice_version-24                        1342            852491 ns/op         4694041 B/op        255 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/256/converted_to_slice-24                    613           1893162 ns/op         2633866 B/op       5623 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/256/converted_to_slice_no_collect-24         753           1648610 ns/op         1331162 B/op       5614 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/256/no_conversion-24                         582           2420985 ns/op         2647333 B/op       6135 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/512/slice_version-24                         561           2110799 ns/op        10444846 B/op        511 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/512/converted_to_slice-24                    260           4830063 ns/op         5190920 B/op      11256 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/512/converted_to_slice_no_collect-24         298           4140319 ns/op         2667256 B/op      11246 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/512/no_conversion-24                         220           5666513 ns/op         5215528 B/op      12280 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/1024/slice_version-24                        258           4613372 ns/op        23003218 B/op       1023 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/1024/converted_to_slice-24                    96          12786406 ns/op        12602173 B/op      22522 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/1024/converted_to_slice_no_collect-24         93          11399285 ns/op         5352931 B/op      22510 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/1024/no_conversion-24                         81          14625242 ns/op        12637485 B/op      24570 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/2048/slice_version-24                        100          10706652 ns/op        50233485 B/op       2048 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/2048/converted_to_slice-24                    44          27355051 ns/op        27185636 B/op      45055 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/2048/converted_to_slice_no_collect-24         50          24393264 ns/op        10747742 B/op      45038 allocs/op
Benchmark_deque_merge_sort_slice_conversion/[]bigStruct/2048/no_conversion-24                         37          30531614 ns/op        27203482 B/op      49149 allocs/op
PASS
ok      github.com/ngicks/go-example-compare-search     258.078s
```
