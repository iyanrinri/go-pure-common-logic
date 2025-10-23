# API Documentation

## Endpoints

### POST /api/v1/two-sum

Solves the Two Sum problem by finding two numbers in an array that add up to a target value.

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