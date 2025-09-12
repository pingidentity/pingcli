# Acceptance Testing Strategy

This document outlines the comprehensive testing strategy for the PingCLI, based on analysis of existing test patterns and CLI testing best practices. All new commands, connectors, and functionality should follow these testing patterns to ensure consistency, reliability, and maintainability.

## Overview

The CLI uses Go's built-in testing framework with acceptance tests that interact with real Ping Identity environments (PingOne and PingFederate). Tests are organized by command/functionality area and follow consistent naming and structure patterns. The testing strategy encompasses unit tests for command logic, integration tests for API interactions, and end-to-end tests for complete CLI workflows.

## Test Organization

### File Structure
- **Command tests**: `<command>_test.go` files co-located with command implementations
- **Connector/Integration tests**: Located in `internal/connector/<service>/` directories
- **Test utilities**: Located in `internal/testing/` with specialized helpers:
  - `testutils/` - General test utilities and API client setup
  - `testutils_cobra/` - Cobra command testing helpers
  - `testutils_koanf/` - Configuration testing helpers
  - `testutils_resource/` - Resource-specific test data generators
  - `testutils_terraform/` - Terraform plan validation helpers

### Package Structure
- Tests are placed in `<package>_test` packages (e.g., `platform_test`, `config_test`)
- Import test helpers from `internal/testing/testutils*` packages
- Use shared client configurations from `testutils.GetClientInfo()`
- Container-based testing for PingFederate integration via Docker

## Core Testing Patterns

### 1. Standard CLI Command Tests

Every CLI command should implement these core test functions:

#### **Basic Execution Tests**
Tests that commands execute successfully with valid parameters:
```go
func Test<Command>Cmd_Execute(t *testing.T) {
    // Test command executes without error with minimal valid parameters
    // Use testutils_cobra.ExecutePingcli() for command execution
    // Use testutils.CheckExpectedError() to validate no error occurred
}
```

#### **Argument Validation Tests**
Tests command argument validation and error handling:
```go
func Test<Command>Cmd_TooManyArgs(t *testing.T) {
    // Test command rejects invalid argument counts
    expectedErrorPattern := `^failed to execute 'pingcli <command>': command accepts X arg\(s\), received Y$`
    err := testutils_cobra.ExecutePingcli(t, "<command>", "extra-arg")
    testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

func Test<Command>Cmd_InvalidFlag(t *testing.T) {
    // Test command rejects invalid flags
    expectedErrorPattern := `^unknown flag: --invalid$`
    err := testutils_cobra.ExecutePingcli(t, "<command>", "--invalid")
    testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
```

#### **Flag Validation Tests**
Tests all supported flags with valid and invalid values:
```go
func Test<Command>Cmd_<FlagName>Flag(t *testing.T) {
    // Test each flag with valid values
    // Test flag combinations and interdependencies
    // Test required flag groups (marked with cobra's MarkFlagsRequiredTogether)
}

func Test<Command>Cmd_<FlagName>FlagInvalid(t *testing.T) {
    // Test flags with invalid values
    // Validate appropriate error messages
}
```

### 2. Integration and API Testing Strategy

#### **Authentication Testing**
- Test all supported authentication methods for each service
- Validate proper error handling for invalid credentials
- Test credential flag combinations and requirements

```go
func Test<Command>Cmd_PingOneWorkerFlags(t *testing.T) {
    // Test PingOne worker authentication with valid credentials
    // Use environment variables for actual credentials
    err := testutils_cobra.ExecutePingcli(t, "<command>",
        "--pingone-worker-environment-id", os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
        "--pingone-worker-client-id", os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
        "--pingone-worker-client-secret", os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
        "--pingone-region-code", os.Getenv("TEST_PINGONE_REGION_CODE"))
    testutils.CheckExpectedError(t, err, nil)
}

func Test<Command>Cmd_PingFederateBasicAuthFlags(t *testing.T) {
    // Test PingFederate basic authentication
    // Use container-based testing for PingFederate
}
```

