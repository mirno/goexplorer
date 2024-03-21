# 2. Use Conventional Commits

Date: 03-21-2024

## Status

Proposed

## Context

To improve the readability and navigability of our project's history, we need a standardized commit message format. Conventional Commits offer a lightweight convention for creating descriptive commit messages that are easy to understand and can be automatically parsed for generating release notes or versioning.

## Decision

We will adopt the Conventional Commits specification for all commit messages. This involves:

- Prefixing commit messages with types such as `feat`, `fix`, `docs`, `style`, `refactor`, `test`, and `chore`.
- Including a scope when relevant to specify the part of the codebase affected.
- Writing a concise description of the changes made.

## Consequences

- Improved commit log readability.
- The ability to automate changelog generation and semantic versioning.
- Requires education and discipline from all contributors to adhere to the specification.
