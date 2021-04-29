// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START documentai_v1beta3_generated_DocumentProcessorService_ProcessDocument_sync]

package main

import (
	"context"

	documentai "cloud.google.com/go/documentai/apiv1beta3"
	documentaipb "google.golang.org/genproto/googleapis/cloud/documentai/v1beta3"
)

func main() {
	// import documentaipb "google.golang.org/genproto/googleapis/cloud/documentai/v1beta3"

	ctx := context.Background()
	c, err := documentai.NewDocumentProcessorClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &documentaipb.ProcessRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ProcessDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END documentai_v1beta3_generated_DocumentProcessorService_ProcessDocument_sync]