#### **Export/Import Testing**
For commands that export or import configurations:
- Test minimal exports (specific services/resources)
- Test comprehensive exports (all available data)
- Test different output formats (HCL, Terraform)
- Validate exported data completeness and accuracy

```go
func Test<Command>Cmd_ExportFormats(t *testing.T) {
    // Test each supported export format
    // Validate output file structure and content
    // Test format-specific features and limitations
}
```

#### **Container Integration Testing**
For PingFederate-related functionality:
- Use Docker containers for consistent test environments
- Test against known container configurations
- Validate SSL/TLS certificate handling

```go
// Container tests should be marked for container execution
func Test<Command>Cmd_ContainerIntegration(t *testing.T) {
    // This test requires the PingFederate container to be running
    // Use testutils.GetClientInfo() to get configured PingFederate client
}
```

### 3. Test Environment Setup and Validation

The CLI testing framework uses environment variables and helper functions to validate test requirements before execution.

#### **Essential Environment Variables**

**PingOne Integration Testing**
Required for all PingOne-related tests:
- `TEST_PINGONE_ENVIRONMENT_ID` - PingOne environment for testing
- `TEST_PINGONE_REGION_CODE` - PingOne region (e.g., 'NA', 'EU', 'AP')
- `TEST_PINGONE_WORKER_CLIENT_ID` - Worker application client ID with admin roles
- `TEST_PINGONE_WORKER_CLIENT_SECRET` - Worker application client secret

**PingFederate Container Testing**
Required for PingFederate integration tests:
- `TEST_PING_IDENTITY_DEVOPS_USER` - DevOps program username
- `TEST_PING_IDENTITY_DEVOPS_KEY` - DevOps program access key
- `TEST_PING_IDENTITY_ACCEPT_EULA` - Must be set to "YES" to accept EULA

#### **Test Setup Patterns**

**Standard CLI Test Setup**
```go
func Test<Command>Cmd_<Scenario>(t *testing.T) {
    // Initialize configuration system for testing
    testutils_koanf.InitKoanfs(t)
    
    // Create temporary directories as needed
    outputDir := t.TempDir()
    
    // Execute CLI command with test parameters
    err := testutils_cobra.ExecutePingcli(t, "command", "subcommand", "--flag", "value")
    testutils.CheckExpectedError(t, err, nil)
}
```

**Integration Test Setup**
```go
func Test<Command>Cmd_Integration(t *testing.T) {
    // Get configured API clients
    clientInfo := testutils.GetClientInfo(t)
    
    // Use actual API clients for integration testing
    // clientInfo.PingOneApiClient, clientInfo.PingFederateApiClient
}
```

#### **Container Environment Management**

**PingFederate Container Lifecycle**
The CLI includes Make targets for managing PingFederate containers:

```bash
# Start PingFederate container for testing
make spincontainer

# Run tests requiring PingFederate
make test

# Clean up container after testing
make removetestcontainer
```

**Container Test Requirements**
- Docker must be running and accessible
- DevOps program credentials must be configured
- Container health checks must pass before tests execute
- Tests should clean up any configuration changes made during testing

**Container Test Patterns**
```go
func TestPingFederateIntegration(t *testing.T) {
    // Container tests should validate container availability
    clientInfo := testutils.GetClientInfo(t)
    if clientInfo.PingFederateApiClient == nil {
        t.Skip("PingFederate container not available")
    }
    
    // Use configured client for testing
    // Test should be idempotent and not affect other tests
}
```

#### **CLI-Specific Test Considerations**

**Configuration Management**
- Use `testutils_koanf.InitKoanfs(t)` to initialize the configuration system
- Test both command-line flags and configuration file scenarios
- Validate configuration precedence (flags override config files)

**Output Format Testing**
- Test different output formats (JSON, table, quiet modes)
- Validate machine-readable output for automation scenarios  
- Test human-readable output formatting

**Error Message Testing**
- Validate specific error message patterns using regex
- Test error scenarios for invalid inputs, network issues, authentication failures
- Ensure error messages provide actionable guidance to users

