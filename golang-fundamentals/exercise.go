package main

import "fmt"

type cost struct{
	day int 
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := make([]float64, costs[len(costs) - 1].day + 1)
	fmt.Println(costsByDay)
	for i := 0; i < len(costs); i++ {
		currentDay := costs[i].day
		costsByDay[currentDay] += costs[i].value
	}
	return costsByDay
}


func main() {
	temp := getCostsByDay([]cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	})
	
	fmt.Println(temp)
}
