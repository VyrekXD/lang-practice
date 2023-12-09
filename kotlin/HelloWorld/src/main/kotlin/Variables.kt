// const val are used as a global constant in the whole file, they are declared before the code executes
const val everywhere = "for all time, always."

fun main() {
	// IMPORTANT: val are constants, var can mutate

	// Variables cannot contain reserved keywords
	val hello = "world"
	println(hello)

	// Variables can be typed, stricting types into the variable
	val name: String = "Milo"
	val age: Int = 17

	// println(name)
	// println(age)

	// Strings can contain template strings, just like TS
	// Curly braces are used when you do an operation just like ${age + 18}, if not just use $age
	println("Name: $name Age: $age")

	// Just like ts you dont specify the type of a variable first, you do the keyword
	// ts: let idk: int = 10

	// int saves smaller numbers
	val number: Int = 10

	// long saves bigger numbers, needs an L at the end
	val longernum: Long = 100L

	// if u want decimals use doubles
	val decimals: Double = 3.1416

	// if u want bigger decimals use floats, needs an L at the end
	val floats: Float = 3.1416F

	// booleans use boolean
	val bool: Boolean = false

	// strings use String
	val text: String = "hello"

	// you can also declare single chars, they need '' not ""
	val chars: Char = 'A'

	// IMPORTANT: Kotlin has type inference

	val color = "red"
	// you cant do color = null, if you try to do color = null you will get a null pointer exception
	println(color.uppercase())

	// to make a variable "nullable" you need to use the ?
	var brand: String? = "dior"
	println(brand)
	brand = null

	// to do something with a variable IF exists use ? just like ts
	//	println(brand?.length)
}