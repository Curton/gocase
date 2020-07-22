/**
 * @author Covey Liu, covey@liukedun.com
 * @date 2020/7/9 15:18
 */

package _struct

type user struct {
	name  string
	email string
}

type admin struct {
	user
	name string
}
