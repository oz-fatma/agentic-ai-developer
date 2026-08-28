# Day 51: Auth & Security — Authentication Basics

## 1. Hash Passwords

Used `bcrypt` — never store plaintext passwords.

## 2. Registration Flow

Created users with unique emails and hashed passwords.

## 3. Login Flow

Verified credentials with constant-time compare.

## 4. Protect Routes

Required authentication on selected handlers.

## Run

```bash
cd learn/go
go run ./day51
```

## Summary

| Concept | Takeaway |
|---|---|
| Hash Passwords | Used `bcrypt` — never store plaintext passwords |
| Registration Flow | Created users with unique emails and hashed passwords |
| Login Flow | Verified credentials with constant-time compare |
| Protect Routes | Required authentication on selected handlers |
