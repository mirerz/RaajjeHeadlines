// lib/ui/raahi/benethe_card_widget.dart
import 'package:flutter/material.dart';
import 'dart:math';
import 'benethe_stack.dart';

class BenetheCardWidget extends StatefulWidget {
  final String cardSlug;
  final String imageUrl;

  const BenetheCardWidget({
    Key? key,
    required this.cardSlug,
    required this.imageUrl,
  }) : super(key: key);

  @override
  _BenetheCardWidgetState createState() => _BenetheCardWidgetState();
}

class _BenetheCardWidgetState extends State<BenetheCardWidget> with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  late Animation<double> _animation;
  bool _isFront = false;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      duration: const Duration(milliseconds: 800),
      vsync: this,
    );

    _animation = Tween<double>(begin: 0, end: 1).animate(
      CurvedAnimation(parent: _controller, curve: Curves.easeInOutBack),
    );
  }

  @override
  void dispose() {
    _controller.dispose();
    super.initState();
  }

  void _toggleFlip() {
    if (_isFront) {
      _controller.reverse();
    } else {
      _controller.forward();
    }
    setState(() {
      _isFront = !_isFront;
    });
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: _toggleFlip,
      child: AnimatedBuilder(
        animation: _animation,
        builder: (context, child) {
          final angle = _animation.value * pi;
          final isBack = angle <= pi / 2;

          return RepaintBoundary(
            child: Transform(
              transform: Matrix4.identity()
                ..setEntry(3, 2, 0.0012) // Prestige depth updated
                ..rotateY(angle),
              alignment: Alignment.center,
              child: isBack
                  ? _buildBackSide()
                  : Transform(
                      transform: Matrix4.identity()..rotateY(pi),
                      alignment: Alignment.center,
                      child: buildBenetheCard(widget.cardSlug, widget.imageUrl),
                    ),
            ),
          );
        },
      ),
    );
  }

  Widget _buildBackSide() {
    return Container(
      width: double.infinity,
      height: double.infinity,
      decoration: BoxDecoration(
        color: const Color(0xFF000B1E), // Midnight Blue
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: const Color(0xFFF4B400), width: 3), // Thick Gold Border
        image: const DecorationImage(
          image: NetworkImage('https://assets.729holdings.news/benethe/back_mandala.png'),
          fit: BoxFit.cover,
          opacity: 0.8,
        ),
      ),
      child: Stack(
        alignment: Alignment.center,
        children: [
          // Centered 729 Circular Emblem Layer
          Container(
            padding: const EdgeInsets.all(24),
            decoration: BoxDecoration(
              shape: BoxType.circle,
              border: Border.all(color: const Color(0xFFF4B400).withOpacity(0.3), width: 1),
            ),
            child: const Icon(Icons.brightness_high, color: Color(0xFFF4B400), size: 80),
          ),
          Positioned(
            bottom: 40,
            child: Text(
              "BENETHÉ",
              style: TextStyle(
                color: const Color(0xFFF4B400),
                letterSpacing: 12,
                fontSize: 10,
                fontWeight: FontWeight.w900,
                shadows: [
                  Shadow(blurRadius: 4, color: Colors.black),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
