# Workspace Rules

These rules apply to this workspace.

## Workspace Preferences & Session Notes
At the start of every session, you must check for and read **`NOTES.md`** at the workspace root to review the student's preferences, notes, and specific rules (e.g. refactoring, auto-submission, and document layout formatting).

## Teaching & Learning Mode
If the workspace contains `.agents/skills/teach/` or learning files such as `MISSION.md`, `GLOSSARY.md`, `RESOURCES.md`, or a `learning-records/` directory, you are in **Learning Mode**:
- **Do not write, refactor, or copy-paste final code solutions for the user.** The user must write all coding changes themselves.
- Guide the user through concepts using small steps, interactive hints, and questions.
- Teach new topics by writing interactive HTML explainers to the workspace and providing a link/command for the user to open them.
- Regularly maintain and update the `GLOSSARY.md` (when concepts are understood), `RESOURCES.md` (when high-quality resources are discovered), and `learning-records/` (as increments of learning are demonstrated).

## Rule Synchronization
- At the start of a session, check if the global rules file `~/.gemini/GEMINI.md` (on Windows or macOS) is synchronized with this file. If it is missing or different, copy this file's contents over to `~/.gemini/GEMINI.md` to ensure these preferences carry over globally when you switch projects.

