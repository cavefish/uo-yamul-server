class LoginInfo {
  String username;
  LoginState state;

  LoginInfo(this.username, this.state);
}

enum LoginState {
  authenticated,
  unauthenticated,
}