### 4. Export and Data Validation Testing

#### **Export Functionality Testing**
For platform export and similar data extraction commands:

**Export Format Validation**
```go
func Test<Command>Cmd_ExportFormat<Format>(t *testing.T) {
    // Test each supported export format (HCL, Terraform, etc.)
    // Validate format-specific output structure
    // Test format conversion accuracy
}
```

**Export Completeness Testing**
```go
func TestExportCompleteness(t *testing.T) {
    // Use testutils.ValidateImportBlocks() to verify exported data
    // Test that all expected resources are exported
    // Validate resource relationships and dependencies
}
```

**Service-Specific Export Testing**
```go
func Test<Command>Cmd_Service<ServiceName>(t *testing.T) {
    // Test export of specific services (PingOne SSO, MFA, Protect, etc.)
    // Validate service-specific resource types and configurations
    // Test service filtering and selection
}
```

### 5. Configuration and Profile Testing

#### **Profile Management Testing**
For config-related commands that manage profiles and settings:

**Profile CRUD Operations**
```go
func TestConfigProfile_<Operation>(t *testing.T) {
    // Test profile creation, reading, updating, deletion
    // Validate profile persistence and retrieval
    // Test profile switching and activation
}
```

**Configuration Key Management**
```go
func TestConfigKey_<Operation>(t *testing.T) {
    // Test setting and getting configuration values
    // Validate configuration validation and type checking
    // Test configuration inheritance and precedence
}
```

**Authentication Flow Testing**
```go
func TestAuth_<Flow>(t *testing.T) {
    // Test login/logout flows
    // Validate token persistence and refresh
    // Test authentication error handling
}
```

### 6. Request and Custom API Testing

#### **Custom Request Testing**
For request command functionality that makes custom API calls:

**HTTP Method Testing**
```go
func TestRequest_<Method>(t *testing.T) {
    // Test GET, POST, PUT, DELETE, PATCH operations
    // Validate request formatting and response handling
    // Test authentication integration with custom requests
}
```

**API Compatibility Testing**
```go
func TestRequest_APICompatibility(t *testing.T) {
    // Test against known API endpoints
    // Validate response parsing and error handling
    // Test API versioning and compatibility matrices
}
```

#### **Plugin System Testing**
For plugin management functionality:

**Plugin Lifecycle Testing**
```go
func TestPlugin_<Operation>(t *testing.T) {
    // Test plugin add, list, remove operations
    // Validate plugin discovery and loading
    // Test plugin execution and integration
}
```

## Advanced Testing Patterns

### 1. CLI-Specific Testing Considerations

#### **Command Chaining and Workflow Testing**
Test complex workflows that involve multiple CLI commands:
```go
func TestWorkflow_<Scenario>(t *testing.T) {
    // Test sequences like: login -> export -> configure -> logout
    // Validate state persistence between commands
    // Test failure recovery and cleanup
}
```

#### **Interactive vs Non-Interactive Testing**
```go
func TestInteractive_<Command>(t *testing.T) {
    // Test commands with interactive prompts
    // Use testutils.WriteStringToPipe() for input simulation
    // Validate prompt text and response handling
}
```

#### **Cross-Platform Compatibility**
- Test path handling differences (Windows vs Unix)
- Validate file permission handling
- Test shell integration and autocompletion

### 2. Performance and Scale Testing

#### **Large Data Set Testing**
```go
func Test<Command>_LargeDataSet(t *testing.T) {
    // Test commands with large numbers of resources
    // Validate memory usage and processing time
    // Test pagination and batching logic
}
```

#### **Concurrent Execution Testing**
```go
func Test<Command>_Concurrent(t *testing.T) {
    // Test multiple CLI instances running simultaneously
    // Validate file locking and state management
    // Test API rate limiting and retry logic
}
```

### 3. Error Handling and Recovery Testing

#### **Network Error Simulation**
```go
func Test<Command>_NetworkErrors(t *testing.T) {
    // Test behavior with network timeouts
    // Test API unavailability scenarios
    // Validate retry logic and exponential backoff
}
```

