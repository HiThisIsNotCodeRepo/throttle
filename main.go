package main

import (
	"fmt"
	"sort"
)

func droppedRequests(requestTime []int32) int32 {
	sort.Slice(requestTime, func(i, j int) bool { return requestTime[i] < requestTime[j] })
	var tenSecondDrop bool
	var oneSecondDrop bool
	var sixtySecondDrop bool
	var dropCounter int32
	var oneSecond oneSecondPipe
	var tenSecond tenSecondPipe
	var sixtySecond sixtySecondPipe
	for _, v := range requestTime {
		sixtySecondDrop = sixtySecond.add(v)
		if sixtySecondDrop == true {
			dropCounter++
			//fmt.Printf("Ten second drop occurred:%v,drop number:%v\n", dropCounter, v)
		}
		tenSecondDrop = tenSecond.add(v)
		if tenSecondDrop == true && sixtySecondDrop == false {
			dropCounter++
			//fmt.Printf("Ten second drop occurred:%v,drop number:%v\n", dropCounter, v)
		}
		oneSecondDrop = oneSecond.add(v)
		if oneSecondDrop == true && tenSecondDrop == false && sixtySecondDrop == false {
			dropCounter++
			//fmt.Printf("One second drop occurred:%v,drop number:%v\n", dropCounter, v)
		}
		sixtySecondDrop = false
		tenSecondDrop = false
		oneSecondDrop = false

	}
	return dropCounter
}

type oneSecondPipe struct {
	data []int32
}

func (this *oneSecondPipe) add(num int32) bool {
	if len(this.data) == 0 {
		this.data = append(this.data, num)
		return false
	} else {
		if this.data[0] == num {
			if len(this.data) == 3 {
				return true
			} else {
				this.data = append(this.data, num)
				return false
			}
		} else {
			this.data = []int32{num}
			return false
		}
	}
}

type tenSecondPipe struct {
	data []int32
}

func (this *tenSecondPipe) add(num int32) bool {
	if len(this.data) == 0 {
		this.data = append(this.data, num)
		return false
	} else {
		if num-this.data[0] < 10 {
			this.data = append(this.data, num)
			if len(this.data) > 20 {
				return true
			} else {
				return false
			}
		} else {
			lastIndex := lastIndexOf(this.data[0], this.data)
			this.data = this.data[lastIndex:]
			this.data = append(this.data, num)
			if len(this.data) > 20 {
				return true
			} else {
				return false
			}
		}
	}
}

type sixtySecondPipe struct {
	data []int32
}

func (this *sixtySecondPipe) add(num int32) bool {
	if len(this.data) == 0 {
		this.data = append(this.data, num)
		return false
	} else {
		if num-this.data[0] < 60 {
			this.data = append(this.data, num)
			if len(this.data) > 60 {
				return true
			} else {
				return false
			}
		} else {
			lastIndex := lastIndexOf(this.data[0], this.data)
			this.data = this.data[lastIndex:]
			this.data = append(this.data, num)
			if len(this.data) > 60 {
				return true
			} else {
				return false
			}
		}
	}
}

func lastIndexOf(value int32, source []int32) int {
	for tempIndex, tempValue := range source {
		if tempValue > value {
			return tempIndex
		}
	}
	return 0
}

func main() {

	//intArr := []int32{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 11, 11, 11, 11}
	//intArr := []int32{1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 13, 13, 13, 13, 13, 13, 13, 13, 14, 14, 14, 14, 14, 14, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 16, 16, 16, 16, 16, 16, 16, 16, 17, 17, 17, 17, 17, 17, 18, 18, 18, 18, 18, 18, 19, 19, 19, 19, 19, 19, 20, 20, 20, 20, 20, 20, 20}
	//intArr := []int32{
	//	1, 1, 1, 1,
	//	2, 2, 2,
	//	3, 3, 3, 3,
	//	4,
	//	5, 5, 5,
	//	6, 6, 6, 6,
	//	7, 7, 7,
	//	8, 8, 8, 8,
	//	9, 9, 9, 9, 9,
	//	10, 10,
	//	11, 11, 11, 11, 11, 11,
	//	12, 12, 12, 12, 12, 12, 12,
	//	13, 13, 13, 13,
	//	14, 14, 14, 14, 14,
	//	16, 16, 16, 16, 16, 16,
	//	17, 17, 17,
	//	18, 18, 18, 18, 18, 18, 18, 18,
	//	19, 19, 19, 19, 19, 19, 19,
	//	20, 20, 20, 20, 20,
	//}
	intArr := []int32{
		100, 100,
		101,
		102, 102,
		103, 103,
		105, 105,
		106, 106, 106,
		107, 107,
		109, 109, 109, 109,
		110, 110, 110,
		111, 111,
		112,
		114,
		115, 115,
		116, 116, 116,
		117,
		118, 118,
		120, 120,
		121, 121,
		122, 122,
		123, 123,
		124, 124,
		125,
		127,
		128, 128,
		129,
		131,
		133,
		134,
		135, 135, 135,
		136,
		137, 137, 137,
		138, 138,
		140, 140, 140,
		141, 141,
		143,
		144,
		145,
		146,
		149, 149, 149, 149, 149,
		151, 151,
		152, 152,
		154, 154,
		155,
		156, 156,
		157, 158,
		158,
		159, 159,
		160, 160, 160,
		162, 162, 162,
		164,
		166,
		167, 167,
		169, 169,
		172, 172, 172, 172,
		174,
		175, 175, 175,
		176,
		177, 177, 177,
		179, 179,
		180, 180, 180, 180,
		181, 181, 181,
		182, 182,
		183, 183, 183,
		184, 184,
		187, 187,
		188, 188, 188, 188, 189,
		189, 189,
		190,
		191, 191,
		192, 192,
		193, 195,
		195, 195, 195,
		197,
		198, 198,
		199, 199,
		200, 200, 200}

	fmt.Printf("arr length:%v\n", len(intArr))
	fmt.Println(droppedRequests(intArr))
}
