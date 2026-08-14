# Day 16: SOLID — SRP and OCP

**Language:** Dart

## 1. SRP Refactor — Split Two Reasons to Change

**Before (violates SRP):**

```dart
class User {
  String name;
  void saveToDatabase() { /* DB logic */ }
  void sendWelcomeEmail() { /* email logic */ }
}
```

**After:**

```dart
class User {
  final String name;
  User(this.name);
}

class UserRepository {
  void save(User user) { /* DB only */ }
}

class EmailService {
  void sendWelcome(User user) { /* email only */ }
}
```

## 2. OCP Extension — Strategy Instead of Switch

**Before (edit switch everywhere):**

```dart
double calculateShipping(String type, double weight) {
  switch (type) {
    case 'standard': return weight * 2;
    case 'express': return weight * 5;
    default: throw ArgumentError('Unknown');
  }
}
// Adding 'overnight' requires editing this function
```

**After (open for extension):**

```dart
abstract class ShippingStrategy {
  double calculate(double weight);
}

class StandardShipping implements ShippingStrategy {
  @override
  double calculate(double weight) => weight * 2;
}

class ExpressShipping implements ShippingStrategy {
  @override
  double calculate(double weight) => weight * 5;
}

// Add OvernightShipping — new class, no edit to existing
class OvernightShipping implements ShippingStrategy {
  @override
  double calculate(double weight) => weight * 10;
}
```

## 3. Before/After — Change Count

| Add new shipping type | Before | After |
|---|---|---|
| Files to edit | 1 switch + all callers | 1 new class only |
| Risk of breaking existing | High | Low |

## 4. Smell Check — SRP Violation

**Smell:** `ReportGenerator` generates PDF, sends email, and saves to DB.

**Fix:** Split into `ReportGenerator`, `ReportExporter`, `ReportDelivery`.

**Takeaway:** One reason to change per class; extend via new types, not edits.
