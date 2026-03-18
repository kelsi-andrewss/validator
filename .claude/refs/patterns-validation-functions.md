# Patterns: Validation Functions

### FieldLevel interface

Validators receive FieldLevel with Field(), Param(), GetTag(). Return bool.

Example: `baked_in.go`

### StructLevel interface

Struct validators get Current(), Parent(), ReportError()

Example: `struct_level.go`
