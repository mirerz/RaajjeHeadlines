package com.seventwonine.executive.viewmodel

import androidx.compose.runtime.State
import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.seventwonine.commonui.network.SentinelNetworkClient
import com.seventwonine.commonui.models.RaahiResponse
import kotlinx.coroutines.launch

data class RaahiState(
    val name: String = "",
    val isCalculating: Boolean = false,
    val finalForce: Int? = null,
    val synthesis: String? = null,
    val error: String? = null
)

class RaahiViewModel : ViewModel() {
    private val _state = mutableStateOf(RaahiState())
    val state: State<RaahiState> = _state

    fun onNameChange(newName: String) {
        _state.value = _state.value.copy(name = newName)
    }

    fun calculateDestiny() {
        val currentName = _state.value.name
        if (currentName.isBlank()) return

        _state.value = _state.value.copy(isCalculating = true, error = null)

        viewModelScope.launch {
            try {
                val response = SentinelNetworkClient.api.calculateRaahi(mapOf("name" to currentName))
                if (response.isSuccessful && response.body() != null) {
                    val body = response.body()!!
                    _state.value = _state.value.copy(
                        isCalculating = false,
                        finalForce = body.final_force,
                        synthesis = body.synthesis
                    )
                } else {
                    _state.value = _state.value.copy(isCalculating = false, error = "Handshake Interrupted")
                }
            } catch (e: Exception) {
                _state.value = _state.value.copy(isCalculating = false, error = "Network Failure: ${e.message}")
            }
        }
    }
}

