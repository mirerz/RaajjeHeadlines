package com.seventwonine.commonui.network

import com.seventwonine.commonui.models.RaahiResponse
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.POST

interface SentinelApi {
    @POST("v1/raahi/calculate")
    suspend fun calculateRaahi(@Body request: Map<String, String>): Response<RaahiResponse>
}
