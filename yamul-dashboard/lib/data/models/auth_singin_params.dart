class AuthLoginParams {
  final String username;
  final String password;

  AuthLoginParams(this.username, this.password);

  @override
  String toString() {
    return 'AuthSignInParams{username: $username}';
  }
}
