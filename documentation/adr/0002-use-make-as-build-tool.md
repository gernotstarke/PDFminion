# 2. Use-make-as-build-tool

## Status
Accepted

## Date
2024-11-18

## Context
We needed to choose a build tool for our Go application PDFminion that would handle:
- Cross-platform compilation
- Multiple build targets
- Release packaging
- Installation/uninstallation
- Development workflows

The main alternatives considered were:
- Make
- Shell scripts
- Go's built-in build commands

## Decision
We decided to use Make as our primary build tool.

## Reasons

### Pros
1. **Ubiquity**
- Make is installed by default on most Unix-like systems
- Well-understood by most developers
- Extensive documentation available
- Long history of reliability

2. **Platform Independence**
- Works on all major platforms (Linux, macOS, Windows via WSL)
- Consistent behavior across environments

3. **Dependency Management**
- Built-in dependency tracking
- Efficient rebuilds (only rebuilds what's necessary)
- Clear visualization of build dependencies

4. **Simplicity**
- Declarative syntax
- No additional dependencies needed
- Easy to maintain and modify
- Self-documenting through target names

5. **Flexibility**
- Can execute any shell command
- Easy to add new targets
- Supports both simple and complex build processes
- Can integrate with other tools seamlessly


### Cons
1. **Windows Compatibility**
- Requires WSL or MinGW on Windows
- May create friction for Windows developers

2. **Syntax**
- Tab-based syntax can be error-prone
- Learning curve for complex features
- Limited string manipulation capabilities

3. **Error Handling**
- Basic error handling capabilities
- Can be verbose for complex error scenarios

4. **Debugging**
- Limited built-in debugging facilities
- Can be hard to troubleshoot complex makefiles

## Consequences

### Positive
1. **Development Workflow**
- Simple commands like `make mac` or `make release`
- Easy to remember and use
- Quick to execute
- Consistent across team members

2. **Maintenance**
- Single file (Makefile) contains all build logic
- Easy to add new targets and modify existing ones
- Version control friendly
- Self-documenting

3. **Integration**
- Easy integration with CI/CD pipelines
- Works well with existing Go tools
- Can be extended with shell scripts if needed

### Negative
1. **Team Requirements**
- Team members need basic Make knowledge
- Windows developers need additional setup
- May need documentation for complex targets

2. **Scaling**
- Complex build processes may become hard to maintain
- Limited modularity compared to modern build tools
- May need to supplement with scripts for complex tasks

## Notes
- Maintain clear documentation of available make targets

