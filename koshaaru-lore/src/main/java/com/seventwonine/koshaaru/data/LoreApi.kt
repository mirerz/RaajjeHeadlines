package com.seventwonine.koshaaru.data

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import retrofit2.http.GET
import retrofit2.http.Path
import retrofit2.http.Query

@Serializable
data class LoreHotspotResponse(
    val status: String,
    val data: List<LoreHotspotDto>
)

@Serializable
data class LoreDetailResponse(
    val status: String,
    val data: LoreHotspotDto
)

@Serializable
data class LoreHotspotDto(
    val id: String,
    val name: String,
    @SerialName("lore_data") val loreData: String,
    val category: String,
    @SerialName("is_gemini_verified") val isGeminiVerified: Boolean,
    @SerialName("last_awakening") val lastAwakening: String
)

interface LoreService {
    @GET("api/lore")
    suspend fun getHotspots(): LoreHotspotResponse

    @GET("api/lore/{id}")
    suspend fun getLoreDetail(@Path("id") id: String): LoreDetailResponse

    @GET("api/lore/search")
    suspend fun searchLore(@Query("q") query: String): LoreHotspotResponse
}
