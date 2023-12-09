fun main() {
	val name1 = "Vyrek"
	val name2 = "Marcock"
	// you can "construct" a string with String("Vyrek".toCharArray())
	val name3 = "Vyrek"

	// You can compare a string with this two but the best way is ===
	//	println(name1 == name3)
	//	println(name1.equals(name3))
	println(name1 === name3)
	println(name1 === name3)
}