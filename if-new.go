package main

import (
	"fmt"
	"math"
)


func checksqrt(x, y, limit float64) float64 {
	if v:=math.Pow(x,y); v < limit {
		return v;
	}else{ //else should not be separated from if block.....
	//if v < limit return v else return limit.... instead of else just returning the limit.
	//return limit; // We can use else keyword too....
		return limit;
	}
}
func main(){
	fmt.Println(checksqrt(3,2,10))
	fmt.Println(checksqrt(3,3,20))
}
