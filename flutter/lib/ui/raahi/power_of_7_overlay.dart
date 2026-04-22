import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

class PowerOf7Overlay extends StatelessWidget {
  final String backgroundUrl;
  final String cardTitleEn;
  final String cardTitleDv;
  final String transliteration;

  const PowerOf7Overlay({
    Key? key,
    required this.backgroundUrl,
    required this.cardTitleEn,
    required this.cardTitleDv,
    required this.transliteration,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: Alignment.topCenter,
      children: [
        // 1. Base Imagery (Design DNA)
        Image.network(
          backgroundUrl,
          fit: BoxFit.cover,
          width: double.infinity,
          height: double.infinity,
        ),

        // 2. The 'Safe Zone' Header (Programmatic SVG Overlay)
        Positioned(
          top: 40,
          child: Column(
            children: [
              Text(
                cardTitleEn,
                style: const TextStyle(
                  color: Color(0xFFF4B400),
                  fontWeight: FontWeight.bold,
                  fontSize: 16,
                  letterSpacing: 2,
                ),
              ),
              const SizedBox(height: 4),
              Text(
                "($transliteration)",
                style: TextStyle(
                  color: Colors.white.withOpacity(0.7),
                  fontStyle: FontStyle.italic,
                  fontSize: 12,
                ),
              ),
              const SizedBox(height: 12),
              // Accurate Noto Thaana Vector (SVG-like Rendering)
              RepaintBoundary(
                child: Text(
                  cardTitleDv,
                  style: const TextStyle(
                    color: Color(0xFFF4B400),
                    fontFamily: 'DivehiGlobal',
                    fontWeight: FontWeight.w900,
                    fontSize: 42,
                    shadows: [
                      Shadow(blurRadius: 15, color: Colors.black),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }
}
