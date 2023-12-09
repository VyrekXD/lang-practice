package com.discord.bot.extensions.commands.`fun`

import com.kotlindiscord.kord.extensions.commands.Arguments
import com.kotlindiscord.kord.extensions.commands.application.slash.converters.impl.optionalNumberChoice
import com.kotlindiscord.kord.extensions.commands.converters.impl.optionalInt
import com.kotlindiscord.kord.extensions.extensions.Extension
import com.kotlindiscord.kord.extensions.extensions.publicSlashCommand
import dev.kord.common.Color
import dev.kord.rest.builder.message.embed

val DiceColors: Map<Number, String> = mapOf(
	4 to "24A146",
	6 to "27BCD1",
	8 to "9334E6",
	10 to "E13294",
	12 to "D93025",
	20 to "F36D00"
)

val DiceUrls: Map<Number, String> = mapOf(
	4 to "https://www.google.com/logos/fnbx/polyhedral_dice/d4_blank.png",
	6 to "https://www.google.com/logos/fnbx/polyhedral_dice/d6_blank.png",
	8 to "https://www.google.com/logos/fnbx/polyhedral_dice/d8_blank.png",
	10 to "https://www.google.com/logos/fnbx/polyhedral_dice/d10_blank.png",
	12 to "https://www.google.com/logos/fnbx/polyhedral_dice/d12_blank.png",
	20 to "https://www.google.com/logos/fnbx/polyhedral_dice/d20_blank.png"
)

class Dice: Extension() {
	override val name = "dice"

	override suspend fun setup() {
		publicSlashCommand(::DiceArguments) {
			name = "dice"
			description = "Get a random number from different types of dice"

			action {
				val type = arguments.type?.toInt() ?: 6
				val quantity = arguments.quantity ?: 1
				val results: MutableList<String> = mutableListOf()

				if (quantity > 1) {
					for (i in 0 until quantity) {
						results.add("Dice `${i + 1}` result is **${(1..type).random()}**")
					}
				} else {
					results.add("Your number is **${(1..type).random()}**")
				}

				respond {
					embed {
						author {
							name = user.asUser().username
							icon = user.asUser().avatar?.cdnUrl?.toUrl()
						}

						thumbnail {
							url = DiceUrls[type]!!
						}

						color = Color(DiceColors[type]!!.toInt(16))

						description = results.joinToString("\n")

						footer {
							text =
								"Numbers chosen from 1 to $type ${if (quantity > 1) "for $quantity dices" else ""}"
						}
					}
				}
			}
		}
	}

	inner class DiceArguments: Arguments() {
		private val diceChoices: MutableMap<String, Long> = mutableMapOf(
			"4" to 4,
			"6 (Default)" to 6,
			"8" to 8,
			"10" to 10,
			"12" to 12,
			"20" to 20
		)

		val type by optionalNumberChoice {
			name = "type"
			description = "Type of dice to roll"

			choices = diceChoices
		}

		val quantity by optionalInt {
			name = "quantity"
			description = "Quantity of dices to roll (All dices will be the same type)"

			minValue = 1
			maxValue = 10
		}
	}
}