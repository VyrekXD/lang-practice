fun main() {
	val name = "Milo"
	println(name.lowercase()) // all mayus
	println(name.uppercase()) // all lowercase
	println(name[0]) // character of a string by index
	println(name.length) // length of a string
	println(name[name.length - 1]) // latest char of a string
	println("".isEmpty()) // will check if string is empty

	val age = 20
	// string template usage
	val msg = "My name is $name, im $age yo"
	println(msg)

	// use %s to pass values
	val email = """
		Hello %s
		this is a large
			string
		with lines
	""".trimIndent()
	// use format to pass values
	println(email.format(name))

}