#### **Authentication Error Testing**
```go
func Test<Command>_AuthenticationErrors(t *testing.T) {
    // Test expired tokens and credentials
    // Test insufficient permissions
    // Validate error message clarity and actionability
}
```

#### **File System Error Testing**
```go
func Test<Command>_FileSystemErrors(t *testing.T) {
    // Test disk space issues
    // Test permission denied scenarios
    // Test file corruption and recovery
}
```

### 4. Security and Credential Testing

#### **Credential Storage Testing**
```go
func TestCredentialStorage_<Scenario>(t *testing.T) {
    // Test secure storage of tokens and credentials
    // Validate encryption and access controls
    // Test credential cleanup and expiration
}
```

#### **Sensitive Data Handling**
```go
func TestSensitiveData_<Command>(t *testing.T) {
    // Test that sensitive data is not logged or exposed
    // Validate redaction in error messages and output
    // Test secure temporary file handling
}
```

### 5. CLI Feature Deprecation and Backward Compatibility

When deprecating CLI flags, commands, or changing behavior:

#### **Deprecation Strategy**
- **Simultaneous Support**: Both deprecated and new functionality must work during the deprecation period
- **Gradual Migration**: Users should be able to migrate incrementally without breaking changes
- **Clear Warnings**: Deprecated flags/commands should generate helpful deprecation warnings
- **Documentation**: Update both help text and user documentation

#### **Required Deprecation Tests**
```go
func Test<Command>_Deprecation_<Flag>(t *testing.T) {
    // Test deprecated flag works (legacy usage)
    // Test new flag works (current usage)
    // Test both flags work together (migration period)
    // Test migration path from deprecated to new flag
    // Validate deprecation warnings are generated
}
```

#### **Backward Compatibility Validation**
- Existing user scripts and workflows must continue to work unchanged
- No functional regression during deprecation period
- Smooth migration path from old to new CLI interfaces
- Proper handling of edge cases during transition

### 6. Multi-Service Integration Testing

For commands that work across multiple Ping Identity services:
- Test service interoperability and data consistency
- Test partial service availability scenarios
- Validate cross-service authentication and authorization
- Test service-specific configuration and feature differences

### 7. Configuration Migration and Upgrade Testing

#### **Configuration File Format Regression Testing**
When modifying configuration file formats, schema, or default values:

**Legacy Configuration Compatibility**
```go
func TestConfig_LegacyFormatCompatibility(t *testing.T) {
    // Test that existing configuration files continue to work
    // Load pre-upgrade configuration samples
    // Validate that all settings are properly migrated
    // Test that no user data is lost during upgrade
}
```

**Configuration Schema Evolution**
```go
func TestConfig_SchemaEvolution_v<OldVersion>_to_v<NewVersion>(t *testing.T) {
    // Test specific schema version transitions
    // Validate field renames, type changes, and removals
    // Test default value handling for new fields
    // Verify backward compatibility warnings are shown
}
```

**Configuration File Migration**
```go
func TestConfig_Migration_<Scenario>(t *testing.T) {
    // Create legacy configuration file
    legacyConfig := `
    {
        "old_field": "value",
        "deprecated_setting": true
    }`
    
    configFile := testutils.WriteStringToPipe(t, legacyConfig)
    defer configFile.Close()
    
    // Test that CLI can load and migrate the configuration
    err := testutils_cobra.ExecutePingcli(t, "config", "migrate", "--config-file", configFile.Name())
    testutils.CheckExpectedError(t, err, nil)
    
    // Validate migrated configuration
    // Check that new format is used
    // Verify all data was preserved
}
```

## Configuration Regression Testing Strategy

### Configuration File Format Testing

#### **Version Compatibility Matrix Testing**
Create comprehensive tests for configuration file compatibility across versions:

