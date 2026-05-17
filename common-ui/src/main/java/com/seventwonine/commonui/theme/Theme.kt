package com.seventwonine.commonui.theme

import androidx.compose.foundation.isSystemInDarkTheme
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.darkColorScheme
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable

private val DarkColorScheme = darkColorScheme(
    primary = BioluminescentCyan,
    secondary = SolarAmber,
    tertiary = MaritimeIndigo,
    background = MaritimeIndigo,
    surface = SurfaceDark,
    onPrimary = MaritimeIndigo,
    onSecondary = MaritimeIndigo,
    onTertiary = TextPrimary,
    onBackground = TextPrimary,
    onSurface = TextPrimary,
)

private val LightColorScheme = lightColorScheme(
    primary = BioluminescentCyan,
    secondary = SolarAmber,
    tertiary = MaritimeIndigo,
    // Note: Project constitution prioritizes a premium dark experience
)

@Composable
fun SentinelTheme(
    darkTheme: Boolean = isSystemInDarkTheme(),
    content: @Composable () -> Unit
) {
    val colorScheme = if (darkTheme) DarkColorScheme else LightColorScheme

    MaterialTheme(
        colorScheme = colorScheme,
        typography = Typography,
        content = content
    )
}
