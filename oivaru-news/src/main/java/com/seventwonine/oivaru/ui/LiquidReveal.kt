package com.seventwonine.oivaru.ui

import androidx.compose.animation.*
import androidx.compose.animation.core.*
import androidx.compose.foundation.layout.Box
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.graphicsLayer

@Composable
fun LiquidReveal(
    index: Int,
    content: @Composable () -> Unit
) {
    var visible by remember { mutableStateOf(false) }
    
    LaunchedEffect(Unit) {
        // Staggered entry based on index
        kotlinx.coroutines.delay(index * 100L)
        visible = true
    }

    AnimatedVisibility(
        visible = visible,
        enter = fadeIn(animationSpec = tween(600)) + 
                expandVertically(animationSpec = tween(600, easing = EaseOutExpo)) +
                scaleIn(initialScale = 0.8f, animationSpec = tween(600, easing = EaseOutBack)),
        exit = fadeOut()
    ) {
        Box(
            modifier = Modifier.graphicsLayer {
                // Subtle perspective tilt during entry
                rotationX = if (visible) 0f else 10f
            }
        ) {
            content()
        }
    }
}