```go
func TestConfigCompatibility_Matrix(t *testing.T) {
    testCases := []struct {
        version     string
        configFile  string
        expectError bool
        description string
    }{
        {
            version:     "v1.0.0",
            configFile:  "testdata/configs/v1.0.0-sample.json",
            expectError: false,
            description: "Legacy v1.0.0 format should still work",
        },
        {
            version:     "v1.5.0", 
            configFile:  "testdata/configs/v1.5.0-with-new-fields.yaml",
            expectError: false,
            description: "v1.5.0 format with new fields",
        },
        {
            version:     "v2.0.0",
            configFile:  "testdata/configs/v2.0.0-breaking-changes.yaml",
            expectError: false,
            description: "v2.0.0 format after breaking changes",
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.description, func(t *testing.T) {
            err := testutils_cobra.ExecutePingcli(t, "config", "validate", "--config-file", tc.configFile)
            if tc.expectError {
                testutils.CheckExpectedError(t, err, &expectedErrorPattern)
            } else {
                testutils.CheckExpectedError(t, err, nil)
            }
        })
    }
}
```

#### **Configuration Migration Testing**
Test automatic migration of configuration files during upgrades:

```go
func TestConfigMigration_AutoUpgrade(t *testing.T) {
    // Create a temporary directory for config testing
    configDir := t.TempDir()
    
    // Create legacy configuration file
    legacyConfigPath := filepath.Join(configDir, "config.yaml")
    legacyConfig := `
version: "1.0"
profiles:
  default:
    pingone:
      environment_id: "old-format-env-id"  
      client_credentials:
        client_id: "test-client"
        client_secret: "test-secret"
      region: "NA"`
    
    err := os.WriteFile(legacyConfigPath, []byte(legacyConfig), 0644)
    require.NoError(t, err)
    
    // Test that CLI automatically migrates the config on first use
    err = testutils_cobra.ExecutePingcli(t, "config", "get", 
        "--config-file", legacyConfigPath,
        "profiles.default.pingone.environment_id")
    testutils.CheckExpectedError(t, err, nil)
    
    // Read the config file back and verify it was migrated
    migratedConfig, err := os.ReadFile(legacyConfigPath)
    require.NoError(t, err)
    
    // Verify new format is used
    assert.Contains(t, string(migratedConfig), `version: "2.0"`)
    assert.Contains(t, string(migratedConfig), `worker_environment_id:`) // New field name
}
```

#### **Configuration Backup and Recovery Testing**
```go
func TestConfigUpgrade_BackupAndRecovery(t *testing.T) {
    configDir := t.TempDir()
    originalConfigPath := filepath.Join(configDir, "config.yaml")
    backupConfigPath := filepath.Join(configDir, "config.yaml.backup")
    
    // Create original configuration
    originalConfig := generateLegacyConfig("1.5.0")
    err := os.WriteFile(originalConfigPath, []byte(originalConfig), 0644)
    require.NoError(t, err)
    
    // Test that upgrade creates backup
    err = testutils_cobra.ExecutePingcli(t, "config", "upgrade", 
        "--config-file", originalConfigPath,
        "--create-backup")
    testutils.CheckExpectedError(t, err, nil)
    
    // Verify backup was created
    assert.FileExists(t, backupConfigPath)
    
    // Verify backup contains original content
    backupContent, err := os.ReadFile(backupConfigPath)
    require.NoError(t, err)
    assert.Equal(t, originalConfig, string(backupContent))
    
    // Test recovery from backup
    err = testutils_cobra.ExecutePingcli(t, "config", "restore", 
        "--backup-file", backupConfigPath,
        "--config-file", originalConfigPath)
    testutils.CheckExpectedError(t, err, nil)
}
```

### Configuration Schema Evolution Testing

#### **Field Rename Testing**
```go
func TestConfigSchema_FieldRenames(t *testing.T) {
    // Test configuration files with renamed fields
    renamedFields := map[string]string{
        "environment_id": "worker_environment_id",
        "client_id":      "worker_client_id", 
        "client_secret":  "worker_client_secret",
        "region":         "region_code",
    }
    
    for oldField, newField := range renamedFields {
        t.Run(fmt.Sprintf("%s_to_%s", oldField, newField), func(t *testing.T) {
            // Create config with old field name
            legacyConfig := fmt.Sprintf(`
