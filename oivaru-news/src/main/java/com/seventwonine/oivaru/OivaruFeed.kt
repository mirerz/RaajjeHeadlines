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

import com.seventwonine.commonui.theme.SentinelTheme
import com.seventwonine.commonui.theme.BioluminescentCyan
import com.seventwonine.commonui.theme.SolarAmber
import com.seventwonine.commonui.theme.GlassBackground
import com.seventwonine.commonui.theme.GlassBorder

@Composable
fun OivaruFeed(modifier: Modifier = Modifier) {
    val stories = listOf(
        NewsStory("SENTINEL PROTOCOL ACTIVATED", "729 AI confirms the launch of the second-tier news aggregation node.", "2m ago", "CORE", ""),
        NewsStory("MARITIME INDIGO PEAK", "Bioluminescent levels in the Maldivian basin reach record highs.", "15m ago", "NATURE", ""),
        NewsStory("CYBER-GOVERNANCE UPDATE", "New RBAC protocols enforced across the Editorial Nexus.", "1h ago", "TECH", ""),
        NewsStory("LORE AWAKENING IN PROGRESS", "Deep-scan algorithms detecting legacy hull DNA in the central kosh.", "3h ago", "LORE", "")
    )

    SentinelTheme {
        Box(modifier = modifier.fillMaxSize().background(MaterialTheme.colorScheme.background)) {
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
}

import com.seventwonine.commonui.components.GlassCard

@Composable
fun LeadStoryCard() {
    Box(
        modifier = Modifier
            .fillMaxWidth()
            .height(400.dp)
    ) {
        // Placeholder for Hero Image with a gradient
        Box(
            modifier = Modifier
                .fillMaxSize()
                .background(
                    brush = androidx.compose.ui.graphics.Brush.verticalGradient(
                        colors = listOf(
                            MaterialTheme.colorScheme.tertiary,
                            MaterialTheme.colorScheme.background
                        )
                    )
                )
        )
        
        // Glassmorphism Overlay
        GlassCard(
            modifier = Modifier
                .align(Alignment.BottomCenter)
                .fillMaxWidth()
                .padding(16.dp)
        ) {
            Column(modifier = Modifier.padding(24.dp)) {
                Text(
                    text = "OIVARU: THE PULSE",
                    color = MaterialTheme.colorScheme.primary,
                    style = MaterialTheme.typography.labelSmall,
                    fontWeight = FontWeight.Bold,
                    letterSpacing = 2.sp
                )
                Spacer(modifier = Modifier.height(8.dp))
                Text(
                    text = "MALDIVES DEPLOYS FIRST AGENTIC NEWSROOM",
                    color = MaterialTheme.colorScheme.onSurface,
                    style = MaterialTheme.typography.displaySmall,
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
            .padding(horizontal = 24.dp, vertical = 16.dp)
    ) {
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Text(
                text = story.category, 
                color = MaterialTheme.colorScheme.secondary, 
                style = MaterialTheme.typography.labelSmall, 
                fontWeight = FontWeight.Bold
            )
            Text(
                text = story.timestamp, 
                color = MaterialTheme.colorScheme.onSurface.copy(alpha = 0.4f), 
                style = MaterialTheme.typography.labelSmall
            )
        }
        Spacer(modifier = Modifier.height(8.dp))
        Text(
            text = story.title,
            color = MaterialTheme.colorScheme.onSurface,
            style = MaterialTheme.typography.titleLarge,
            fontWeight = FontWeight.Bold
        )
        Spacer(modifier = Modifier.height(4.dp))
        Text(
            text = story.summary,
            color = MaterialTheme.colorScheme.onSurface.copy(alpha = 0.7f),
            style = MaterialTheme.typography.bodyMedium,
            maxLines = 2
        )
        Spacer(modifier = Modifier.height(16.dp))
        HorizontalDivider(
            color = MaterialTheme.colorScheme.primary.copy(alpha = 0.1f), 
            thickness = 0.5.dp
        )
    }
}
