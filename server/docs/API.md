# Kaipaste API Documentation

Welcome to the **Kaipaste API**. This document describes the endpoints and usage for interacting programmatically with Kaipaste to securely share code and text.

---

## Base URL

```
https://kaipaste.****.com/api
```

Replace `kaipaste.example.com` with your deployment domain.

---

## Authentication

- Most endpoints require authentication via an API key or token.
- Token is stored in cookies for web clients.

---

## Endpoints

### 1. Create a paste

**POST** `/paste`

Create a new code/text paste.

**Request Body**

```json
{
  "title": "Hello World Example",
  "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
  "language": "go",
  "password": null,
  "burnAfterRead": false,
  "expiresAt": "2025-12-12T12:00:00Z",
  "views": 0,
  "isPublic": true
}
```

**Response**

```json
{
  "pasteId": "b7346fc0-d340-11f0-9940-8c8d28d905b2",
  "success": true
}
```

---

### 2. Retrieve a paste

**GET** `/paste/{id}`

Fetch the content for a specific paste.

**Optional query:**

- If the paste is password protected, pass: `?password=your_password`

**Response**

```json
{
  "paste": {
    "id": "b7346fc0-d340-11f0-9940-8c8d28d905b2",
    "title": "Hello World Example",
    "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
    "language": "go",
    "password": null,
    "burnAfterRead": false,
    "expiresAt": "2025-12-12T15:00:00+03:00",
    "createdAt": "2025-12-07T10:45:35.352005+03:00",
    "views": 0,
    "userId": null,
    "isPublic": true
  }
}
```

---

### 3. Delete a paste

**DELETE** `/paste/{id}`

Permanently deletes a paste.

**Response**

```json
{
  "success": true
}
```

---

### 4. List Your pastes

**GET** `/paste/mine`

Returns a list of pastes created by the authenticated user.

**Response**

```json
{
  pastes: [
  {
    "id":"b7346fc0-d340-11f0-9940-8c8d28d905b2",
    "title": "Hello World Example",
    "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
    "language": "go",
    "password": null,
    "burnAfterRead": false,
    "expiresAt": "2025-12-12T15:00:00+03:00",
    "createdAt": "2025-12-07T10:45:35.352005+03:00",
    "views": 0,
    "userId": "b7346fc0-d340-11f0-9940-8c8d28d905b2",
    "isPublic": true
  },
  {
    "id": "b7346fc0-d340-11f0-9940-8c8d28d905b2",
    "title": "Hello World Example",
    "content": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
    "language": "go",
    "password": null,
    "burnAfterRead": false,
    "expiresAt": "2025-12-12T15:00:00+03:00",
    "createdAt": "2025-12-07T10:45:35.352005+03:00",
    "views": 0,
    "userId": "b7346fc0-d340-11f0-9940-8c8d28d905b2",
    "isPublic": true
  }
    ]
  success: true
  ...
}
```

---

## Status Codes

- `200 OK` – Success
- `201 Created` – New resource created
- `400 Bad Request` – Invalid input
- `401 Unauthorized` – Invalid or missing token
- `404 Not Found` – paste does not exist or expired
- `500 Internal Server Error` – Server-side issue

---


## The documentation is being developed and improved continuously.

For feedback or help, reach out via [GitHub Issues](https://github.com/JiruGutema/Kaipaste/issues).

---
