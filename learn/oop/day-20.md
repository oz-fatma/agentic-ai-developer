# Day 20: Capstone Model & Review

**Language:** Dart · **Domain:** Online shop orders

---

## 1. Capstone Domain — Cohesive Model

```dart
// --- Value object (immutable) ---
class Money {
  final double amount;
  final String currency;
  const Money(this.amount, this.currency);

  Money add(Money other) {
    if (currency != other.currency) throw ArgumentError('Currency mismatch');
    return Money(amount + other.amount, currency);
  }
}

// --- Entity ---
class Product {
  final String id;
  final String name;
  final Money price;
  Product({required this.id, required this.name, required this.price});
}

class CartItem {
  final Product product;
  int quantity;
  CartItem(this.product, this.quantity);
  Money get subtotal => Money(
    product.price.amount * quantity,
    product.price.currency,
  );
}

// --- Contract (DIP) ---
abstract class PaymentMethod {
  bool pay(Money amount);
}

class CreditCardPayment implements PaymentMethod {
  @override
  bool pay(Money amount) {
    print('Charged ${amount.amount} ${amount.currency}');
    return true;
  }
}

// --- Strategy (OCP) ---
abstract class DiscountStrategy {
  Money apply(Money subtotal);
}

class NoDiscount implements DiscountStrategy {
  Money apply(Money subtotal) => subtotal;
}

// --- Service (SRP) ---
class OrderService {
  final PaymentMethod _payment;
  final DiscountStrategy _discount;

  OrderService(this._payment, this._discount);

  bool checkout(List<CartItem> items) {
    var total = const Money(0, 'USD');
    for (final item in items) {
      total = total.add(item.subtotal);
    }
    total = _discount.apply(total);
    return _payment.pay(total);
  }
}
```

## 2. Apply Principles

| Principle | Where |
|---|---|
| Encapsulation | `Money` immutable; cart totals via methods |
| Abstraction | `PaymentMethod`, `DiscountStrategy` contracts |
| SRP | `OrderService` checkout only; payment separate |
| OCP | New discount = new class, no edit to `OrderService` |
| DIP | `OrderService` depends on abstractions |

## 3. Optional Pattern — Strategy (used)

`DiscountStrategy` — clear benefit for multiple discount types.

Factory skipped — only one payment path in capstone; plain constructor is enough.

## 4. Reflection — 20 Days

**Learned:**
- Days 1–5: Think in objects; classes, constructors, responsibilities
- Days 6–10: Encapsulate state; enforce invariants; immutable vs mutable
- Days 11–15: Inheritance, polymorphism, interfaces, composition
- Days 16–20: SOLID, Factory/Strategy/Observer, refactoring

**Practice next:**
- Build a small Flutter app using these Dart classes
- Write unit tests for invariants (`Loan`, `Money`, `OrderService`)
- Refactor one real project god-class into smaller objects

---

## OOP Track Complete — 20/20

| Phase | Days | Focus |
|---|---|---|
| OOP Fundamentals | 1–5 | Classes, objects, constructors |
| Encapsulation & Abstraction | 6–10 | Private state, invariants |
| Inheritance & Polymorphism | 11–15 | extends, implements, composition |
| SOLID, Patterns & Practice | 16–20 | Professional design |

**Language:** Dart · **Ready for Flutter integration**
