package main

import "fmt"


func test_pointer(){

	b := 255         // b mang gia tri 255, o dia chi 0xaaaa
	var a *int = &b  // a mang gia tri 0xaaaa, la mot con tro

	// thay doi *a cx nhu thay doi b 
	// co the truyen con tro vao func : func change(val *int) { }
	// luc goi thi truyen vao ham dia chi bien change(&x)

	fmt.Printf("Type of a is %T\n", a)

	fmt.Println("address of b is", a)
	fmt.Println("address of b is", &b)

	fmt.Println("value of b is", b)
	fmt.Println("value of b is", *a)

	// khoi tao con tro bang new 
	size := new(int)
	*size = 85
	fmt.Println("size = ", *size)

}


// struct define
type Employee struct {
	name 	 string
	salary   int
	currency string
}

// method of struct Employee
func (e Employee) methodDisplaySalary() {     // '(e Employee)' goi la receiver type
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

// cai tren la value receiver -> thi khong the thay doi gia tri cua obj do 
// -> sinh ra pointer receiver -> co the thay doi dc obj
func (e *Employee) change_salary(salary_new int) {
	e.salary = salary_new
}

// func thi khong ho tro cung ten khac kieu, method thi co .
// ben canh do go k phai oop -> can co struct + method de lm dieu nay 
func funcDisplaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

// for test
func test_struct(){

	// cho struct ao day dc nhung khong cho methodd vao day dc 

	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}

	emp1.methodDisplaySalary()
	fmt.Println()
	funcDisplaySalary(emp1)
	fmt.Println()
	emp1.change_salary(1000)
	emp1.methodDisplaySalary()

}


func test_map() {

	person_age := map[string]int {   // var person_age := err
		"Jack": 24,
		"long":10,
	}

	// fmt.Print(person_age)	

	value, exits_key := person_age["long1"]

	if exits_key  {
		fmt.Print(value)	
	} else {
		fmt.Print("name not found !\n")
		fmt.Print(value)  // default 0 
	}

	// map la kieu tham chieu newPersonAge := person_age (2 map cung tro toi vung nho 0xaaaa ) => thay doi 1/2 thi cung thay doi ca 2 

}


// test interface 
type SalaryCalculator interface {  
    CalculateSalary() int     // method() return_type
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId    int
    basicpay int
}

//salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing  
the salaries of the individual employees  
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func test_interface(){
	pemp1 := Permanent{
        empId:    1,
        basicpay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicpay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicpay: 3000,
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)
}


// ============================================ //
// =======Concurrency, Goroutines, Channels=====//
// Parallel sẽ xử lí nhiều tác vụ cùng một thời điểm
// Concurrency sẽ xử lí nhiều tác vụ một lần   -> can Mutexes, Semaphores, Locks. -> golang : Goroutines 


func main() {
	// test_map()
	// test_pointer()
	// test_struct()
	// test_interface()
}