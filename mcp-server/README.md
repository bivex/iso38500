# ISO 38500 Governance MCP Server

A Model Context Protocol (MCP) server that exposes the ISO 38500 Governance Framework SDK as AI assistant tools. This allows AI assistants like Claude to interact with enterprise IT governance systems through standardized MCP tools.

## Overview

The MCP server provides a standardized interface for AI assistants to:
- Create and manage applications in governance portfolios
- Establish governance agreements and compliance frameworks
- Evaluate applications for risk and compliance
- Monitor governance metrics and KPIs
- Run enterprise governance demonstrations

## Features

### 🛠️ Available Tools

#### Application Management
- **`create_application`** - Create new applications in the portfolio
- **`list_applications`** - View all applications with status details

#### Portfolio Management
- **`create_portfolio`** - Create application portfolios
- **`add_to_portfolio`** - Add applications to portfolios
- **`list_portfolios`** - View all portfolios and their applications

#### Governance Framework
- **`create_governance_agreement`** - Establish governance agreements
- **`evaluate_application`** - Assess application compliance and risk
- **`evaluate_portfolio`** - Evaluate entire portfolio health
- **`monitor_governance`** - Track KPIs and risk indicators

#### Enterprise Demo
- **`run_enterprise_demo`** - Execute complete enterprise governance scenario

## Installation

1. **Clone and build the SDK:**
```bash
git clone https://github.com/iso38500/iso38500-governance-sdk.git
cd iso38500-governance-sdk
go mod tidy
go build ./...
```

2. **Build the MCP server:**
```bash
cd mcp-server
go mod tidy
go build -o mcp-server main.go
```

## Configuration

The MCP server uses in-memory repositories by default. For production use, you can configure:
- Database repositories (PostgreSQL, MySQL)
- External service integrations
- Custom governance policies

## Usage

### As an MCP Server

The server communicates via JSON-RPC over stdin/stdout:

```bash
./mcp-server
```

### Integration with Claude Desktop

Add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "governance": {
      "command": "/path/to/mcp-server",
      "args": []
    }
  }
}
```

### Example Tool Usage

**Create an Application:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "create_application",
    "arguments": {
      "id": "erp-system-001",
      "name": "Enterprise Resource Planning",
      "description": "Core ERP system for business operations",
      "version": "2024.1.0"
    }
  }
}
```

**Evaluate Application Risk:**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/call",
  "params": {
    "name": "evaluate_application",
    "arguments": {
      "application_id": "erp-system-001",
      "evaluator": "AI Governance Assistant"
    }
  }
}
```

## Tool Specifications

### create_application
Creates a new application in the governance portfolio.

**Parameters:**
- `id` (string, required): Unique application identifier
- `name` (string, required): Application name
- `description` (string, required): Application description
- `version` (string, optional): Application version (default: "1.0.0")

### create_portfolio
Creates a new application portfolio.

**Parameters:**
- `id` (string, required): Unique portfolio identifier
- `name` (string, required): Portfolio name
- `description` (string, required): Portfolio description
- `owner` (string, required): Portfolio owner

### add_to_portfolio
Adds an application to a portfolio.

**Parameters:**
- `portfolio_id` (string, required): Portfolio identifier
- `application_id` (string, required): Application identifier

### create_governance_agreement
Creates a governance agreement for an application.

**Parameters:**
- `id` (string, required): Unique agreement identifier
- `application_id` (string, required): Application identifier
- `title` (string, required): Agreement title

### evaluate_application
Evaluates an application for governance compliance.

**Parameters:**
- `application_id` (string, required): Application to evaluate
- `evaluator` (string, optional): Name of evaluator (default: "MCP Assistant")

**Returns:** Risk level, technical health score, business value, recommendations

### evaluate_portfolio
Evaluates an entire portfolio for governance compliance.

**Parameters:**
- `portfolio_id` (string, required): Portfolio to evaluate

**Returns:** Application counts, risk distribution, portfolio health metrics

### monitor_governance
Monitors governance metrics for an application.

**Parameters:**
- `agreement_id` (string, required): Governance agreement identifier

**Returns:** KPI measurements, risk indicators, compliance status

### list_applications
Lists all applications in the portfolio.

**Returns:** Detailed list of all applications with status and metadata

### list_portfolios
Lists all portfolios.

**Returns:** Detailed list of all portfolios with applications and metadata

### run_enterprise_demo
Runs the complete enterprise governance demonstration.

**Returns:** Summary of enterprise governance scenario execution

## MCP Protocol

The server implements the Model Context Protocol specification:

- **Protocol Version:** 2024-11-05
- **Transport:** JSON-RPC 2.0 over stdin/stdout
- **Capabilities:** Tools with list and call operations

### Protocol Messages

**Initialize:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {}
}
```

**List Tools:**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list",
  "params": {}
}
```

**Call Tool:**
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "tool_name",
    "arguments": {...}
  }
}
```

## Architecture

```
┌─────────────────┐
│  AI Assistant   │
│   (Claude)      │
├─────────────────┤
│  MCP Protocol   │ ← JSON-RPC over stdin/stdout
├─────────────────┤
│ MCP Server      │ ← Tool execution
├─────────────────┤
│ Governance SDK  │ ← Business logic
├─────────────────┤
│ Repositories    │ ← Data persistence
└─────────────────┘
```

## Development

### Building
```bash
go build -o mcp-server main.go
```

### Testing
```bash
go test ./...
```

### Debugging
Enable debug logging:
```bash
DEBUG=1 ./mcp-server
```

## Security Considerations

- **Input Validation:** All tool inputs are validated
- **Error Handling:** Sensitive errors are sanitized
- **Access Control:** Repository-level access controls apply
- **Audit Logging:** All governance actions are logged

## Contributing

1. Follow the existing code structure and patterns
2. Add comprehensive tests for new tools
3. Update documentation for API changes
4. Ensure MCP protocol compliance

## License

MIT License - see LICENSE file for details.

## Related Projects

- [ISO 38500 Governance SDK](https://github.com/iso38500/iso38500-governance-sdk)
- [Model Context Protocol](https://modelcontextprotocol.io/)
- [Claude Desktop](https://docs.anthropic.com/claude/docs/desktop)
