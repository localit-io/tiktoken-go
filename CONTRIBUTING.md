# Contributing Guide

We’re excited that you’re interested in contributing to this project! Whether it's fixing a bug, proposing a feature, or improving documentation, we welcome your input. Below are a few simple guidelines to help you get started.

## Reporting Issues

If you found a bug or want to request a feature, please use our [GitHub issue tracker](https://github.com/localit-io/tiktoken-go/issues).

- For **bugs**, include as much detail as possible:
  - What you were doing when it happened
  - Your OS and Go version
  - Any relevant error output or logs

- For **feature requests**, describe the use case clearly and explain how the new feature would help solve your problem.

If you're reporting a **security-related issue**, you can either open a GitHub issue or contact us directly at <info@wola.io> if it’s sensitive.

## Submitting Code

If you’re planning to submit a code contribution:

1. **Start with an issue**  
   Open an issue first to describe the bug or feature you're working on. This helps avoid duplicate efforts and allows us to coordinate better. Mention in the issue that you’re working on it so it can be assigned to you.

2. **Fork and branch**  
   Fork the repository and create a new branch for your changes. Keep each feature or fix in its own branch to keep pull requests focused and easy to review.

3. **Include tests**  
   Significant changes should include tests to maintain code quality and prevent regressions.

4. **Write clear commit messages**  
   Please follow [good commit message practices](https://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html). This helps keep the history clean and makes reviewing and debugging easier.

5. **Open a pull request**  
   Once your changes are ready, push them to your fork and open a [pull request](https://help.github.com/articles/creating-a-pull-request).  
   **Note:** We use “squash and merge” for all pull requests, so don’t worry about the number of commits in your PR—your work will be merged as a single commit.

## Coding Style

Before submitting a pull request, please take a moment to review the existing code and try to keep your changes consistent with the overall style of the project. We aim to follow Go best practices wherever they make sense.

General guidelines:
- Keep line width below **120 characters** when possible and reasonable.
- We follow a combination of established style guides, including:
    - [Google Go Style Guide](https://google.github.io/styleguide/go/)
    - [Uber Go Style Guide](https://github.com/uber-go/guide)
    - [Effective Go](https://go.dev/doc/effective_go)

If you're unsure about something, feel free to refer to those resources — or open an issue to ask.

Thanks for contributing!

