package com.discord.bot.extensions.commands.`fun`

import com.kotlindiscord.kord.extensions.commands.Arguments
import com.kotlindiscord.kord.extensions.commands.converters.impl.user
import com.kotlindiscord.kord.extensions.extensions.Extension
import com.kotlindiscord.kord.extensions.extensions.publicSlashCommand

class Slap : Extension() {
	override var name = "slap"

	override suspend fun setup() {
		publicSlashCommand(::SlapArgs) {
			name = "slap"
			description = "Get slapped!"

			action {
				respond {
					content = "_Slaps ${arguments.target.mention} around a bit with a smelly trout!_"
				}
			}
		}
	}

	inner class SlapArgs: Arguments() {
		val target by user {
			name = "user"
			description = "User to slap"
		}
	}
}