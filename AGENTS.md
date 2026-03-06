# AGENTS.md — MindSpec Project
<!-- BEGIN mindspec:managed -->

This project uses [MindSpec](https://github.com/mrmaxsteel/mindspec), a spec-driven development framework.

## Workflow

Run `mindspec instruct` for mode-appropriate operating guidance.

## Build & Test

```bash
make build    # Build binary
make test     # Run all tests
```

## Modes

This project follows a strict spec-driven workflow with human gates:

1. **Spec** — define the problem and acceptance criteria (no code)
2. **Plan** — break the spec into implementation beads (no code)
3. **Implement** — write code against the approved plan
4. **Review** — verify implementation meets acceptance criteria

Transition between modes using `mindspec approve spec|plan` and `mindspec complete`.

## Conventions

- Every functional change must reference a spec in `.mindspec/docs/specs/`
- In Spec and Plan modes, only documentation may be created or modified — no code changes
- Working tree must be clean before switching modes
- Run `mindspec doctor` to verify project structure health
<!-- END mindspec:managed -->
