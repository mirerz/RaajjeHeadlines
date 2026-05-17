package com.seventwonine.executive.ui

import androidx.compose.animation.*
import androidx.compose.animation.core.*
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.BasicTextField
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.seventwonine.commonui.theme.SentinelTheme
import com.seventwonine.executive.viewmodel.RaahiViewModel

@Composable
fun RaahiOnboarding(onFinish: () -> Unit) {
    val viewModel: RaahiViewModel = viewModel()
    val state by viewModel.state

    SentinelTheme {
        Box(
            modifier = Modifier
                .fillMaxSize()
                .background(Color(0xFF0D1117))
                .padding(32.dp),
            contentAlignment = Alignment.Center
        ) {
            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                verticalArrangement = Arrangement.Center
            ) {
                Text(
                    text = "INITIALIZE HANDSHAKE",
                    style = MaterialTheme.typography.labelSmall,
                    color = MaterialTheme.colorScheme.primary,
                    letterSpacing = 4.sp
                )

                Spacer(modifier = Modifier.height(48.dp))

                AnimatedContent(
                    targetState = state.finalForce == null,
                    transitionSpec = {
                        fadeIn(animationSpec = tween(500)) with fadeOut(animationSpec = tween(500))
                    },
                    label = "OnboardingStep"
                ) { isInputting ->
                    if (isInputting) {
                        InputStep(
                            name = state.name,
                            isCalculating = state.isCalculating,
                            onNameChange = { viewModel.onNameChange(it) },
                            onCalculate = { viewModel.calculateDestiny() }
                        )
                    } else {
                        ResultStep(
                            force = state.finalForce!!,
                            synthesis = state.synthesis ?: "Analyzing patterns...",
                            onJoin = onFinish
                        )
                    }
                }
            }
        }
    }
}

@Composable
fun InputStep(
    name: String,
    isCalculating: Boolean,
    onNameChange: (String) -> Unit,
    onCalculate: () -> Unit
) {
    Column(horizontalAlignment = Alignment.CenterHorizontally) {
        Box(
            modifier = Modifier
                .fillMaxWidth()
                .height(64.dp)
                .border(1.dp, Color.White.copy(alpha = 0.1f), RoundedCornerShape(4.dp))
                .padding(horizontal = 16.dp),
            contentAlignment = Alignment.CenterStart
        ) {
            if (name.isEmpty()) {
                Text(
                    text = "ENTER FULL NAME",
                    style = MaterialTheme.typography.bodyLarge,
                    color = Color.White.copy(alpha = 0.3f)
                )
            }
            BasicTextField(
                value = name,
                onValueChange = onNameChange,
                modifier = Modifier.fillMaxWidth(),
                textStyle = MaterialTheme.typography.bodyLarge.copy(
                    color = Color.White,
                    textAlign = TextAlign.Start
                ),
                singleLine = true
            )
        }

        Spacer(modifier = Modifier.height(24.dp))

        Button(
            onClick = onCalculate,
            enabled = !isCalculating && name.isNotBlank(),
            modifier = Modifier
                .fillMaxWidth()
                .height(56.dp),
            shape = RoundedCornerShape(4.dp),
            colors = ButtonDefaults.buttonColors(
                containerColor = MaterialTheme.colorScheme.primary,
                disabledContainerColor = Color.DarkGray
            )
        ) {
            if (isCalculating) {
                CircularProgressIndicator(modifier = Modifier.size(24.dp), color = Color.Black)
            } else {
                Text("CALCULATE RAAHI", fontWeight = FontWeight.Bold, color = Color.Black)
            }
        }
    }
}

@Composable
fun ResultStep(
    force: Int,
    synthesis: String,
    onJoin: () -> Unit
) {
    Column(horizontalAlignment = Alignment.CenterHorizontally) {
        Text(
            text = "YOUR FINAL FORCE",
            style = MaterialTheme.typography.labelSmall,
            color = Color.White.copy(alpha = 0.5f)
        )
        Text(
            text = force.toString(),
            style = MaterialTheme.typography.displayLarge,
            color = MaterialTheme.colorScheme.primary,
            fontSize = 80.sp,
            fontWeight = FontWeight.Black
        )

        Spacer(modifier = Modifier.height(32.dp))

        Box(
            modifier = Modifier
                .fillMaxWidth()
                .background(Color.White.copy(alpha = 0.05f), RoundedCornerShape(8.dp))
                .padding(24.dp)
        ) {
            Text(
                text = synthesis,
                style = MaterialTheme.typography.bodyMedium,
                color = Color.White.copy(alpha = 0.8f),
                lineHeight = 24.sp,
                textAlign = TextAlign.Center
            )
        }

        Spacer(modifier = Modifier.height(48.dp))

        Button(
            onClick = onJoin,
            modifier = Modifier
                .fillMaxWidth()
                .height(56.dp),
            shape = RoundedCornerShape(4.dp),
            colors = ButtonDefaults.buttonColors(containerColor = Color.White)
        ) {
            Text("JOIN THE REPUBLIC", fontWeight = FontWeight.Bold, color = Color.Black)
        }
    }
}
