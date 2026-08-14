# Day 7: Invariants and Validation

**Language:** Dart · **Domain:** Library loan

## 1. List Invariants — `Loan`

1. Due date must be in the future at creation
2. Member and book must not be null
3. A book cannot be on two active loans at once (enforced by `Book.borrow()`)

## 2. Preserve on Every Change

```dart
class Loan {
  final Member member;
  final Book book;
  DateTime _dueDate;

  Loan({
    required this.member,
    required this.book,
    required DateTime dueDate,
  }) : _dueDate = dueDate {
    if (dueDate.isBefore(DateTime.now())) {
      throw ArgumentError('Due date must be in the future');
    }
    if (!book.isAvailable) {
      throw StateError('Book is not available');
    }
    book.borrow();
  }

  DateTime get dueDate => _dueDate;

  void extend(int days) {
    if (days <= 0) throw ArgumentError('Days must be positive');
    _dueDate = _dueDate.add(Duration(days: days));
  }

  bool get isOverdue => DateTime.now().isAfter(_dueDate);
}
```

## 3. Fail Fast

```dart
void deposit(double amount) {
  if (amount <= 0) throw ArgumentError('Amount must be positive');
  _balance += amount;
}
```

## 4. Informal Tests (Edge Cases)

| Test | Expected |
|---|---|
| Create loan with past due date | `ArgumentError` |
| Extend by 0 days | `ArgumentError` |
| Borrow unavailable book | `StateError` |
| Deposit negative amount | `ArgumentError` |

**Takeaway:** Invariants are promises — every mutator must protect them.
