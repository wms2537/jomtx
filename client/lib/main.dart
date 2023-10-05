import 'package:jomtx/router/route_information_parser.dart';
import 'package:jomtx/router/router_delegate.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

void main() async {
  // usePathUrlStrategy();
  // await initHiveForFlutter();
  runApp(const MyApp());
}

class MyCustomScrollBehavior extends MaterialScrollBehavior {
  // Override behavior methods and getters like dragDevices
  @override
  Set<PointerDeviceKind> get dragDevices => {
        PointerDeviceKind.touch,
        PointerDeviceKind.mouse,
      };
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'Caj-Up',
      theme: ThemeData(
        // primarySwatch: Colors.amber,
        colorScheme: ColorScheme.fromSwatch(
          primarySwatch: Colors.amber,
          backgroundColor: const Color(0xffe0e0e0),
          brightness: Brightness.light,
        ),
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      scrollBehavior: MyCustomScrollBehavior(),
      routeInformationParser: AppRouteInformationParser(),
      routerDelegate: AppRouterDelegate(),
    );
  }
}
