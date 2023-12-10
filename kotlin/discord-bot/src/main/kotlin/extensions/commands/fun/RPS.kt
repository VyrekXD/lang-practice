package com.discord.bot.extensions.commands.`fun`

import com.discord.bot.util.randomColor
import com.kotlindiscord.kord.extensions.commands.Arguments
import com.kotlindiscord.kord.extensions.commands.application.slash.converters.impl.numberChoice
import com.kotlindiscord.kord.extensions.commands.application.slash.converters.impl.stringChoice
import com.kotlindiscord.kord.extensions.extensions.Extension
import com.kotlindiscord.kord.extensions.extensions.publicSlashCommand
import dev.kord.rest.builder.message.embed

val RPSResponses: Set<String> = setOf(
	"Aww man... I lost, i chose **%s** and you chose **%s**, _%s loses to %s_.",
	"Yay! I won, i chose **%s** and you chose **%s**, _%s beats %s_.",
	 "Haha, we were tied, _%s ties with %s_."
)

val RPSNames: Set<String> = setOf(
	"Rock",
	"Paper",
	"Scissors"
)

class RPS: Extension() {
	override val name = "rockpaperscissors"

	override suspend fun setup() {
		publicSlashCommand(::RPSArguments) {
			name = "rockpaperscissors"
			description = "Play a game of rock paper scissors with me"

			action {
				val num = (0..2).random()
				val input = arguments.choice.toInt()
				val res = if ((num == 0 && input == 1) || (num == 1 && input == 2) || (num == 2 && input == 0)) {
					// user wins
					0
				} else if ((num == 0 && input == 2) || (num == 1 && input == 0) || num == 2 && input == 1) {
					// bot wins
					1
				} else {
					// tie
					2
				}

				respond {
					embed {
						author {
							name = user.asUser().username
							icon = user.asUser().avatar?.cdnUrl?.toUrl()
						}

						color = randomColor()

						description = String.format(
							RPSResponses.elementAt(res),
							RPSNames.elementAt(num),
							RPSNames.elementAt(input),
							RPSNames.elementAt(num),
							RPSNames.elementAt(input),
						)
					}
				}
			}
		}
	}

	inner class RPSArguments: Arguments() {
		private val rpsChoices: MutableMap<String, Long> = mutableMapOf(
			"Rock" to 0,
			"Paper" to 1,
			"Scissors" to 2,
		)

		val choice by numberChoice {
			name = "choice"
			description = "Choose between rock, paper or scissors."

			choices = rpsChoices
		}
	}
}