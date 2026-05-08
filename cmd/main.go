package main

import (
	"fmt"
	"time"

	"cuckoo"
)

func main() {
	const n = 1 << 20 // 2^20

	// create cuckoo hash with 2^20 cells
	c := cuckoo.NewCuckoo(20)

	// Insert
	t0 := time.Now()
	for i := 1; i <= n; i++ {
		//fmt.Printf("%d, %d\n", cuckoo.Key(i), cuckoo.Value(i))
		c.Insert(cuckoo.Key(i), cuckoo.Value(i))
	}
	dt := time.Since(t0)
	fmt.Printf("Inserted %d items in %s\n", n, dt)
	fmt.Printf("cuckoo len: %d\n", c.Len())
	// Search sequentially
	t1 := time.Now()
	found := 0
	for i := 1; i <= n; i++ {
		if _, ok := c.Search(cuckoo.Key(i)); ok {
			found++
		}
	}
	dt2 := time.Since(t1)
	fmt.Printf("Searched %d items in %s, found %d\n", n, dt2, found)

	c.PrintFirstNBuckets(10)

	t2 := time.Now()
	iteratedCount := 0
	var sum cuckoo.Value // 顺便做一个简单的累加测试，防止编译器把空循环优化掉

	c.ForRange(func(k cuckoo.Key, v cuckoo.Value) {
		iteratedCount++
		sum += v
	})

	dt3 := time.Since(t2)
	fmt.Printf("Iterated over %d items in %s (Sum of values: %d)\n", iteratedCount, dt3, sum)

	for i := 1; i <= 10; i++ {
		if _, ok := c.GetHashPosition(cuckoo.Key(i)); ok {
			//found++
		}
	}

}
