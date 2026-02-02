# Contributing to Thrustlab GTM Skills

Thank you for your interest in contributing to Thrustlab! This guide will help you create high-quality Claude skills for GTM teams.

## Getting Started

1. **Fork the repository**
2. **Clone your fork**
   ```bash
   git clone https://github.com/your-username/thrustlab.git
   cd thrustlab
   ```
3. **Create a new branch**
   ```bash
   git checkout -b skill/your-skill-name
   ```

## Skill Development Guidelines

### Directory Structure

Each skill should have its own directory:

```
skills/
└── your-skill-name/
    ├── SKILL.md (required)
    ├── README.md (optional - for complex skills)
    └── templates/ (optional - supporting files)
```

### SKILL.md Format

Every skill must have a `SKILL.md` file with YAML frontmatter:

```yaml
---
name: your-skill-name
description: Clear description of what the skill does and when to use it
argument-hint: "expected arguments (if any)"
---

# Skill Instructions

Detailed instructions for Claude on how to execute this skill...
```

#### Frontmatter Fields

| Field | Required | Description |
|-------|----------|-------------|
| `name` | Yes | Lowercase, hyphens only, max 64 chars |
| `description` | Yes | Helps Claude decide when to use the skill |
| `argument-hint` | No | Hints for expected arguments |
| `disable-model-invocation` | No | Set to `true` for manual invocation only |
| `user-invocable` | No | Set to `false` to hide from user menus |
| `allowed-tools` | No | Restrict which tools Claude can use |

### Writing Effective Skills

#### 1. Clear Purpose
- Start with a clear explanation of what the skill does
- Explain when it should be used
- Define the expected outcomes

#### 2. Structured Instructions
```markdown
## Required Information
List what information Claude needs to gather from the user

## Execution Steps
1. Step-by-step instructions
2. Be specific and actionable
3. Include error handling guidance

## Output Format
Describe what the final output should look like
```

#### 3. Best Practices
- Use clear section headings (##, ###)
- Include examples where helpful
- Specify required MCP servers or tools
- Provide validation criteria
- Add helpful notes and tips

#### 4. GTM Focus
Since this is a GTM skills library, ensure your skill:
- Solves a real GTM problem
- Saves time on repetitive tasks
- Follows GTM best practices
- Works well with other skills in the library

### Testing Your Skill

Before submitting, test your skill:

1. **Install locally**
   ```bash
   cp -r skills/your-skill-name ~/.claude/skills/
   ```

2. **Test invocation**
   - Try manual invocation: `/your-skill-name`
   - Test with various arguments
   - Verify Claude follows all instructions

3. **Edge cases**
   - Test with missing information
   - Test with unusual inputs
   - Verify error handling

4. **Documentation**
   - Ensure all prerequisites are documented
   - Verify examples work as described
   - Check that argument hints are helpful

### Submission Process

1. **Create a pull request**
   - Clear title: "Add [skill-name] skill"
   - Describe what the skill does
   - List any prerequisites (MCP servers, etc.)
   - Include usage examples

2. **PR Description Template**
   ```markdown
   ## Skill Name
   [Your skill name]

   ## Purpose
   [What problem does this solve?]

   ## Prerequisites
   - [ ] Notion MCP Server (if applicable)
   - [ ] Other requirements

   ## Usage Example
   ```
   /your-skill-name example arguments
   ```

   ## Testing Checklist
   - [ ] Tested locally in Claude Desktop
   - [ ] Tested with various inputs
   - [ ] Documentation is complete
   - [ ] Examples work as described
   ```

3. **Review process**
   - Maintainers will review for quality and GTM relevance
   - Address any feedback
   - Once approved, your skill will be merged!

## Skill Ideas

Need inspiration? Here are some GTM skills that would be valuable:

### Research & Analysis
- Competitive analysis automation
- Market sizing calculator
- TAM/SAM/SOM analyzer
- Trend research aggregator

### Strategy & Planning
- Positioning statement generator
- Messaging framework builder
- Channel strategy planner
- Launch timeline creator

### Content & Enablement
- Battle card generator
- Sales deck outliner
- Case study template
- Product one-pager creator

### Customer Intelligence
- ICP (Ideal Customer Profile) builder
- Persona interview analyzer
- Win/loss analysis
- Customer journey mapper

### Execution & Tracking
- Launch checklist generator
- OKR tracker setup
- Campaign planner
- Metrics dashboard creator

## Code of Conduct

- Be respectful and constructive in discussions
- Focus on GTM use cases and practical value
- Help others learn and improve their skills
- Give credit where credit is due

## Questions?

- Open an issue for questions or discussions
- Tag maintainers for urgent matters
- Check existing skills for examples

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for helping make GTM operations more efficient for everyone! 🚀
