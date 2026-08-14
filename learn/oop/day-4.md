# Day 4: Behavior and Responsibilities

**Language:** Dart · **Domain:** Library → Order-style collaboration

## 1. Single Responsibility — Split God Object

**Before (god object):**

```dart
// BAD: one class does everything
class LibrarySystem {
  String memberName;
  List<String> books = [];
  void addBook() {}
  void sendEmailReminder() {}
  void calculateFine() {}
}
```

**After (split):**

```dart
class Member {
  final String name;
  Member(this.name);
}

class Loan {
  final Member member;
  final Book book;
  DateTime dueDate;
  Loan(this.member, this.book, this.dueDate);
}

class FineCalculator {
  double calculate(Loan loan) {
    final daysLate = DateTime.now().difference(loan.dueDate).inDays;
    return daysLate > 0 ? daysLate * 2.0 : 0;
  }
}
```

## 2. Tell, Don't Ask

**Ask (bad):**

```dart
if (!book.isAvailable) throw StateError('Unavailable');
book.isAvailable = false;
```

**Tell (good):**

```dart
book.borrow(); // object owns the rule
```

## 3. Commands vs Queries

| Method | Type | Changes state? |
|---|---|---|
| `borrow()` | Command | Yes |
| `returnBook()` | Command | Yes |
| `info()` | Query | No |
| `isOverdue()` | Query | No |

## 4. Collaboration

```dart
class LoanService {
  Loan checkout(Member member, Book book) {
    book.borrow(); // tell Book
    return Loan(member, book, DateTime.now().add(Duration(days: 14)));
  }

  void checkin(Loan loan) {
    loan.book.returnBook();
  }
}
```

**Takeaway:** Each class has one job; objects collaborate via clear methods.
