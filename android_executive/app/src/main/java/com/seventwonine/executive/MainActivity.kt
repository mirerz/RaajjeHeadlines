package com.seventwonine.executive

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.seventwonine.commonui.theme.SentinelTheme
import com.seventwonine.executive.ui.RaahiOnboarding
import com.seventwonine.executive.ui.SentinelSurface
import com.seventwonine.koshaaru.KoshaaruBrowser
import com.seventwonine.oivaru.OivaruFeed

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            SentinelTheme {
                var currentView by remember { mutableStateOf("OIVARU") }
                var isOnboarded by remember { mutableStateOf(false) }

                if (!isOnboarded) {
                    RaahiOnboarding(onFinish = { isOnboarded = true })
                } else {
                    Scaffold(
                    bottomBar = {
                        NavigationBar(
                            containerColor = MaterialTheme.colorScheme.surface,
                            contentColor = MaterialTheme.colorScheme.primary
                        ) {
                            NavigationBarItem(
                                selected = currentView == "OIVARU",
                                onClick = { currentView = "OIVARU" },
                                label = { Text("PULSE", style = MaterialTheme.typography.labelSmall) },
                                icon = { /* Icon placeholder */ },
                                colors = NavigationBarItemDefaults.colors(
                                    selectedIconColor = MaterialTheme.colorScheme.primary,
                                    unselectedIconColor = MaterialTheme.colorScheme.onSurface.copy(alpha = 0.5f)
                                )
                            )
                            NavigationBarItem(
                                selected = currentView == "LORE",
                                onClick = { currentView = "KOSHAARU" },
                                label = { Text("LORE", style = MaterialTheme.typography.labelSmall) },
                                icon = { /* Icon placeholder */ }
                            )
                            NavigationBarItem(
                                selected = currentView == "SOVEREIGNTY",
                                onClick = { currentView = "SOVEREIGNTY" },
                                label = { Text("REPUBLIC", style = MaterialTheme.typography.labelSmall) },
                                icon = { /* Icon placeholder */ }
                            )
                        }
                    }
                ) { innerPadding ->
                    Box(modifier = Modifier.padding(innerPadding).fillMaxSize()) {
                        // Background Shader Surface
                        SentinelSurface()
                        
                        Column(modifier = Modifier.fillMaxSize()) {
                            // Unified Executive Header
                            Row(
                                modifier = Modifier
                                    .fillMaxWidth()
                                    .padding(24.dp),
                                horizontalArrangement = Arrangement.SpaceBetween,
                                verticalAlignment = Alignment.CenterVertically
                            ) {
                                Column {
                                    val headerText = when(currentView) {
                                        "OIVARU" -> "އަދުގެ"
                                        "KOSHAARU" -> "ކޮށާރު"
                                        else -> "ޖުމްހޫރީ"
                                    }
                                    Text(
                                        text = headerText,
                                        color = MaterialTheme.colorScheme.primary,
                                        style = MaterialTheme.typography.displayMedium,
                                        fontWeight = FontWeight.ExtraBold,
                                    )
                                    Text(
                                        text = "SENTINEL INTERFACE [V2.2]",
                                        style = MaterialTheme.typography.labelSmall,
                                        color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.4f),
                                        letterSpacing = 2.sp
                                    )
                                }
                            }

                            // Dynamic Content Switcher
                            Box(modifier = Modifier.weight(1f)) {
                                when(currentView) {
                                    "OIVARU" -> OivaruFeed()
                                    "KOSHAARU" -> KoshaaruBrowser()
                                    "SOVEREIGNTY" -> com.seventwonine.executive.ui.SovereigntyDashboard()
                                }
                            }


                            // Footer Telemetry (Sentinel Standard)
                            Row(
                                modifier = Modifier
                                    .fillMaxWidth()
                                    .padding(24.dp),
                                horizontalArrangement = Arrangement.SpaceBetween
                            ) {
                                TelemetryItem("NETWORK", "1.04_LINK")
                                TelemetryItem("STATUS", "STABLE")
                                TelemetryItem("VIBE", "OPTIMIZED")
                            }
                        }
                    }
                    }
                }
            }
        }
    }
}

@Composable
fun TelemetryItem(label: String, value: String) {
    Column {
        Text(
            text = label, 
            color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.4f), 
            style = MaterialTheme.typography.labelSmall,
            fontSize = 8.sp,
            fontWeight = FontWeight.Bold
        )
        Text(
            text = value, 
            color = MaterialTheme.colorScheme.onBackground, 
            style = MaterialTheme.typography.bodyLarge,
            fontWeight = FontWeight.Medium
        )
    }
}
