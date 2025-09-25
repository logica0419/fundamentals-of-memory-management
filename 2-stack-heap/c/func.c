int b(int arg) {
  int b = 1;
  return arg + b;
}

int a(int arg) {
  int a = arg + 2;
  return b(a);
}

int main() {
  int v = 5;
  v = a(v);
}
