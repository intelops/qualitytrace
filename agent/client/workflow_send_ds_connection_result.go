package client

import (
	"context"
	"fmt"

	"github.com/intelops/qualityTrace/agent/proto"
	"github.com/intelops/qualityTrace/agent/telemetry"
)

func (c *Client) SendDataStoreConnectionResult(ctx context.Context, response *proto.DataStoreConnectionTestResponse) error {
	client := proto.NewOrchestratorClient(c.conn)

	response.AgentIdentification = c.sessionConfig.AgentIdentification
	response.Metadata = telemetry.ExtractMetadataFromContext(ctx)

	_, err := client.SendDataStoreConnectionTestResult(ctx, response)
	if err != nil {
		return fmt.Errorf("could not send otlp connection result request: %w", err)
	}

	return nil
}
