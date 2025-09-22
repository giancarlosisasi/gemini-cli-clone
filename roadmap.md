# Gemini CLI Clone - 1 Week Sprint

## Project Overview
Build a **minimal but functional** interactive chat CLI with Gemini API in 1 week (14-21 hours total, 2-3 hours daily).

**Focus**: Core chat functionality, basic TUI, structured output
**Target**: Working CLI that you can actually use and learn from

## Tech Stack (Minimal)

### Essential Dependencies
- **Cobra** - CLI framework
- **Viper + godotenv** - Config management
- **Google AI Go SDK** - Gemini API (https://github.com/googleapis/go-genai)
- **bubbletea** - Interactive TUI
- **lipgloss** - Basic styling
- **glamour** - Markdown rendering

### Optional (if time permits)
- **chroma** - Syntax highlighting
- **fatih/color** - Simple colors

## Daily Sprint Plan

### Day 1 (Monday): Foundation & Setup
**Time**: 2-3 hours | **Goal**: Working CLI skeleton with config

#### Tasks:
- [ ] Initialize Go project with basic structure
- [ ] Install core dependencies (Cobra, Viper, godotenv)
- [ ] Create `.env` and config system
- [ ] Basic Cobra commands: `root`, `version`, `chat`
- [ ] Test config loading works

#### Deliverable:
```bash
./gemini-cli version
./gemini-cli chat --help
```

---

### Day 2 (Tuesday): Gemini API Integration  
**Time**: 2-3 hours | **Goal**: API connection working

#### Tasks:
- [ ] Create `internal/gemini` package
- [ ] Integrate Google AI Go SDK
- [ ] API key validation and basic client
- [ ] Simple request/response test
- [ ] Error handling for API calls

#### Deliverable:
```bash
./gemini-cli chat --test  # Should connect and get response
```

---

### Day 3 (Wednesday): Basic Interactive Shell
**Time**: 2-3 hours | **Goal**: Simple TUI chat loop

#### Tasks:
- [ ] Basic bubbletea model for chat
- [ ] Input/output loop with message history
- [ ] Simple text rendering (no fancy formatting yet)
- [ ] Special commands: `/exit`, `/clear`, `/help`
- [ ] Handle Ctrl+C gracefully

#### Deliverable:
```bash
./gemini-cli chat  # Interactive shell that works!
```

---

### Day 4 (Thursday): Output Formatting
**Time**: 2-3 hours | **Goal**: Pretty output with basic styling

#### Tasks:
- [ ] Integrate glamour for markdown rendering
- [ ] Add lipgloss styling (colors, borders)
- [ ] Implement `--output` flag (plain, pretty, json)
- [ ] Basic message formatting and layout
- [ ] Loading spinner during API calls

#### Deliverable:
Beautiful formatted responses with markdown support

---

### Day 5 (Friday): Polish & Basic Features
**Time**: 2-3 hours | **Goal**: Production-ready basic version

#### Tasks:
- [ ] Add conversation context (message history)
- [ ] Better error messages and handling
- [ ] Configuration validation
- [ ] Basic session management (in-memory)
- [ ] Help documentation and examples

#### Deliverable:
Fully functional chat CLI ready for daily use

---

### Weekend (Optional): Extras & Testing
**Time**: 1-2 hours | **Goal**: Nice-to-haves

#### Optional Tasks:
- [ ] Add syntax highlighting with chroma
- [ ] Session export (JSON/markdown)
- [ ] Windows Terminal optimizations
- [ ] Basic tests for core functions
- [ ] README with setup instructions

---

## Simplified Project Structure

```
gemini-cli/
├── cmd/
│   └── root.go          # Cobra setup + chat command
├── internal/
│   ├── config/          # Config loading (viper + godotenv)
│   ├── gemini/          # API client wrapper
│   └── ui/              # Bubbletea models
├── .env.example
├── config.yaml.example
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Essential Features Only

### Must Have (Days 1-5):
- ✅ Interactive chat with Gemini API
- ✅ Config management (.env + yaml)
- ✅ Basic TUI with bubbletea
- ✅ Markdown rendering
- ✅ Multiple output formats
- ✅ Error handling

### Nice to Have (Weekend):
- 🎯 Syntax highlighting
- 🎯 Session persistence
- 🎯 Advanced styling
- 🎯 Tests

### Cut for Time:
- ❌ Advanced session management
- ❌ Plugin system
- ❌ Cross-platform builds
- ❌ Comprehensive testing
- ❌ Documentation beyond basic README

## Daily Time Breakdown

### 2 Hour Sessions:
- **30 mins**: Setup/review previous day
- **90 mins**: Core development
- **10 mins**: Test and commit

### 3 Hour Sessions:
- **30 mins**: Setup/review
- **2 hours**: Core development + testing
- **30 mins**: Documentation/cleanup

## Success Metrics (End of Week)

**Minimum Viable Product**:
- [ ] Can chat with Gemini API interactively
- [ ] Pretty formatted responses
- [ ] Basic error handling
- [ ] Configurable via .env file
- [ ] Works on Windows

**Stretch Goals**:
- [ ] Syntax highlighting in code blocks
- [ ] Session export functionality
- [ ] Polished UI with good styling
- [ ] Basic test coverage

## Quick Start Commands

### Day 1 Setup:
```bash
mkdir gemini-cli && cd gemini-cli
go mod init gemini-cli
go get github.com/spf13/cobra@latest
go get github.com/spf13/viper@latest
go get github.com/joho/godotenv@latest
echo "GEMINI_API_KEY=your_key_here" > .env
```

### Development Commands:
```bash
go run main.go version
go run main.go chat
go build -o gemini-cli.exe  # Windows
```

## Risk Mitigation (1 Week)

**If Behind Schedule**:
- **Day 3**: Skip fancy TUI, use simple CLI input/output
- **Day 4**: Use basic color instead of glamour
- **Day 5**: Skip session management, focus on stability

**Common Time Sinks to Avoid**:
- Perfect project structure (keep it simple)
- Over-engineering config system
- Complex error handling (basic is fine)
- Advanced TUI features
- Comprehensive testing

## Learning Focus

This 1-week sprint teaches you:
- **Cobra CLI patterns** (Day 1)
- **API integration in Go** (Day 2)  
- **Basic TUI with bubbletea** (Day 3)
- **Terminal formatting** (Day 4)
- **CLI polish and UX** (Day 5)

**Result**: Working knowledge of Go CLI development + a useful tool you built yourself!

Ready to start? Let me know if you want me to create the Day 1 starter code! 🚀

## Epic 1: Foundation & Project Setup

### Story 1.1: Project Structure & Dependencies  
- [ ] Initialize Go module: `go mod init gemini-cli-clone`
- [ ] Create directory structure:
  ```
  gemini-cli-clone/
  ├── cmd/            # CLI commands
  ├── internal/       # Private application code
  │   ├── config/     # Configuration management
  │   ├── gemini/     # API client
  │   ├── ui/         # TUI components
  │   └── chat/       # Chat logic
  ├── pkg/            # Public packages
  ├── .env.example    # Environment template
  └── configs/        # Config file templates
  ```
- [ ] Install core dependencies with `go get`
- [ ] Set up `.gitignore` for Go + Windows development
- [ ] Create basic `Makefile` with Windows batch alternatives

### Story 1.2: Configuration System
- [ ] Create `internal/config` package
- [ ] Implement `.env` file loading with `godotenv`
- [ ] Set up Viper for config file management (YAML/JSON)
- [ ] Support config hierarchy: `.env` → `config.yaml` → CLI flags
- [ ] Add config validation and default values
- [ ] Create `config init` command to generate templates

### Story 1.3: Basic CLI Structure
- [ ] Set up Cobra root command
- [ ] Implement global flags: `--config`, `--output`, `--debug`
- [ ] Add version command with build info
- [ ] Create help system with examples
- [ ] Add output format flag: `--output=plain|json|pretty`

**Windows Focus**: Ensure paths work with Windows backslashes, test in PowerShell and CMD
**Acceptance Criteria**: CLI loads config from .env + config files, responds to basic commands

## Epic 2: Gemini API Integration

### Story 2.1: API Client Setup
- [ ] Create `internal/gemini` package
- [ ] Integrate Google AI Go SDK
- [ ] Implement API key validation and testing
- [ ] Create client wrapper with error handling
- [ ] Add request/response logging for debugging
- [ ] Implement connection health check

### Story 2.2: Chat Session Management  
- [ ] Create `internal/chat` package for session logic
- [ ] Implement conversation context management
- [ ] Add message history storage (in-memory for now)
- [ ] Support conversation reset/clear functionality
- [ ] Add token counting and limits awareness
- [ ] Implement streaming response handling

### Story 2.3: Error Handling & Resilience
- [ ] Create custom error types for API issues
- [ ] Implement retry logic with exponential backoff
- [ ] Add rate limiting awareness
- [ ] Handle network connectivity issues gracefully
- [ ] Create user-friendly error messages
- [ ] Add debug mode with detailed API logs

**Windows Focus**: Test API calls work correctly from Windows (firewall, proxy considerations)
**Acceptance Criteria**: Successfully connect to Gemini API and handle basic chat requests

## Epic 3: Interactive Shell & TUI

### Story 3.1: Basic Interactive Shell
- [ ] Create `chat` command using bubbletea
- [ ] Implement basic input/output loop
- [ ] Add welcome message and help text  
- [ ] Support multi-line input (Ctrl+Enter or similar)
- [ ] Add special commands: `/help`, `/exit`, `/clear`, `/debug`
- [ ] Handle graceful shutdown (Ctrl+C)

### Story 3.2: Enhanced TUI Experience
- [ ] Create `internal/ui` package with bubbletea models
- [ ] Add loading spinner during API requests
- [ ] Implement message history display with scrolling
- [ ] Add input validation and visual feedback
- [ ] Support copy/paste functionality
- [ ] Create responsive layout that adapts to terminal size

### Story 3.3: Output Formatting & Styling
- [ ] Integrate glamour for markdown rendering
- [ ] Add chroma for syntax highlighting in code blocks
- [ ] Use lipgloss for consistent styling and themes
- [ ] Implement different output modes (plain/rich/json)
- [ ] Add color customization options
- [ ] Support Windows Terminal features (if available)

**Windows Focus**: 
- Test in CMD, PowerShell, Windows Terminal
- Ensure Unicode/emoji support works correctly
- Handle Windows-specific terminal limitations gracefully

**Acceptance Criteria**: Smooth interactive chat experience with proper formatting and visual feedback

## Epic 4: Advanced Chat Features

### Story 4.1: Session Management
- [ ] Add conversation persistence (save/load sessions)
- [ ] Implement named sessions: `/session save <name>`, `/session load <name>`
- [ ] Add session listing and management commands
- [ ] Support session export (JSON, markdown, plain text)
- [ ] Add automatic session backup functionality
- [ ] Implement session search and filtering

### Story 4.2: Output Customization & Export
- [ ] Add `--output` flag support: `plain`, `json`, `pretty`, `markdown`
- [ ] Implement structured JSON output for programmatic use
- [ ] Add export functionality for conversations
- [ ] Support custom output templates
- [ ] Add statistics display (tokens used, response time, etc.)
- [ ] Implement quiet mode for scripting

### Story 4.3: Developer Experience Features
- [ ] Add comprehensive debug mode with API request/response logging
- [ ] Implement configuration validation and diagnostics
- [ ] Add performance metrics and timing information
- [ ] Support verbose logging levels
- [ ] Add health check command for troubleshooting
- [ ] Implement plugin system foundation (for future extensions)

**Windows Focus**:
- Ensure file paths work correctly on Windows
- Test session files in Windows-specific locations (%APPDATA%)
- Handle Windows file locking and permissions

**Acceptance Criteria**: Advanced session management with multiple output formats and developer-friendly features

## Epic 5: Polish & Distribution

### Story 5.1: Testing & Quality Assurance
- [ ] Add unit tests for core packages (config, gemini, chat)
- [ ] Implement integration tests with mock API responses
- [ ] Add TUI component testing with bubbletea test helpers
- [ ] Create end-to-end testing scenarios
- [ ] Set up test coverage reporting and CI/CD
- [ ] Add performance benchmarks for API calls

### Story 5.2: Documentation & User Experience
- [ ] Create comprehensive README with setup instructions
- [ ] Add usage examples and common workflows
- [ ] Document all CLI commands and flags
- [ ] Create troubleshooting guide for Windows users
- [ ] Add contribution guidelines and development setup
- [ ] Record demo GIFs/videos for key features

### Story 5.3: Build & Release
- [ ] Set up goreleaser for cross-platform builds
- [ ] Create Windows installer/package (MSI or Chocolatey)
- [ ] Add auto-update functionality (optional)
- [ ] Set up release automation with GitHub Actions
- [ ] Create installation scripts for different platforms
- [ ] Add security scanning and vulnerability checks

**Windows-Specific Tasks**:
- Test on different Windows versions (10, 11)
- Ensure compatibility with Windows Defender
- Test with different terminal emulators
- Create Windows-specific installation guide

**Acceptance Criteria**: Production-ready CLI with comprehensive testing, documentation, and distribution

## Development Phases

### Phase 1 (Week 1): Foundation & Configuration
**Focus**: Get basic CLI structure working with proper config management
- Complete Epic 1: Project setup, dependencies, config system
- **Deliverable**: CLI that loads config from .env and config files
- **Windows Testing**: Verify paths and config loading work in Windows

### Phase 2 (Week 2): API Integration  
**Focus**: Connect to Gemini API and handle basic requests
- Complete Epic 2: API client, error handling, basic chat logic
- **Deliverable**: CLI can make successful API calls to Gemini
- **Windows Testing**: Ensure API calls work from Windows environment

### Phase 3 (Weeks 3-4): Interactive Shell
**Focus**: Build the core interactive experience
- Complete Epic 3: TUI with bubbletea, formatting, syntax highlighting
- **Deliverable**: Functional interactive chat with rich formatting
- **Windows Testing**: Test in CMD, PowerShell, Windows Terminal

### Phase 4 (Week 5): Advanced Features
**Focus**: Session management and output customization
- Complete Epic 4: Sessions, export functionality, developer features  
- **Deliverable**: Feature-complete chat CLI with advanced options
- **Windows Testing**: File operations and session persistence

### Phase 5 (Week 6): Polish & Release
**Focus**: Testing, documentation, and distribution
- Complete Epic 5: Testing, docs, builds, Windows-specific packaging
- **Deliverable**: Production-ready release with Windows installer

## Learning Objectives (CLI/TUI Development)

### Week 1-2: CLI Fundamentals
- Learn Cobra framework patterns and best practices
- Understand Viper configuration management
- Master Go project structure for CLI tools
- Learn environment variable and config file handling

### Week 3-4: TUI Development  
- Master bubbletea framework and The Elm Architecture
- Learn terminal UI component composition
- Understand terminal capabilities and limitations
- Learn styling and theming with lipgloss

### Week 5-6: Advanced CLI Concepts
- Learn cross-platform build and distribution
- Understand testing strategies for CLI/TUI apps
- Master structured output and formatting
- Learn Windows-specific CLI considerations

## Success Metrics

### Technical Milestones
- [ ] Successfully authenticate and chat with Gemini API
- [ ] Smooth interactive TUI experience with proper error handling
- [ ] Multi-format output support (plain, JSON, pretty)
- [ ] Session persistence and management
- [ ] Windows Terminal integration with rich formatting
- [ ] Cross-platform builds (Windows primary, Linux/macOS secondary)

### Learning Outcomes
- [ ] Comfortable with Cobra and Viper for CLI development
- [ ] Proficient in bubbletea for terminal UI development
- [ ] Understanding of Go project structure for CLI tools
- [ ] Knowledge of cross-platform considerations
- [ ] Experience with CLI testing strategies

## Windows Development Tips

### Terminal Considerations
- **Windows Terminal**: Modern terminal with full Unicode support
- **PowerShell**: Default shell, good color support
- **CMD**: Limited features, fallback compatibility
- **Git Bash**: Unix-like environment on Windows

### Development Setup
- Use **air** for live reloading during development
- Configure **golangci-lint** with Windows paths
- Test with Windows Defender real-time protection enabled
- Use **goreleaser** for cross-platform builds

### Windows-Specific Features
- Handle backslash path separators correctly
- Support Windows Terminal features (hyperlinks, etc.)
- Consider Windows file permissions and user directories
- Test with different Windows versions and terminals

## Getting Started

### Prerequisites
```bash
# Required tools
go version # 1.21 or higher
git --version
```

### Quick Start
```bash
# 1. Initialize project
mkdir gemini-cli-clone && cd gemini-cli-clone
go mod init gemini-cli-clone

# 2. Install dependencies
go get github.com/spf13/cobra@latest
go get github.com/spf13/viper@latest
go get github.com/joho/godotenv@latest

# 3. Create .env file
echo "GEMINI_API_KEY=your_api_key_here" > .env

# 4. Start with basic CLI structure
```

This focused roadmap will give you hands-on experience with modern Go CLI/TUI development while building something immediately useful!