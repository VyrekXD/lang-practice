package com.discord.bot.extensions.commands

import com.kotlindiscord.kord.extensions.extensions.Extension
import com.kotlindiscord.kord.extensions.extensions.publicSlashCommand

class Ping: Extension() {
	override var name = "ping"

	override suspend fun setup() {
		publicSlashCommand() {
			name = "ping"
			description = "Pong!"

			action {
				respond {
					content = "Pong! `${bot.kordRef.gateway.averagePing}`ms"
				}
			}
		}
	}
}