enum AuthStatus { authenticated, unauthenticated }

class AuthState {
  final AuthStatus status;
  final String? user;

  AuthState._({required this.status});

  factory AuthState.authenticated() {
    return AuthState._(status: AuthStatus.authenticated);
  }

  factory AuthState.unauthenticated() {
    return AuthState._(status: AuthStatus.unauthenticated);
  }
}
