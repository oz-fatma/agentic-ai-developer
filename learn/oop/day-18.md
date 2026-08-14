# Day 18: Design Patterns — Factory & Strategy

**Language:** Dart

## 1. Factory — Centralize Creation

```dart
abstract class Notifier {
  void send(String message);
}

class EmailNotifier implements Notifier {
  void send(String msg) => print('Email: $msg');
}

class PushNotifier implements Notifier {
  void send(String msg) => print('Push: $msg');
}

class NotifierFactory {
  static Notifier create(String type) {
    switch (type) {
      case 'email': return EmailNotifier();
      case 'push': return PushNotifier();
      default: throw ArgumentError('Unknown notifier: $type');
    }
  }
}

final notifier = NotifierFactory.create('email');
```

## 2. Strategy — Interchangeable Algorithms

```dart
abstract class DiscountStrategy {
  double apply(double price);
}

class NoDiscount implements DiscountStrategy {
  double apply(double price) => price;
}

class PercentDiscount implements DiscountStrategy {
  final double percent;
  PercentDiscount(this.percent);
  double apply(double price) => price * (1 - percent / 100);
}

class Cart {
  final DiscountStrategy _discount;
  Cart(this._discount);
  double total(double subtotal) => _discount.apply(subtotal);
}
```

## 3. Replace Conditionals

```dart
// Instead of:
// if (type == 'student') price *= 0.8;
// else if (type == 'senior') price *= 0.7;

final cart = Cart(PercentDiscount(20)); // strategy injected
```

## 4. When NOT to Use a Pattern

**Skip Factory when:** Only one implementation exists — `Notifier()` is enough.

**Skip Strategy when:** One algorithm, no variation — plain method is clearer.

**Rule:** Pattern must reduce complexity, not add ceremony.

**Takeaway:** Factory for creation; Strategy for swappable algorithms — use with judgment.
