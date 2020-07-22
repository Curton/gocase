/**
 * @author Covey Liu, covey@liukedun.com
 * @date 2020/7/9 15:18
 */

package _struct

import (
	"fmt"
	"testing"
)

func TestS(t *testing.T) {
	a := admin{
		user: user{
			name:  "aaa",
			email: "ccc",
		},
		name: "bbb",
	}

	fmt.Println(a.name)
	fmt.Println(a.user.name)
	fmt.Println(a.user.email)
	fmt.Println(a.email)

}
