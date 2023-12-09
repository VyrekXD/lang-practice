// here we bring PI to our code without needing to use kotlin.math everytime
import kotlin.math.PI
import kotlin.math.abs

// you can also bring everything from a package
//import kotlin.math.*

fun main() {
	val num = -10
	val number = 2

	// operators
	println(num + number)
	println(num - number)
	println(num / number)
	println(num * number)
	println(num % number)

	// included math
	println(PI)
	println(abs(num))

	var newNum = 12
	// to increment by 1 use ++, it also modifies the variable
	// with newNum++ it prints the actual value and then modifies the variable
	println(newNum++) // prints 12 and becomes 13
	println(newNum) // prints 13
	// with ++newNum it first modifies the variable then it prints
	println(++newNum) // becomes 14 and prints 14
	// -- also works the same way
	println(--newNum) // becomes 13 and prints 13

	
}