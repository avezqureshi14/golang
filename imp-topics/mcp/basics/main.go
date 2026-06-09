package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//////////////////////
// MCP TYPES
//////////////////////

type Request struct {
	Jsonrpc string         `json:"jsonrpc"`
	ID      any            `json:"id"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params,omitempty"`
}

type Response struct {
	Jsonrpc string    `json:"jsonrpc"`
	ID      any       `json:"id,omitempty"`
	Result  any       `json:"result,omitempty"`
	Error   *RPCError `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//////////////////////
// SERVER LOGIC
//////////////////////

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Bytes()

		var req Request
		if err := json.Unmarshal(line, &req); err != nil {
			writeError(nil, -32700, "invalid JSON")
			continue
		}

		switch req.Method {

		//////////////////////
		// INITIALIZE
		//////////////////////
		case "initialize":
			write(req.ID, map[string]any{
				"protocolVersion": "0.1",
				"serverInfo": map[string]any{
					"name":    "go-mcp-single",
					"version": "1.0.0",
				},
				"capabilities": map[string]any{
					"tools": true,
				},
			})

		//////////////////////
		// LIST TOOLS
		//////////////////////
		case "tools/list":
			write(req.ID, map[string]any{
				"tools": []map[string]any{
					{
						"name":        "add",
						"description": "Add two numbers",
						"inputSchema": map[string]any{
							"type": "object",
							"properties": map[string]any{
								"a": map[string]any{"type": "number"},
								"b": map[string]any{"type": "number"},
							},
							"required": []string{"a", "b"},
						},
					},
				},
			})

		//////////////////////
		// CALL TOOL
		//////////////////////
		case "tools/call":
			params := req.Params

			name, ok := params["name"].(string)
			if !ok {
				writeError(req.ID, -32602, "invalid tool name")
				continue
			}

			switch name {

			case "add":
				args, ok := params["arguments"].(map[string]any)
				if !ok {
					writeError(req.ID, -32602, "invalid arguments")
					continue
				}

				a, aok := args["a"].(float64)
				b, bok := args["b"].(float64)

				if !aok || !bok {
					writeError(req.ID, -32602, "invalid numbers")
					continue
				}

				result := a + b + 1000

				write(req.ID, map[string]any{
					"content": []map[string]any{
						{
							"type": "text",
							"text": fmt.Sprintf("%v", result),
						},
					},
				})

			default:
				writeError(req.ID, -32601, "tool not found")
			}

		default:
			writeError(req.ID, -32601, "method not found")
		}
	}
}

//////////////////////
// HELPERS
//////////////////////

func write(id any, result any) {
	resp := Response{
		Jsonrpc: "2.0",
		ID:      id,
		Result:  result,
	}
	send(resp)
}

func writeError(id any, code int, msg string) {
	resp := Response{
		Jsonrpc: "2.0",
		ID:      id,
		Error: &RPCError{
			Code:    code,
			Message: msg,
		},
	}
	send(resp)
}

func send(resp Response) {
	data, _ := json.Marshal(resp)
	fmt.Println(string(data)) // IMPORTANT: stdout = protocol
}
