package main

import (
	"fmt"
)

func main() {
	has := hasGroupsSizeX([]int{1, 1})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1})
	fmt.Println("hax groups size x: t% ", has)
}

func hasGroupsSizeX(deck []int) bool {
	num := len(deck)
	if num >= 2 {
		//		fmt.Println("num is " + strconv.Itoa(num))
		for factor := 2; factor <= num; factor++ {
			if num%factor == 0 {
				usedMap := map[int]bool{}
				for i, value := range deck {
					_, used := usedMap[i]
					if !used {
						count := 1
						for j := i + 1; j < num; j++ {
							if deck[j] == value {
								usedMap[j] = true
								count++
							}
							if count == factor {
								usedMap[i] = true
								//								fmt.Println("meet current factor: %d", factor)
								break
							}
						}
					}
					//					fmt.Println("current used map len: %d, %d", i, len(usedMap))
					if i == num-1 && len(usedMap) == num {
						//						fmt.Println("factor is " + strconv.Itoa(factor))
						return true
					}
				}
			}
		}
	}
	return false
}
