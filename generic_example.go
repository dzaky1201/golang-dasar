package main

import "fmt"

func SumGen[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}

	return total
}

func SumNumbers1(m map[string]int64)int64{
	var total int64
	for _, v := range m {
		total += v
	}

	return total
}

// type constraint
type Number interface {
	int64 | float64
}

func SumNumbers2[K comparable, V Number](m map[K]V)V{
	var total V
	for _, v := range m {
		total += v
	}

	return total
}

type UserModel[T Number]struct{
	Name string
	Scores []T
}

func (m *UserModel[int64]) SetScores(scores []int64){
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64){
	m.Scores = scores
}


func main() {
	total1 := SumGen([]int{1, 2, 3, 4, 5})
	total2 := SumGen([]float32{2.5,7.2})
	total3 := SumGen([]float64{1.23,6.33,12.6})

	fmt.Println("total :", total1)
	fmt.Println("total :", total2)
	fmt.Println("total :", total3)

	ints := map[string]int64 {"first":64,"second":12}
	floats := map[string]float64 {"first":64.45,"second":12.32}
	booleans := map[bool]int64 {true:1,false:0}
	fmt.Printf("Generic sums with constraint: %v, %v and %v\n", SumNumbers2(ints), SumNumbers2(floats), SumNumbers2(booleans))

	m1 := UserModel[int64]{
		Name: "Agung",
		Scores: []int64{1,2,3},
	}
	fmt.Println("scores :", m1.Scores)

	m2 := UserModel[float64]{
		Name: "Jhon",
		Scores: []float64{1.45,2.54,3.89},
	}
	fmt.Println("scores :", m2.Scores)
}