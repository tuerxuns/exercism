# Notes

_Scratchpad for session notes, user preferences, and working observations._

## User Preferences

- **Read Teach Skill first**: At the start of every session, after reading `NOTES.md`, always read `.agents/skills/teach/SKILL.md` and its format files (`LEARNING-RECORD-FORMAT.md`, `RESOURCES-FORMAT.md`, `MISSION-FORMAT.md`) before doing any teaching or updating any learning files.
- **Read workspace state files**: After reading the teach skill, also read `MISSION.md`, `GLOSSARY.md`, `RESOURCES.md`, and list the latest entries in `learning-records/` to understand current progress.
- **Submit on confirmation**: Once an exercise is complete and all refactoring is done, wait for the user's explicit confirmation before running `exercism submit`.
- **Refactor & Review First**: Before submitting to Exercism or staging/pushing to Git, always check if the code can be refactored to be cleaner, more concise, or more idiomatic.
- **Glossary Organization**: Categorize glossary terms into distinct subheadings: Syntax & Basics, Data Types & Collections, Control Flow, Logic & Operations, Strings & Formatting, and Programming Paradigms & Patterns.

## Explainer Template & Design System

- **Unified Style**: All exercise `explainer.html` files must follow the glassmorphic multi-card layout template saved at [explainer-template.html](file:///Users/alex/Exercism/go/explainer-template.html) (also copied under [.agents/skills/teach/explainer-template.html](file:///Users/alex/Exercism/go/.agents/skills/teach/explainer-template.html)).
- **Key Constraints**:
  - Keep main content inside a `.container` with `max-width: 840px`.
  - Use custom CSS variables in `:root` to preserve the exercise's color identity (e.g. green for Hello World, red for Blackjack, purple for Tech Palace) while keeping formatting and dimensions uniform.
  - Implement a header with a glowing backdrop blur effect.
  - Divide sections using individual `<section>` cards with hover transform effects.
  - Standardize subheadings (`h2`, `h3`), callouts, code blocks (prevent layout breaks via `minmax(0, 1fr)` columns), and self-check quiz containers.
  - Standardize the footer format: `Built with ❤️ for [Exercise Name]. Track: Go | Exercise: [Exercise Name]`.
- **Modern Go Features**: Always use modern Go idioms and features in explainers, regardless of what the exercise README or instructions show. Exercism's official docs can be outdated. Key examples:
  - Prefer `for _, v := range slice` over manual index loops where possible.
  - Use `for i := range n` (Go 1.22+) instead of `for i := 0; i < n; i++` for simple counters.
  - Always check the `go.mod` version in the exercise directory to confirm which features are available.
