// Copyright 2021 JMSilvaDev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package apiclient is a library that contains methods to manipulate the Form3API
/* Usage:
1 - Instantiate a client:
	client = apiclient.NewClient()
2 - Methods:
	Fetch: client.Fetch(account-uuid)
	Delete: client.Delete(account-uuid, version)
	NewAccount: account := client.NewAccount()
	Create: client.Delete(account)
*/
package apiclient
