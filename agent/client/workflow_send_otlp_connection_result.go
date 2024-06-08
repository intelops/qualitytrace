package client

import (
	"context"
	"fmt"

	"github.com/intelops/qualitytrace/agent/proto"
	"github.com/intelops/qualitytrace/agent/telemetry"
)

func (c *Client) SendOTLPConnectionResult(ctx context.Context, response *proto.OTLPConnectionTestResponse) error {
	client := proto.NewOrchestratorClient(c.conn)

	response.AgentIdentification = c.sessionConfig.AgentIdentification
	response.Metadata = telemetry.ExtractMetadataFromContext(ctx)

	_, err := client.SendOTLPConnectionTestResult(ctx, response)
	if err != nil {
		return fmt.Errorf("could not send otlp connection result request: %w", err)
	}

	return nil
}
