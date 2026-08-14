# Day 12: Method Overriding

**Language:** Dart

## 1. Override a Method

```dart
class Animal {
  String speak() => '...';
}

class Dog extends Animal {
  @override
  String speak() => 'Woof!';
}

class Cat extends Animal {
  @override
  String speak() => 'Meow!';
}
```

## 2. Call Super

```dart
class Book extends MediaItem {
  final String author;

  Book({required super.title, required super.id, required this.author});

  @override
  String describe() {
    return '${super.describe()} by $author'; // extend base behavior
  }
}
```

## 3. Substitution Demo — Polymorphism

```dart
void main() {
  final List<Animal> animals = [Dog(), Cat(), Dog()];

  for (final animal in animals) {
    print(animal.speak()); // Woof! Meow! Woof! — runtime dispatch
  }
}
```

```dart
final List<MediaItem> catalog = [
  Book(title: 'Dart', id: 'B1', author: 'Chris'),
  DVD(title: 'OOP Film', id: 'D1', durationMinutes: 90),
];

for (final item in catalog) {
  print(item.describe()); // correct subclass version runs
}
```

## 4. Document Override

```dart
// Override: Book adds author info to base description
@override
String describe() => '${super.describe()} by $author';
```

**Takeaway:** Base type reference, subclass behavior at runtime — that's polymorphism.
