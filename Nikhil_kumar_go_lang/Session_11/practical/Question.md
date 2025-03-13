# Document Versioning System

A robust document versioning system that allows users to undo and redo changes made to documents. The system efficiently stores document versions and allows users to navigate through previous and next versions.

## Implementation

The system is implemented using:
- **Doubly Linked List**: For maintaining document versions
- **Stack**: For tracking undo/redo operations

## Requirements

1. Implement a struct `VersionManager` with the following methods:
   * `AddVersion(content string)`: Adds a new version of the document.
   * `Undo() string`: Moves back to the previous version, returning the current version content.
   * `Redo() string`: Moves forward to the next version if Undo was previously called, returning the new version content.
   * `GetCurrentVersion() string`: Returns the current version content.

2. **Constraints:**
   * A **Doubly Linked List** should be used to store the document versions.
   * A **Stack** should be used to keep track of undone versions for redoing actions.
   * If a new version is added **after undoing**, the redo stack should be cleared.
   * The system should allow at most **N versions** (if N=5, the oldest version should be removed when adding a new one).

## Example Usage

```go
vm := VersionManager{}
// Adding Versions 
vm.AddVersion("Version 1") 
vm.AddVersion("Version 2") 
vm.AddVersion("Version 3") 
fmt.Println(vm.GetCurrentVersion()) // Output: "Version 3"

vm.Undo() 
fmt.Println(vm.GetCurrentVersion()) // Output: "Version 2"

vm.Undo() 
fmt.Println(vm.GetCurrentVersion()) // Output: "Version 1"

vm.Redo() 
fmt.Println(vm.GetCurrentVersion()) // Output: "Version 2"

vm.AddVersion("Version 4") // This clears the redo stack 
fmt.Println(vm.GetCurrentVersion()) // Output: "Version 4"

vm.Undo() 
fmt.Println(vm.GetCurrentVersion()) // Output: "Version 2"
```

## Expected Implementation

* Use a **Doubly Linked List** for document version traversal.
* Use a **Stack** to keep track of undone versions.
* When undoing, push the current version to the redo stack.
* When adding a new version after undoing, clear the redo stack.
