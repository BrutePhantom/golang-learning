package main 

import "fmt"

func main(){
//Imported these universally for all programs 


// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.


	var a int
	fmt.Scan(&a)

	if a >0{
		fmt.Println("Positive")
	}else if a<0{
		fmt.Println("Negative")
	}else{
		fmt.Println("Zero")
	}

// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================

	var b int 
	fmt.Scan(&b)
	if b>=18{
		fmt.Println("Aap vote kr sakte hain")
	}else{
		fmt.Println("Aap vote nhi kr sakte h")
	}


// =========================================================

// 3. Write a program that checks whether a given number is even or odd.

// =========================================================

	var c int
	fmt.Scan(&c)
	if c % 2 == 0{
		fmt.Println("Even hai ye number")
	}else{
		fmt.Println("Odd hai ye number")
	}


// =========================================================

// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================

	var d int
	fmt.Scan(&d)
	if d >= 33{
		fmt.Println("Bach gya launde")
	}else{
		fmt.Println("Tu tho fail hogya bhai")
	}


// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================

	var e int
	fmt.Scan(&e)
	if (e % 4 == 0 && e % 100 != 0) || (e % 400 == 0){
		fmt.Println("Leap year hai ye")
	}else{
		fmt.Println("Leap year nhi hai ye")
	}


// =========================================================


// -----------------------------------------------------------
// Section B: Medium (Logic Building)
// -----------------------------------------------------------

// 1. Write a program to calculate grade based on marks using if-else if conditions.
//    90+ → Grade A
//    75–89 → Grade B
//    60–74 → Grade C
//    Below 60 → Grade F

// =========================================================

	var score int 
	fmt.Scan(&score)
	if score >= 90{
		fmt.Println("Grade A")
	}else if score >= 80{
		fmt.Println("Grade B")
	}else if score >= 70{
		fmt.Println("Grade C")
	}else if score >= 60{
		fmt.Println( "Grade D")
	}else{
		fmt.Println("Grade F")
	}


// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================

	var m1 int 
	var m2 int

	fmt.Scan(&m1,&m2)
	if m1 > m2{
		fmt.Printf("%d is bada hai",m1)
	}else if (m2>m1){
		fmt.Printf("%d is bada hai",m2)
	}else{
		fmt.Println("Dono barabar hai")
	}


// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================

	var meow string
	fmt.Scan(&meow)
	switch meow{
	case "a":
		fmt.Println("Vowel hai")
	case "e":
		fmt.Println("Vowel hai")
	case "i":
		fmt.Println("Vowel hai")
	case "o":
		fmt.Println("Vowel hai")
	case "u":
		fmt.Println("Vowel hai")
	case "A":
		fmt.Println("Vowel hai")
	case "E":
		fmt.Println("Vowel hai")
	case "I":
		fmt.Println("Vowel hai")
	case "O":
		fmt.Println("Vowel hai")
	case "U":
		fmt.Println("Vowel hai")
	default:
		fmt.Println("Yo tho consonant hai")
	}


// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================

	var day string
	fmt.Scan(&day)
	switch day{
		case "Monday":
			fmt.Println("Weekday hai")
		case "Tuesday":
			fmt.Println("Weekday hai")
		case "Wednesday":
			fmt.Println("Weekday hai")
		case "Thursday":
			fmt.Println("Weekday hai")
		case "Friday":
			fmt.Println("Weekday hai")
		case "Saturday":
			fmt.Println("Weekend hai")
		case "Sunday":
			fmt.Println("Weekend hai")
		default:
			fmt.Println("Invalid day")
		}


// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================


	var n int
	fmt.Scan(&n)
	if n % 3 == 0 && n % 5 == 0{
		fmt.Println("Dono se divisible hai")
	}else{
		fmt.Println("Dono se divisible nhi h")
	}

// =========================================================


// // -----------------------------------------------------------
// Section C: Hard (Real-World Logic & Nesting)
// -----------------------------------------------------------

// Note: Use the following syntax to take input from the user in Go:
//
// var variableName datatype
// fmt.Print("Enter something: ")
// fmt.Scan(&variableName)
//
// Example:
// var age int
// fmt.Print("Enter your age: ")
// fmt.Scan(&age)


// =========================================================

// 1. Write a program to simulate a login system:
//    - Ask the user to enter a username and password.
//    - If both match predefined values, print “Login successful”.
//    - If username is correct but password is wrong, print “Invalid password”.
//    - Otherwise, print “User not found”.

// Hint: You can compare input strings using the == operator.


// =========================================================


	var userkanaam string
	var accountkichabi string

	const username string = "prabhav"
	const password string = "MeowisOP123#"

	fmt.Println("Apne account ka nam enter kre")
	fmt.Scan(&userkanaam)
	fmt.Println("Apne account ka password enter kre")
	fmt.Scan(&accountkichabi)

	if userkanaam == username && accountkichabi == password{
		fmt.Println("Login successful , Aapka swagat hai")
	}else if userkanaam == username && accountkichabi != password{
		fmt.Println("Invalid password , krapya chabi badle")
	}else{
		fmt.Println("User not found , bhag ja")
	}

// =========================================================

// 2. Write a program that checks whether a given triangle is:
//    - Equilateral (all sides equal)
//    - Isosceles (two sides equal)
//    - Scalene (all sides different)
//
//    - Ask the user to enter the lengths of all three sides.
//    - Use nested if statements to determine the triangle type.

// =========================================================

	var side1 int
	var side2 int
	var side3 int
	fmt.Println("triangle ke 3 sides ke lengths enter kre")
	fmt.Scan(&side1,&side2,&side3)
	if side1 == side2{
		if side2 == side3{
			fmt.Println("ye tho equilateral triangle hai")
		}else{
			fmt.Println("ye tho isosceles triangle hai")
		}
	}else if side2 == side3{
		fmt.Println("ye tho isosceles triangle hai")
	}else if side1 == side3{
		fmt.Println(" ye tho isosceles triangle hai")
	}else{
		fmt.Println("scalene triangle hai ye tho")
	}


// =========================================================

// 3. Write a program that simulates a menu system using a switch statement:
//    - Display a simple menu:
//         1 → “Start Game”
//         2 → “Load Game”
//         3 → “Exit”
//    - Ask the user to enter their choice (1, 2, or 3).
//    - Use a switch statement to print the corresponding message.
//    - If input doesn’t match any option, print “Invalid option”.

// =========================================================


	var game int
	fmt.Println("Kya krna h batao:")
	fmt.Println("1 → Start Game")
	fmt.Println("2 → Load Game")
	fmt.Println("3 → Exit")
	fmt.Scan(&game)
	switch game{
	case 1:
		fmt.Println("Game start ho rha hai")
	case 2:
		fmt.Println("Game load ho rha hai")
	case 3:
		fmt.Println("Game band hora h")
	default:
		fmt.Println("bawla h ke , dobara try kr")
	}

// =========================================================

// 4. Write a program that uses a switch without an expression to classify temperature:
//    - Ask the user to input the current temperature (integer).
//    - Based on the value, print one of the following:
//         Below 0 → “Freezing”
//         0–15 → “Cold”
//         16–30 → “Warm”
//         Above 30 → “Hot”

// Hint: Use logical operators (>, <, <=, >=) inside switch cases.

// =========================================================

	var temp int
	fmt.Println("Abhi temperature kitna hai")
	fmt.Scan(&temp)
	switch {
	case temp < 0:
		fmt.Println("bht jada thand hai")
	case temp <= 15:
		fmt.Println("thand hai")
	case temp <= 30:
		fmt.Println("medium hai")
	case temp > 30:
		fmt.Println("garam hai")
	default:
		fmt.Println("Kripya thermometer lae")
	}


// =========================================================

// 5. Write a program to check admission eligibility for a student based on nested conditions:
//    - Ask the user to input total marks (in percentage).
//    - Ask if the student passed Math and Science (true/false).
//    - Conditions:
//         Minimum marks: 60%
//         Must have passed both Math and Science
//    - If both conditions are true, print “Eligible for admission”.
//    - Otherwise, print the specific reason for rejection.

// =========================================================

	var marks int
	var Math bool
	var Science bool
	fmt.Println("percentage bta bhai")
	fmt.Scan(&marks)
	fmt.Println("maths mein pass h ya fail (true/false)")
	fmt.Scan(&Math)
	fmt.Println("science mein pass h ya fail (true/false)")
	fmt.Scan(&Science)
	if marks >= 60{
		if Math == true && Science == true{
			fmt.Println("admission mil jaega")
		}else if Math == false && Science == true{
			fmt.Println("Maths me fail ho , isliye admission nhi milega")
		}else if Math == true && Science == false{
			fmt.Println("Science me fail ho , isliye admission nhi milega")
		}else{
			fmt.Println("Sorry aap thode jada gawar ho , isliye admission nhi milega")
		}
	}


// =========================================================


// -----------------------------------------------------------
// Bonus (Challenge Zone)
// -----------------------------------------------------------

// 1. Write a program that takes three numbers as input and prints the largest using only conditional statements.

// Hint: Use nested if or multiple if-else conditions.


// =========================================================

	var num1 int
	var num2 int
	var num3 int
	fmt.Println("3 number kripya krke enter kre")
	fmt.Scan(&num1,&num2,&num3)
	if num1 >= num2{
		if num1 >= num3{
			fmt.Printf("%d bada hai",num1)
		}else{
			fmt.Printf("%d bada hai",num3)
		}
	}else{
		if num2 >= num3{
			fmt.Printf("%d bada hai",num2)
		}else{
			fmt.Printf("%d bada hai",num3)
		}
	}


// =========================================================

// 2. Create a small calculator using a switch statement:
//    - Ask the user to input two numbers and an operator (+, -, *, /).
//    - Perform the operation based on the chosen operator.
//    - Display the result.
//    - If the operator is invalid, print “Invalid operator”.

// =========================================================

	var num11 int
	var num21 int
	var operator string
	fmt.Println("Calculator me apka swagat hai")
	fmt.Println("Pehla number daale")
	fmt.Scan(&num11)
	fmt.Println("Dusra number daale")
	fmt.Scan(&num21)
	fmt.Println("Operator daale (+, -, *, /)")
	fmt.Scan(&operator)
	switch operator{
	case "+":
		fmt.Printf("Result: %d",num11+ num21)
	case "-":
		fmt.Printf("Result: %d",num11 - num21)
	case "*":
		fmt.Printf("Result: %d",num11 * num21)
	case "/":
		if num21 != 0{
			fmt.Printf("Result: %d",num11 / num21)
		}else{
			fmt.Println("Zero se divide nhi kr skte")
		}
	default:
		fmt.Println("Invalid operator")
	}


// =========================================================

// 3. Write a program that determines whether a given year is a century leap year or not.
//    - Ask the user to enter a year.
//    - A year is a century leap year if it is divisible by 400.
//    - A year is a normal leap year if it is divisible by 4 but not by 100.
//    - Otherwise, it is not a leap year.

// =========================================================

	var year int
	fmt.Println("Year enter kro")
	fmt.Scan(&year)
	if year % 400 == 0{
		fmt.Println("Century leap year hai ye")
	}else if year % 4 == 0 && year % 100 != 0{
		fmt.Println("Normal leap year hai ye")
	}else{
		fmt.Println("Leap year nhi hai ye")
	}


// =========================================================

// 4. Write a program that uses fallthrough in a switch to demonstrate cumulative conditions:
//    - Example: Enter a number and print messages for all cases up to that number using fallthrough.
//    - For example, if number = 2, output should print case 1 and case 2 messages.

// =========================================================

	var meow1 int
	fmt.Println("Ek number enter kro (1-5)")
	fmt.Scan(&meow1)
	switch meow1{
	case 1:
		fmt.Println("Case 1 m apka swagat hai")
		fallthrough
	case 2:
		fmt.Println("Case 2 m apka swagat hai")
		fallthrough
	case 3:
		fmt.Println("Case 3 m apka swagat hai")
		fallthrough
	case 4:
		fmt.Println("Case 4 m apka swagat hai")
		fallthrough
	case 5:
		fmt.Println("Case 5 m apka swagat hai")
	default:
		fmt.Println("Invalid case number")
	}


// =========================================================

// 5. Write a program that checks if a student qualifies for a scholarship:
//    - Ask for the student’s marks (percentage).
//    - Ask for attendance percentage.
//    - Ask if the student has any backlogs (true/false).
//    - Conditions:
//         Must have 85% or higher marks
//         Attendance above 90%
//         No backlog
//    - If all conditions are true, print “Scholarship Approved”.
//    - Otherwise, print “Scholarship Denied”.

// =========================================================

	var marks1 int
	var attendance int
	var backlog bool
	fmt.Println("percentage daalo")
	fmt.Scan(&marks1)
	fmt.Println("Attendance daalo")
	fmt.Scan(&attendance)
	fmt.Println("backlog lagi hai kya (true/false)")
	fmt.Scan(&backlog)
	if marks1 >= 85{
		if attendance > 90{
			if backlog == false{
				fmt.Println("Scholarship mil jaegi aapko")
			}else{
				fmt.Println("Scholarship Denied , back clear kro apni")
			}
		}else{
			fmt.Println("Scholarship Denied ,class attend nhi krte ho")
		}
	}else{
		fmt.Println("Scholarship Denied, marks kam hai")
	}
}
