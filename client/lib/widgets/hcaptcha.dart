import 'package:jomtx/utils/dart_ui.dart' as ui;
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:universal_html/html.dart' as html;
import 'package:webview_flutter/webview_flutter.dart';

class Captcha extends StatefulWidget {
  const Captcha({Key? key}) : super(key: key);
  @override
  State<StatefulWidget> createState() {
    return CaptchaState();
  }
}

class CaptchaState extends State<Captcha> {
  late WebViewController webViewController;
  @override
  void initState() {
    super.initState();
    if (kIsWeb) {
      ui.platformViewRegistry.registerViewFactory(
          'hcaptcha-html',
          (int viewId) => html.IFrameElement()
            ..src = "https://hcaptcha.wmtech.cc"
            // ..addEventListener('Captcha', (html.Event event) {
            //   print(event);
            //   var data = (event as html.MessageEvent).data;
            //   print(data);
            //   print(data['sender']);
            //   print(data['message']);
            //   // setState(() {
            //   //   //...
            //   // });
            // })
            ..style.border = 'none'
            ..width = '300'
            ..height = '80');
    } else {
      webViewController = WebViewController()
        ..setJavaScriptMode(JavaScriptMode.unrestricted)
        ..addJavaScriptChannel(
          'Captcha',
          onMessageReceived: (JavaScriptMessage message) async {
            // message contains the 'h-captcha-response' token.
            // Send it to your server in the login or other
            // data for verification via /siteverify
            // see: https://docs.hcaptcha.com/#server
            // print(message.message);
            Navigator.of(context).pop(message.message);
          },
        )
        ..setBackgroundColor(const Color(0x00000000))
        ..setNavigationDelegate(
          NavigationDelegate(
            onProgress: (int progress) {
              // Update loading bar.
            },
            onPageStarted: (String url) {},
            onPageFinished: (String url) {},
            onWebResourceError: (WebResourceError error) {},
          ),
        )
        ..loadRequest(Uri.parse("https://hcaptcha.wmtech.cc"));
    }
    html.window.onMessage.forEach((element) {
      if (!mounted) return;
      Navigator.of(context).pop(element.data);
    });
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      content: SizedBox(
        width: 400,
        child: Center(
          child: kIsWeb
              ? const HtmlElementView(viewType: 'hcaptcha-html')
              : WebViewWidget(controller: webViewController),
        ),
      ),
    );
  }
}
