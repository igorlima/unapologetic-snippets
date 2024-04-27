---
layout: default
title: Authentication
nav_order: 1
parent: Algorithms and Data Structures
has_children: true
has_toc: true
permalink: /docs/algorithms-and-data-structures/authentication
---

# Authentication

- [go back]({% link docs/algorithms-data-structures/index.md %}#authentication)

## JWT [^1]

JWT authentication and session authentication are ways to authenticate users of your web app.

JWT stands for JSON Web Token, and it is a commonly used stateless user authentication standard used to securely transmit information between client and server in a JSON format.

A JWT is encoded by default, not encrypted. It is digitally signed with a secret key that only the server knows. Additionally, it has the capability to be encrypted.

### Structure of JWT [^1]

A JSON Web Token consists of 3 parts separated by a period.
The header, the payload, and the signature.
Each section is base64 encoded.

![image](https://github.com/igorlima/unapologetic-snippets/assets/1886786/0bb79cff-3ea0-4000-bff0-948c90d814e6)

__Header__

The header consists of token type, which is JWT, and the signing algorithm used, such as HMAC SHA256 or RSA.

```json
{
  "typ": "JWT",
  "alg": "HS256"
}
```

__Payload__

A payload carries claims, which are statements about the user along with additional data. These claims typically include the time the token was issued (`iat`) and its expiration time (`exp`), as tokens should have a limited lifespan.

```json
{
  "iss": "example_issuer",
  "sub": "user_id123",
  "exp": 1644768000,
  "iat": 1644744000
}
```

__Signature__

The signature is most important part of a JWT.

It is calculated using the header, the payload, and the secret, which are fed to the signing algorithm to use.

```plaintext
signature = HMAC-SHA256(base64urlEncode(header) + "." + base64urlEncode(payload), secret_salt)
```



------ ------

[^1]: [JWT explained in 4 minutes (With Visuals)](https://dev.to/jaypmedia/jwt-explained-in-4-minutes-with-visuals-g3n)
