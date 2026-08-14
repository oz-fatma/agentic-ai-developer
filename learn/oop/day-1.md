# Day 1: What Is Object-Oriented Programming?

**Language for this track:** Dart

## 1. Compare Styles — Procedural vs Class

**Feature:** Deposit money and check bank account balance.

**Procedural version:**

```dart
int balance = 0;

void deposit(int amount) {
  balance += amount;
}

int getBalance() {
  return balance;
}

void main() {
  deposit(100);
  deposit(50);
  print(getBalance()); // 150
}
```

**Class version:**

```dart
class BankAccount {
  int balance = 0;

  void deposit(int amount) {
    balance += amount;
  }

  int getBalance() {
    return balance;
  }
}

void main() {
  final account = BankAccount();
  account.deposit(100);
  account.deposit(50);
  print(account.getBalance()); // 150
}
```

**Differences:**

| Procedural | OOP (class) |
|---|---|
| `balance` is separate global state | `balance` lives inside `BankAccount` |
| Functions operate on external data | Methods belong to the object |
| Hard to manage multiple accounts | Each `BankAccount()` has its own balance |
| Data and behavior are scattered | Data and behavior are grouped together |

## 2. Identify Objects — Real-World Domain

**Domain:** Library

| Object | Responsibilities |
|---|---|
| **Book** | Store title, author, ISBN; know if available or borrowed |
| **Member** | Store name, member ID; borrow and return books |
| **Librarian** | Check out books, check in returns, manage catalog |
| **Loan** | Link a member to a book; track due date and return status |

## 3. Vocabulary Drill

**Class:** A blueprint that defines what data and actions objects of that type have.  
*Example:* `class Book { String title; void displayInfo() {} }`

**Object:** A concrete instance created from a class.  
*Example:* `final book = Book();` — one specific book in memory.

**Attribute:** Data stored on an object (also called a field or property).  
*Example:* `book.title` or `account.balance`

**Method:** A function defined on a class that describes what the object can do.  
*Example:* `account.deposit(100)` or `book.displayInfo()`

## 4. Language Choice

**Language:** Dart

**Why this language:**
- Clean, readable class syntax — good for learning OOP
- Used in Flutter for mobile/web apps
- Strong typing helps catch mistakes early
- Same language for OOP practice and future UI work if needed
