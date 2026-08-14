# Day 11: Inheritance Basics

**Language:** Dart · **Domain:** Library media types

## 1. Base Class — Shared Fields/Methods

```dart
class MediaItem {
  final String title;
  final String id;
  bool isAvailable = true;

  MediaItem({required this.title, required this.id});

  void borrow() {
    if (!isAvailable) throw StateError('Not available');
    isAvailable = false;
  }

  void returnItem() => isAvailable = true;

  String describe() => '$title ($id)';
}
```

## 2. Extend — Two Subclasses

```dart
class Book extends MediaItem {
  final String author;
  Book({required super.title, required super.id, required this.author});

  @override
  String describe() => '${super.describe()} by $author';
}

class DVD extends MediaItem {
  final int durationMinutes;
  DVD({required super.title, required super.id, required this.durationMinutes});

  @override
  String describe() => '${super.describe()} — $durationMinutes min';
}
```

## 3. Reuse Inherited Methods

```dart
final book = Book(title: 'OOP', id: 'B001', author: 'Jane');
final dvd = DVD(title: 'Dart Course', id: 'D001', durationMinutes: 120);

book.borrow(); // inherited from MediaItem
dvd.borrow();  // same
```

## 4. Shallow Hierarchy — One Level Only

```
MediaItem (base)
├── Book
└── DVD
```

No `DigitalBook extends Book extends MediaItem` tower — keep it flat.

**Takeaway:** Inheritance shares behavior; keep hierarchies shallow.
