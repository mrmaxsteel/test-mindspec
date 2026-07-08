# CLAUDE.md
<!-- BEGIN mindspec:managed -->

**IMPORTANT**: You MUST read and follow [AGENTS.md](AGENTS.md) as your primary behavioral instructions. AGENTS.md is the canonical source of project conventions, workflow rules, and development guidance shared across all coding agents.

Run `mindspec instruct` for mode-appropriate operating guidance. This is emitted automatically by the SessionStart hook.

## Skills

### Spec lifecycle gates

| Skill | Purpose |
|:------|:--------|
| `/ms-spec-create` | Create a new specification (enters Spec Mode) |
| `/ms-spec-approve` | Approve spec → Plan Mode |
| `/ms-plan-approve` | Approve plan → Implementation Mode |
| `/ms-impl-approve` | Approve implementation → Idle |

### Bead lifecycle

| Skill | Purpose |
|:------|:--------|
| `/ms-bead-impl` | Stage the impl prompt (Phase A) + dispatch the subagent (Phase B) |
| `/ms-bead-fix` | Dispatch a fix-up subagent with the consolidated change list |

### Review panel

| Skill | Purpose |
|:------|:--------|
| `/ms-panel-run` | Step 0 writes the panel dir + BRIEF + `panel.json`; then launch 6 reviewers and collect verdicts |
| `/ms-panel-tally` | Single decision authority: decision matrix, artifact gates, consolidation, halt-recovery |

### Orchestrators

| Skill | Purpose |
|:------|:--------|
| `/ms-bead-cycle` | Single bead end-to-end: pick+claim → impl → panel → fix → re-panel → merge |
| `/ms-spec-autopilot` | Whole spec: cycle every bead until the spec is done |
| `/ms-spec-final-review` | Final panel of the whole spec branch vs main, before `/ms-impl-approve` |

## Bead-loop guardrails (mindspec)

See **AGENTS.md § Bead-loop guardrails (mindspec)** for the canonical orchestrator rules and subagent prompt fences (only the cycle runs `mindspec complete`, after the panel gate passes; never raw `git merge bead/<id>`; one `git push` at end-of-spec; subagents make exactly one commit, tests must PASS). Surviving skills reference that section rather than re-stating it.
<!-- END mindspec:managed -->


<!-- BEGIN BEADS INTEGRATION v:1 profile:minimal hash:6cd5cc61 -->
## Beads Issue Tracker

This project uses **bd (beads)** for issue tracking. Run `bd prime` to see full workflow context and commands.

### Quick Reference

```bash
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --claim  # Claim work
bd close <id>         # Complete work
```

### Rules

- Use `bd` for ALL task tracking — do NOT use TodoWrite, TaskCreate, or markdown TODO lists
- Run `bd prime` for detailed command reference and session close protocol
- Use `bd remember` for persistent knowledge — do NOT use MEMORY.md files

**Architecture in one line:** issues live in a local Dolt DB; sync uses `refs/dolt/data` on your git remote; `.beads/issues.jsonl` is a passive export. See https://github.com/gastownhall/beads/blob/main/docs/SYNC_CONCEPTS.md for details and anti-patterns.

## Agent Context Profiles

The managed Beads block is task-tracking guidance, not permission to override repository, user, or orchestrator instructions.

- **Conservative (default)**: Use `bd` for task tracking. Do not run git commits, git pushes, or Dolt remote sync unless explicitly asked. At handoff, report changed files, validation, and suggested next commands.
- **Minimal**: Keep tool instruction files as pointers to `bd prime`; use the same conservative git policy unless active instructions say otherwise.
- **Team-maintainer**: Only when the repository explicitly opts in, agents may close beads, run quality gates, commit, and push as part of session close. A current "do not commit" or "do not push" instruction still wins.

## Session Completion

This protocol applies when ending a Beads implementation workflow. It is subordinate to explicit user, repository, and orchestrator instructions.

1. **File issues for remaining work** - Create beads for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **Handle git/sync by active profile**:
   ```bash
   # Conservative/minimal/default: report status and proposed commands; wait for approval.
   git status

   # Team-maintainer opt-in only, unless current instructions forbid it:
   git pull --rebase
   git push
   git status
   ```
5. **Hand off** - Summarize changes, validation, issue status, and any blocked sync/commit/push step

**Critical rules:**
- Explicit user or orchestrator instructions override this Beads block.
- Do not commit or push without clear authority from the active profile or the current user request.
- If a required sync or push is blocked, stop and report the exact command and error.
<!-- END BEADS INTEGRATION -->
