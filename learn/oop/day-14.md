# Day 14: Composition over Inheritance

**Language:** Dart

## 1. Spot Misuse — Forced Inheritance

```dart
// AWKWARD: Penguin extends Bird but can't fly
class Bird {
  void fly() => print('Flying');
}
class Penguin extends Bird {
  @override
  void fly() => throw UnsupportedError('Penguins cannot fly');
}
```

**Problem:** Subclass breaks base contract — LSP violation.

## 2. Compose Instead

```dart
class FlyBehavior {
  void fly() => print('Flying');
}

class NoFly {
  void fly() => print('Cannot fly');
}

class Bird {
  final dynamic movement; // or a FlyBehavior interface
  Bird(this.movement);
  void move() => movement.fly();
}

final sparrow = Bird(FlyBehavior());
final penguin = Bird(NoFly());
```

## 3. Behavior Sharing via Composition

```dart
class Logger {
  void log(String msg) => print('[LOG] $msg');
}

class OrderService {
  final Logger _logger;
  OrderService(this._logger);

  void placeOrder(String item) {
    _logger.log('Order placed: $item');
  }
}

class PaymentService {
  final Logger _logger;
  PaymentService(this._logger);

  void charge(double amount) {
    _logger.log('Charged: \$$amount');
  }
}
// Share Logger without inheritance
```

## 4. Rule of Thumb

| Use inheritance when | Use composition when |
|---|---|
| True **is-a** relationship | **Has-a** relationship |
| Subclass honors base contract | You want to reuse behavior only |
| Shallow hierarchy | Inheritance feels forced |
| Shared identity/type | Mixing unrelated capabilities |

**Takeaway:** Prefer composition when reuse is the only reason for inheritance.
