package com.discord.bot

import com.discord.bot.extensions.commands.`fun`.*
import com.discord.bot.extensions.commands.*
import com.discord.bot.extensions.events.*

import com.kotlindiscord.kord.extensions.ExtensibleBot
import com.kotlindiscord.kord.extensions.utils.env
import com.kotlindiscord.kord.extensions.utils.envOrNull
import dev.kord.common.entity.PresenceStatus
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent

private val TOKEN = env("TOKEN")

@OptIn(PrivilegedIntent::class)
suspend fun main() {
	val bot = ExtensibleBot(TOKEN) {
		applicationCommands {
			defaultGuild(envOrNull("TESTING_GUILD"))
		}

		extensions {
			// Commands
			add(::Slap)
			add(::Ping)
			add(::Dice)

			// Events
			add(::MessageCreate)
		}

		intents(true) {
			+Intent.MessageContent
		}

		presence {
			status = PresenceStatus.Invisible

			listening("Kotlin bot goes brrr")
		}
	}

	println("Commands registered :3")

	bot.start()
}