package com.raajje.headlines.core_data

import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.delay
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext
import java.net.URL
import javax.xml.parsers.DocumentBuilderFactory

/**
 * MaadhaaRepository: Data integration for Maadhaa27 Live.
 * Aggregates political news and audio broadcasts.
 */
class MaadhaaRepository {

    // Newsdata.io integration (Placeholder API Key)
    private val newsDataApiKey = "pub_placeholder"
    
    /**
     * Fetches political news for the "Side-by-Side" balanced layout.
     * Hits the new Cloud Run /api/feed endpoint.
     */
    fun getBalancedNewsStream(): Flow<Pair<List<MaadhaaNewsItem>, List<MaadhaaNewsItem>>> = flow {
        while (true) {
            val news = try {
                fetchProductionFeed()
            } catch (e: Exception) {
                emptyList()
            }

            // Splitting into balanced columns for Maadhaa27 UI
            val leftWing = news.filterIndexed { index, _ -> index % 2 == 0 }
            val rightWing = news.filterIndexed { index, _ -> index % 2 != 0 }
            
            emit(Pair(leftWing, rightWing))
            delay(300000) // 5 minutes refresh
        }
    }

    private suspend fun fetchProductionFeed(): List<MaadhaaNewsItem> = withContext(Dispatchers.IO) {
        val items = mutableListOf<MaadhaaNewsItem>()
        try {
            val url = URL(NetworkConfig.FEED_ENDPOINT)
            val connection = url.openConnection() as java.net.HttpURLConnection
            connection.requestMethod = "GET"
            
            if (connection.responseCode == 200) {
                val response = connection.inputStream.bufferedReader().use { it.readText() }
                val jsonArray = org.json.JSONArray(response)
                for (i in 0 until jsonArray.length()) {
                    val obj = jsonArray.getJSONObject(i)
                    items.add(MaadhaaNewsItem(
                        id = obj.getString("id"),
                        title = obj.getString("rephrased_headline_dv"),
                        source = obj.optString("section", "National"),
                        summary = obj.optString("category", "General")
                    ))
                }
            }
        } catch (e: Exception) {
            // Handle error
        }
        items
    }

    /**
     * Fetches the latest Clubhouse episodes and live stream status.
     * Integrates with an RSS feed for episode archives.
     */
    fun getAudioHubData(rssUrl: String = "https://example.com/clubhouse/rss"): Flow<AudioHubState> = flow {
        while (true) {
            val episodes = try {
                fetchRssEpisodes(rssUrl)
            } catch (e: Exception) {
                // Fallback to simulated data if network fails
                List(8) { i ->
                    ClubhouseEpisode(
                        id = "ep_$i",
                        title = "Maadhaa27 Episode ${200 - i} (Cached)",
                        duration = "45:00",
                        date = "Oct ${10 - i}, 2024",
                        audioUrl = ""
                    )
                }
            }
            emit(AudioHubState(
                isLive = true, // Would typically check a live status API
                liveTitle = "LIVE: National Sovereignty Discussion",
                episodes = episodes.take(8)
            ))
            delay(600000) // 10 minutes refresh
        }
    }

    private suspend fun fetchRssEpisodes(url: String): List<ClubhouseEpisode> = withContext(Dispatchers.IO) {
        val episodes = mutableListOf<ClubhouseEpisode>()
        try {
            val factory = DocumentBuilderFactory.newInstance()
            val builder = factory.newDocumentBuilder()
            val doc = builder.parse(URL(url).openStream())
            val items = doc.getElementsByTagName("item")
            
            for (i in 0 until items.length) {
                val item = items.item(i)
                val title = item.childNodes.item(0).textContent // Simplified XML parsing
                episodes.add(ClubhouseEpisode(
                    id = "rss_$i",
                    title = title,
                    duration = "Unknown",
                    date = "Recent",
                    audioUrl = ""
                ))
            }
        } catch (e: Exception) {
            // Handle parsing error
        }
        episodes
    }
}

data class MaadhaaNewsItem(
    val id: String,
    val title: String,
    val source: String,
    val summary: String
)

data class ClubhouseEpisode(
    val id: String,
    val title: String,
    val duration: String,
    val date: String,
    val audioUrl: String
)

data class AudioHubState(
    val isLive: Boolean,
    val liveTitle: String?,
    val episodes: List<ClubhouseEpisode>
)
