# Contributing to Our Project

We're excited that you're interested in contributing to our project! To maintain our project's integrity and workflow efficiency, we adhere to certain development practices and standards. Please take a moment to familiarize yourself with this guide before contributing.

## Code of Conduct

Please note that this project is released with a Contributor Code of Conduct. By participating in this project, you agree to abide by its terms. (Link to the Code of Conduct)

## Trunk Based Development (TBD)

We use Trunk Based Development as our branching strategy:

- **Main Branch**: All development work is merged back to the main branch as soon as it's ready.
- **Short-lived Feature Branches**: If you're working on a feature or a fix, create a branch from the main trunk. Keep it short-lived (merge it back within a couple of days).
- **Feature Toggles**: For larger features, consider using feature toggles to merge your work without affecting the production.

For more details, refer to our [ADR on Trunk Based Development](decisions/002-trunk-based-development.md).

## Commit Messages

We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for our commit messages:

- Start your commit message with a type (`feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`).
- Include a scope if applicable.
- Use the imperative mood for the subject line.
- Separate the subject from the body with a blank line and use the body to explain the "what" and "why" of the commit.

Example: `feat(authentication): implement OAuth login flow`

For more details, refer to our [ADR on Conventional Commits](decisions/003-conventional-commits.md).

## Pull Requests

- Ensure your code adheres to the project's coding standards.
- Write meaningful commit messages following the Conventional Commits guidelines.
- Keep your pull requests concise and focused. They should relate to a single issue or feature.
- Include any necessary tests and documentation updates with your changes.

## Reporting Bugs

- Use the issue tracker to report bugs.
- Describe the bug in detail, including steps to reproduce, expected behavior, and actual behavior.
- If possible, include screenshots or code snippets.

## Requesting Enhancements

- Use the issue tracker to suggest enhancements.
- Be clear and concise in your description of the proposed feature or improvement.
- Explain how this enhancement would benefit the project.

Thank you for contributing to our project! Your efforts help us build a better product.

