# Day 5: OOP Fundamentals — Practice

**Language:** Dart · **Domain:** Library loan system

## 1. Mini Domain — 3 Classes

```dart
class Book {
  final String title;
  bool _available = true;
  Book(this.title);
  bool get isAvailable => _available;
  void borrow() {
    if (!_available) throw StateError('Not available');
    _available = false;
  }
  void returnBook() => _available = true;
}

class Member {
  final String id;
  final String name;
  Member({required this.id, required this.name});
}

class Loan {
  final Member member;
  final Book book;
  final DateTime dueDate;
  Loan({required this.member, required this.book, required this.dueDate});
}

class LoanService {
  Loan checkout(Member member, Book book) {
    book.borrow();
    return Loan(
      member: member,
      book: book,
      dueDate: DateTime.now().add(const Duration(days: 14)),
    );
  }

  void checkin(Loan loan) => loan.book.returnBook();
}
```

## 2. Demo Scenarios

```dart
void main() {
  final book = Book('Effective Dart');
  final member = Member(id: 'M001', name: 'Ayşe');
  final service = LoanService();

  // Scenario 1: checkout
  final loan = service.checkout(member, book);
  print('${member.name} borrowed ${book.title}');

  // Scenario 2: try double borrow — fails
  try {
    book.borrow();
  } catch (e) {
    print('Error: $e');
  }

  // Scenario 3: return
  service.checkin(loan);
  print('Book available: ${book.isAvailable}');

  // Scenario 4: second member borrows
  final member2 = Member(id: 'M002', name: 'Mehmet');
  service.checkout(member2, book);
  print('${member2.name} borrowed ${book.title}');
}
```

## 3. Enforce Creation Rules

```dart
class Member {
  final String id;
  final String name;
  Member({required this.id, required this.name}) {
    if (id.isEmpty) throw ArgumentError('ID required');
    if (name.trim().length < 2) throw ArgumentError('Name too short');
  }
}
```

## 4. Review Checklist

- [x] `Book` — availability and borrow/return logic
- [x] `Member` — identity data with validation
- [x] `Loan` — links member + book + due date
- [x] No public raw fields for critical state (`_available` is private)
- [x] Each class has a clear responsibility

**Phase 1 complete (Days 1–5).**
