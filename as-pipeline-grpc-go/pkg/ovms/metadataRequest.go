/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: BSD-3-Clause
 **********************************************************************/

package ovms

import (
	"context"
	"log"
	"time"

	grpc_client "aicsd/grpc-client"
)

func ModelMetadataRequest(client grpc_client.GRPCInferenceServiceClient, modelName string, modelVersion string) *grpc_client.ModelMetadataResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create status request for a given model
	modelMetadataRequest := grpc_client.ModelMetadataRequest{
		Name:    modelName,
		Version: modelVersion,
	}
	// Submit modelMetadata request to server
	modelMetadataResponse, err := client.ModelMetadata(ctx, &modelMetadataRequest)
	if err != nil {
		log.Fatalf("Couldn't get server model metadata: %v", err)
	}
	return modelMetadataResponse
}
