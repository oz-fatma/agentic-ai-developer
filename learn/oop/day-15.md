# Day 15: Inheritance & Polymorphism — Practice

**Language:** Dart · **Domain:** Payment notifications

## 1. Design Challenge — Contract + Implementations

```dart
abstract class Notifier {
  void notify(String recipient, String message);
}

class EmailNotifier implements Notifier {
  @override
  void notify(String recipient, String message) {
    print('Email to $recipient: $message');
  }
}

class PushNotifier implements Notifier {
  @override
  void notify(String recipient, String message) {
    print('Push to $recipient: $message');
  }
}

class SmsNotifier implements Notifier {
  @override
  void notify(String recipient, String message) {
    print('SMS to $recipient: $message');
  }
}
```

## 2. Polymorphic Loop

```dart
void sendBulkReminder(List<Notifier> notifiers, String member, String msg) {
  for (final notifier in notifiers) {
    notifier.notify(member, msg); // polymorphic dispatch
  }
}
```

## 3. Composition Pass

```dart
class ReminderService {
  final List<Notifier> _notifiers;
  final Logger _logger; // composed, not inherited

  ReminderService(this._notifiers, this._logger);

  void remindOverdue(Member member, Book book) {
    final msg = 'Return "${book.title}" — overdue!';
    for (final n in _notifiers) {
      n.notify(member.name, msg);
    }
    _logger.log('Reminder sent to ${member.name}');
  }
}
```

## 4. Critique — Avoided Deep Inheritance

**Avoided:** `EmailReminder extends Reminder extends Notification extends Message`

**Used instead:** `Notifier` interface + composition with `Logger`

**Why:** Notification types share behavior (send message), not identity hierarchy.

**Phase 3 complete (Days 11–15).**
