import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../domain/models/user.dart';
import '../../domain/models/user_params.dart';
import '../../repositories/auth_repository.dart';

final authProvider = AsyncNotifierProvider<AuthNotifier, User?>(() {
  final authRepository = AuthRepository();
  return AuthNotifier(authRepository);
});

class AuthNotifier extends AsyncNotifier<User?> {
  final AuthRepository _authRepository;

  AuthNotifier(this._authRepository);

  @override
  Future<User?> build() async {
    return null;
  }

  Future<void> signUp(UserParams user) async {
    state = const AsyncLoading();
    try {
      final response = await _authRepository.signUp(user);
      state = AsyncData(response?['user']);
    } catch (e, stackTrace) {
      state = AsyncError(e.toString(), stackTrace);
    }
  }
}

