import 'package:flutter/material.dart';
import 'package:google_nav_bar/google_nav_bar.dart';

class CameraView extends StatefulWidget {
  const CameraView({super.key, required this.camera});

  final String camera;

  @override
  State<CameraView> createState() => _CameraViewState();
}

class _CameraViewState extends State<CameraView> {
  @override
  Widget build(BuildContext context) {

    return Scaffold(
      bottomNavigationBar: GNav(
        
        onTabChange:  (index){
          print(index);
        },
        tabs: const [
        GButton(
          icon: Icons.home,
          text: ' Friends',
        
        ),
        GButton(
          icon: Icons.camera,
          text: ' Yourself',
          ),
        GButton(
          icon: Icons.search,
          text: ' Adventure',
          ),
        ],
        ),
    );


    return const Placeholder();
  }
}