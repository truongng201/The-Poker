export function EmailValidation(email) {
  const re = /\S+@\S+\.\S+/;
  return re.test(email);
}

export function PasswordValidation(password) {
  const re = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,32}$/;
  return re.test(password);
}

export function UsernameValidation(username) {
  const re = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,32}$/;
  return re.test(username);
}
