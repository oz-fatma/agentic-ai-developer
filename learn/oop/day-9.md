# Day 9: Immutable vs Mutable Design

**Language:** Dart

## 1. Mutable Path

```dart
class MutablePoint {
  int x;
  int y;
  MutablePoint(this.x, this.y);

  void move(int dx, int dy) {
    x += dx;
    y += dy;
  }
}

final p = MutablePoint(0, 0);
p.move(3, 4);
print('(${p.x}, ${p.y})'); // (3, 4) — same object changed
```

## 2. Immutable Path

```dart
class ImmutablePoint {
  final int x;
  final int y;
  const ImmutablePoint(this.x, this.y);

  ImmutablePoint move(int dx, int dy) {
    return ImmutablePoint(x + dx, y + dy); // new instance
  }
}

final p = ImmutablePoint(0, 0);
final p2 = p.move(3, 4);
print('(${p.x}, ${p.y})');  // (0, 0) — original unchanged
print('(${p2.x}, ${p2.y})'); // (3, 4)
```

## 3. Compare Trade-offs

| | Mutable | Immutable |
|---|---|---|
| Ease of use | Simple updates in place | Must reassign reference |
| Debugging | Harder — state changes anywhere | Easier — state never changes |
| Sharing | Risky — shared reference can mutate | Safe — share freely |
| Flutter/UI | Entities with identity | Value objects, state models |

## 4. When to Choose

- **Mutable:** Objects with identity that persist (`Member`, `Book` availability)
- **Immutable:** Value objects (`Money`, `DateRange`, `Coordinates`)
- **Flutter tip:** Prefer immutable data classes for UI state; use `copyWith` pattern

**Takeaway:** Mix both deliberately — not everything needs to be mutable.
