# Pitfalls: Custom validators

- Must return bool not error
- RegisterValidation is NOT thread-safe — register before use
- Overwriting baked-in validators breaks standard behavior
