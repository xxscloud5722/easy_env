class Session {
  saveToken(value: string) {
    sessionStorage.setItem('token', value);
  }

  getToken() {
    return sessionStorage.getItem('token') || undefined;
  }
}

export default new Session();
