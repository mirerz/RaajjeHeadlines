import 'package:flutter/material.dart';

class ThemeTokens {
  // --- 🎨 Core Palette (Design DNA) ---
  static const Color midnightBlue = Color(0xFF000B1E);
  static const Color prestigeGold = Color(0xFFF4B400);
  static const Color alertRed = Color(0xFFFF4500);
  static const Color cyberCyan = Color(0xFF00FFFF);
  static const Color intensePurple = Color(0xFF4B0082);
  static const Color glassOverlay = Color(0x33FFFFFF);

  // --- 🖋️ Typography Matrix (Divehi Global) ---
  static const TextStyle headlineDv = TextStyle(
    fontFamily: 'DivehiGlobal',
    fontWeight: FontWeight.w900,
    fontSize: 28,
    height: 1.6,
    letterSpacing: -0.5,
    color: Colors.white,
  );

  static const TextStyle subheaderDv = TextStyle(
    fontFamily: 'DivehiGlobal',
    fontWeight: FontWeight.w700,
    fontSize: 18,
    height: 1.4,
    color: prestigeGold,
  );

  static const TextStyle tickerText = TextStyle(
    fontFamily: 'DivehiGlobal',
    fontWeight: FontWeight.bold,
    fontSize: 14,
    letterSpacing: 0.8,
    color: cyberCyan,
  );

  // --- 📐 Layout Tokens ---
  static const double cardRadius = 12.0;
  static const double prestigeDepth = 0.0012;
}
