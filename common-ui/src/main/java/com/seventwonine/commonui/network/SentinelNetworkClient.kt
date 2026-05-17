package com.seventwonine.commonui.network

import com.jakewharton.retrofit2.converter.kotlinx.serialization.asConverterFactory
import kotlinx.serialization.json.Json
import okhttp3.MediaType.Companion.toMediaType
import retrofit2.Retrofit

object SentinelNetworkClient {
    // 10.0.2.2 is the special IP to access the host machine's localhost from the Android Emulator
    private const val BASE_URL = "http://10.0.2.2:3001/"

    private val json = Json {
        ignoreUnknownKeys = true
        coerceInputValues = true
    }

    private val retrofit = Retrofit.Builder()
        .baseUrl(BASE_URL)
        .addConverterFactory(json.asConverterFactory("application/json".toMediaType()))
        .build()

    fun <T> createService(serviceClass: Class<T>): T {
        return retrofit.create(serviceClass)
    }

    val api: SentinelApi by lazy { createService(SentinelApi::class.java) }

    inline fun <reified T> create(): T = createService(T::class.java)
}
