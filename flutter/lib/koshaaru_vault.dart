import 'package:flutter/material.dart';

class KoshaaruVaultPage extends StatelessWidget {
  const KoshaaruVaultPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF070915), // Endheri Black Coral
      appBar: AppBar(
        title: const Text("ENDHERI ODI VAULT", style: TextStyle(letterSpacing: 3, fontSize: 16)),
        centerTitle: true,
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Container(
              width: 200,
              height: 200,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                border: Border.all(color: const Color(0xFF00FFFF).withOpacity(0.5), width: 2),
                boxShadow: [
                  BoxShadow(
                    color: const Color(0xFF00FFFF).withOpacity(0.2), // Bioluminescent Cyan
                    blurRadius: 50,
                    spreadRadius: 10,
                  )
                ],
              ),
              child: const Icon(Icons.directions_boat, size: 80, color: Color(0xFF00FFFF)),
            ),
            const SizedBox(height: 40),
            const Text(
              "COMPACT VAULT VIEW",
              style: TextStyle(
                color: Color(0xFF00FFFF),
                fontSize: 10,
                fontWeight: FontWeight.bold,
                letterSpacing: 3,
              ),
            ),
            const SizedBox(height: 8),
            const Text(
              "DHOVEMI SAGA",
              style: TextStyle(
                color: Colors.white,
                fontSize: 24,
                fontWeight: FontWeight.w900,
                letterSpacing: 5,
              ),
            ),
            const SizedBox(height: 16),
            Text(
              "Multimodal Lore Matrix Offline.\nHyperOS Optimized. Connect to Data Connect to sync hotspots.",
              textAlign: TextAlign.center,
              style: TextStyle(color: Colors.grey[400], fontSize: 13, height: 1.5, letterSpacing: 1),
            ),
            const SizedBox(height: 40),
            OutlinedButton.icon(
              onPressed: () {},
              icon: const Icon(Icons.download, color: Color(0xFF00FFFF)),
              label: const Text("SYNC HERITAGE", style: TextStyle(color: Color(0xFF00FFFF), fontWeight: FontWeight.bold)),
              style: OutlinedButton.styleFrom(
                side: const BorderSide(color: Color(0xFF00FFFF)),
                padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 12),
                shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(30)),
              ),
            )
          ],
        ),
      ),
    );
  }
}
