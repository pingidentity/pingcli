# Pull Request Checklist

The following provides the steps to check/run to prepare for creating a PR to the `main` branch. PRs that follow these checklists will merge faster than PRs that do not.

*Note: This checklist is designed to support both human contributors and automated code review tools.*

## For Automated Code Review

This checklist includes specific verification criteria marked with *Verification* that can be programmatically checked to support both manual and automated review processes

## PR Planning & Structure

- [ ] **PR Scope**. To ensure maintainer reviews are as quick and efficient as possible, please separate support for different features into separate PRs. For example, support for a new `pingcli platform export` service can go in the same PR, however support for new command functionality should be separated. It's acceptable to merge related changes into the same PR where structural changes are being made.
  - *Verification*: Check that files modified are logically related (same command directory, related functionality)

- [ ] **PR Title**. To assist the maintainers in assessing PRs for priority, please provide a descriptive title of the functionality being supported. For example: `Add support for PingFederate export` or `Fix authentication flow for PingOne SSO`
  - *Verification*: Title should be descriptive and match the type of changes (Add/Update/Fix/Remove)

- [ ] **PR Description**. Please follow the provided PR description template and check relevant boxes. Include a clear description of:
  - What functionality is being added/changed
  - Why the change is needed (e.g., to fix an issue - include the issue number as reference)
  - Any breaking changes to CLI commands or configuration
  - *Verification*: Check that PR description template boxes are completed and description sections are filled

## Code Development

### Architecture & Design

- [ ] **Code implementation**. New code should follow the established CLI architecture patterns.
  - *Verification*: 
    - New commands are in `cmd/<command>/` directories
    - Command implementations follow the cobra command pattern
    - Internal functionality is organized in `internal/` packages
    - Connector implementations are in `internal/connector/<service>/`

- [ ] **SDK Usage**. All Ping Identity API interactions must use the appropriate Go SDKs rather than direct API calls
  - *Verification*: 
    - PingOne API calls use `github.com/patrickcping/pingone-go-sdk-v2/` packages or the `github.com/pingidentity/pingone-go-client/` packages
    - PingFederate API calls use `github.com/pingidentity/pingfederate-go-client/` packages
    - No direct HTTP calls to Ping Identity APIs (check for `http.Client`, `http.Get`, `http.Post`, etc.)

### Code Quality

- [ ] **Installation**. Ensure dependencies are properly maintained and the CLI installs successfully:

```shell
make install
```
*Verification*: Run command and verify exit code 0

- [ ] **Code Formatting**. Ensure code is properly formatted:

```shell
make fmt
```
*Verification*: Run command and verify no files are modified (clean git status)

- [ ] **Code Linting**. Run all linting checks to ensure code quality and consistency:

```shell
make vet
make importfmtlint
make golangcilint
```
*Verification*: Commands must exit with code 0

This includes:
- Go vet checks
- golangci-lint for comprehensive static analysis
- Import organization checks

## Testing

### Unit Tests

- [ ] **Unit Tests**. Where a code function performs work internally to a module, but has an external scope (i.e., a function with an initial capital letter `func MyFunction`), unit tests should ideally be created. Not all functions require a unit test, if in doubt please ask:

```shell
make test
```
*Verification*: Run command and verify exit code 0

### Integration Tests

- [ ] **Integration Tests**. Where new commands or connectors are being created, or existing functionality is being updated, integration tests should be created or modified according to the [acceptance test strategy](/contributing/acceptance-test-strategy.md)
  - *Verification*:
    - New commands have corresponding `*_test.go` files in same directory
    - Test files follow naming convention `Test<Command>_*` or `Test<Function>_*`
    - Tests include both success and error scenarios
    - Connector tests validate API integration functionality
    - Configuration dependencies are not hardcoded into the tests (but created as pre-requisites), to allow them to be run by any developer on their own tenant

- [ ] **Test Environment**. Ensure you have access to a PingOne trial or licensed environment for integration testing. The following environment variables must be set for full test execution:
  - `TEST_PINGONE_ENVIRONMENT_ID`
  - `TEST_PINGONE_REGION_CODE` 
  - `TEST_PINGONE_WORKER_CLIENT_ID`
  - `TEST_PINGONE_WORKER_CLIENT_SECRET`
  
  For PingFederate container tests:
  - `TEST_PING_IDENTITY_DEVOPS_USER`
  - `TEST_PING_IDENTITY_DEVOPS_KEY`
  - `TEST_PING_IDENTITY_ACCEPT_EULA=YES`

