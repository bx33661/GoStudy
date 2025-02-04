package main

type Employee struct {
	firstName, lastName string
	age, salary         int
}

// 使用指针可以更高效地传递结构体
//emp := &Employee{"Sam", "Anderson", 55, 6000}
//fmt.Println(emp.firstName)  // 可以直接使用emp.firstName，而不需要(*emp).firstName
