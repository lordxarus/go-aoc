package main

import "testing"

func TestToNumeral(t *testing.T) {
	in := "twothree41eight"
	want := "23418"
	got := toNumeral(in)
	if want != got {
		t.Fatalf("in: %s want: %s got: %s", in, want, got)
	}
}

func TestParseByHand(t *testing.T) {
	in := []string{
		"fnm3oneightsdn",
		"3nine2gkbpnvninegmgqdgr",
	}
	want := []int{
		38,
		39,
	}
	for i, s := range in {
		got := firstAndLast(s)
		if want[i] != got {
			t.Fatalf("want: %d got: %d", want[i], got)
		}
	}

}

func TestParseMany(t *testing.T) {
	in := []string{
		"nkzjrdqrmpztpqninetwofour1znnkd",
		"3four4",
		"three3ninefive",
		"822",
		"3n",
		"three",
		"fourtwo",
		"cqqlrjkjbrggmccctbjzcjqktfngmmkftddponemqj17",
		"dtwoneeight6llzcxssgrdfjmjvfbvtwo9",
		"twoqdjpssix75tonetwothree",
		"3341",
		"kksh2",
		"",
		"foursixfivecsvpxqbct7threelrgbxz",
		"132twofour",
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		//
		"cdfkddlnine34frsshqsevensevenfourqf",
		"s5sevenxrdfr4mhpstgbjcfqckronesix",
		"2xmdmtgcjhd8eighttwo",
		"six9jhnloneightf",
		"seven4nineeight175ninenine",
		"ninegbsscnbtcnzhsevenfmfvmv3lrbthree3",
		"fnm3oneightsdn",
		"89clszfjqlffive5s",
		"tpxmjxd28onefrn9hnmztsmxsmctpdxjh",
		"4jqszj91ninegvzrsctzl",
		"3nine2gkbpnvninegmgqdgr",
		"sbkncfseven18qdvtssf17jn4",
		"4jddtplseven",
		"nine8qslr",
		"pvgksfone72foureight",
		"1eightfive",
		"eighttwoxlsjdxxfive1",
		"2ctfonekpns",
		"kpkpx1onenine",
		"xfgxnvdxj1",
	}
	want := []int{
		91,
		34,
		35,
		82,
		33,
		33,
		42,
		17,
		29,
		23,
		31,
		22,
		-1,
		43,
		14,
		29,
		83,
		13,
		24,
		42,
		14,
		76,
		94,
		56,
		22,
		68,
		79,
		93,
		38,
		85,
		29,
		49,
		39,
		74,
		47,
		98,
		18,
		15,
		81,
		21,
		19,
		11,
	}
	out := parse(in)
	for i, n := range out {
		if want[i] != n {
			t.Fatalf("want: %d, got: %d, cand: %s", want[i], n, in[i])
		}
	}

}

func TestFuzzy(t *testing.T) {

}