profiles:
  test:
    pingone:
      %s: "test-value"`, oldField)
            
            configFile := testutils.WriteStringToPipe(t, legacyConfig)
            defer configFile.Close()
            
            // Test that old field name still works (with deprecation warning)
            err := testutils_cobra.ExecutePingcli(t, "config", "validate", "--config-file", configFile.Name())
            testutils.CheckExpectedError(t, err, nil)
            
            // Test that new field name works
            newConfig := fmt.Sprintf(`
profiles:
  test:
    pingone:
      %s: "test-value"`, newField)
            
            newConfigFile := testutils.WriteStringToPipe(t, newConfig)
            defer newConfigFile.Close()
            
            err = testutils_cobra.ExecutePingcli(t, "config", "validate", "--config-file", newConfigFile.Name())
            testutils.CheckExpectedError(t, err, nil)
        })
    }
}
```

#### **Default Value Evolution Testing**
```go
func TestConfigSchema_DefaultValueEvolution(t *testing.T) {
    testCases := []struct {
        configVersion string
        expectedDefaults map[string]interface{}
        description string
    }{
        {
            configVersion: "1.0",
            expectedDefaults: map[string]interface{}{
                "output_format": "table",
                "color_output": true,
                "region_code": "NA",
            },
            description: "v1.0 defaults",
        },
        {
            configVersion: "2.0", 
            expectedDefaults: map[string]interface{}{
                "output_format": "json",  // Changed default
                "color_output": "auto",   // Changed from bool to string
                "region_code": "NA",
                "verify_tls": true,       // New field with default
            },
            description: "v2.0 defaults with new fields",
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.description, func(t *testing.T) {
            // Create minimal config for version
            minimalConfig := fmt.Sprintf(`version: "%s"`, tc.configVersion)
            configFile := testutils.WriteStringToPipe(t, minimalConfig)
            defer configFile.Close()
            
            // Test that defaults are applied correctly
            for key, expectedValue := range tc.expectedDefaults {
                output, err := testutils_cobra.ExecutePingcliWithOutput(t, "config", "get", 
                    "--config-file", configFile.Name(), key)
                testutils.CheckExpectedError(t, err, nil)
                assert.Contains(t, output, fmt.Sprintf("%v", expectedValue))
            }
        })
    }
}
```

### Configuration Test Data Management

#### **Test Configuration File Generators** 
```go
// In internal/testing/testutils_config/
func GenerateLegacyConfig(version string, profiles map[string]interface{}) string {
    // Generate configuration files in legacy formats
    // Support different versions and profile structures
    config := map[string]interface{}{
        "version": version,
        "profiles": profiles,
    }
    
    yamlData, _ := yaml.Marshal(config)
    return string(yamlData)
}

func GenerateCurrentConfig(profiles map[string]interface{}) string {  
    // Generate configuration files in current format
    return GenerateLegacyConfig("2.0", profiles)
}

func CreateTestProfile(envID, clientID, clientSecret, region string) map[string]interface{} {
    // Create standardized test profile structure
    return map[string]interface{}{
        "pingone": map[string]interface{}{
            "worker_environment_id": envID,
            "worker_client_id": clientID,
            "worker_client_secret": clientSecret,
            "region_code": region,
        },
    }
}
```

#### **Configuration Validation Helpers**
```go
func ValidateConfigMigration(t *testing.T, beforePath, afterPath string) {
    // Helper to validate that configuration migration preserved all data
    beforeConfig := loadConfigFile(t, beforePath)
    afterConfig := loadConfigFile(t, afterPath)
    
    // Check that no functional settings were lost
    validateProfilesEquivalent(t, beforeConfig.Profiles, afterConfig.Profiles)
    
    // Verify new format compliance
    assert.True(t, isValidCurrentFormat(afterConfig))
    
    // Validate that functionality still works
    validateConfigFunctionality(t, afterPath)
}

func CompareConfigValues(t *testing.T, config1, config2 interface{}, ignoredFields []string) {
    // Deep compare configuration values
    // Ignore version metadata and timestamps
    // Validate functional equivalence
    
    normalized1 := normalizeConfig(config1, ignoredFields)
    normalized2 := normalizeConfig(config2, ignoredFields)
    
    assert.Equal(t, normalized1, normalized2, "Configuration values should be functionally equivalent")
}
```

