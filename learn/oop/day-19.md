# Day 19: Observer & Refactoring Legacy OOP

**Language:** Dart

## 1. Observer Lite

```dart
class LoanEvents {
  final List<void Function(String)> _listeners = [];

  void subscribe(void Function(String) listener) {
    _listeners.add(listener);
  }

  void notify(String event) {
    for (final listener in _listeners) {
      listener(event);
    }
  }
}

class Library {
  final LoanEvents events = LoanEvents();

  void checkout(Member member, Book book) {
    book.borrow();
    events.notify('${member.name} borrowed ${book.title}');
  }
}

// Usage
library.events.subscribe((e) => print('LOG: $e'));
library.events.subscribe((e) => print('EMAIL: Overdue check for $e'));
```

## 2. Legacy Refactor — God Class to Objects

**Before:**

```dart
class LibraryApp {
  List<String> books = [];
  List<String> members = [];
  void addBook(String t) => books.add(t);
  void addMember(String n) => members.add(n);
  void borrow(int bi, int mi) { /* 30 lines */ }
  void printReport() { /* 20 lines */ }
}
```

**After — extract 3 objects:**

```dart
class BookCatalog { /* manages books */ }
class MemberRegistry { /* manages members */ }
class LoanService { /* borrow/return logic */ }
class LibraryApp {
  final catalog = BookCatalog();
  final members = MemberRegistry();
  late final loans = LoanService(catalog, members);
}
```

## 3. Characterization — Before Changing

Document existing behavior:
- [ ] Member can borrow if book available
- [ ] Same book cannot be borrowed twice
- [ ] Report lists all books and status

## 4. Small Steps

1. Extract `Book` class — verify tests pass
2. Extract `Member` class — verify
3. Extract `LoanService` — verify
4. Remove god class — verify

**Takeaway:** Characterize first, extract objects, verify after each step.
