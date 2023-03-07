package main

import (
	"fmt"
)

func main() {
	time := []int{2}
	totalTrips := 1
	ans := minimumTime(time, totalTrips)
	fmt.Println(ans)
}

func minimumTime(time []int, totalTrips int) int64 {
	lowerBound, upperBound := int64(1), findMin(time)*int64(totalTrips)
	for lowerBound < upperBound {
		mid := (lowerBound + upperBound) / 2
		if canFinish(&time, totalTrips, mid) {
			upperBound = mid
		} else {
			lowerBound = mid + 1
		}
	}
	return lowerBound
}

func canFinish(time *[]int, totalTrips int, givenTime int64) bool {
	tripsInMidTime := int64(0)
	for _, v := range *time {
		tripsInMidTime += givenTime / int64(v)
	}
	return tripsInMidTime >= int64(totalTrips)
}

func findMin(time []int) int64 {
	currMin := time[0]
	for _, v := range time {
		if v < currMin {
			currMin = v
		}
	}
	return int64(currMin)
}
