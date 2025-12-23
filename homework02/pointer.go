package main

// 题目 1
func taskOne(i *int) {
	*i = *i + 10
}

// 题目 2
func taskTwo(slice *[]int) {
	for i, v := range *slice {
		(*slice)[i] = v * 2
	}
}

// func main() {
// 	i := 3
// 	taskOne(&i)
// 	fmt.Println(i)

// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Println(slice)
// 	taskTwo(&slice)
// 	fmt.Println(slice)
// }
