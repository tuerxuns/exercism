# Notes

_Scratchpad for session notes, user preferences, and working observations._

## User Preferences

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
