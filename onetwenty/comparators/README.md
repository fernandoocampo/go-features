# comparable interface

The types are comparable whether they follow below rules:

1. Boolean, numeric, string, pointer, and channel types are strictly comparable.
2. Struct types are strictly comparable if all their field types are strictly comparable.
3. Array types are strictly comparable if their array element types are strictly comparable.
4. Type parameters are strictly comparable if all types in their type set are strictly comparable.