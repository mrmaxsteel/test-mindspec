# CLAUDE.md
<!-- mindspec:managed -->

See [AGENTS.md](AGENTS.md) for project conventions shared across all coding agents.

Run `mindspec instruct` for mode-appropriate operating guidance. This is emitted automatically by the SessionStart hook.

## Custom Commands

| Command | Purpose |
|:--------|:--------|
| `/ms-explore` | Enter, promote, or dismiss an Explore Mode session |
| `/ms-spec-init` | Initialize a new specification (enters Spec Mode) |
| `/ms-spec-approve` | Approve spec → Plan Mode |
| `/ms-plan-approve` | Approve plan → Implementation Mode |
| `/ms-impl-approve` | Approve implementation → Idle |
| `/ms-spec-status` | Check current mode and active spec/bead state |
