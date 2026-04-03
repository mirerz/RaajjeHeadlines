import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

void main() {
  runApp(const RaajjeHeadlinesApp());
}

class RaajjeHeadlinesApp extends StatelessWidget {
  const RaajjeHeadlinesApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Raajjé HEADLINES',
      themeMode: ThemeMode.system, // Device Theme Collaboration
      theme: ThemeData(
        brightness: Brightness.light,
        scaffoldBackgroundColor: const Color(0xFFF8FAFC),
        primaryColor: const Color(0xFFFF1493),
        cardColor: Colors.white,
        appBarTheme: const AppBarTheme(
          backgroundColor: Colors.transparent,
          titleTextStyle: TextStyle(color: Color(0xFFFF1493), fontSize: 20, fontWeight: FontWeight.w900, fontStyle: FontStyle.italic),
        ),
      ),
      darkTheme: ThemeData(
        brightness: Brightness.dark,
        primaryColor: const Color(0xFFFF1493),
        scaffoldBackgroundColor: const Color(0xFF0B0E23),
        cardColor: const Color(0xFF1A1A2E),
      ),
      home: const NewsFeedPage(),
    );
  }
}

class NewsFeedPage extends StatefulWidget {
  const NewsFeedPage({super.key});

  @override
  State<NewsFeedPage> createState() => _NewsFeedPageState();
}

class _NewsFeedPageState extends State<NewsFeedPage> {
  List<Article> articles = [];
  bool isLoading = true;
  TimeOfDay pickDropTime = const TimeOfDay(hour: 13, minute: 30);

  @override
  void initState() {
    super.initState();
    fetchNews();
  }

