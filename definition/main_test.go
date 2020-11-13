package definition

import (
	"log"
	"testing"
)

func multiplyElementWithFirst(datas []int) int {
	res := 0          //Fault : Not checking whether or not there are element in here
	first := datas[0] //Error : There is no data at all (Panic)

	for _, val := range datas {
		res += val * first
	}
	return res //Failure : Result is none
}

func doPow2(val int) (int, error) {
	res := val * 2  //Fault : multiplied by 2 instead of itself
	return res, nil //Error : res is not a power of two
}

func TestDefinition(t *testing.T) {
	pow2Res, _ := doPow2(3) //Failure : Silent error, wrong output
	log.Println(pow2Res)
	multRes := multiplyElementWithFirst([]int{1, 2, 3, 4}) //Will run normally
	log.Println(multRes)
	multResErr := multiplyElementWithFirst([]int{}) //Failure : Fail stop panic
	log.Println(multResErr)
}
