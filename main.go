/*
Copyright © 2021 VATSAL PAREKH <vparekh@redhat.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "github.com/vatsalparekh/cowin-query/cmd"

var (
	pincode   string
	inputDate string
)

func main() {
	// Parse the input pincode and inputDate from CLI args
	// Sanitize the inputs
	// Query API-Setu with the input filters
	// Find our 18+ slots from the Query data
	// Output the results
	// https://apisetu.gov.in/public/marketplace/api/cowin
	cmd.Execute()
}
