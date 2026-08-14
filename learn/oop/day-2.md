# Day 2: Classes and Objects

**Language:** Dart · **Domain:** Library

## 1. Define a Class

```dart
class Book {
  String title;
  String author;
  bool isAvailable;

  Book(this.title, this.author, {this.isAvailable = true});
}
```

## 2. Create Instances

```dart
final book1 = Book('Clean Code', 'Robert Martin');
final book2 = Book('Dart in Action', 'Chris Buckett');
```

Two objects from the same class — different state.

## 3. Call Methods

```dart
class Book {
  String title;
  String author;
  bool isAvailable;

  Book(this.title, this.author, {this.isAvailable = true});

  void borrow() {
    if (!isAvailable) throw StateError('Book already borrowed');
    isAvailable = false;
  }

  void returnBook() {
    isAvailable = true;
  }

  String info() => '$title by $author (${isAvailable ? "available" : "borrowed"})';
}

book1.borrow();
print(book1.info()); // Clean Code by Robert Martin (borrowed)
print(book2.info()); // Dart in Action by Chris Buckett (available)
```

## 4. Inspect State — Independent Instances

```dart
void main() {
  final a = Book('1984', 'George Orwell');
  final b = Book('1984', 'George Orwell'); // same values, different objects

  a.borrow();
  print(a.isAvailable); // false
  print(b.isAvailable); // true — independent identity
  print(identical(a, b)); // false
}
```

**Takeaway:** Class = blueprint. Each instance has its own state.
