package com.seventwonine.koshaaru

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

data class LoreHotspot(
    val id: String,
    val name: String,
    val description: String,
    val density: Float
)

@Composable
fun KoshaaruBrowser(modifier: Modifier = Modifier) {
    val hotspots = listOf(
        LoreHotspot("1", "Endheri Odi", "Tactical migration patterns of regional fleets.", 0.85f),
        LoreHotspot("2", "Majlis Pulse", "Sentiment analysis of Legislative shifts.", 0.62f),
        LoreHotspot("3", "Blue Economy", "Deep-dive into maritime resilience projects.", 0.94f)
    )

    Column(
        modifier = modifier
            .fillMaxSize()
            .background(Color(0xFF070915))
            .padding(16.dp)
    ) {
        Text(
            text = "KOSHAARU LORE BROWSER",
            color = Color.Cyan,
            fontSize = 12.sp,
            fontWeight = FontWeight.Bold,
            modifier = Modifier.padding(bottom = 16.dp)
        )

        LazyColumn(verticalArrangement = Arrangement.spacedBy(16.dp)) {
            items(hotspots) { hotspot ->
                LoreCard(hotspot)
            }
        }
    }
}

@Composable
fun LoreCard(hotspot: LoreHotspot) {
    Column(
        modifier = Modifier
            .fillMaxWidth()
            .background(Color(0xFF1A1A2E))
            .padding(16.dp)
    ) {
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.SpaceBetween
        ) {
            Text(text = hotspot.name, color = Color.White, fontWeight = FontWeight.Bold, fontSize = 18.sp)
            Text(text = "${(hotspot.density * 100).toInt()}% DENSITY", color = Color.Magenta, fontSize = 10.sp)
        }
        Spacer(modifier = Modifier.height(8.dp))
        Text(text = hotspot.description, color = Color.White.copy(alpha = 0.7f), fontSize = 14.sp)
    }
}
