You are building a document versioning system where each document has multiple versions. Users can undo and redo changes made to the document. The system needs to efficiently store document versions and allow users to navigate through previous and next versions.

To achieve this, implement the VersionManager in Golang using a Doubly Linked List (for maintaining document versions) and a Stack (for tracking undo/redo operations).

Requirements:
Implement a struct VersionManager with the following methods:

AddVersion(content string): Adds a new version of the document.
Undo() string: Moves back to the previous version, returning the current version content.
Redo() string: Moves forward to the next version if Undo was previously called, returning the new version content.
GetCurrentVersion() string: Returns the current version content.
Constraints:

A Doubly Linked List should be used to store the document versions.
A Stack should be used to keep track of undone versions for redoing actions.
If a new version is added after undoing, the redo stack should be cleared.
The system should allow at most N versions (if N=5, the oldest version should be removed when adding a new one).
Example Usage
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
Expected Implementation:
Use a Doubly Linked List for document version traversal.
Use a Stack to keep track of undone versions.
When undoing, push the current version to the redo stack.
When adding a new version after undoing, clear the redo stack.