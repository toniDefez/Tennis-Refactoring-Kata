_Support this and all my katas via [Patreon](https://www.patreon.com/EmilyBache)_

# Tennis Refactoring Kata (Go Version)

This is a Go-only implementation of Emily Bache's [Tennis Refactoring Kata](https://sammancoaching.org/kata_descriptions/tennis.html).

You can learn more about this exercise in [this YouTube video](https://youtu.be/XifUs1FhWRc), where Emily explains its purpose and how to approach it. If you’re using this version, you are likely interested in writing clean, testable, and idiomatic Go code.

---

## 🏗️ The Scenario

You’ve joined a consulting company. One of your colleagues has been working on a system for the Tennis Society. The contract includes 10 hours of billable work, and your colleague has already used 8.5 hours.

Unfortunately, they’ve fallen ill. They claim the work is complete and that all tests pass. Your manager has asked you to take over and spend up to an hour tidying up the code so she can invoice the client for the full 10 hours.

Your task:
- Review and refactor the code (without changing the tests).
- Make notes to provide feedback on your colleague's design.
- Be prepared to explain the value of the refactoring you've done, beyond just fulfilling billable time.

---

## 💡 Refactoring Focus

Start with the `TennisGame1` implementation. It has some design smells and hardcoded elements (like `"player1"` and `"player2"`) that should be addressed.

The provided test suite is fast and comprehensive. **You should not change the tests**, only run them frequently while you refactor.

---

## ✅ Goals

- Improve readability and maintainability
- Apply Go idioms and good design practices
- Use encapsulation where it makes sense (e.g. defining a `Player` struct)
- Reduce duplication and complexity
- Fix hardcoded values
- Optionally: separate formatting logic from domain logic

---

## 🧪 Running the Tests

Assuming you have Go installed:

```bash
go test ./...

