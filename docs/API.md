# API Documentation

## Endpoints

### POST /api/v1/two-sum

Solves the Two Sum problem by finding two numbers in an array that add up to a target value.

### POST /api/v1/palindrome

Checks if a given text is a palindrome. Ignores case, spaces, and non-alphanumeric characters.

#### Request

**Content-Type:** `application/json`

```json
{
    "nums": [2, 7, 11, 15],
    "target": 9
}
```

**Parameters:**
- `nums` (array of integers): Array of numbers to search through
- `target` (integer): Target sum value

#### Response

**Success (200 OK):**
```json
{
    "indices": [0, 1]
}
```

**Error (400 Bad Request):**
```json
{
    "error": "no solution found"
}
```

#### Examples

**Example 1:**
```bash
curl -X POST http://localhost:8081/api/v1/two-sum \
  -H "Content-Type: application/json" \
  -d '{"nums": [2, 7, 11, 15], "target": 9}'
```
Response: `{"indices": [0, 1]}`

**Example 2:**
```bash
curl -X POST http://localhost:8081/api/v1/two-sum \
  -H "Content-Type: application/json" \
  -d '{"nums": [3, 2, 4], "target": 6}'
```
Response: `{"indices": [1, 2]}`

**Example 3:**
```bash
curl -X POST http://localhost:8081/api/v1/two-sum \
  -H "Content-Type: application/json" \
  -d '{"nums": [3, 3], "target": 6}'
```
Response: `{"indices": [0, 1]}`

#### Error Cases

- Array with less than 2 elements
- No valid solution exists
- Invalid JSON format
- Missing required fields

---

## POST /api/v1/palindrome

Checks if a given text is a palindrome, ignoring case, spaces, and non-alphanumeric characters.

#### Request

**Content-Type:** `application/json`

```json
{
    "text": "A man, a plan, a canal: Panama"
}
```

**Parameters:**
- `text` (string): Text to check for palindrome properties

#### Response

**Success (200 OK):**
```json
{
    "is_palindrome": true,
    "clean_text": "amanaplanacanalpanama"
}
```

#### Examples

**Example 1: Simple palindrome**
```bash
curl -X POST http://localhost:8081/api/v1/palindrome \
  -H "Content-Type: application/json" \
  -d '{"text": "racecar"}'
```
Response: `{"is_palindrome": true, "clean_text": "racecar"}`

**Example 2: Complex palindrome with punctuation**
```bash
curl -X POST http://localhost:8081/api/v1/palindrome \
  -H "Content-Type: application/json" \
  -d '{"text": "A man, a plan, a canal: Panama"}'
```
Response: `{"is_palindrome": true, "clean_text": "amanaplanacanalpanama"}`

**Example 3: Not a palindrome**
```bash
curl -X POST http://localhost:8081/api/v1/palindrome \
  -H "Content-Type: application/json" \
  -d '{"text": "hello world"}'
```
Response: `{"is_palindrome": false, "clean_text": "helloworld"}`

**Example 4: Case insensitive**
```bash
curl -X POST http://localhost:8081/api/v1/palindrome \
  -H "Content-Type: application/json" \
  -d '{"text": "Taco Cat"}'
```
Response: `{"is_palindrome": true, "clean_text": "tacocat"}`