package course

import "fmt"

const (
	first = iota + 1
	space
	increment
	second = 2 << iota
)

// iota gets reset in different const blocks
const (
	third = iota
)

// primative and collections
func pac() {
	val := new(string)
	*val = "jimmy"
	var pff = val
	fmt.Println(pff)

	firstName := "tim"

	fmt.Println(firstName)

	const c int = 2
	const x float32 = 1
	fmt.Println(c + 2)

	fmt.Println(float32(c) + x)
	fmt.Println(first, space, increment, second)
	fmt.Println(third)

	//collections

	//array

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	arr[1] = 4
	fmt.Println(arr[1])

	arr2 := [3]int{1, 2, 3}
	arr2[1] = 4
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}

	slice := arr3[:] // : means create a slice that takes from the beginning to the end.
	fmt.Println(slice)

	arr3[1] = 42
	slice[2] = 27
	fmt.Println(arr3, slice)

	slice2 := []int{3, 2, 1} //compiler handles the init of the array that slice is pointing to

	fmt.Println(slice2)

	slice2 = append(slice2, 4) // go is handling the underlying array, so it fits our needs
	// for big sets of data, there are functions to pre allocates data.
	// but in most cases, append will be just fine,
	fmt.Println(slice2)

	s1 := slice2[1:]
	s2 := slice2[:2]
	s3 := slice2[1:2]

	fmt.Println(s1, s2, s3)

	// map
	m := map[string]int{"foo": 42} // [string] = key, int = value type
	fmt.Println(m)

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	// struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
		Car       string
	}

	var u user
	fmt.Println(u) // zero values are shown.

	u.ID = 1
	u.FirstName = "Tim"
	u.LastName = "my"
	u.Car = "BMW"
	fmt.Println(u)

	u2 := user{ID: 2, FirstName: "Harvey", LastName: "Dent"}
	fmt.Println(u2)
}
