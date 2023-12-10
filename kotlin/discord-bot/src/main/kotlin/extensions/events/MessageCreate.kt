package com.discord.bot.extensions.events

import com.kotlindiscord.kord.extensions.extensions.Extension
import com.kotlindiscord.kord.extensions.extensions.event
import com.kotlindiscord.kord.extensions.utils.respond
import dev.kord.core.event.message.MessageCreateEvent

class MessageCreate: Extension() {
	override var name = "MessageCreate"

	override suspend fun setup() {
		event<MessageCreateEvent> {
			action {
				if (event.message.content.lowercase() == "vyrek" ||
					event.message.content.contains("538421122920742942")) {
					event.message.respond("gay :3")
				} else if (event.message.content.lowercase() == "marcock" ||
					event.message.content.contains("507367752391196682")) {
					event.message.respond("cock :3")
				} else if (event.message.content.lowercase() == "marcrock") {
					event.message.respond("marcock* :3")
				}
			}
		}
	}
}