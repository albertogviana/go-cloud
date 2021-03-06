// Copyright 2018 The Go Cloud Development Kit Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memblob_test

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/blob"
	"gocloud.dev/blob/memblob"
)

func Example() {

	// Create an in-memory bucket.
	b := memblob.OpenBucket(nil)

	// Now we can use b to read or write files to the container.
	ctx := context.Background()
	err := b.WriteAll(ctx, "my-key", []byte("hello world"), nil)
	if err != nil {
		log.Fatal(err)
	}
	data, err := b.ReadAll(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// Output:
	// hello world
}

func Example_open() {
	// Open creates a *blob.Bucket from a URL.
	b, err := blob.OpenBucket(context.Background(), "mem://")
	if err != nil {
		log.Fatal(err)
	}

	// Now we can use b to read or write files to the container.
	ctx := context.Background()
	err = b.WriteAll(ctx, "my-key", []byte("hello world"), nil)
	if err != nil {
		log.Fatal(err)
	}
	data, err := b.ReadAll(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// Output:
	// hello world
}