- [ ] **Container Tests**. If working with PingFederate functionality, ensure container-based integration tests work properly:

```shell
make spincontainer
# Run your tests
make removetestcontainer
```
*Verification*: Container starts healthy and tests execute successfully

## Documentation

### Code Documentation

- [ ] **Command Documentation**. Each cobra command should have appropriate help text and usage examples
  - *Verification*: 
    - All commands have `Short` and `Long` descriptions
    - Command flags have clear descriptions
    - Complex commands include usage examples in their help text

- [ ] **Custom Errors**. If required, implement appropriate custom error or warning messages for better user experience when API errors or validation errors occur. Include instruction on how the reader can address the root of the warning or error. Most API level errors do not need custom error handling.
  - *Verification*: Custom error functions include actionable guidance for users

- [ ] **Configuration Documentation**. Changes to configuration options should be documented appropriately
  - *Verification*: New configuration keys are documented in `/docs/tool-configuration/configuration-key.md`

### Examples

- [ ] **CLI Examples**. New or modified commands should have appropriate usage examples in documentation
  - *Verification*:
    - Examples exist in relevant documentation files under `docs/`
    - Command examples demonstrate both basic and advanced usage
    - Examples are tested and work with the current implementation

- [ ] **Plugin Examples**. If working with plugin functionality, ensure examples are updated in the `examples/plugin/` directory
  - *Verification*: Plugin examples compile successfully and demonstrate proper usage patterns

### Documentation Updates

- [ ] **README Updates**. If adding new functionality, ensure the main README.md is updated with relevant information
  - *Verification*: 
    - New commands are documented in the Commands section
    - Installation instructions remain accurate
    - Configuration examples reflect any new requirements

- [ ] **Documentation Review**. Ensure all documentation changes are clear, accurate, and follow the existing style
  - *Verification*: Documentation is well-structured and follows consistent formatting

## Security & Compliance

- [ ] **Security Scan**. Ensure your code passes security scanning (this will be automatically checked in CI, but you can run locally if gosec is installed)
  - *Verification*: No obvious security issues like hardcoded secrets, unsafe file operations, or command injection vectors

- [ ] **Sensitive Data**. Ensure no sensitive data (API keys, tokens, etc.) are committed to the repository
  - *Verification*: 
    - No API keys, passwords, or tokens in code or test files
    - Sensitive test data uses environment variables
    - Configuration files use appropriate masking for sensitive values
    - No `.env` files or similar containing credentials

- [ ] **Input Validation**. Implement appropriate input validation for all user-provided data
  - *Verification*: 
    - Command flags include appropriate validation
    - File paths are validated and sanitized
    - API inputs are validated before sending to services

## Final Checks

- [ ] **All Make Targets**. Run the comprehensive development check (excluding time-intensive tests):

```shell
make devchecknotest
```
*Verification*: Run command and verify exit code 0

- [ ] **CI Compatibility**. Verify your changes will pass automated CI checks by ensuring all the above steps pass locally
  - *Verification*: All previous verification steps completed successfully

- [ ] **Breaking Changes**. If your PR introduces breaking changes to CLI commands, configuration, or output formats, ensure they are:
  - Clearly documented in the PR description
  - Included in the changelog entry
  - Follow the project's versioning strategy
  - *Verification*: 
    - Breaking changes are documented in PR description
    - Migration guidance is provided for users

## Additional Notes

- The maintainers may run additional tests in different PingOne regions and with different Ping Identity service configurations
- Large PRs may take longer to review - consider breaking them into smaller, focused changes
- If you're unsure about any step, please ask questions in your PR or create an issue for discussion
- For CLI changes that affect user workflows, consider backwards compatibility and migration paths

---

## Documentation-Only Changes

If you are making documentation-only changes (guides, examples, or help text), you can use this simplified checklist:

- [ ] **Guide Updates**. New or updated guides should be clear, well-structured, and include practical CLI examples

- [ ] **Example Updates**. Ensure any CLI command examples are syntactically correct and demonstrate current functionality

- [ ] **Help Text Updates**. Verify that command help text is accurate and follows consistent formatting

- [ ] **Configuration Documentation**. If updating configuration docs, ensure examples match current CLI behavior

Documentation changes are generally merged quicker than code changes as there is less to review.
