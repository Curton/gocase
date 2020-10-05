/**
 * @author Covey Liu, covey@liukedun.com
 * @date 2020/10/5 11:11
 */

package _struct

import "fmt"

type TT struct {
	t int
}

func (t TT) NewT(tValue int) TT {
	return TT{t: tValue}

}
func main() {
	t := TT.NewT(*new(TT), 999)
	fmt.Println(t)

	t2 := new(TT)
	fmt.Println(t2)
	t3 := t2.NewT(888)
	fmt.Println(t3)
}
