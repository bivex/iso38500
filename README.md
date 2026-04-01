# ISO 38500 Governance SDK

IT governance standard reference and implementation toolkit based on ISO/IEC 38500.

## Components

### Governance SDK (`iso38500-governance-sdk/`)

Go-based SDK implementing ISO 38500 governance principles with Docker support.

```bash
cd iso38500-governance-sdk
docker-compose up
```

### MCP Server (`mcp-server/`)

Model Context Protocol server for governance evaluation and assessment.

```bash
cd mcp-server
go build -o mcp-server .
./mcp-server
```

## ISO 38500 Principles

The six key principles of IT governance covered:

1. **Responsibility** - Clear roles and accountability
2. **Strategy** - IT aligned with business objectives
3. **Acquisition** - Sound IT investment decisions
4. **Performance** - IT supports business processes
5. **Conformance** - Compliance with obligations
6. **Human Behavior** - Respect for human factors

## License

See [LICENSE](LICENSE) for details.
