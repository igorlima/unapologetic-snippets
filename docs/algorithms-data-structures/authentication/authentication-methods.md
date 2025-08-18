---
layout: default
title: Authentication Methods
nav_order: 1
parent: Authentication
grand_parent: Algorithms and Data Structures
has_children: false
has_toc: true
permalink: /docs/algorithms-and-data-structures/authentication/authentication-mehtods
---

- [go back]({% link docs/algorithms-data-structures/authentication/index.md %})

<details markdown="block">
  <summary>
    Table of contents
  </summary>
  {: .text-delta }
1. TOC
{:toc}
</details>

# Authentication Methods

Below will break down six common authentication methods — **Basic Auth, Cookies, Tokens, API Keys, OTP, and SSO** — and explain when (and why) to use each. [^1]

![img](https://github.com/user-attachments/assets/49d7ec07-ebc6-4c4f-becd-9e68483c432e)

## Basic Authentication

Basic authentication works by sending the encoded username and password to the server with every request. These credentials are included in the request headers.

The username and password are combined into a single string and encoded using Base64. This encoded string is then sent in the Authorization header of the HTTP request.

![img](https://github.com/user-attachments/assets/c514e258-313a-4eed-bc5a-e8cb5f2fad0c)


## Cookie-Based Authentication

Cookie-based authentication works by generating a session ID when the user logs in, which is then stored as a cookie in the user’s browser. For subsequent requests, the browser automatically includes this cookie, allowing the server to verify the user’s identity.

The generated session ID is given an expiry time to ensure that a single cookie cannot be used for infinity.

![img](https://github.com/user-attachments/assets/700c7fcb-3f67-4493-91ba-7d08a62b808b)


## Token-Based Authentication

In this authentication method, the server responds to a user’s login request by issuing a token. The token is self-contained, meaning it carries all the information needed for verification, and it’s designed to be tamper-proof.

In token-based authentication, the server doesn’t need to store any session data — only a public key (if using signed tokens like JWTs) to verify the token provided by the client. This contrasts with session-based authentication, where the server must maintain session state.

Once the client obtains a token, it can include it in the request headers to make authenticated requests. This makes token-based authentication stateless and well-suited for scalable, distributed systems.

![img](https://github.com/user-attachments/assets/09a737c1-28f6-450c-b19c-26b1d1b62df7)

> The most popular token based authentication is Json Web Tokens (or JWT).

## API Key-Based Authentication

API key-based authentication is a method for authenticating users or applications accessing an API (Application Programming Interface).

An API key is a unique identifier (usually a long alphanumeric string) issued to a developer or application to access an API. It acts as a simple access token, identifying the calling project or app.

API keys are issued when developers register with an API provider and request access. These keys are typically associated with specific projects and can be configured with varying levels of access control. API providers use them to monitor usage and enforce rate limits, often for billing, analytics, or security purposes.

![img](https://github.com/user-attachments/assets/9d271fd5-b709-44b1-a647-5777882ca25d)

There are several different ways in which a client can send an API key to the server:
1. **HTTP Header** — `Authorization: Api-Key <your_api_key> or x-api-key: <your_api_key>`
1. **Query Parameter** — `https://api.example.com/data?api_key=<your_api_key>`
1. **Request Body** (for POST requests) — `{ "api_key": "<your_api_key>" }`

## OTP Based Authentication

Another widely used authentication method in modern software applications is OTP-based authentication (One-Time Password). Unlike traditional fixed passwords, OTP authentication relies on a unique, time-sensitive code that is valid for only one login session or transaction.

When a user attempts to log in, the server generates a one-time password with a short expiry window (typically 30 seconds to a few minutes). This OTP is sent to the user via a trusted, pre-verified channel — such as a registered mobile number (via SMS) or email address. The user must enter the correct OTP to complete the login process.

This method adds an additional layer of security by:
- Reducing the risk of password reuse and phishing attacks
- Ensuring that only the rightful user (with access to the trusted device/email) can log in
- Once the OTP is verified and the user is successfully authenticated, the system typically issues a session token or cookie to maintain the user’s authenticated state for the duration of the session. This token is then used for subsequent requests, avoiding the need to re-enter credentials repeatedly.

> 🔐 OTP authentication is often used as a form of two-factor authentication (2FA) when combined with a traditional password, further strengthening account security.


## Single Sign On (SSO)

Single Sign-On (SSO) is a way to let users log in just once and then access multiple apps or systems without having to log in again each time. After the user signs in through the SSO provider, a token is created and passed to other trusted services, so the user can move between them smoothly without being asked to enter their credentials again.

![img](https://github.com/user-attachments/assets/048f83b7-6679-4589-bfe0-119db9282ebe)


Here is how it works:
- The user visits the website and tries to log in.
- The website checks if the user is already logged in through the presence of some tokens or cookies. If it is, the request is processed.
- If not, the user is redirected to the SSO login page.
- The user logs in through the SSO system.
- SSO verifies credentials with the identity provider (e.g., Active Directory).
- Upon successful verification, SSO sends authentication info to the website.
- The website grants access and uses tokens/cookies to maintain authentication across other apps or pages.

------ ------

[^1]: [6 Authentication Methods I Wish I Knew Before Building My First Web App](https://medium.com/write-a-catalyst/6-authentication-methods-i-wish-i-knew-before-building-my-first-web-app-abb969e96dc6)
