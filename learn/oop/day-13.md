# Day 13: Interfaces and Abstract Contracts

**Language:** Dart

## 1. Define a Contract

```dart
abstract class PaymentMethod {
  String get name;
  bool pay(double amount);
}
```

## 2. Implement Twice

```dart
class CreditCard implements PaymentMethod {
  @override
  String get name => 'Credit Card';

  @override
  bool pay(double amount) {
    print('Paid \$$amount via credit card');
    return true;
  }
}

class PayPal implements PaymentMethod {
  @override
  String get name => 'PayPal';

  @override
  bool pay(double amount) {
    print('Paid \$$amount via PayPal');
    return true;
  }
}
```

## 3. Depend on Contract — Not Concrete Class

```dart
class CheckoutService {
  bool processPayment(PaymentMethod method, double total) {
    if (total <= 0) return false;
    return method.pay(total);
  }
}
```

## 4. Swap Implementations

```dart
void main() {
  final checkout = CheckoutService();

  checkout.processPayment(CreditCard(), 49.99);
  checkout.processPayment(PayPal(), 19.99);
  // Same caller — different implementations
}
```

**Dart note:** Use `abstract class` or `interface` keyword (Dart 3) for pure contracts.

**Takeaway:** Program to `PaymentMethod`, not `CreditCard` — swap without changing callers.
