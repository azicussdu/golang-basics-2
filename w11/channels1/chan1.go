package main

func sumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func concurrentSum(nums []int) int {
	ch := make(chan int)
	go func() {
		ch <- sumSlice(nums[:len(nums)/3])
	}()
	go func() {
		ch <- sumSlice(nums[len(nums)/3 : 2*len(nums)/3])
	}()
	go func() {
		ch <- sumSlice(nums[2*len(nums)/3:])
	}()
	sum1 := <-ch
	sum2 := <-ch
	sum3 := <-ch
	return sum1 + sum2 + sum3
}

//func main() {
//	ch := make(chan int)
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			ch <- i
//			time.Sleep(1 * time.Second)
//		}
//		close(ch)
//	}()
//
//	for v := range ch {
//		fmt.Println("Received:", v)
//	}
//}