## Test Configuration Patterns

### CLI Command Test Helpers

Use consistent command testing patterns:
```go
func test<Command>Cmd_<Scenario>(t *testing.T) {
    // Initialize configuration system
    testutils_koanf.InitKoanfs(t)
    
    // Set up test data/environment as needed
    outputDir := t.TempDir()
    
    // Execute command with specific parameters
    err := testutils_cobra.ExecutePingcli(t, "command", "subcommand",
        "--flag1", "value1",
        "--flag2", "value2")
    
    // Validate results
    testutils.CheckExpectedError(t, err, expectedErrorPattern)
}
```

### Test Data Generation

Implement test data generators in `internal/testing/testutils_resource/`:
```go
func Generate<Resource>TestData() <ResourceType> {
    // Generate consistent test data for resources
    // Use deterministic values for repeatability
    // Include edge cases and boundary conditions
}
```

### Integration Test Helpers

Implement API integration helpers:
```go
func setupTestEnvironment(t *testing.T) *connector.ClientInfo {
    clientInfo := testutils.GetClientInfo(t)
    
    // Validate required services are available
    if clientInfo.PingOneApiClient == nil {
        t.Skip("PingOne API client not configured")
    }
    
    return clientInfo
}
```

## Quality Standards

### Test Coverage Requirements

All CLI commands and major functionality must have:
- [ ] **Basic execution testing** with valid parameters
- [ ] **Argument validation testing** (too many/few arguments)
- [ ] **Flag validation testing** with valid and invalid values
- [ ] **Required flag group testing** (flags that must be used together)
- [ ] **Help flag testing** (--help, -h)
- [ ] **Error condition testing** with expected error patterns
- [ ] **Authentication testing** for all supported auth methods
- [ ] **Integration testing** with real API endpoints (where applicable)

#### **Additional Requirements for CLI Changes**
When modifying existing commands:
- [ ] **Deprecation testing** for any removed or changed flags/commands
- [ ] **Migration path validation** from old to new CLI interfaces
- [ ] **Dual functionality testing** during deprecation periods
- [ ] **Warning validation** for deprecated flag/command usage
- [ ] **Backward compatibility testing** for existing user workflows

#### **Configuration Migration Requirements**
When changing configuration file formats or schema:
- [ ] **Version compatibility matrix testing** across supported versions
- [ ] **Automatic migration testing** for seamless upgrades
- [ ] **Backup and recovery testing** to prevent data loss
- [ ] **Field rename/restructure testing** with proper deprecation warnings
- [ ] **Default value evolution testing** to ensure consistency
- [ ] **Configuration validation testing** for all supported formats

#### **Export/Import Command Requirements**
For data export/import commands:
- [ ] **Format validation testing** for all supported output formats
- [ ] **Completeness testing** using `testutils.ValidateImportBlocks()`
- [ ] **Service filtering testing** (specific services vs. all services)
- [ ] **Overwrite behavior testing** (existing vs. new output directories)

### Test Reliability

- Use `t.Parallel()` for tests that can run concurrently
- Implement proper cleanup for temporary files and directories  
- Use `t.TempDir()` for temporary directory creation
- Handle container lifecycle properly for PingFederate tests
- Validate test stability across multiple runs
- Use deterministic test data and avoid random values

### Test Organization and Maintenance

- Follow naming convention: `Test<Command>Cmd_<Scenario>`
- Keep test functions focused on single scenarios
- Use descriptive test and variable names
- Group related tests in the same file as the command implementation
- Document complex test scenarios and setup requirements
- Regular review and update of test patterns as CLI evolves
- Ensure container tests work in CI/CD environments

