# Pitfalls: Tag parsing

- Pipe must be escaped as 0x7C to avoid OR expansion
- Comma must be escaped as 0x2C
- Circular aliases cause stack overflow
