package com.seventwonine.koshaaru.ui

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.seventwonine.commonui.network.SentinelNetworkClient
import com.seventwonine.koshaaru.data.LoreHotspotDto
import com.seventwonine.koshaaru.data.LoreService
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

sealed class LoreUiState {
    object Loading : LoreUiState()
    data class Success(val hotspots: List<LoreHotspotDto>) : LoreUiState()
    data class Error(val message: String) : LoreUiState()
}

class LoreViewModel : ViewModel() {
    private val loreService = SentinelNetworkClient.create<LoreService>()

    private val _uiState = MutableStateFlow<LoreUiState>(LoreUiState.Loading)
    val uiState: StateFlow<LoreUiState> = _uiState.asStateFlow()

    init {
        fetchLore()
    }

    fun fetchLore() {
        viewModelScope.launch {
            _uiState.value = LoreUiState.Loading
            try {
                val response = loreService.getHotspots()
                if (response.status == "success") {
                    _uiState.value = LoreUiState.Success(response.data)
                } else {
                    _uiState.value = LoreUiState.Error("Server returned failure")
                }
            } catch (e: Exception) {
                _uiState.value = LoreUiState.Error(e.localizedMessage ?: "Unknown error occurred")
            }
        }
    }
}