## Examples

### Complete CLI Command Test Structure
```go
// Basic execution test
func Test<Command>Cmd_Execute(t *testing.T) {
    testutils_koanf.InitKoanfs(t) 
    outputDir := t.TempDir()
    
    err := testutils_cobra.ExecutePingcli(t, "command", "subcommand",
        "--output-directory", outputDir,
        "--overwrite")
    testutils.CheckExpectedError(t, err, nil)
}

// Argument validation tests  
func Test<Command>Cmd_TooManyArgs(t *testing.T) { /* ... */ }
func Test<Command>Cmd_InvalidFlag(t *testing.T) { /* ... */ }

// Flag testing
func Test<Command>Cmd_<FlagName>Flag(t *testing.T) { /* ... */ }
func Test<Command>Cmd_<FlagName>FlagInvalid(t *testing.T) { /* ... */ }

// Integration tests
func Test<Command>Cmd_PingOneIntegration(t *testing.T) { /* ... */ }
func Test<Command>Cmd_PingFederateContainerIntegration(t *testing.T) { /* ... */ }

// For commands with deprecated functionality
func Test<Command>Cmd_Deprecation_<Flag>(t *testing.T) {
    // Legacy flag usage test
    // New flag usage test 
    // Migration path test
    // Warning validation test
}
```

### Complete Export Command Test Structure
```go
func TestPlatformExportCmd_Execute(t *testing.T) { /* Basic execution */ }
func TestPlatformExportCmd_ServiceGroupFlag(t *testing.T) { /* Service filtering */ }
func TestPlatformExportCmd_ExportFormatFlag(t *testing.T) { /* Format testing */ }
func TestPlatformExportCmd_PingOneWorkerFlags(t *testing.T) { /* Authentication */ }
func TestPlatformExportCmd_ContainerIntegration(t *testing.T) { /* Container tests */ }
```

### Complete Integration Test Structure
```go
func TestConnector_<Service>_Export(t *testing.T) {
    clientInfo := testutils.GetClientInfo(t)
    resource := &connector.<Service>Resource{ClientInfo: clientInfo}
    
    // Test export functionality
    testutils.ValidateImportBlocks(t, resource, expectedBlocks)
}
```

### Complete Configuration Migration Test Structure
```go
func TestConfigMigration_v<X>_to_v<Y>(t *testing.T) {
    // Setup legacy configuration
    legacyConfig := testutils_config.GenerateLegacyConfig("1.0", testProfiles)
    configPath := filepath.Join(t.TempDir(), "config.yaml")
    
    // Test migration process
    err := testutils_cobra.ExecutePingcli(t, "config", "upgrade", "--config-file", configPath)
    testutils.CheckExpectedError(t, err, nil)
    
    // Validate migration results
    testutils_config.ValidateConfigMigration(t, legacyConfigPath, configPath)
}

func TestConfigCompatibility_AllVersions(t *testing.T) {
    versions := []string{"1.0.0", "1.5.0", "2.0.0", "2.1.0"}
    
    for _, version := range versions {
        t.Run(fmt.Sprintf("version_%s", version), func(t *testing.T) {
            configFile := fmt.Sprintf("testdata/configs/sample-%s.yaml", version)
            
            // Test that configuration loads without error
            err := testutils_cobra.ExecutePingcli(t, "config", "validate", "--config-file", configFile)
            testutils.CheckExpectedError(t, err, nil)
            
            // Test that functionality works with this configuration
            err = testutils_cobra.ExecutePingcli(t, "platform", "export", 
                "--config-file", configFile,
                "--dry-run")
            testutils.CheckExpectedError(t, err, nil)
        })
    }
}
```

This comprehensive testing strategy ensures that all CLI functionality is thoroughly validated, providing confidence in the reliability and correctness of the PingCLI across different environments and use cases. The configuration migration testing ensures users can seamlessly upgrade between CLI versions without losing their settings or experiencing breaking changes.