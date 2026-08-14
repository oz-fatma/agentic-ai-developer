# Day 10: Encapsulation & Abstraction — Practice

**Language:** Dart · **Domain:** Hardened library model

## 1. Harden Day 5 Model — All Fields Private

```dart
class Book {
  final String title;
  bool _available = true;

  Book(this.title);

  bool get isAvailable => _available;

  void borrow() {
    if (!_available) throw StateError('Book not available');
    _available = false;
  }

  void returnBook() => _available = true;
}
```

## 2. Invariant Suite on Every Mutator

```dart
class Member {
  final String id;
  final String _name;
  final List<Loan> _activeLoans = [];

  Member({required this.id, required String name}) : _name = name {
    if (id.isEmpty) throw ArgumentError('ID required');
    if (name.trim().length < 2) throw ArgumentError('Name too short');
  }

  String get name => _name;
  List<Loan> get activeLoans => List.unmodifiable(_activeLoans);

  void addLoan(Loan loan) {
    if (_activeLoans.length >= 5) throw StateError('Loan limit reached');
    _activeLoans.add(loan);
  }
}
```

## 3. Abstraction Pass — Domain Language

```dart
class Library {
  void checkout(Member member, Book book) { /* ... */ }
  void checkin(Loan loan) { /* ... */ }
  void sendOverdueReminders() { /* ... */ }
}
```

## 4. Before/After — Illegal Actions Now Impossible

1. Cannot set `book.isAvailable = false` directly — must call `borrow()`
2. Cannot give a member 100 loans — limit enforced in `addLoan()`
3. Cannot create a loan with empty member ID
4. Cannot mutate another member's loan list from outside
5. Cannot apply a 150% discount on price — validation in `applyDiscount()`

**Phase 2 complete (Days 6–10).**
