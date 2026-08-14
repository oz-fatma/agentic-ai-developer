# Day 17: SOLID — LSP, ISP, and DIP

**Language:** Dart

## 1. LSP — Subtypes Must Substitute Bases

**Violation:**

```dart
class Rectangle {
  double width, height;
  void setWidth(double w) => width = w;
  void setHeight(double h) => height = h;
}
class Square extends Rectangle {
  @override
  void setWidth(double w) { width = height = w; } // surprises callers
}
```

**Fix:** Don't inherit Square from Rectangle — use separate types or composition.

## 2. ISP — Split Fat Interface

**Before:**

```dart
abstract class Worker {
  void work();
  void eat();
  void sleep();
}
class Robot implements Worker {
  void work() {}
  void eat() => throw UnsupportedError('Robots do not eat');
  void sleep() => throw UnsupportedError('Robots do not sleep');
}
```

**After:**

```dart
abstract class Workable { void work(); }
abstract class Eatable { void eat(); }

class Human implements Workable, Eatable {
  void work() {}
  void eat() {}
}

class Robot implements Workable {
  void work() {}
}
```

## 3. DIP — Depend on Abstraction

```dart
abstract class PaymentGateway {
  bool charge(double amount);
}

class StripeGateway implements PaymentGateway {
  @override
  bool charge(double amount) { /* Stripe API */ return true; }
}

class OrderService {
  final PaymentGateway _gateway; // abstraction, not Stripe directly
  OrderService(this._gateway);

  bool checkout(double total) => _gateway.charge(total);
}

// Inject concrete type at app startup
final service = OrderService(StripeGateway());
```

## 4. Violation Hunt + Fix

| Principle | Violation | Fix |
|---|---|---|
| LSP | Square extends Rectangle | Separate classes |
| ISP | Robot forced to `eat()` | Split interfaces |
| DIP | Service imports Stripe directly | Inject `PaymentGateway` |

**Takeaway:** SOLID together — honest subtypes, small interfaces, inject abstractions.
