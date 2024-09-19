package slice

import (
	"fmt"
	"testing"
)

// 3. range
//    1. range迭代string
//    2. range迭代map

func TestStruct(t *testing.T) {
	// p和people不是一个值，所以如果for range内部有写操作的话，最好传递struct指针
	t.Run("for-range内部有写操作，需要传struct指针", func(t *testing.T) {
		type person struct {
			name   string
			age    byte
			isDead bool
		}

		t.Run("", func(t *testing.T) {
			p1 := person{name: "zzy", age: 100}
			p3 := person{name: "px", age: 20}
			people := []person{p1, p3}
			go func(people []person) {
				for _, p := range people {
					if p.age < 50 {
						p.isDead = true
					}
				}
			}(people)

			for _, p := range people {
				if p.isDead {
					fmt.Println("who is younger?", p.name)
				}
			}
		})

		t.Run("", func(t *testing.T) {
			p1 := &person{name: "zzy", age: 100}
			p3 := &person{name: "px", age: 20}

			people := []*person{p1, p3}
			go func(people []*person) {
				for _, p := range people {
					if p.age < 50 {
						p.isDead = true
					}
				}
			}(people)

			for _, p := range people {
				if p.isDead {
					t.Logf("who is younger? %s", p.name)
				}
			}
		})
	})
}
