package main
import "fmt"

func average(nums ...float64) float64 {
	avg := 0.0
	for _,val:=range nums {
		avg = avg + float64(val)
	}
	avg = avg/ float64(len(nums))
	return avg
}
func main()  {
	// fmt.Println(average(3,5,7))
	a := average(3,5,7)
	fmt.Println(a)
}