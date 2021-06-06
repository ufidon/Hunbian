package object

import "testing"

func Test串键(测 *testing.T) {
	猴1 := &C串型{Z值: "这只猴子名字叫孙悟空！"}
	猴2 := &C串型{Z值: "这只猴子名字叫孙悟空！"}
	男1 := &C串型{Z值: "男人来自火星"}
	男2 := &C串型{Z值: "男人来自火星"}

	if 猴1.O耦键() != 猴2.O耦键() {
		测.Errorf("异键同值串")
	}

	if 男1.O耦键() != 男2.O耦键() {
		测.Errorf("异键同值串")
	}

	if 猴1.O耦键() == 男1.O耦键() {
		测.Errorf("异键同值串")
	}
}

func Test贰键(测 *testing.T) {
	真1 := &E贰型{Z值: true}
	真2 := &E贰型{Z值: true}
	假1 := &E贰型{Z值: false}
	假2 := &E贰型{Z值: false}

	if 真1.O耦键() != 真2.O耦键() {
		测.Errorf("异键真贰")
	}

	if 假1.O耦键() != 假2.O耦键() {
		测.Errorf("异键假贰")
	}

	if 真1.O耦键() == 假1.O耦键() {
		测.Errorf("真假同键")
	}
}

func Test整键(测 *testing.T) {
	甲1 := &Z整型{Z值: 1}
	甲2 := &Z整型{Z值: 1}
	乙1 := &Z整型{Z值: 2}
	乙2 := &Z整型{Z值: 2}

	if 甲1.O耦键() != 甲2.O耦键() {
		测.Errorf("同整异键")
	}

	if 乙1.O耦键() != 乙2.O耦键() {
		测.Errorf("同整异键")
	}

	if 甲1.O耦键() == 乙1.O耦键() {
		测.Errorf("异整同键")
	}
}
