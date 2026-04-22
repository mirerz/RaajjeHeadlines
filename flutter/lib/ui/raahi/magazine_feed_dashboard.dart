import 'package:flutter/material.dart';
import 'package:flutter_staggered_grid_view/flutter_staggered_grid_view.dart';
import 'benethe_card_widget.dart';
import 'theme_tokens.dart';
import 'live_intelligence_ticker.dart';

class MagazineFeedDashboard extends StatelessWidget {
  final List<Map<String, String>> mockCards = [
    {'slug': 'fool', 'image': 'https://assets.729holdings.news/benethe/fool_bg.png'},
    {'slug': 'magician', 'image': 'https://assets.729holdings.news/benethe/magician_bg.png'},
    {'slug': 'star', 'image': 'https://assets.729holdings.news/benethe/star_bg.png'},
    {'slug': 'empress', 'image': 'https://assets.729holdings.news/benethe/empress_bg.png'},
    {'slug': 'high_priestess', 'image': 'https://assets.729holdings.news/benethe/high_priestess_bg.png'},
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: ThemeTokens.midnightBlue,
      body: SafeArea(
        child: Column(
          children: [
            // --- 📡 Sentinel Pulse Header ---
            Container(
              padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 24),
              decoration: BoxDecoration(
                border: Border(bottom: BorderSide(color: ThemeTokens.intensePurple.withOpacity(0.5))),
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text("adhu.space", style: TextStyle(color: ThemeTokens.cyberCyan, fontWeight: FontWeight.w900, letterSpacing: 2)),
                  const Icon(Icons.radar, color: ThemeTokens.alertRed, size: 18),
                ],
              ),
            ),
            const LiveIntelligenceTicker(),
            Expanded(
              child: MasonryGridView.count(
                padding: const EdgeInsets.all(16),
                crossAxisCount: 2,
                mainAxisSpacing: 20,
                crossAxisSpacing: 20,
                itemCount: mockCards.length,
                itemBuilder: (context, index) {
                  // Staggered layout ratios from Stitch DNA
                  final double height = (index % 3 == 0) ? 320 : (index % 2 == 0 ? 240 : 280);
                  return Container(
                    height: height,
                    decoration: BoxDecoration(
                      boxShadow: [
                        BoxShadow(
                          color: ThemeTokens.cyberCyan.withOpacity(0.05),
                          blurRadius: 20,
                          spreadRadius: 2,
                        )
                      ],
                    ),
                    child: BenetheCardWidget(
                      cardSlug: mockCards[index]['slug']!,
                      imageUrl: mockCards[index]['image']!,
                    ),
                  );
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}
