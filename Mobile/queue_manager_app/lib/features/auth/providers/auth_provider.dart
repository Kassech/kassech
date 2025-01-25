import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/config/const/local_storage_constants.dart';
import 'package:queue_manager_app/core/services/local_storage_service.dart';

import '../models/user.dart';
import '../models/user_params.dart';
import '../repositories/auth_repository.dart';

final authProvider = AsyncNotifierProvider<AuthNotifier, User?>(() {
  final authRepository = AuthRepository();
  return AuthNotifier(authRepository);
});

class AuthNotifier extends AsyncNotifier<User?> {
  final AuthRepository _authRepository;

  AuthNotifier(this._authRepository);

  @override
  Future<User?> build() async {
    return await _checkAuth();
  }

  Future<void> signUp(UserParams user) async {
    state = const AsyncLoading();
    try {
      final response = await _authRepository.signUp(user);
      state = AsyncData(response);
    } on DioException catch (e) {
      state = AsyncError(e.response?.data['message'] ?? e.message, e.stackTrace);
    }  catch (e, stackTrace) {
      state = AsyncError(e.toString(), stackTrace);
    }
  }

  Future<void> login({
    String? phoneNumber,
    String? email,
    required String password,
  }) async {
    state = const AsyncLoading();
    try {
      final response = await _authRepository.login(
        phoneNumber: phoneNumber,
        email: email,
        password: password,
      );
      state = AsyncData(response);
    }  on DioException catch (e) {
      state = AsyncError(e.response ?? 'something went wrong' , e.stackTrace);
    } catch (e, stackTrace) {
      state = AsyncError(e.toString(), stackTrace);
    }
  }

  Future<void> logout() async {
    state = const AsyncLoading();
    try {
      await _authRepository.logout();
      state = const AsyncData(null);
    } on DioException catch (e) {
      state = AsyncError(e.response?.data['message'] ?? e.message, e.stackTrace);
    } catch (e, stackTrace) {
      state = AsyncError(e.toString(), stackTrace);
    }
  }

  Future<User?> _checkAuth() async {
    state = const AsyncLoading();
    try {
      final userJson =
          await LocalStorageService().getString(LocalStorageConstants.userKey);
      if (userJson != null) {
        final user = User.fromJson(jsonDecode(userJson));
        state = AsyncData(user);
        return user;
      }

      state = const AsyncData(null);
      return null;
    } catch (e, stackTrace) {
      state = AsyncError(e, stackTrace);
      return null;
    }
  }
}
