package com.seventwonine.oivaru

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.blur
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

data class NewsStory(
    val title: String,
    val summary: String,
    val timestamp: String,
    val category: String,
    val imageUrl: String
)

@Composable
fun OivaruFeed(modifier: Modifier = Modifier) {
    val stories = listOf(
        NewsStory("SENTINEL PROTOCOL ACTIVATED", "729 AI confirms the launch of the second-tier news aggregation node.", "2m ago", "CORE", ""),
        NewsStory("MARITIME INDIGO PEAK", "Bioluminescent levels in the Maldivian basin reach record highs.", "15m ago", "NATURE", ""),
        NewsStory("CYBER-GOVERNANCE UPDATE", "New RBAC protocols enforced across the Editorial Nexus.", "1h ago", "TECH", "")
    )

    Box(modifier = modifier.fillMaxSize()) {
        LazyColumn(
            modifier = Modifier.fillMaxSize(),
            contentPadding = PaddingValues(bottom = 24.dp)
        ) {
            item {
                LeadStoryCard()
            }
            items(stories) { story ->
                SecondaryStoryItem(story)
            }
        }
    }
}

@Composable
fun LeadStoryCard() {
    Box(
        modifier = Modifier
            .fillMaxWidth()
            .height(400.dp)
    ) {
        // Placeholder for Hero Image
        Box(modifier = Modifier.fillMaxSize().background(Color.DarkGray))
        
        // Glassmorphism Overlay
        Box(
            modifier = Modifier
                .align(Alignment.BottomCenter)
                .fillMaxWidth()
                .padding(16.dp)
                .clip(RoundedCornerShape(12.dp))
                .background(Color.White.copy(alpha = 0.1f))
                .padding(24.dp)
        ) {
            Column {
                Text(
                    text = "OIVARU: THE PULSE",
                    color = Color(0xFF00FFFF),
                    fontSize = 12.sp,
                    fontWeight = FontWeight.Bold,
                    letterSpacing = 2.sp
                )
                Spacer(modifier = Modifier.height(8.dp))
                Text(
                    text = "MALDIVES DEPLOYS FIRST AGENTIC NEWSROOM",
                    color = Color.White,
                    fontSize = 28.sp,
                    fontWeight = FontWeight.ExtraBold,
                    lineHeight = 34.sp
                )
            }
        }
    }
}

@Composable
fun SecondaryStoryItem(story: NewsStory) {
    Column(
        modifier = Modifier
            .fillMaxWidth()
            .padding(horizontal = 16.dp, vertical = 12.dp)
    ) {
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Text(text = story.category, color = Color(0xFFFFA500), fontSize = 10.sp, fontWeight = FontWeight.Bold)
            Text(text = story.timestamp, color = Color.White.copy(alpha = 0.4f), fontSize = 10.sp)
        }
        Spacer(modifier = Modifier.height(4.dp))
        Text(
            text = story.title,
            color = Color.White,
            fontSize = 18.sp,
            fontWeight = FontWeight.Bold
        )
        Text(
            text = story.summary,
            color = Color.White.copy(alpha = 0.6f),
            fontSize = 14.sp,
            maxLines = 2
        )
        Spacer(modifier = Modifier.height(12.dp))
        HorizontalDivider(color = Color.Cyan.copy(alpha = 0.1f), thickness = 0.5.dp)
    }
}
