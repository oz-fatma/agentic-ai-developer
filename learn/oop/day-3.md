# Day 3: Constructors and Initialization

**Language:** Dart · **Domain:** Library

## 1. Write a Constructor — Require Essential Fields

```dart
class Book {
  final String title;
  final String author;

  Book({required this.title, required this.author});
}
```

## 2. Validate Inputs — Fail Fast

```dart
class Book {
  final String title;
  final String author;
  final double price;

  Book({
    required this.title,
    required this.author,
    required this.price,
  }) {
    if (title.trim().isEmpty) throw ArgumentError('Title cannot be empty');
    if (author.trim().isEmpty) throw ArgumentError('Author cannot be empty');
    if (price < 0) throw ArgumentError('Price cannot be negative');
  }
}
```

## 3. Default Values — Optional Fields

```dart
class Book {
  final String title;
  final String author;
  final bool isAvailable;
  final int edition;

  Book({
    required this.title,
    required this.author,
    this.isAvailable = true,
    this.edition = 1,
  });
}
```

## 4. Named Constructor / Factory

```dart
class Book {
  final String title;
  final String author;
  final bool isAvailable;

  Book({required this.title, required this.author, this.isAvailable = true});

  // Named constructor for a common case
  Book.unknown({required this.title})
      : author = 'Unknown Author',
        isAvailable = true;

  // Factory for digital copies
  factory Book.digital({required String title, required String author}) {
    return Book(title: '$title (Digital)', author: author, isAvailable: true);
  }
}

final b1 = Book(title: 'OOP Guide', author: 'Jane Doe');
final b2 = Book.unknown(title: 'Old Manuscript');
final b3 = Book.digital(title: 'Flutter Basics', author: 'Google');
```

**Takeaway:** Constructors enforce invariants at creation — no half-built objects.
