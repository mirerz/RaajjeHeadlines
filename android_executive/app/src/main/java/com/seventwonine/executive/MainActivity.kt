package com.seventwonine.executive

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.compose.runtime.*
import com.seventwonine.oivaru.OivaruFeed
import com.seventwonine.executive.ui.SentinelSurface
import com.seventwonine.koshaaru.KoshaaruBrowser

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            Box(modifier = Modifier.fillMaxSize()) {
                // Background Shader Surface
                SentinelSurface()
                
                // Main Content Layout
                Column(modifier = Modifier.fillMaxSize()) {
                    var currentView by remember { mutableStateOf("OIVARU") }

                    // Header Area
                    Row(
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(24.dp),
                        horizontalArrangement = Arrangement.SpaceBetween,
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        Column {
                            Text(
                                text = if (currentView == "OIVARU") "އަދުގެ" else "ކޮށާރު",
                                color = Color.Cyan,
                                fontSize = 32.sp,
                                fontWeight = FontWeight.Bold,
                            )
                            Text(
                                text = "SENTINEL INTERFACE [V2.0]",
                                color = Color.White.copy(alpha = 0.6f),
                                fontSize = 8.sp,
                                letterSpacing = 1.sp
                            )
                        }
                        
                        // Tab Switcher
                        Row(modifier = Modifier.padding(top = 8.dp)) {
                            Text(
                                text = "PULSE",
                                color = if (currentView == "OIVARU") Color.Cyan else Color.Gray,
                                fontSize = 10.sp,
                                fontWeight = FontWeight.Bold,
                                modifier = Modifier.clickable { currentView = "OIVARU" }
                            )
                            Spacer(modifier = Modifier.width(16.dp))
                            Text(
                                text = "LORE",
                                color = if (currentView == "KOSHAARU") Color.Cyan else Color.Gray,
                                fontSize = 10.sp,
                                fontWeight = FontWeight.Bold,
                                modifier = Modifier.clickable { currentView = "KOSHAARU" }
                            )
                        }
                    }

                    // Feed Content
                    Box(modifier = Modifier.weight(1f)) {
                        if (currentView == "OIVARU") {
                            OivaruFeed()
                        } else {
                            KoshaaruBrowser()
                        }
                    }

                    // Footer Telemetry
                    Row(
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(24.dp),
                        horizontalArrangement = Arrangement.SpaceBetween
                    ) {
                        TelemetryItem("INSTITUTIONS", "1.04")
                        TelemetryItem("SOVEREIGNTY", "READY")
                        TelemetryItem("SENTINEL LUME", "98%")
                    }
                }
            }
        }
    }
}

@Composable
fun TelemetryItem(label: String, value: String) {
    Column {
        Text(text = label, color = Color.White.copy(alpha = 0.4f), fontSize = 6.sp, fontWeight = FontWeight.Bold)
        Text(text = value, color = Color.White, fontSize = 14.sp, fontWeight = FontWeight.Medium)
    }
}
