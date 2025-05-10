# Result Management System

A Go implementation of a flexible student result management system that supports multiple student types with different grading logics.

## Overview

This system demonstrates the use of Go's powerful interface and struct capabilities to model different types of students (Engineering and Arts) with their own unique grading mechanisms while maintaining a clean, unified API.

## Key Features

- Support for multiple student types (Engineering, Arts)
- Type-specific grading policies and GPA calculations
- Comprehensive error handling and input validation
- Clean API for managing students and their academic records

## Implementation Details

### Core Components

1. **Interfaces**
   - `Student`: Common operations for all student types
   - `Grading`: Encapsulates grading logic that varies by student type

2. **Structs**
   - `BaseStudent`: Common student attributes (composition pattern)
   - `EngineeringStudent` & `ArtsStudent`: Type-specific implementations
   - `ResultManagementSystem`: Manages the collection of students

3. **Validation & Error Handling**
   - Custom error types for different failure scenarios
   - Input validation at all levels
   - Proper error propagation

## SOLID Principles Application

The implementation follows SOLID principles:

- **Single Responsibility**: Each struct has a focused purpose
- **Open/Closed**: System can be extended with new student types without modifying existing code
- **Liskov Substitution**: All student implementations follow the Student interface contract
- **Interface Segregation**: Interfaces are focused and minimal
- **Dependency Inversion**: High-level modules depend on abstractions

## Grading Logic

- **Engineering Students**:
  - Must pass each subject individually (grade ≥ 40)
  - GPA calculated on a 10-point scale

- **Arts Students**:
  - Must achieve a minimum average grade (≥ 35)
  - GPA calculated on a 4-point scale

## Best Practices Implemented

- No magic numbers (constants used throughout)
- Strategy pattern for varying behaviors
- Proper encapsulation and data protection
- Comprehensive error handling
- Clean separation of concerns
