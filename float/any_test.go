/**
 * @author Covey Liu, covey@liukedun.com
 * @date 2020/7/22 17:54
 */

package float

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	size := 139
	origin := float64(100)
	total := float64(0)
	for i := 0; i < size; i++ {
		total += origin
		origin *= 1.005
	}
	total += origin
	fmt.Println(total / float64(size+1))
	fmt.Println(origin)
}