  Future<void> _selectPickDropTime() async {
    final TimeOfDay? picked = await showTimePicker(
      context: context,
      initialTime: pickDropTime,
      builder: (context, child) => Theme(data: ThemeData.dark().copyWith(
        primaryColor: const Color(0xFFFF1493),
        colorScheme: const ColorScheme.dark(primary: Color(0xFFFF1493), secondary: Color(0xFF00FFFF)),
      ), child: child!),
    );
    if (picked != null && picked != pickDropTime) {
      setState(() => pickDropTime = picked);
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          backgroundColor: const Color(0xFFFF1493),
          content: Text("Reminder set for ${picked.format(context)}", style: const TextStyle(fontWeight: FontWeight.bold)),
        ),
      );
    }
  }

  Future<void> fetchNews() async {
    try {
      final response = await http.get(Uri.parse('http://localhost:3000/api/v1/feed'));
      if (response.statusCode == 200) {
        final List<dynamic> data = json.decode(response.body);
        setState(() {
          articles = data.map((json) => Article.fromJson(json)).toList();
          isLoading = false;
        });
      }
    } catch (e) {
      debugPrint("Fetch error: $e");
      setState(() => isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Image.file(
          File('C:/Users/Surface/.gemini/antigravity/brain/tempmediaStorage/media__1775183124968.png'),
          height: 32,
          fit: BoxFit.contain,
        ),
        centerTitle: true,
        backgroundColor: Colors.transparent,
        elevation: 0,
        actions: [
          IconButton(icon: const Icon(Icons.search, color: Color(0xFF00FFFF)), onPressed: () {}),
          IconButton(icon: const Icon(Icons.person_outline, color: Color(0xFF00FFFF)), onPressed: () {}),
        ],
      ),
      body: Column(
        children: [
          _buildUtilityWidgets(),
          Expanded(
            child: isLoading 
              ? const Center(child: CircularProgressIndicator(color: Color(0xFFFF1493))) 
              : ListView.builder(
                  padding: const EdgeInsets.all(16),
                  itemCount: articles.length,
                  itemBuilder: (context, index) => _buildArticleCard(articles[index]),
                ),
          ),
        ],
      ),
    );
  }

  Widget _buildUtilityWidgets() {
    return Container(
      height: 120,
      padding: const EdgeInsets.symmetric(vertical: 10),
      child: ListView(
        scrollDirection: Axis.horizontal,
        padding: const EdgeInsets.symmetric(horizontal: 16),
        children: [
          _buildGlassWidget(
            icon: Icons.calendar_month,
            title: "14 Ramadan 1447",
            subtitle: "Lunar Cycle",
            color: const Color(0xFF00CED1),
            onTap: () {},
          ),
          _buildGlassWidget(
            icon: Icons.mosque,
            title: "Next: Maghrib",
            subtitle: "18:12 (Malé)",
            color: const Color(0xFF00FFFF),
            onTap: () {},
          ),
          _buildGlassWidget(
            icon: Icons.notifications_active,
            title: "Pick/Drop",
            subtitle: "Time: ${pickDropTime.format(context)}",
            color: const Color(0xFFFF4500),
            onTap: _selectPickDropTime,
          ),
        ],
      ),
    );
  }

  Widget _buildGlassWidget({required IconData icon, required String title, required String subtitle, required Color color, required VoidCallback onTap}) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        width: 160,
        margin: const EdgeInsets.only(right: 12),
        padding: const EdgeInsets.all(12),
        decoration: BoxDecoration(
          color: const Color(0xFF1E1E3F).withOpacity(0.8),
          borderRadius: BorderRadius.circular(20),
          border: Border.all(color: color.withOpacity(0.5), width: 1.5),
          boxShadow: [BoxShadow(color: color.withOpacity(0.2), blurRadius: 10)],
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Icon(icon, color: color, size: 24),
            const Spacer(),
            Text(title, style: const TextStyle(fontWeight: FontWeight.w900, fontSize: 13, color: Colors.white)),
            Text(subtitle, style: TextStyle(fontSize: 10, color: Colors.grey[400], letterSpacing: 0.5)),
          ],
        ),
      ),
    );
  }

  Widget _buildArticleCard(Article article) {
    return Card(
      margin: const EdgeInsets.only(bottom: 24),
      color: const Color(0xFF1A1A2E),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(20),
        side: BorderSide(color: Colors.white.withOpacity(0.05)),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          if (article.visualURL.isNotEmpty) _buildVisualHeader(article),
          Padding(
            padding: const EdgeInsets.all(20.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: [
                    if (article.isBreaking) _buildBadge("BREAKING", const Color(0xFFFF4500)),
                    if (article.isPremium) _buildBadge("MEMBER ONLY", const Color(0xFFFFFF00)),
                  ],
                ),
                const SizedBox(height: 12),
                Text(
                  article.rephrasedHeadlineDv,
                  style: const TextStyle(fontSize: 22, fontWeight: FontWeight.w900, color: Colors.white, height: 1.3),
                  textAlign: TextAlign.right,
                ),
                if (article.isPremium && article.subscriberBriefing.isNotEmpty) ...[
                  const SizedBox(height: 16),
                  _buildSectionHeader("NEWS BRIEFING", const Color(0xFF00FFFF)),
                  Text(article.subscriberBriefing, style: const TextStyle(color: Colors.white70, fontSize: 13, height: 1.6), textAlign: TextAlign.right),
                ],
                const SizedBox(height: 16),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(article.category.toUpperCase(), style: const TextStyle(color: Color(0xFF00FFFF), fontSize: 10, fontWeight: FontWeight.bold, letterSpacing: 1.5)),
                    const Icon(Icons.arrow_forward_ios, color: Color(0xFFFF1493), size: 14),
                  ],
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildVisualHeader(Article article) {
    final bool isGoogle = article.visualSource == 'google';
    final Color accent = isGoogle ? Colors.red : const Color(0xFFB23AEE);
    return Container(
      height: 180,
      width: double.infinity,
      decoration: BoxDecoration(
        color: Colors.black,
        borderRadius: const BorderRadius.vertical(top: Radius.circular(20)),
        image: DecorationImage(
          image: NetworkImage(isGoogle ? "https://img.youtube.com/vi_webp/placeholder/maxresdefault.webp" : "https://placeholder.meta.com/social_visual.png"),
          fit: BoxFit.cover,
          opacity: 0.4,
        ),
      ),
      child: Center(
        child: Icon(isGoogle ? Icons.play_circle_fill : Icons.camera_alt, color: accent, size: 50),
      ),
    );
  }

  Widget _buildSectionHeader(String title, Color color) {
    return Container(
      margin: const EdgeInsets.only(bottom: 8),
      padding: const EdgeInsets.only(bottom: 4),
      decoration: BoxDecoration(border: Border(bottom: BorderSide(color: color.withOpacity(0.3), width: 1))),
      child: Text(title, style: TextStyle(color: color, fontSize: 10, fontWeight: FontWeight.black, letterSpacing: 1.2)),
    );
  }

  Widget _buildBadge(String text, Color color) {
    return Container(
      margin: const EdgeInsets.only(left: 6),
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 4),
      decoration: BoxDecoration(
        color: color,
        borderRadius: BorderRadius.circular(4),
        boxShadow: [BoxShadow(color: color.withOpacity(0.4), blurRadius: 8)],
      ),
      child: Text(text, style: const TextStyle(color: Colors.white, fontSize: 9, fontWeight: FontWeight.black)),
    );
  }
}

class Article {
  final String id;
  final String rephrasedHeadlineDv;
  final String category;
  final String visualSource;
  final String visualURL;
  final String subscriberBriefing;
  final String subscriberExplainer;
  final bool isBreaking;
  final bool isPremium;
  final bool isAd;

  Article({
    required this.id,
    required this.rephrasedHeadlineDv,
    required this.category,
    required this.visualSource,
    required this.visualURL,
    required this.subscriberBriefing,
    required this.subscriberExplainer,
    required this.isBreaking,
    required this.isPremium,
    required this.isAd,
  });

  factory Article.fromJson(Map<String, dynamic> json) {
    return Article(
      id: json['id']?.toString() ?? "",
      rephrasedHeadlineDv: json['rephrased_headline_dv'] ?? "",
      category: json['category'] ?? "General",
      visualSource: json['visual_source'] ?? "internal",
      visualURL: json['visual_url'] ?? "",
      subscriberBriefing: json['subscriber_briefing'] ?? "",
      subscriberExplainer: json['subscriber_explainer'] ?? "",
      isBreaking: json['is_breaking'] ?? false,
      isPremium: json['is_premium'] ?? false,
      isAd: json['is_ad'] ?? false,
    );
  }
}
