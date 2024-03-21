# 1. Use Trunk Based Development

Date: 03-21-2024

## Status

Proposed

## Context

The project needs a version control strategy that supports our continuous integration and deployment goals. Trunk Based Development (TBD) is a source-control branching model focused on short-lived branches and simplified merging, which can enhance our workflow by reducing complexity and supporting rapid releases.

## Decision

We will adopt Trunk Based Development as our primary branching strategy. This involves:

- Developers branching off from the main trunk for new features or fixes.
- Keeping branches short-lived (ideally merged within a day or two).
- Using feature toggles if necessary to merge incomplete features without affecting production.

## Consequences

- Reduced complexity in managing branches.
- Easier collaboration and integration among team members.
- Potential need for increased discipline in code quality and testing to ensure the trunk remains stable.
