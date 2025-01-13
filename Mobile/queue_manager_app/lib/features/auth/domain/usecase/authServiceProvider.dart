import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/auth/domain/entitites/authStatus.dart';

class AuthService extends StateNotifier<AuthState> {
  AuthService() : super(AuthState.unauthenticated());

  Future<void> login(String user) async {
    // Perform login logic here
    state = AuthState.authenticated();
  }

  Future<void> logout() async {
    // Perform logout logic here
    state = AuthState.unauthenticated();
  }
}

final authServiceProvider =
    StateNotifierProvider<AuthService, AuthState>((ref) {
  return AuthService();
});
