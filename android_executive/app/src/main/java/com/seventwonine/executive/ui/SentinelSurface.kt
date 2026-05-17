package com.seventwonine.executive.ui

import android.net.Uri
import android.view.Choreographer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.viewinterop.AndroidView
import com.google.ar.sceneform.SceneView
import com.google.ar.sceneform.math.Quaternion
import com.google.ar.sceneform.math.Vector3
import com.google.ar.sceneform.rendering.ModelRenderable
import com.google.ar.sceneform.Node

@Composable
fun SentinelSurface(modifier: Modifier = Modifier) {
    var modelRenderable by remember { mutableStateOf<ModelRenderable?>(null) }
    
    AndroidView(
        modifier = modifier.fillMaxSize(),
        factory = { context ->
            val sceneView = SceneView(context)
            // Optional: Transparent background so it blends well
            sceneView.backgroundColor = com.google.ar.sceneform.rendering.Color(0.027f, 0.035f, 0.082f) // Endheri Black Coral #070915
            
            // Load a placeholder 3D Model (.glb)
            // Replace with Uri.parse("android.resource://${context.packageName}/raw/your_model") for local assets
            ModelRenderable.builder()
                .setSource(context, Uri.parse("https://storage.googleapis.com/ar-answers-in-search-models/static/Tiger/model.glb"))
                .setIsFilamentGltf(true)
                .build()
                .thenAccept { renderable ->
                    modelRenderable = renderable
                    
                    val node = Node().apply {
                        setParent(sceneView.scene)
                        this.renderable = renderable
                        localPosition = Vector3(0f, -0.5f, -2.5f) // Positioned nicely behind the UI
                    }
                    
                    // Dynamic ambient rotation for the 3D Background
                    var angle = 0f
                    sceneView.scene.addOnUpdateListener {
                        angle += 0.5f
                        node.localRotation = Quaternion.axisAngle(Vector3(0f, 1f, 0f), angle)
                    }
                }
                .exceptionally { 
                    null 
                }
                
            sceneView
        },
        onRelease = { sceneView ->
            sceneView.destroy()
        }
    )
}
