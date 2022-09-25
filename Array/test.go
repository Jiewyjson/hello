package main

func main() {
	var numbers = make([]int, 3, 4) //(0,0,0)
	numbers = append(numbers, 2)    //(0,0,0,2)
	numbers = append(numbers, 3)    //(0,0,0,2,3)
	s1 := numbers[3:]               //(2,3)
	s2 := []int{11, 22, 33}
	copy(s1, s2)

	for _, value := range s1 {
		println(value)
	}
	println("--------")
	for _, value := range s2 {
		println(value)
	}
}
