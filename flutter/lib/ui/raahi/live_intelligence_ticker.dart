import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'theme_tokens.dart';

class LiveIntelligenceTicker extends StatefulWidget {
  final String wsUrl;
  const LiveIntelligenceTicker({Key? key, this.wsUrl = 'wss://api.adhu.space/ws/ticker'}) : super(key: key);

  @override
  _LiveIntelligenceTickerState createState() => _LiveIntelligenceTickerState();
}

class _LiveIntelligenceTickerState extends State<LiveIntelligenceTicker> {
  late WebSocketChannel _channel;

  @override
  void initState() {
    super.initState();
    _channel = WebSocketChannel.connect(Uri.parse(widget.wsUrl));
  }

  @override
  void dispose() {
    _channel.sink.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 40,
      width: double.infinity,
      color: Colors.black,
      child: StreamBuilder(
        stream: _channel.stream,
        builder: (context, snapshot) {
          final String message = snapshot.hasData 
              ? snapshot.data.toString() 
              : "SIGNAL CONNECTING...";

          return RepaintBoundary(
            child: ListView(
              scrollDirection: Axis.horizontal,
              children: [
                Center(
                  child: Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 20),
                    child: Text(
                      "🚨 $message",
                      style: ThemeTokens.tickerText,
                    ),
                  ),
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
