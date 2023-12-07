package advent

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Day7(input string) (int, int) {
	tot := 0
	rankCard := func(hand string, p2 bool) int {
		counts := map[rune]int{}
		for _, c := range hand {
			counts[c]++
		}
		jacks := counts['C']
		counts['C'] = 0
		byCount := map[int]int{}
		for _, v := range counts {
			byCount[v]++
		}
		if byCount[5] == 1 {
			return 7
		}
		if byCount[4] == 1 {
			if p2 && jacks > 0 {
				return 7
			}
			return 6
		}
		if byCount[3] == 1 {
			// full house with jacks not possible
			if byCount[2] == 1 {
				return 5
			}
			// 3+1 is 4
			if p2 && jacks == 1 {
				return 6
			}
			// 3 +2 is 5
			if p2 && jacks == 2 {
				return 7
			}
			return 4
		}
		if byCount[2] == 2 {
			if p2 && jacks == 1 {
				return 5
			}
			return 3
		}
		if byCount[2] == 1 {
			if p2 && jacks == 1 {
				return 4
			}
			if p2 && jacks == 2 {
				return 6
			}
			if p2 && jacks == 3 {
				return 7
			}
			return 2
		}
		if p2 && jacks == 1 {
			return 2
		}
		if p2 && jacks == 2 {
			return 4
		}
		if p2 && jacks == 3 {
			return 6
		}
		if p2 && jacks >= 4 {
			return 7
		}
		return 1
	}
	type hand struct {
		cards  string
		bet    int
		score  int
		score2 int
	}
	hands := []*hand{}
	for _, l := range Lines(input) {
		parts := strings.Split(l, " ")
		cards := parts[0]
		bet, _ := strconv.Atoi(parts[1])
		// make alphabetical
		cards = strings.ReplaceAll(cards, "T", "B")
		cards = strings.ReplaceAll(cards, "J", "C")
		cards = strings.ReplaceAll(cards, "Q", "D")
		cards = strings.ReplaceAll(cards, "K", "E")
		cards = strings.ReplaceAll(cards, "A", "F")
		hands = append(hands, &hand{
			cards:  cards,
			bet:    bet,
			score:  rankCard(cards, false),
			score2: rankCard(cards, true),
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]
		if a.score != b.score {
			return a.score < b.score
		}
		for i := 0; i < 5; i++ {
			if a.cards[i] != b.cards[i] {
				return a.cards[i] < b.cards[i]
			}
		}
		panic("same?")
	})
	for i, h := range hands {
		tot += (i + 1) * h.bet
	}
	tot2 := 0
	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]
		if a.score2 != b.score2 {
			return a.score2 < b.score2
		}
		ac := strings.ReplaceAll(a.cards, "C", "1")
		bc := strings.ReplaceAll(b.cards, "C", "1")
		for i := 0; i < 5; i++ {
			if ac[i] != bc[i] {
				return ac[i] < bc[i]
			}
		}
		panic("same?")
	})
	for i, h := range hands {
		//fmt.Println(i, h.cards, h.score2)
		tot2 += (i + 1) * h.bet
	}
	return tot, tot2
}

func TestDay7_Ex1(t *testing.T) {
	a, b := Day7(exD7)
	assert.Equal(t, 6440, a)
	assert.Equal(t, 5905, b)
}

func TestDay7_Actual(t *testing.T) {
	a, b := Day7(inputD7)
	fmt.Printf("Day 7 Part 1: %d\n", a)
	fmt.Printf("Day 7 Part 2: %d\n", b)
}

const exD7 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

