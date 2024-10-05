import 'package:flutter/material.dart';
import 'package:google_nav_bar/google_nav_bar.dart';
import '../pages/HomePage.dart'; // Import your pages here
import '../pages/AdventurPage.dart';
import '../pages/YourselfPage.dart';
class CameraView extends StatefulWidget {
  const CameraView({super.key, required this.camera});

  final String camera;

  @override
  State<CameraView> createState() => _CameraViewState();
}

class _CameraViewState extends State<CameraView> {
  int _selectedIndex = 0;

  // List of pages to switch between
  final List<Widget> _pages = [
    HomePage(),
    YourselfPage(),
    AdventurePage(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _pages[_selectedIndex], // Display the selected page
      
      bottomNavigationBar: GNav(
        selectedIndex: _selectedIndex,
        onTabChange: (index) {
          setState(() {
            _selectedIndex = index; // Update the selected tab index
          });
        },
        tabs: const [
          GButton(
            icon: Icons.home,
            text: 'Friends',
          ),
          GButton(
            icon: Icons.camera,
            text: 'Yourself',
          ),
          GButton(
            icon: Icons.search,
            text: 'Adventure',
          ),
        ],
      ),
    );
  }
}