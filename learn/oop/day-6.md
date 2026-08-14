# Day 6: Encapsulation — Hiding Internal State

**Language:** Dart · **Domain:** Library

## 1. Make State Private

```dart
class BankAccount {
  double _balance = 0; // private

  void deposit(double amount) {
    if (amount <= 0) throw ArgumentError('Amount must be positive');
    _balance += amount;
  }

  double get balance => _balance; // read-only access
}
```

## 2. Break External Writes

```dart
// Before: account.balance = -999; // possible with public field

// After:
final account = BankAccount();
// account._balance = -999; // compile error — private
// account.balance = -999;  // compile error — no setter
account.deposit(100); // only valid path
```

## 3. Controlled Updates

```dart
class Product {
  final String name;
  double _price;

  Product(this.name, this._price);

  double get price => _price;

  void applyDiscount(double percent) {
    if (percent < 0 || percent > 100) {
      throw ArgumentError('Discount must be 0–100');
    }
    _price = _price * (1 - percent / 100);
  }
}
```

## 4. Read Access — Safe Getters

```dart
class Book {
  final String title;
  bool _available = true;

  Book(this.title);

  bool get isAvailable => _available; // read-only view

  void borrow() {
    if (!_available) throw StateError('Unavailable');
    _available = false;
  }
}
```

**Takeaway:** Hide fields; expose behavior. Callers cannot break invariants from outside.
