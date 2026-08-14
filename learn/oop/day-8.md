# Day 8: Abstraction and Interfaces of Thought

**Language:** Dart · **Domain:** Library checkout

## 1. Raise the Level — Intent-Revealing Method

**Before (low-level steps):**

```dart
book.isAvailable = false;
member.loans.add(loan);
print('Loan created');
```

**After (named intent):**

```dart
library.checkout(member, book);
```

## 2. Stable Abstraction — Interface

```dart
abstract class Notifiable {
  void send(String message);
}

class EmailNotifier implements Notifiable {
  @override
  void send(String message) => print('Email: $message');
}

class SmsNotifier implements Notifiable {
  @override
  void send(String message) => print('SMS: $message');
}
```

## 3. Caller Depends on Abstraction

```dart
class Library {
  final Notifiable notifier;

  Library(this.notifier);

  void checkout(Member member, Book book) {
    book.borrow();
    notifier.send('${member.name}, you borrowed "${book.title}"');
  }
}
```

## 4. Document Intent — Good Names

| Bad | Good |
|---|---|
| `process()` | `checkout()` |
| `updateFlag()` | `markAsReturned()` |
| `calc()` | `calculateFine()` |
| `doStuff()` | `sendOverdueReminder()` |

**Takeaway:** Name methods after outcomes; depend on contracts, not concrete steps.
