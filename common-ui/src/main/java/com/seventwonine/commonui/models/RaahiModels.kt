package com.seventwonine.commonui.models

import kotlinx.serialization.Serializable

@Serializable
data class RaahiResponse(
    val total_value: Int,
    val final_force: Int,
    val synthesis: String,
    val status: String
)
