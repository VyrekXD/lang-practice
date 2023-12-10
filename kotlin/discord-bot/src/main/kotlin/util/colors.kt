package com.discord.bot.util

import dev.kord.common.Color

fun randomColor(): Color {
	return Color((0..0xffffff + 1).random())
}