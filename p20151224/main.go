package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kyoh86/qiita/util"
)

func main() {
	nofbzz, err := SekaiNoNabeatsu(100)
	fmt.Println("まともに言えた数：")
	for _, n := range nofbzz {
		fmt.Println(n)
	}
	fmt.Println("アホまたは犬：")
	fmt.Println(err)
}

// SekaiNoNabeatsu は、3の倍数だけアホになり、5の倍数だけ犬っぽくなります
func SekaiNoNabeatsu(num int) ([]int, error) {
	var ret []int
	var errs = new(util.Errors)

	for i := 0; i < num; i++ {
		fizz := i % 3
		buzz := i % 5
		switch {
		case fizz == 0 && buzz == 0:
			errs.Push(errors.New("くぅ〜ん" + strings.Repeat("！", i/5/3)))
		case fizz == 0 && buzz != 0:
			errs.Push(errors.New("さぁ〜ん" + strings.Repeat("！", i/3)))
		case fizz != 0 && buzz == 0:
			errs.Push(errors.New("わん" + strings.Repeat("！", i/5)))
		default:
			ret = append(ret, i)
		}
	}
	return ret, errs.Err()
}
