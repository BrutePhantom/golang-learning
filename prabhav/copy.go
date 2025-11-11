
//class ki majdori jo arpit bhaiya ne karai aaj k session m 


package main 

import "fmt"

func main() {

	//Majdoori No.1
	fmt.Println("Hello, World!")

	fmt.Println();

	//Majdoori No.2
	var s string = "Prabhav"

	var st = "Prabhav"

	st1 := "Prabhav"

	var(
		i ="Prabhav"
		j ="Prabhav"
	)

	const meow = "Prabhav"

	fmt.Println(s, st, st1,i,j,meow)

	fmt.Println();

	//Majdoori No.3
	var b int = 20
	if b >= 18{
		fmt.Println("You are not a child anymore")
	}

	fmt.Println();

	//Majdoori No.4
	var age int = 10
	if age>=18{
		fmt.Println("You are not a child anymore")
	}else{
		fmt.Println("Abhi bache ho tum")
	}

	fmt.Println();

	//Majdoori No.5
	var age1 int = 65
	if age1<18{
		fmt.Println("abhi aap bbache ho")
	}else if age1>=18 && age1<65{
		fmt.Println("aap bache ab nhi rahe")
	}else{
		fmt.Println("dadaji namaste")
	}

	fmt.Println();

	//Majdoori No.6
	var day = "Monday"
	switch day {
	case "Monday":
		fmt.Println("Aaj se JIMMM")
	case "Wednesday":
		fmt.Println("Aaj thoda light krunga")
	case "Friday":
		fmt.Println("Aaj tho friday h aaj kon krta h")
	case "Sunday":
		fmt.Println("Kal monday h aaj tho soounga")
	default:
		fmt.Println("Aaj kuch ni")
	}

	fmt.Println();

	//Majdoori No.7
	for i := 0; i <= 11; i++ {
		fmt.Println("Mera Priya Mitra No.-", i)
	}

	fmt.Println();

	//Majdoori No.8
	ik := 0
	for ik <= 10 {
		fmt.Println("Meow")
		ik++
	}

	fmt.Println();

	//Majdoori No.9
	nums := []int{10,20,30}
	for _,value := range nums {
		fmt.Println(value)
	}
}