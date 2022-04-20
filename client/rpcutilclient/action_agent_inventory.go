// generated code; DO NOT EDIT
//
// Client for Choria RPC Agent 'rpcutil' Version 0.25.1 generated using Choria version 0.25.1

package rpcutilclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/choria-io/go-choria/protocol"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
	"github.com/choria-io/go-choria/providers/agent/mcorpc/ddl/agent"
	"github.com/choria-io/go-choria/providers/agent/mcorpc/replyfmt"
)

// AgentInventoryRequester performs a RPC request to rpcutil#agent_inventory
type AgentInventoryRequester struct {
	r    *requester
	outc chan *AgentInventoryOutput
}

// NewAgentInventoryOutput creates a new instance of AgentInventoryOutput
func NewAgentInventoryOutput(details *ResultDetails, data json.RawMessage) (*AgentInventoryOutput, error) {
	output := &AgentInventoryOutput{
		reply:   make(map[string]interface{}),
		details: details,
	}

	return output, json.Unmarshal(data, &output.reply)
}

// AgentInventoryOutput is the output from the agent_inventory action
type AgentInventoryOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// NewAgentInventoryResult creates a new instance of AgentInventoryResult
func NewAgentInventoryResult(ddl *agent.DDL) *AgentInventoryResult {
	return &AgentInventoryResult{ddl: ddl}
}

// AgentInventoryResult is the result from a agent_inventory action
type AgentInventoryResult struct {
	ddl        *agent.DDL
	stats      *rpcclient.Stats
	outputs    []*AgentInventoryOutput
	rpcreplies []*replyfmt.RPCReply
	mu         sync.Mutex
}

// RecordResult records a new output into the result
func (d *AgentInventoryResult) RecordResult(output *AgentInventoryOutput, reply *replyfmt.RPCReply) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.outputs = append(d.outputs, output)
	d.rpcreplies = append(d.rpcreplies, reply)
}

func (d *AgentInventoryResult) RenderResults(w io.Writer, format RenderFormat, displayMode DisplayMode, verbose bool, silent bool, colorize bool, log Log) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.stats == nil {
		return fmt.Errorf("result stats is not set, result was not completed")
	}

	results := &replyfmt.RPCResults{
		Agent:   d.stats.Agent(),
		Action:  d.stats.Action(),
		Replies: d.rpcreplies,
		Stats:   d.stats,
	}

	addl, err := d.ddl.ActionInterface(d.stats.Action())
	if err != nil {
		return err
	}

	switch format {
	case JSONFormat:
		return results.RenderJSON(w, addl)
	case TableFormat:
		return results.RenderTable(w, addl)
	case TXTFooter:
		results.RenderTXTFooter(w, verbose)
		return nil
	default:
		return results.RenderTXT(w, addl, verbose, silent, replyfmt.DisplayMode(displayMode), colorize, log)
	}
}

// Stats is the rpc request stats
func (d *AgentInventoryResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *AgentInventoryOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *AgentInventoryOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *AgentInventoryOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// ParseAgentInventoryOutput parses the result value from the AgentInventory action into target
func (d *AgentInventoryOutput) ParseAgentInventoryOutput(target interface{}) error {
	j, err := d.JSON()
	if err != nil {
		return fmt.Errorf("could not access payload: %s", err)
	}

	err = json.Unmarshal(j, target)
	if err != nil {
		return fmt.Errorf("could not unmarshal JSON payload: %s", err)
	}

	return nil
}

// Do performs the request
func (d *AgentInventoryRequester) Do(ctx context.Context) (*AgentInventoryResult, error) {
	dres := NewAgentInventoryResult(d.r.client.ddl)

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		// filtered by expr filter
		if r == nil {
			return
		}

		output, err := NewAgentInventoryOutput(NewResultDetails(pr.SenderID(), int(r.Statuscode), r.Statusmsg, pr.Time()), r.Data)
		if err != nil {
			d.r.client.errorf("Could not decode reply from %s: %s", pr.SenderID(), err)
		}

		// caller wants a channel so we dont return a resulset too (lots of memory etc)
		// this is unused now, no support for setting a channel
		if d.outc != nil {
			d.outc <- output
			return
		}

		// else prepare our result set
		dres.RecordResult(output, &replyfmt.RPCReply{
			Sender:   pr.SenderID(),
			RPCReply: r,
		})
	}

	res, err := d.r.do(ctx, handler)
	if err != nil {
		return nil, err
	}

	dres.stats = res

	return dres, nil
}

// AllOutputs provide access to all outputs
func (d *AgentInventoryResult) AllOutputs() []*AgentInventoryOutput {
	return d.outputs
}

// EachOutput iterates over all results received
func (d *AgentInventoryResult) EachOutput(h func(r *AgentInventoryOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Agents is the value of the agents output
//
// Description: List of agents on the server
func (d *AgentInventoryOutput) Agents() []interface{} {
	val := d.reply["agents"]

	return val.([]interface{})

}
