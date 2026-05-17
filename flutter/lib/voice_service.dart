import 'dart:typed_data';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:audioplayers/audioplayers.dart';

class VoiceService {
  final AudioPlayer _audioPlayer = AudioPlayer();
  // Fetch your Cloud Run URL from your environment variables
  final String backendUrl = const String.fromEnvironment('BACKEND_URL', defaultValue: 'http://localhost:3000/v1');

  Future<void> synthesizeAndPlay(String text) async {
    try {
      final response = await http.post(
        Uri.parse('$backendUrl/synthesize'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'text': text}),
      );

      if (response.statusCode == 200) {
        // Play the raw byte stream directly
        Uint8List audioBytes = response.bodyBytes;
        await _audioPlayer.play(BytesSource(audioBytes));
      } else {
        print('Backend Error: ${response.statusCode}');
      }
    } catch (e) {
      print('Failed to stream audio: $e');
    }
  }
}
