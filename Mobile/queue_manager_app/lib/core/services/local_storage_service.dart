import 'package:shared_preferences/shared_preferences.dart';

import '../../config/const/local_storage_constants.dart';

class LocalStorageService {
  static final LocalStorageService _instance = LocalStorageService._internal();
  factory LocalStorageService() => _instance;

  LocalStorageService._internal();

  SharedPreferences? _prefs;
  String? _cachedToken;

  /// Initialize SharedPreferences once (called in the app's startup logic)
  Future<void> init() async {
    _prefs = await SharedPreferences.getInstance();
  }

  /// Save a string value
  Future<void> saveString(String key, String value) async {
    final result =  await _prefs?.setString(key, value);

    print('Saved $key: $value');
    if (key == LocalStorageConstants.accessTokenKey) {
      _cachedToken = value;
    }
  }

  /// Retrieve a string value
  String? getString(String key) {
    if (key == LocalStorageConstants.accessTokenKey && _cachedToken != null) {
      return _cachedToken;
    }

    return _prefs?.getString(key);
  }

  /// Save the auth token
  Future<void> saveToken(String token) async {
    await saveString(LocalStorageConstants.accessTokenKey, token);
  }

  /// Get the auth token
  String? getToken() {
    return getString(LocalStorageConstants.accessTokenKey);
  }

  /// Remove a specific key
  Future<void> remove(String key) async {
    await _prefs?.remove(key);

    if (key == LocalStorageConstants.accessTokenKey) {
      _cachedToken = null;
    }
  }

  /// Clear all data
  Future<void> clear() async {
    await _prefs?.clear();

    _cachedToken = null;
  }
}
