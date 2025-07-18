// File generated by FlutterFire CLI.
// ignore_for_file: type=lint
import 'package:firebase_core/firebase_core.dart' show FirebaseOptions;
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, kIsWeb, TargetPlatform;

/// Default [FirebaseOptions] for use with your Firebase apps.
///
/// Example:
/// ```dart
/// import 'firebase_options.dart';
/// // ...
/// await Firebase.initializeApp(
///   options: DefaultFirebaseOptions.currentPlatform,
/// );
/// ```
class DefaultFirebaseOptions {
  static FirebaseOptions get currentPlatform {
    if (kIsWeb) {
      return web;
    }
    switch (defaultTargetPlatform) {
      case TargetPlatform.android:
        return android;
      case TargetPlatform.iOS:
        return ios;
      case TargetPlatform.macOS:
        return macos;
      case TargetPlatform.windows:
        return windows;
      case TargetPlatform.linux:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for linux - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      default:
        throw UnsupportedError(
          'DefaultFirebaseOptions are not supported for this platform.',
        );
    }
  }

  static const FirebaseOptions web = FirebaseOptions(
    apiKey: 'AIzaSyCf6Lh7G2sFwqpSZUBKoaXdNFXTt-SoytY',
    appId: '1:729434706009:web:d26d0de793726f8c57e148',
    messagingSenderId: '729434706009',
    projectId: 'kassechtransportaion',
    authDomain: 'kassechtransportaion.firebaseapp.com',
    storageBucket: 'kassechtransportaion.firebasestorage.app',
    measurementId: 'G-8HQ6864J49',
  );

  static const FirebaseOptions android = FirebaseOptions(
    apiKey: 'AIzaSyDR_ajkCa2c_ocesiyXNuGPXQUPFKFD_TQ',
    appId: '1:729434706009:android:7905c333aa47a30357e148',
    messagingSenderId: '729434706009',
    projectId: 'kassechtransportaion',
    storageBucket: 'kassechtransportaion.firebasestorage.app',
  );

  static const FirebaseOptions ios = FirebaseOptions(
    apiKey: 'AIzaSyBXrzHIcpwlJLr_MVct2E25_4n-F5EcrTw',
    appId: '1:729434706009:ios:dac147255008305e57e148',
    messagingSenderId: '729434706009',
    projectId: 'kassechtransportaion',
    storageBucket: 'kassechtransportaion.firebasestorage.app',
    iosBundleId: 'com.example.queueManagerApp',
  );

  static const FirebaseOptions macos = FirebaseOptions(
    apiKey: 'AIzaSyBXrzHIcpwlJLr_MVct2E25_4n-F5EcrTw',
    appId: '1:729434706009:ios:dac147255008305e57e148',
    messagingSenderId: '729434706009',
    projectId: 'kassechtransportaion',
    storageBucket: 'kassechtransportaion.firebasestorage.app',
    iosBundleId: 'com.example.queueManagerApp',
  );

  static const FirebaseOptions windows = FirebaseOptions(
    apiKey: 'AIzaSyCf6Lh7G2sFwqpSZUBKoaXdNFXTt-SoytY',
    appId: '1:729434706009:web:46e1ddb666cca76957e148',
    messagingSenderId: '729434706009',
    projectId: 'kassechtransportaion',
    authDomain: 'kassechtransportaion.firebaseapp.com',
    storageBucket: 'kassechtransportaion.firebasestorage.app',
    measurementId: 'G-S7HYW89BZY',
  );
}