const inputD7 = `Q97J7 740
65KJ6 889
6664J 275
JKK44 856
TJ285 722
482JT 338
KQ7T7 949
5K89Q 977
KJ949 437
366T3 204
T5797 847
3K7K3 953
AKAQA 323
88728 241
27AJA 29
69466 132
QT589 378
388Q5 7
39333 625
556J5 936
72ATQ 610
3Q6QJ 962
657QQ 929
85478 25
8AA48 195
A5AAK 44
33232 510
77227 824
77748 502
AJAAA 945
A3AAA 30
AKJ29 721
4666T 724
78Q8Q 261
Q7A95 951
342TQ 796
K7K77 523
39K39 128
82888 782
J466A 70
25222 463
TT222 418
9T2T9 848
3JKKJ 832
2TJ2T 36
47777 840
494TK 332
77KTK 35
A9A9A 219
3ATT3 217
5J8JQ 763
AA224 144
82TT2 220
55J2J 623
96999 930
T9T9T 268
J6AAJ 756
22J2A 626
A6652 519
55JKK 225
6JJ56 257
6Q996 384
22J22 498
A4888 188
43K7J 309
627QA 904
55K8A 160
5JJK5 924
QAAAJ 712
64967 417
J4744 822
7AK2A 747
6Q7Q7 859
53294 986
J88J8 897
25T8T 19
99939 223
337J3 749
TT32T 798
T7777 33
3278J 168
4J3QT 787
5TA5J 741
494J4 255
66TT6 415
7KAKA 664
448J9 792
T898T 124
49444 74
J66J6 410
AAQA7 785
7J673 573
T4TTT 40
K7277 992
Q947J 162
A4AQQ 581
73A37 881
KAA99 243
J8369 63
7K773 799
QQQJJ 503
3A3A7 556
A4A7A 943
94949 32
22QA2 744
Q665K 57
Q33QQ 920
33TT3 343
88555 814
QQ85T 466
7AQ4K 791
8T888 590
2T78K 835
39939 854
A8888 458
K462T 146
T6A72 527
42AQ8 42
KKKKQ 143
JTTT8 694
K3JK7 766
54627 153
Q8A62 213
A9999 193
Q2772 400
6JJ6J 322
A66A8 834
8Q3J5 467
49489 959
88598 896
2K228 545
5A939 917
66K66 414
2T666 233
J4Q7J 58
A8A5Q 899
T5547 708
75535 14
9AT8A 597
383J3 584
A2KJK 471
T8822 211
54649 512
8K668 113
QAQAQ 587
6666J 296
7K3J6 941
74QTA 447
6JK66 380
788Q8 973
33J4A 857
8Q99A 873
666Q8 589
AJJA7 52
66983 692
9555K 319
249QQ 371
AT7T4 525
924JA 631
3K883 273
J6333 830
5QAT2 838
88J77 254
5T6K3 644
22AA5 388
5J5QT 807
AATT7 194
66659 197
239JQ 534
922QQ 733
84Q5K 385
66TJA 893
TTT7T 643
66646 393
K5A68 672
7KKKJ 133
234J5 216
J4J44 134
77773 236
AT8K7 784
88QQQ 551
42Q5J 330
42AA4 720
77587 350
385A2 27
32JQ2 264
2A323 802
8J7Q8 533
6Q6Q2 704
6656T 176
8AQTJ 158
89TJ8 293
37668 286
26662 688
9KK47 560
2262J 391
64545 382
42Q93 509
J8Q88 968
QA798 588
Q4444 609
Q55QQ 811
A5555 87
93838 570
A9354 431
8QAJ5 887
222J8 665
T6Q45 680
7866K 95
33323 919
55A95 705
K3KKK 760
KKKQQ 918
TTQT7 686
A38A9 454
Q38K5 370
KKKKJ 465
6Q336 368
33K33 931
T2TTT 870
QKAQK 682
2K635 157
696QQ 69
9J29J 450
A9K6A 971
KQK5K 691
QJQA2 500
6655A 351
A6AAJ 778
83J88 661
38TQ5 244
J3JTT 816
K9949 506
KT8AJ 96
5947A 304
Q8848 345
KJ537 89
K62Q5 300
93T44 743
K44KT 60
559KJ 565
87JA8 155
AKATA 464
J9423 335
943J4 948
4Q365 752
3Q5A3 636
9T998 82
AT486 731
KQK8Q 409
JKTTK 647
5J9TA 877
444A7 357
K7K57 987
37273 656
J999J 106
J2922 998
2492A 633
4489A 394
777J7 452
3T5J4 477
7ATT7 260
J69T7 79
885J6 77
3JA55 552
J9TK9 815
777Q7 874
K5232 221
KQKK4 697
A3333 8
8KA84 453
5QKJK 212
9KAKK 574
355Q5 844
Q4Q44 484
5552Q 761
J654K 600
43KJ5 813
3AA9T 329
JAA88 923
4A4K4 738
922Q2 413
22JJ4 886
43425 359
342T7 292
44K87 139
88656 657
T54TT 821
5T5T3 354
AQ367 759
J8584 485
Q5QAA 303
983K4 161
22262 487
74KQ4 177
22293 478
588JJ 831
QA2A6 864
9Q4A7 56
K433J 196
5J468 730
JAA24 862
53365 190
229TA 379
34J74 206
K587A 627
T4TQT 64
86868 520
69KJ6 674
J7T63 871
KK33K 486
AJJTK 576
JA777 416
J4Q88 852
4K9QA 884
K3393 10
J4444 828
A9JAJ 389
999TT 605
Q5Q66 137
QQK2K 101
QQ7J3 93
4TTT7 507
55757 670
A32A7 333
722TT 624
5TKJ2 53
KKK36 980
8Q4JQ 619
35JT3 428
22J2Q 441
9564K 481
T6K4T 395
K25KK 181
99992 620
93778 289
Q67A8 622
222A5 693
372Q5 577
5Q7A3 298
AA5AJ 540
278Q7 191
9Q999 451
77JAA 461
6A7J5 839
3A9A9 72
33363 270
J3Q49 970
22Q2K 432
AA666 772
2K2J2 667
KTQA2 521
64994 817
793T6 37
25255 381
AK64T 882
6AQ8T 655
J2QQ6 328
K2KK6 408
97J88 488
QA4QQ 103
J44JJ 955
29992 511
J4884 174
6A229 240
ATJQ2 781
QAK7T 653
JKQQK 285
9AT49 985
43253 942
76Q8Q 472
88388 475
576AK 438
8QA53 637
69TK9 420
2T77K 566
22343 67
52558 9
2376J 742
82KA9 703
3TTTT 377
8T88T 726
6J79K 392
3Q333 591
55655 272
3JKQA 716
22266 779
J854J 535
A2TTT 771
99636 991
J2TTT 214
68QT8 837
4JT67 164
K88TA 710
ATQ35 711
99T99 342
333J2 430
5K255 208
JJ22J 252
8KT24 171
4287J 80
K7878 494
K37J7 851
Q97T4 860
QJ5K5 12
4KTQ4 658
ATTAA 898
562KJ 806
4T8Q6 308
JQ4K4 48
T8844 163
66677 50
QQQQ8 777
23KJK 988
8QA56 130
9329Q 299
8K85T 412
82K27 810
347A9 958
KQ8J2 198
7JA49 735
JK6Q8 582
766TK 641
7939K 580
Q6888 59
2A2A2 282
JAQQ5 652
99K88 364
44443 757
97739 483
268JA 85
24444 256
QTQTT 62
AAA4A 606
2AKKT 532
875KT 127
9226J 346
97977 316
KK4KK 258
66355 659
8K6TJ 167
59J7A 46
9TJ4K 231
33Q35 150
555J5 187
T6QTT 954
88588 433
34333 305
4AJ8Q 698
TTATT 183
9KJJ5 373
KA65A 459
849KA 396
J2272 950
6Q999 45
TT444 455
9AQQA 611
6275A 449
Q7KTJ 489
K22KJ 599
J8298 571
J5JK7 676
JJ922 426
53T38 406
7K477 790
82822 539
57434 344
8T555 15
97997 946
29675 262
3T322 554
K5T7J 186
7KKKK 598
997J7 572
Q22QJ 47
52773 713
JJ8JJ 557
4JT44 367
5Q4A9 684
8656T 707
J4JK4 999
7K777 1
36388 921
AT487 630
7Q89K 495
QJQQQ 422
8K8K2 51
A5A75 827
85822 102
44K46 809
86686 938
A77AA 230
T7JTT 479
6Q746 964
K8KTK 732
995Q7 281
2798Q 1000
QQ5QQ 325
AA76J 68
3QA4A 868
J7257 337
8Q55Q 516
JAAA7 594
6Q2A4 6
66696 885
T429J 603
T7878 518
47TJT 75
JTJ95 823
J24Q3 638
44477 470
7A779 994
8J268 907
77797 543
26822 462
22279 818
6JA6A 209
73T37 602
AA55A 727
4JA84 237
4T4AJ 915
68786 916
TTJT5 714
J2J22 339
KK88Q 596
TTJ4J 900
K898K 966
666T6 555
8J3QT 90
33366 863
QQJKQ 267
3QQQ4 207
29TQT 253
33553 112
22AAA 34
Q9842 315
89KK6 290
3969K 805
54844 245
55765 20
7A3J6 156
375K2 83
33J3Q 775
77KQQ 224
Q7254 28
QA87T 729
T4772 326
33AAA 909
33T2T 905
8JQT8 201
JJJA8 990
23KAQ 361
AA777 908
35974 960
55559 558
26223 423
6JT58 375
92286 842
7AK33 853
86248 764
338K3 151
44466 849
Q8Q88 719
T23J2 617
A5A7A 579
88688 115
6K64K 320
79524 737
Q3692 249
37366 505
5Q992 526
22242 340
TJ7T7 608
68666 210
3T454 175
T2299 266
TT8TQ 26
A647A 178
33399 841
2J99T 228
22K22 284
T49AK 425
9KJ3K 200
996T4 402
557A5 718
7JT78 903
94AAJ 140
843A8 444
73773 271
K73QA 4
8QA2A 614
QAQQQ 100
89999 639
58Q54 767
26422 173
A7AJK 654
AKTJ7 159
AA44A 311
6592Q 148
KQQ4Q 922
Q6622 935
59599 618
J9646 493
A2T22 548
8A233 215
TTJ99 788
JJ868 880
AT32J 98
AQQAA 634
4TJ99 957
7JJK9 349
78TJA 892
6T423 679
TKK6K 794
JJ6TK 327
JQKK2 891
3QQQQ 404
75362 490
T39Q7 754
89997 469
QK7K7 829
KA3TQ 297
78278 136
A5KKJ 84
AAA35 706
4T94T 825
QQ7Q7 172
T89AT 314
66A96 699
332JJ 31
55656 43
4Q454 947
96896 536
T9274 336
78888 372
A4848 604
3AT46 736
TK999 895
7T445 649
QK28Q 383
7T29J 355
7747J 689
3333J 662
6AT66 616
663T5 110
99K4K 508
28J66 317
3T993 362
969J9 531
J5999 227
8899T 979
67757 846
56666 166
J9427 184
95939 789
69ATA 538
T3A6T 668
93Q56 969
9729J 901
K7KQK 808
6QQQ7 185
94477 768
599AA 474
A766A 865
8886J 746
66AK6 259
T3AJQ 443
3KK63 427
TTTJ6 424
36467 123
TTJJT 169
Q24J4 542
7QQ77 41
Q4QJ4 180
K444K 773
367T3 615
549K2 989
TTJJA 583
75TT7 468
9299Q 872
A9J5A 755
22QQQ 234
T44TT 635
A7AAA 276
T7T7T 858
95595 348
KKJ8Q 434
3K49A 888
AJ5A3 675
J8488 21
4K8TJ 666
K7QQ5 914
TKKKA 751
6JQAJ 61
A33A3 248
9J99A 517
24T99 125
K7A7T 592
94AK9 497
992J2 81
82J88 473
778TK 997
44422 629
337Q3 114
JJJJJ 996
K55K5 715
AAAT3 826
44848 499
JJA33 578
35288 295
TK2JT 291
46977 492
22722 92
77699 269
J9KQT 192
52555 758
JK692 793
85A55 687
4A944 972
4J949 440
22KKK 119
969TT 203
55994 229
KKK77 515
6QQ44 118
462K6 876
3J7QA 685
33335 621
A4339 669
8J969 651
52852 55
TTJ3T 242
J3393 356
AATT5 23
566J5 287
9J49Q 152
7K577 911
5Q7J2 974
AAAJJ 24
KKKKA 491
8Q888 750
T7K3A 165
K3565 66
Q4J33 446
33238 683
T368T 54
67Q7J 445
JJ555 263
A4867 723
843QT 933
T85TT 937
TJT23 334
88AJ5 321
35J82 386
72275 145
4367T 86
99696 436
TKTAA 553
96696 39
T4442 549
493K7 365
66Q66 131
747T4 541
95999 544
Q75QQ 995
54KA9 850
582JJ 421
QQAA7 677
K4444 313
84877 804
K9999 568
7JKK7 117
3A9J7 928
66568 939
83733 302
JQ37T 138
K8JJ9 277
K3A33 801
J523J 695
333JA 435
KK855 239
39J99 235
8T8TT 318
6K66K 104
54444 575
K83KK 952
K835T 944
T4A44 135
75566 324
44A3A 562
74KT2 16
T625J 154
94449 528
6969J 126
TK554 671
5A797 601
77778 129
9K485 88
363JK 546
25T39 407
5545J 358
KK6K6 678
QQQQK 569
QT256 280
39896 288
TT9TT 301
39JAK 753
JQQJJ 251
KA2AA 913
79KT5 978
3TT3T 5
76747 38
A7JA4 501
J9AJ8 883
T7JJJ 725
AQA9J 940
444TK 845
6Q6QT 645
8KAQ5 390
3534A 456
958QQ 49
95482 927
J497J 108
QQ8QK 812
674TK 961
8T5J5 769
2KK28 975
2JJ5J 398
692J9 932
AJ8AA 595
9999J 76
373T3 912
JTTTT 982
6A6AA 783
66J36 681
4A77J 278
84848 265
49865 537
34Q6Q 530
2A262 238
K34JQ 776
3A533 283
2Q27Q 247
A8QTQ 341
A856J 116
57A48 926
742Q8 279
23QQK 331
2QQQJ 797
K9JKK 795
K6KKK 855
32299 613
2QA9T 522
36666 401
Q96QQ 189
66525 894
T2822 376
9K586 843
TA988 890
22292 366
7J23K 734
Q283K 149
83333 564
27JTQ 561
9K23Q 360
AQ2AA 120
QQ7QQ 496
4KKK4 963
337JJ 352
J424K 739
95TA2 429
QTJTT 586
T78TQ 250
6TTTT 878
29K72 353
42944 397
QQQ48 547
T45JA 514
JQK63 646
TAT55 866
88J88 310
333JJ 170
A2AJ2 457
98728 411
AAJ39 65
83T86 387
9K9KK 976
86644 993
8Q79J 748
3484A 94
85878 585
268TK 628
5454J 690
77374 482
32532 612
55T55 879
55553 529
6667T 107
3323K 141
65595 405
26628 71
9T423 312
TKKKT 640
KKKA3 142
7J38K 663
6977A 232
A2Q66 981
4AK76 745
24248 222
2KJK6 607
J7776 147
256Q3 403
99J89 875
46A64 369
JTT87 513
7JJ77 111
JA69A 934
KAAAA 956
82222 836
55Q5Q 559
Q666J 869
TTTAA 701
54J6J 820
8JJ44 99
T3KKT 294
6646Q 632
68QQQ 650
4T52T 550
93K62 593
QA5TJ 17
87J3Q 363
7TA68 419
77J62 728
22244 205
T2864 97
AA6J4 717
J9633 182
K259T 762
8Q882 648
K365A 965
68JJ9 2
76JK7 109
J7888 800
33K8J 696
TT464 786
74TQ2 11
33666 122
36JK4 22
3AJ9Q 399
632KT 765
J9T3J 673
T6AK9 91
JJJ2J 967
KQKK6 3
J2346 867
48338 442
Q9475 774
QQ5K5 199
9T354 861
93888 984
7J877 460
QQQ9Q 73
9T2J3 347
TQQT3 803
QQQ32 13
35883 218
7377J 202
38QTQ 780
92KK9 374
TJ957 819
687J7 770
8K345 274
799AJ 925
7J797 480
2555A 660
8349T 702
JKJT7 105
965QJ 902
A4TAT 524
T333J 906
22A22 983
2J884 246
797Q7 563
2J63T 709
A3833 306
KAJJ9 121
55562 700
T4444 439
25272 567
72822 179
32535 226
KKAJA 307
KJKKJ 78
7Q539 833
J825A 504
J898K 448
J6ATJ 910
K59AK 18
TK3TJ 642
888JT 476`
