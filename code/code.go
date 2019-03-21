package main

import (
	"fmt"
	//"runtime"
	"bufio"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Starting code ...")
	Reader := bufio.NewReader(os.Stdin)
	str, _ := Reader.ReadString('\n')
	str = str[:len(str)-1]
	fmt.Println("%s", str)

	//l := []int{100, 2, 7, -200, 3, 400, -10, 4, 9, 9, -20}
	//s := "fuck you "
	// c, left, right, le := RemoveDuplicates(l)
	// fmt.Println(c, left, right, le)
	//fmt.Println(LengthOfLastWord(s))
	//DNS()
}

//058
func LengthOfLastWord(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	res := 0
	for i := size - 1; i >= 0; i-- {
		fmt.Println(string(s[i]))
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}
	return res
}

//053
func MaxSubArray(l []int) int {
	sum, maxSum := 0, 0
	for _, n := range l {
		fmt.Println(sum, n)
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//027
func RemoveElement(l []int, val int) int {
	i, j := 0, len(l)-1
	for {
		for i < len(l) && l[i] != val {
			i++
		}
		for j >= 0 && l[j] == val {
			j--
		}
		if i >= j {
			break
		}
		l[i], l[j] = l[j], l[i]
		fmt.Println(l)
	}
	return i
}

func BlueSort(l []int) []int {
	size := len(l)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
	return l
}

func RemoveDuplicates(l []int) (length, left, right int, sl []int) {
	left, right, s := 0, 1, len(l)
	for ; right < s; right++ {
		//fmt.Println(l, left, right)
		if l[left] == l[right] {
			continue
		}
		left++
		fmt.Println("*", l, left, right)
		l[left], l[right] = l[right], l[left]
	}
	length = left + 1
	sl = l
	return
}

type List struct {
	size int64
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	Value int
	next  *ListNode
}

// func MergeList(l1, l2 []int) (l []int) {
// 	ln1 := listTolistnode(l1)
// }

// func listTolistnode(l []int) *ListNode {
// 	var ln ListNode
// 	for i, j := range l {
// 		ln.Value = i
// 		ln.next = nil
// 	}
// }

type Cache struct {
	data map[string]*endpoints
	mu   sync.RWMutex
}

type endpoints struct {
	addrs []string
	pos   uint32
}

func (eps *endpoints) next() string {
	pos := atomic.AddUint32(&eps.pos, 1)
	idx := pos % uint32(len(eps.addrs))
	return eps.addrs[idx]
}

func (c *Cache) Fetch(domain string) (addrs string, ok bool) {
	c.mu.RLock()
	// defer c.mu.RUnlock()
	if eps, ok := c.data[domain]; ok {
		c.mu.RUnlock()
		return eps.next(), true
	}
	c.mu.RUnlock()
	c.mu.Lock()
	if addrs, ok := c.lookup(domain); ok {
		eps := &endpoints{
			addrs: addrs,
			pos:   0,
		}
		c.data[domain] = eps
		c.mu.Unlock()
		return eps.next(), true
	}
	c.mu.Unlock()
	return "", false
}

func (c *Cache) lookup(domain string) (addrs []string, ok bool) {
	return []string{"127.0.0.1", "192.168.1.1", "10.0.0.1", "172.16.0.1"}, true
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]*endpoints),
	}
}

func DNS() {
	c := NewCache()
	printAddr := func(domain string) {
		if addr, ok := c.Fetch(domain); ok {
			fmt.Println(domain, "->", addr)
		} else {
			fmt.Println(domain, "->", "NULL")
		}
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			printAddr("foo")
			time.Sleep(100 * time.Microsecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			printAddr("bar")
			time.Sleep(100 * time.Microsecond)
		}
	}()
	wg.Wait()
}

//020
func isValid(s string) bool {
	size := len(s)

	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1
			top++
			fmt.Println(stack, top, string(c))
		case '[', '{':
			stack[top] = c + 2
			top++
			fmt.Println(stack, top, string(c))
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
				fmt.Println(stack, top, string(c))
			} else {
				return false
			}
		}
	}
	return top == 0
}

//014
func LongestCommonPrefix(ss []string) string {
	short := shortestString(ss)
	for i := range ss {
		for m := range short {
			if short[m] != ss[i][m] {
				return short[0:m]
			}
		}
	}
	return short
}
func shortestString(ss []string) string {
	r := ss[0]
	for i := range ss {
		if len(ss[i]) < len(r) {
			r = ss[i]
		}
	}
	return r
}

//009
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

//007
func reverse(x int) int {
	res := 0
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x = x / 10
	}
	return res
}

//001
func twoSum(nums []int, target int) []int {
	for _, j := range nums {
		for _, n := range nums {
			if j+n == target {
				return []int{j, n}
			}
		}
	}
	return nil
}
