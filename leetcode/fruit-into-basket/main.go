package main

func main() {

}

func TotalFruit(tree []int) int {
	fruits := 0
	max := 0
	if tree != nil && len(tree) > 0 {
		ftypes := make(map[int]int)

		for index, f := range tree {
			if len(tree)-index < max {
				break
			}

			fruits = 1
			ftypes = make(map[int]int)
			ftypes[f] = 1

			for i := index + 1; i < len(tree); i++ {
				ftypes[tree[i]] = 1
				if len(ftypes) > 2 {
					if fruits > max {
						max = fruits
					}
					break
				}
				fruits++
			}

			if fruits > max {
				max = fruits
			}
		}
	}
	return max
}
