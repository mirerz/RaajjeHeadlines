package com.seventwonine.executive.ui

import android.graphics.RuntimeShader
import android.os.Build
import androidx.annotation.RequiresApi
import androidx.compose.foundation.Canvas
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.ShaderBrush
import androidx.compose.ui.platform.LocalContext
import com.seventwonine.executive.R
import kotlinx.coroutines.isActive

@RequiresApi(Build.VERSION_CODES.TIRAMISU)
@Composable
fun SentinelSurface(modifier: Modifier = Modifier) {
    val context = LocalContext.current
    val shaderCode = remember {
        context.resources.openRawResource(R.raw.bioluminescent_pulse).bufferedReader().use { it.readText() }
    }
    
    val runtimeShader = remember { RuntimeShader(shaderCode) }
    var time by remember { mutableStateOf(0f) }

    LaunchedEffect(Unit) {
        val startTime = System.currentTimeMillis()
        while (isActive) {
            time = (System.currentTimeMillis() - startTime) / 1000f
            runtimeShader.setFloatUniform("uTime", time)
            kotlinx.coroutines.delay(16) // ~60fps
        }
    }

    Canvas(modifier = modifier.fillMaxSize()) {
        runtimeShader.setFloatUniform("uResolution", size.width, size.height)
        drawRect(ShaderBrush(runtimeShader))
    }
}
