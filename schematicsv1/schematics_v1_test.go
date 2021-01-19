/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schematicsv1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`SchematicsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListSchematicsLocation(listSchematicsLocationOptions *ListSchematicsLocationOptions) - Operation response error`, func() {
		listSchematicsLocationPath := "/v1/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSchematicsLocationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSchematicsLocation with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListSchematicsLocationOptions model
				listSchematicsLocationOptionsModel := new(schematicsv1.ListSchematicsLocationOptions)
				listSchematicsLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListSchematicsLocation(listSchematicsLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListSchematicsLocation(listSchematicsLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListSchematicsLocation(listSchematicsLocationOptions *ListSchematicsLocationOptions)`, func() {
		listSchematicsLocationPath := "/v1/locations"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSchematicsLocationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"country": "Country", "geography": "Geography", "id": "ID", "kind": "Kind", "metro": "Metro", "multizone_metro": "MultizoneMetro", "name": "Name"}]`)
				}))
			})
			It(`Invoke ListSchematicsLocation successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListSchematicsLocation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSchematicsLocationOptions model
				listSchematicsLocationOptionsModel := new(schematicsv1.ListSchematicsLocationOptions)
				listSchematicsLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListSchematicsLocation(listSchematicsLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListSchematicsLocationWithContext(ctx, listSchematicsLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListSchematicsLocation(listSchematicsLocationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListSchematicsLocationWithContext(ctx, listSchematicsLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListSchematicsLocation with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListSchematicsLocationOptions model
				listSchematicsLocationOptionsModel := new(schematicsv1.ListSchematicsLocationOptions)
				listSchematicsLocationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListSchematicsLocation(listSchematicsLocationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListResourceGroup(listResourceGroupOptions *ListResourceGroupOptions) - Operation response error`, func() {
		listResourceGroupPath := "/v1/resource_groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceGroup with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListResourceGroupOptions model
				listResourceGroupOptionsModel := new(schematicsv1.ListResourceGroupOptions)
				listResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListResourceGroup(listResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListResourceGroup(listResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceGroup(listResourceGroupOptions *ListResourceGroupOptions)`, func() {
		listResourceGroupPath := "/v1/resource_groups"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceGroupPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"account_id": "AccountID", "crn": "Crn", "default": false, "name": "Name", "resource_group_id": "ResourceGroupID", "state": "State"}]`)
				}))
			})
			It(`Invoke ListResourceGroup successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceGroupOptions model
				listResourceGroupOptionsModel := new(schematicsv1.ListResourceGroupOptions)
				listResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListResourceGroup(listResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListResourceGroupWithContext(ctx, listResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListResourceGroup(listResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListResourceGroupWithContext(ctx, listResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListResourceGroup with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListResourceGroupOptions model
				listResourceGroupOptionsModel := new(schematicsv1.ListResourceGroupOptions)
				listResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListResourceGroup(listResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchematicsVersion(getSchematicsVersionOptions *GetSchematicsVersionOptions) - Operation response error`, func() {
		getSchematicsVersionPath := "/v1/version"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchematicsVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchematicsVersion with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsVersionOptions model
				getSchematicsVersionOptionsModel := new(schematicsv1.GetSchematicsVersionOptions)
				getSchematicsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetSchematicsVersion(getSchematicsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetSchematicsVersion(getSchematicsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSchematicsVersion(getSchematicsVersionOptions *GetSchematicsVersionOptions)`, func() {
		getSchematicsVersionPath := "/v1/version"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchematicsVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"builddate": "Builddate", "buildno": "Buildno", "commitsha": "Commitsha", "helm_provider_version": "HelmProviderVersion", "helm_version": "HelmVersion", "supported_template_types": {"anyKey": "anyValue"}, "terraform_provider_version": "TerraformProviderVersion", "terraform_version": "TerraformVersion"}`)
				}))
			})
			It(`Invoke GetSchematicsVersion successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetSchematicsVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchematicsVersionOptions model
				getSchematicsVersionOptionsModel := new(schematicsv1.GetSchematicsVersionOptions)
				getSchematicsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetSchematicsVersion(getSchematicsVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetSchematicsVersionWithContext(ctx, getSchematicsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetSchematicsVersion(getSchematicsVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetSchematicsVersionWithContext(ctx, getSchematicsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSchematicsVersion with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetSchematicsVersionOptions model
				getSchematicsVersionOptionsModel := new(schematicsv1.GetSchematicsVersionOptions)
				getSchematicsVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetSchematicsVersion(getSchematicsVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions) - Operation response error`, func() {
		listWorkspacesPath := "/v1/workspaces"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWorkspaces with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(schematicsv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listWorkspacesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions)`, func() {
		listWorkspacesPath := "/v1/workspaces"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkspacesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "limit": 5, "offset": 6, "workspaces": [{"applied_shareddata_ids": ["AppliedShareddataIds"], "catalog_ref": {"dry_run": true, "item_icon_url": "ItemIconURL", "item_id": "ItemID", "item_name": "ItemName", "item_readme_url": "ItemReadmeURL", "item_url": "ItemURL", "launch_url": "LaunchURL", "offering_version": "OfferingVersion"}, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "crn": "Crn", "description": "Description", "id": "ID", "last_health_check_at": "2019-01-01T12:00:00", "location": "Location", "name": "Name", "resource_group": "ResourceGroup", "runtime_data": [{"engine_cmd": "EngineCmd", "engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "log_store_url": "LogStoreURL", "output_values": [{"anyKey": "anyValue"}], "resources": [[{"anyKey": "anyValue"}]], "state_store_url": "StateStoreURL"}], "shared_data": {"cluster_id": "ClusterID", "cluster_name": "ClusterName", "entitlement_keys": [{"anyKey": "anyValue"}], "namespace": "Namespace", "region": "Region", "resource_group_id": "ResourceGroupID"}, "status": "Status", "tags": ["Tags"], "template_data": [{"env_values": [{"hidden": true, "name": "Name", "secure": true, "value": "Value"}], "folder": "Folder", "has_githubtoken": true, "id": "ID", "template_type": "TemplateType", "uninstall_script_name": "UninstallScriptName", "values": "Values", "values_metadata": [{"anyKey": "anyValue"}], "values_url": "ValuesURL", "variablestore": [{"description": "Description", "name": "Name", "secure": true, "type": "Type", "value": "Value"}]}], "template_ref": "TemplateRef", "template_repo": {"branch": "Branch", "full_url": "FullURL", "has_uploadedgitrepotar": false, "release": "Release", "repo_sha_value": "RepoShaValue", "repo_url": "RepoURL", "url": "URL"}, "type": ["Type"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "workspace_status": {"frozen": true, "frozen_at": "2019-01-01T12:00:00", "frozen_by": "FrozenBy", "locked": true, "locked_by": "LockedBy", "locked_time": "2019-01-01T12:00:00"}, "workspace_status_msg": {"status_code": "StatusCode", "status_msg": "StatusMsg"}}]}`)
				}))
			})
			It(`Invoke ListWorkspaces successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListWorkspaces(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(schematicsv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listWorkspacesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListWorkspacesWithContext(ctx, listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListWorkspacesWithContext(ctx, listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListWorkspaces with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := new(schematicsv1.ListWorkspacesOptions)
				listWorkspacesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listWorkspacesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listWorkspacesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListWorkspaces(listWorkspacesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions) - Operation response error`, func() {
		createWorkspacePath := "/v1/workspaces"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkspacePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Github-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Github-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateWorkspace with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoRequest model
				templateRepoRequestModel := new(schematicsv1.TemplateRepoRequest)
				templateRepoRequestModel.Branch = core.StringPtr("testString")
				templateRepoRequestModel.Release = core.StringPtr("testString")
				templateRepoRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusRequest model
				workspaceStatusRequestModel := new(schematicsv1.WorkspaceStatusRequest)
				workspaceStatusRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(schematicsv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.AppliedShareddataIds = []string{"testString"}
				createWorkspaceOptionsModel.CatalogRef = catalogRefModel
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Location = core.StringPtr("testString")
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				createWorkspaceOptionsModel.Tags = []string{"testString"}
				createWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				createWorkspaceOptionsModel.TemplateRef = core.StringPtr("testString")
				createWorkspaceOptionsModel.TemplateRepo = templateRepoRequestModel
				createWorkspaceOptionsModel.Type = []string{"testString"}
				createWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusRequestModel
				createWorkspaceOptionsModel.XGithubToken = core.StringPtr("testString")
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions)`, func() {
		createWorkspacePath := "/v1/workspaces"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkspacePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Github-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Github-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"applied_shareddata_ids": ["AppliedShareddataIds"], "catalog_ref": {"dry_run": true, "item_icon_url": "ItemIconURL", "item_id": "ItemID", "item_name": "ItemName", "item_readme_url": "ItemReadmeURL", "item_url": "ItemURL", "launch_url": "LaunchURL", "offering_version": "OfferingVersion"}, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "crn": "Crn", "description": "Description", "id": "ID", "last_health_check_at": "2019-01-01T12:00:00", "location": "Location", "name": "Name", "resource_group": "ResourceGroup", "runtime_data": [{"engine_cmd": "EngineCmd", "engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "log_store_url": "LogStoreURL", "output_values": [{"anyKey": "anyValue"}], "resources": [[{"anyKey": "anyValue"}]], "state_store_url": "StateStoreURL"}], "shared_data": {"cluster_id": "ClusterID", "cluster_name": "ClusterName", "entitlement_keys": [{"anyKey": "anyValue"}], "namespace": "Namespace", "region": "Region", "resource_group_id": "ResourceGroupID"}, "status": "Status", "tags": ["Tags"], "template_data": [{"env_values": [{"hidden": true, "name": "Name", "secure": true, "value": "Value"}], "folder": "Folder", "has_githubtoken": true, "id": "ID", "template_type": "TemplateType", "uninstall_script_name": "UninstallScriptName", "values": "Values", "values_metadata": [{"anyKey": "anyValue"}], "values_url": "ValuesURL", "variablestore": [{"description": "Description", "name": "Name", "secure": true, "type": "Type", "value": "Value"}]}], "template_ref": "TemplateRef", "template_repo": {"branch": "Branch", "full_url": "FullURL", "has_uploadedgitrepotar": false, "release": "Release", "repo_sha_value": "RepoShaValue", "repo_url": "RepoURL", "url": "URL"}, "type": ["Type"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "workspace_status": {"frozen": true, "frozen_at": "2019-01-01T12:00:00", "frozen_by": "FrozenBy", "locked": true, "locked_by": "LockedBy", "locked_time": "2019-01-01T12:00:00"}, "workspace_status_msg": {"status_code": "StatusCode", "status_msg": "StatusMsg"}}`)
				}))
			})
			It(`Invoke CreateWorkspace successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.CreateWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoRequest model
				templateRepoRequestModel := new(schematicsv1.TemplateRepoRequest)
				templateRepoRequestModel.Branch = core.StringPtr("testString")
				templateRepoRequestModel.Release = core.StringPtr("testString")
				templateRepoRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusRequest model
				workspaceStatusRequestModel := new(schematicsv1.WorkspaceStatusRequest)
				workspaceStatusRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(schematicsv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.AppliedShareddataIds = []string{"testString"}
				createWorkspaceOptionsModel.CatalogRef = catalogRefModel
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Location = core.StringPtr("testString")
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				createWorkspaceOptionsModel.Tags = []string{"testString"}
				createWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				createWorkspaceOptionsModel.TemplateRef = core.StringPtr("testString")
				createWorkspaceOptionsModel.TemplateRepo = templateRepoRequestModel
				createWorkspaceOptionsModel.Type = []string{"testString"}
				createWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusRequestModel
				createWorkspaceOptionsModel.XGithubToken = core.StringPtr("testString")
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateWorkspaceWithContext(ctx, createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateWorkspaceWithContext(ctx, createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateWorkspace with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoRequest model
				templateRepoRequestModel := new(schematicsv1.TemplateRepoRequest)
				templateRepoRequestModel.Branch = core.StringPtr("testString")
				templateRepoRequestModel.Release = core.StringPtr("testString")
				templateRepoRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusRequest model
				workspaceStatusRequestModel := new(schematicsv1.WorkspaceStatusRequest)
				workspaceStatusRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := new(schematicsv1.CreateWorkspaceOptions)
				createWorkspaceOptionsModel.AppliedShareddataIds = []string{"testString"}
				createWorkspaceOptionsModel.CatalogRef = catalogRefModel
				createWorkspaceOptionsModel.Description = core.StringPtr("testString")
				createWorkspaceOptionsModel.Location = core.StringPtr("testString")
				createWorkspaceOptionsModel.Name = core.StringPtr("testString")
				createWorkspaceOptionsModel.ResourceGroup = core.StringPtr("testString")
				createWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				createWorkspaceOptionsModel.Tags = []string{"testString"}
				createWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				createWorkspaceOptionsModel.TemplateRef = core.StringPtr("testString")
				createWorkspaceOptionsModel.TemplateRepo = templateRepoRequestModel
				createWorkspaceOptionsModel.Type = []string{"testString"}
				createWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusRequestModel
				createWorkspaceOptionsModel.XGithubToken = core.StringPtr("testString")
				createWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.CreateWorkspace(createWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions) - Operation response error`, func() {
		getWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspacePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspace with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(schematicsv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions)`, func() {
		getWorkspacePath := "/v1/workspaces/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspacePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applied_shareddata_ids": ["AppliedShareddataIds"], "catalog_ref": {"dry_run": true, "item_icon_url": "ItemIconURL", "item_id": "ItemID", "item_name": "ItemName", "item_readme_url": "ItemReadmeURL", "item_url": "ItemURL", "launch_url": "LaunchURL", "offering_version": "OfferingVersion"}, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "crn": "Crn", "description": "Description", "id": "ID", "last_health_check_at": "2019-01-01T12:00:00", "location": "Location", "name": "Name", "resource_group": "ResourceGroup", "runtime_data": [{"engine_cmd": "EngineCmd", "engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "log_store_url": "LogStoreURL", "output_values": [{"anyKey": "anyValue"}], "resources": [[{"anyKey": "anyValue"}]], "state_store_url": "StateStoreURL"}], "shared_data": {"cluster_id": "ClusterID", "cluster_name": "ClusterName", "entitlement_keys": [{"anyKey": "anyValue"}], "namespace": "Namespace", "region": "Region", "resource_group_id": "ResourceGroupID"}, "status": "Status", "tags": ["Tags"], "template_data": [{"env_values": [{"hidden": true, "name": "Name", "secure": true, "value": "Value"}], "folder": "Folder", "has_githubtoken": true, "id": "ID", "template_type": "TemplateType", "uninstall_script_name": "UninstallScriptName", "values": "Values", "values_metadata": [{"anyKey": "anyValue"}], "values_url": "ValuesURL", "variablestore": [{"description": "Description", "name": "Name", "secure": true, "type": "Type", "value": "Value"}]}], "template_ref": "TemplateRef", "template_repo": {"branch": "Branch", "full_url": "FullURL", "has_uploadedgitrepotar": false, "release": "Release", "repo_sha_value": "RepoShaValue", "repo_url": "RepoURL", "url": "URL"}, "type": ["Type"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "workspace_status": {"frozen": true, "frozen_at": "2019-01-01T12:00:00", "frozen_by": "FrozenBy", "locked": true, "locked_by": "LockedBy", "locked_time": "2019-01-01T12:00:00"}, "workspace_status_msg": {"status_code": "StatusCode", "status_msg": "StatusMsg"}}`)
				}))
			})
			It(`Invoke GetWorkspace successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(schematicsv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceWithContext(ctx, getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceWithContext(ctx, getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspace with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceOptions model
				getWorkspaceOptionsModel := new(schematicsv1.GetWorkspaceOptions)
				getWorkspaceOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspace(getWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceOptions model with no property values
				getWorkspaceOptionsModelNew := new(schematicsv1.GetWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspace(getWorkspaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceWorkspace(replaceWorkspaceOptions *ReplaceWorkspaceOptions) - Operation response error`, func() {
		replaceWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceWorkspacePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceWorkspace with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")

				// Construct an instance of the ReplaceWorkspaceOptions model
				replaceWorkspaceOptionsModel := new(schematicsv1.ReplaceWorkspaceOptions)
				replaceWorkspaceOptionsModel.WID = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.CatalogRef = catalogRefModel
				replaceWorkspaceOptionsModel.Description = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.Name = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				replaceWorkspaceOptionsModel.Tags = []string{"testString"}
				replaceWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				replaceWorkspaceOptionsModel.TemplateRepo = templateRepoUpdateRequestModel
				replaceWorkspaceOptionsModel.Type = []string{"testString"}
				replaceWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusUpdateRequestModel
				replaceWorkspaceOptionsModel.WorkspaceStatusMsg = workspaceStatusMessageModel
				replaceWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ReplaceWorkspace(replaceWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ReplaceWorkspace(replaceWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceWorkspace(replaceWorkspaceOptions *ReplaceWorkspaceOptions)`, func() {
		replaceWorkspacePath := "/v1/workspaces/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceWorkspacePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applied_shareddata_ids": ["AppliedShareddataIds"], "catalog_ref": {"dry_run": true, "item_icon_url": "ItemIconURL", "item_id": "ItemID", "item_name": "ItemName", "item_readme_url": "ItemReadmeURL", "item_url": "ItemURL", "launch_url": "LaunchURL", "offering_version": "OfferingVersion"}, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "crn": "Crn", "description": "Description", "id": "ID", "last_health_check_at": "2019-01-01T12:00:00", "location": "Location", "name": "Name", "resource_group": "ResourceGroup", "runtime_data": [{"engine_cmd": "EngineCmd", "engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "log_store_url": "LogStoreURL", "output_values": [{"anyKey": "anyValue"}], "resources": [[{"anyKey": "anyValue"}]], "state_store_url": "StateStoreURL"}], "shared_data": {"cluster_id": "ClusterID", "cluster_name": "ClusterName", "entitlement_keys": [{"anyKey": "anyValue"}], "namespace": "Namespace", "region": "Region", "resource_group_id": "ResourceGroupID"}, "status": "Status", "tags": ["Tags"], "template_data": [{"env_values": [{"hidden": true, "name": "Name", "secure": true, "value": "Value"}], "folder": "Folder", "has_githubtoken": true, "id": "ID", "template_type": "TemplateType", "uninstall_script_name": "UninstallScriptName", "values": "Values", "values_metadata": [{"anyKey": "anyValue"}], "values_url": "ValuesURL", "variablestore": [{"description": "Description", "name": "Name", "secure": true, "type": "Type", "value": "Value"}]}], "template_ref": "TemplateRef", "template_repo": {"branch": "Branch", "full_url": "FullURL", "has_uploadedgitrepotar": false, "release": "Release", "repo_sha_value": "RepoShaValue", "repo_url": "RepoURL", "url": "URL"}, "type": ["Type"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "workspace_status": {"frozen": true, "frozen_at": "2019-01-01T12:00:00", "frozen_by": "FrozenBy", "locked": true, "locked_by": "LockedBy", "locked_time": "2019-01-01T12:00:00"}, "workspace_status_msg": {"status_code": "StatusCode", "status_msg": "StatusMsg"}}`)
				}))
			})
			It(`Invoke ReplaceWorkspace successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ReplaceWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")

				// Construct an instance of the ReplaceWorkspaceOptions model
				replaceWorkspaceOptionsModel := new(schematicsv1.ReplaceWorkspaceOptions)
				replaceWorkspaceOptionsModel.WID = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.CatalogRef = catalogRefModel
				replaceWorkspaceOptionsModel.Description = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.Name = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				replaceWorkspaceOptionsModel.Tags = []string{"testString"}
				replaceWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				replaceWorkspaceOptionsModel.TemplateRepo = templateRepoUpdateRequestModel
				replaceWorkspaceOptionsModel.Type = []string{"testString"}
				replaceWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusUpdateRequestModel
				replaceWorkspaceOptionsModel.WorkspaceStatusMsg = workspaceStatusMessageModel
				replaceWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ReplaceWorkspace(replaceWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceWorkspaceWithContext(ctx, replaceWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ReplaceWorkspace(replaceWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceWorkspaceWithContext(ctx, replaceWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceWorkspace with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")

				// Construct an instance of the ReplaceWorkspaceOptions model
				replaceWorkspaceOptionsModel := new(schematicsv1.ReplaceWorkspaceOptions)
				replaceWorkspaceOptionsModel.WID = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.CatalogRef = catalogRefModel
				replaceWorkspaceOptionsModel.Description = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.Name = core.StringPtr("testString")
				replaceWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				replaceWorkspaceOptionsModel.Tags = []string{"testString"}
				replaceWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				replaceWorkspaceOptionsModel.TemplateRepo = templateRepoUpdateRequestModel
				replaceWorkspaceOptionsModel.Type = []string{"testString"}
				replaceWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusUpdateRequestModel
				replaceWorkspaceOptionsModel.WorkspaceStatusMsg = workspaceStatusMessageModel
				replaceWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ReplaceWorkspace(replaceWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceWorkspaceOptions model with no property values
				replaceWorkspaceOptionsModelNew := new(schematicsv1.ReplaceWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ReplaceWorkspace(replaceWorkspaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteWorkspace(deleteWorkspaceOptions *DeleteWorkspaceOptions)`, func() {
		deleteWorkspacePath := "/v1/workspaces/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkspacePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["destroy_resources"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke DeleteWorkspace successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.DeleteWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteWorkspaceOptions model
				deleteWorkspaceOptionsModel := new(schematicsv1.DeleteWorkspaceOptions)
				deleteWorkspaceOptionsModel.WID = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.RefreshToken = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.DestroyResources = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.DeleteWorkspace(deleteWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DeleteWorkspaceWithContext(ctx, deleteWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.DeleteWorkspace(deleteWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DeleteWorkspaceWithContext(ctx, deleteWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteWorkspace with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkspaceOptions model
				deleteWorkspaceOptionsModel := new(schematicsv1.DeleteWorkspaceOptions)
				deleteWorkspaceOptionsModel.WID = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.RefreshToken = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.DestroyResources = core.StringPtr("testString")
				deleteWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.DeleteWorkspace(deleteWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteWorkspaceOptions model with no property values
				deleteWorkspaceOptionsModelNew := new(schematicsv1.DeleteWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.DeleteWorkspace(deleteWorkspaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions) - Operation response error`, func() {
		updateWorkspacePath := "/v1/workspaces/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkspacePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateWorkspace with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(schematicsv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.CatalogRef = catalogRefModel
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				updateWorkspaceOptionsModel.Tags = []string{"testString"}
				updateWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				updateWorkspaceOptionsModel.TemplateRepo = templateRepoUpdateRequestModel
				updateWorkspaceOptionsModel.Type = []string{"testString"}
				updateWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusUpdateRequestModel
				updateWorkspaceOptionsModel.WorkspaceStatusMsg = workspaceStatusMessageModel
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions)`, func() {
		updateWorkspacePath := "/v1/workspaces/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateWorkspacePath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"applied_shareddata_ids": ["AppliedShareddataIds"], "catalog_ref": {"dry_run": true, "item_icon_url": "ItemIconURL", "item_id": "ItemID", "item_name": "ItemName", "item_readme_url": "ItemReadmeURL", "item_url": "ItemURL", "launch_url": "LaunchURL", "offering_version": "OfferingVersion"}, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "crn": "Crn", "description": "Description", "id": "ID", "last_health_check_at": "2019-01-01T12:00:00", "location": "Location", "name": "Name", "resource_group": "ResourceGroup", "runtime_data": [{"engine_cmd": "EngineCmd", "engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "log_store_url": "LogStoreURL", "output_values": [{"anyKey": "anyValue"}], "resources": [[{"anyKey": "anyValue"}]], "state_store_url": "StateStoreURL"}], "shared_data": {"cluster_id": "ClusterID", "cluster_name": "ClusterName", "entitlement_keys": [{"anyKey": "anyValue"}], "namespace": "Namespace", "region": "Region", "resource_group_id": "ResourceGroupID"}, "status": "Status", "tags": ["Tags"], "template_data": [{"env_values": [{"hidden": true, "name": "Name", "secure": true, "value": "Value"}], "folder": "Folder", "has_githubtoken": true, "id": "ID", "template_type": "TemplateType", "uninstall_script_name": "UninstallScriptName", "values": "Values", "values_metadata": [{"anyKey": "anyValue"}], "values_url": "ValuesURL", "variablestore": [{"description": "Description", "name": "Name", "secure": true, "type": "Type", "value": "Value"}]}], "template_ref": "TemplateRef", "template_repo": {"branch": "Branch", "full_url": "FullURL", "has_uploadedgitrepotar": false, "release": "Release", "repo_sha_value": "RepoShaValue", "repo_url": "RepoURL", "url": "URL"}, "type": ["Type"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "workspace_status": {"frozen": true, "frozen_at": "2019-01-01T12:00:00", "frozen_by": "FrozenBy", "locked": true, "locked_by": "LockedBy", "locked_time": "2019-01-01T12:00:00"}, "workspace_status_msg": {"status_code": "StatusCode", "status_msg": "StatusMsg"}}`)
				}))
			})
			It(`Invoke UpdateWorkspace successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.UpdateWorkspace(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(schematicsv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.CatalogRef = catalogRefModel
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				updateWorkspaceOptionsModel.Tags = []string{"testString"}
				updateWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				updateWorkspaceOptionsModel.TemplateRepo = templateRepoUpdateRequestModel
				updateWorkspaceOptionsModel.Type = []string{"testString"}
				updateWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusUpdateRequestModel
				updateWorkspaceOptionsModel.WorkspaceStatusMsg = workspaceStatusMessageModel
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.UpdateWorkspaceWithContext(ctx, updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.UpdateWorkspaceWithContext(ctx, updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateWorkspace with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")

				// Construct an instance of the UpdateWorkspaceOptions model
				updateWorkspaceOptionsModel := new(schematicsv1.UpdateWorkspaceOptions)
				updateWorkspaceOptionsModel.WID = core.StringPtr("testString")
				updateWorkspaceOptionsModel.CatalogRef = catalogRefModel
				updateWorkspaceOptionsModel.Description = core.StringPtr("testString")
				updateWorkspaceOptionsModel.Name = core.StringPtr("testString")
				updateWorkspaceOptionsModel.SharedData = sharedTargetDataModel
				updateWorkspaceOptionsModel.Tags = []string{"testString"}
				updateWorkspaceOptionsModel.TemplateData = []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}
				updateWorkspaceOptionsModel.TemplateRepo = templateRepoUpdateRequestModel
				updateWorkspaceOptionsModel.Type = []string{"testString"}
				updateWorkspaceOptionsModel.WorkspaceStatus = workspaceStatusUpdateRequestModel
				updateWorkspaceOptionsModel.WorkspaceStatusMsg = workspaceStatusMessageModel
				updateWorkspaceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.UpdateWorkspace(updateWorkspaceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateWorkspaceOptions model with no property values
				updateWorkspaceOptionsModelNew := new(schematicsv1.UpdateWorkspaceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.UpdateWorkspace(updateWorkspaceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadTemplateTar(uploadTemplateTarOptions *UploadTemplateTarOptions) - Operation response error`, func() {
		uploadTemplateTarPath := "/v1/workspaces/testString/template_data/testString/template_repo_upload"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadTemplateTarPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadTemplateTar with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the UploadTemplateTarOptions model
				uploadTemplateTarOptionsModel := new(schematicsv1.UploadTemplateTarOptions)
				uploadTemplateTarOptionsModel.WID = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.TID = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.File = CreateMockReader("This is a mock file.")
				uploadTemplateTarOptionsModel.FileContentType = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.UploadTemplateTar(uploadTemplateTarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.UploadTemplateTar(uploadTemplateTarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UploadTemplateTar(uploadTemplateTarOptions *UploadTemplateTarOptions)`, func() {
		uploadTemplateTarPath := "/v1/workspaces/testString/template_data/testString/template_repo_upload"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadTemplateTarPath))
					Expect(req.Method).To(Equal("PUT"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"file_value": "FileValue", "has_received_file": false, "id": "ID"}`)
				}))
			})
			It(`Invoke UploadTemplateTar successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.UploadTemplateTar(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UploadTemplateTarOptions model
				uploadTemplateTarOptionsModel := new(schematicsv1.UploadTemplateTarOptions)
				uploadTemplateTarOptionsModel.WID = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.TID = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.File = CreateMockReader("This is a mock file.")
				uploadTemplateTarOptionsModel.FileContentType = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.UploadTemplateTar(uploadTemplateTarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.UploadTemplateTarWithContext(ctx, uploadTemplateTarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.UploadTemplateTar(uploadTemplateTarOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.UploadTemplateTarWithContext(ctx, uploadTemplateTarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UploadTemplateTar with error: Param validation error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the UploadTemplateTarOptions model
				uploadTemplateTarOptionsModel := new(schematicsv1.UploadTemplateTarOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := schematicsService.UploadTemplateTar(uploadTemplateTarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke UploadTemplateTar with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the UploadTemplateTarOptions model
				uploadTemplateTarOptionsModel := new(schematicsv1.UploadTemplateTarOptions)
				uploadTemplateTarOptionsModel.WID = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.TID = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.File = CreateMockReader("This is a mock file.")
				uploadTemplateTarOptionsModel.FileContentType = core.StringPtr("testString")
				uploadTemplateTarOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.UploadTemplateTar(uploadTemplateTarOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UploadTemplateTarOptions model with no property values
				uploadTemplateTarOptionsModelNew := new(schematicsv1.UploadTemplateTarOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.UploadTemplateTar(uploadTemplateTarOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspaceReadme(getWorkspaceReadmeOptions *GetWorkspaceReadmeOptions) - Operation response error`, func() {
		getWorkspaceReadmePath := "/v1/workspaces/testString/templates/readme"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceReadmePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["ref"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["formatted"]).To(Equal([]string{"markdown"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceReadme with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceReadmeOptions model
				getWorkspaceReadmeOptionsModel := new(schematicsv1.GetWorkspaceReadmeOptions)
				getWorkspaceReadmeOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceReadmeOptionsModel.Ref = core.StringPtr("testString")
				getWorkspaceReadmeOptionsModel.Formatted = core.StringPtr("markdown")
				getWorkspaceReadmeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceReadme(getWorkspaceReadmeOptions *GetWorkspaceReadmeOptions)`, func() {
		getWorkspaceReadmePath := "/v1/workspaces/testString/templates/readme"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceReadmePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["ref"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["formatted"]).To(Equal([]string{"markdown"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"readme": "Readme"}`)
				}))
			})
			It(`Invoke GetWorkspaceReadme successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceReadme(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceReadmeOptions model
				getWorkspaceReadmeOptionsModel := new(schematicsv1.GetWorkspaceReadmeOptions)
				getWorkspaceReadmeOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceReadmeOptionsModel.Ref = core.StringPtr("testString")
				getWorkspaceReadmeOptionsModel.Formatted = core.StringPtr("markdown")
				getWorkspaceReadmeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceReadmeWithContext(ctx, getWorkspaceReadmeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceReadmeWithContext(ctx, getWorkspaceReadmeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceReadme with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceReadmeOptions model
				getWorkspaceReadmeOptionsModel := new(schematicsv1.GetWorkspaceReadmeOptions)
				getWorkspaceReadmeOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceReadmeOptionsModel.Ref = core.StringPtr("testString")
				getWorkspaceReadmeOptionsModel.Formatted = core.StringPtr("markdown")
				getWorkspaceReadmeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceReadmeOptions model with no property values
				getWorkspaceReadmeOptionsModelNew := new(schematicsv1.GetWorkspaceReadmeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListWorkspaceActivities(listWorkspaceActivitiesOptions *ListWorkspaceActivitiesOptions) - Operation response error`, func() {
		listWorkspaceActivitiesPath := "/v1/workspaces/testString/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkspaceActivitiesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWorkspaceActivities with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListWorkspaceActivitiesOptions model
				listWorkspaceActivitiesOptionsModel := new(schematicsv1.ListWorkspaceActivitiesOptions)
				listWorkspaceActivitiesOptionsModel.WID = core.StringPtr("testString")
				listWorkspaceActivitiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listWorkspaceActivitiesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listWorkspaceActivitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListWorkspaceActivities(listWorkspaceActivitiesOptions *ListWorkspaceActivitiesOptions)`, func() {
		listWorkspaceActivitiesPath := "/v1/workspaces/testString/actions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWorkspaceActivitiesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"actions": [{"action_id": "ActionID", "message": ["Message"], "name": "Name", "performed_at": "2019-01-01T12:00:00", "performed_by": "PerformedBy", "status": "Status", "templates": [{"end_time": "2019-01-01T12:00:00", "log_summary": {"activity_status": "ActivityStatus", "detected_template_type": "DetectedTemplateType", "discarded_files": 14, "error": "Error", "resources_added": 14, "resources_destroyed": 18, "resources_modified": 17, "scanned_files": 12, "template_variable_count": 21, "time_taken": 9}, "log_url": "LogURL", "message": "Message", "start_time": "2019-01-01T12:00:00", "status": "Status", "template_id": "TemplateID", "template_type": "TemplateType"}]}], "workspace_id": "WorkspaceID", "workspace_name": "WorkspaceName"}`)
				}))
			})
			It(`Invoke ListWorkspaceActivities successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListWorkspaceActivities(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWorkspaceActivitiesOptions model
				listWorkspaceActivitiesOptionsModel := new(schematicsv1.ListWorkspaceActivitiesOptions)
				listWorkspaceActivitiesOptionsModel.WID = core.StringPtr("testString")
				listWorkspaceActivitiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listWorkspaceActivitiesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listWorkspaceActivitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListWorkspaceActivitiesWithContext(ctx, listWorkspaceActivitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListWorkspaceActivitiesWithContext(ctx, listWorkspaceActivitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListWorkspaceActivities with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListWorkspaceActivitiesOptions model
				listWorkspaceActivitiesOptionsModel := new(schematicsv1.ListWorkspaceActivitiesOptions)
				listWorkspaceActivitiesOptionsModel.WID = core.StringPtr("testString")
				listWorkspaceActivitiesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listWorkspaceActivitiesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listWorkspaceActivitiesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListWorkspaceActivitiesOptions model with no property values
				listWorkspaceActivitiesOptionsModelNew := new(schematicsv1.ListWorkspaceActivitiesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspaceActivity(getWorkspaceActivityOptions *GetWorkspaceActivityOptions) - Operation response error`, func() {
		getWorkspaceActivityPath := "/v1/workspaces/testString/actions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceActivityPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceActivity with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceActivityOptions model
				getWorkspaceActivityOptionsModel := new(schematicsv1.GetWorkspaceActivityOptions)
				getWorkspaceActivityOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceActivityOptionsModel.ActivityID = core.StringPtr("testString")
				getWorkspaceActivityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceActivity(getWorkspaceActivityOptions *GetWorkspaceActivityOptions)`, func() {
		getWorkspaceActivityPath := "/v1/workspaces/testString/actions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceActivityPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action_id": "ActionID", "message": ["Message"], "name": "Name", "performed_at": "2019-01-01T12:00:00", "performed_by": "PerformedBy", "status": "Status", "templates": [{"end_time": "2019-01-01T12:00:00", "log_summary": {"activity_status": "ActivityStatus", "detected_template_type": "DetectedTemplateType", "discarded_files": 14, "error": "Error", "resources_added": 14, "resources_destroyed": 18, "resources_modified": 17, "scanned_files": 12, "template_variable_count": 21, "time_taken": 9}, "log_url": "LogURL", "message": "Message", "start_time": "2019-01-01T12:00:00", "status": "Status", "template_id": "TemplateID", "template_type": "TemplateType"}]}`)
				}))
			})
			It(`Invoke GetWorkspaceActivity successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceActivity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceActivityOptions model
				getWorkspaceActivityOptionsModel := new(schematicsv1.GetWorkspaceActivityOptions)
				getWorkspaceActivityOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceActivityOptionsModel.ActivityID = core.StringPtr("testString")
				getWorkspaceActivityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceActivityWithContext(ctx, getWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceActivityWithContext(ctx, getWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceActivity with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceActivityOptions model
				getWorkspaceActivityOptionsModel := new(schematicsv1.GetWorkspaceActivityOptions)
				getWorkspaceActivityOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceActivityOptionsModel.ActivityID = core.StringPtr("testString")
				getWorkspaceActivityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceActivityOptions model with no property values
				getWorkspaceActivityOptionsModelNew := new(schematicsv1.GetWorkspaceActivityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteWorkspaceActivity(deleteWorkspaceActivityOptions *DeleteWorkspaceActivityOptions) - Operation response error`, func() {
		deleteWorkspaceActivityPath := "/v1/workspaces/testString/actions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkspaceActivityPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteWorkspaceActivity with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkspaceActivityOptions model
				deleteWorkspaceActivityOptionsModel := new(schematicsv1.DeleteWorkspaceActivityOptions)
				deleteWorkspaceActivityOptionsModel.WID = core.StringPtr("testString")
				deleteWorkspaceActivityOptionsModel.ActivityID = core.StringPtr("testString")
				deleteWorkspaceActivityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteWorkspaceActivity(deleteWorkspaceActivityOptions *DeleteWorkspaceActivityOptions)`, func() {
		deleteWorkspaceActivityPath := "/v1/workspaces/testString/actions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWorkspaceActivityPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"activityid": "Activityid"}`)
				}))
			})
			It(`Invoke DeleteWorkspaceActivity successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.DeleteWorkspaceActivity(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteWorkspaceActivityOptions model
				deleteWorkspaceActivityOptionsModel := new(schematicsv1.DeleteWorkspaceActivityOptions)
				deleteWorkspaceActivityOptionsModel.WID = core.StringPtr("testString")
				deleteWorkspaceActivityOptionsModel.ActivityID = core.StringPtr("testString")
				deleteWorkspaceActivityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DeleteWorkspaceActivityWithContext(ctx, deleteWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DeleteWorkspaceActivityWithContext(ctx, deleteWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteWorkspaceActivity with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the DeleteWorkspaceActivityOptions model
				deleteWorkspaceActivityOptionsModel := new(schematicsv1.DeleteWorkspaceActivityOptions)
				deleteWorkspaceActivityOptionsModel.WID = core.StringPtr("testString")
				deleteWorkspaceActivityOptionsModel.ActivityID = core.StringPtr("testString")
				deleteWorkspaceActivityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteWorkspaceActivityOptions model with no property values
				deleteWorkspaceActivityOptionsModelNew := new(schematicsv1.DeleteWorkspaceActivityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.DeleteWorkspaceActivity(deleteWorkspaceActivityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RunWorkspaceCommands(runWorkspaceCommandsOptions *RunWorkspaceCommandsOptions) - Operation response error`, func() {
		runWorkspaceCommandsPath := "/v1/workspaces/testString/commands"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runWorkspaceCommandsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunWorkspaceCommands with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the TerraformCommand model
				terraformCommandModel := new(schematicsv1.TerraformCommand)
				terraformCommandModel.Command = core.StringPtr("testString")
				terraformCommandModel.CommandParams = core.StringPtr("testString")
				terraformCommandModel.CommandName = core.StringPtr("testString")
				terraformCommandModel.CommandDesc = core.StringPtr("testString")
				terraformCommandModel.CommandOnError = core.StringPtr("testString")
				terraformCommandModel.CommandDependsOn = core.StringPtr("testString")
				terraformCommandModel.CommandStatus = core.StringPtr("testString")

				// Construct an instance of the RunWorkspaceCommandsOptions model
				runWorkspaceCommandsOptionsModel := new(schematicsv1.RunWorkspaceCommandsOptions)
				runWorkspaceCommandsOptionsModel.WID = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.RefreshToken = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Commands = []schematicsv1.TerraformCommand{*terraformCommandModel}
				runWorkspaceCommandsOptionsModel.OperationName = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Description = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RunWorkspaceCommands(runWorkspaceCommandsOptions *RunWorkspaceCommandsOptions)`, func() {
		runWorkspaceCommandsPath := "/v1/workspaces/testString/commands"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(runWorkspaceCommandsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"activityid": "Activityid"}`)
				}))
			})
			It(`Invoke RunWorkspaceCommands successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.RunWorkspaceCommands(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TerraformCommand model
				terraformCommandModel := new(schematicsv1.TerraformCommand)
				terraformCommandModel.Command = core.StringPtr("testString")
				terraformCommandModel.CommandParams = core.StringPtr("testString")
				terraformCommandModel.CommandName = core.StringPtr("testString")
				terraformCommandModel.CommandDesc = core.StringPtr("testString")
				terraformCommandModel.CommandOnError = core.StringPtr("testString")
				terraformCommandModel.CommandDependsOn = core.StringPtr("testString")
				terraformCommandModel.CommandStatus = core.StringPtr("testString")

				// Construct an instance of the RunWorkspaceCommandsOptions model
				runWorkspaceCommandsOptionsModel := new(schematicsv1.RunWorkspaceCommandsOptions)
				runWorkspaceCommandsOptionsModel.WID = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.RefreshToken = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Commands = []schematicsv1.TerraformCommand{*terraformCommandModel}
				runWorkspaceCommandsOptionsModel.OperationName = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Description = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.RunWorkspaceCommandsWithContext(ctx, runWorkspaceCommandsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.RunWorkspaceCommandsWithContext(ctx, runWorkspaceCommandsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke RunWorkspaceCommands with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the TerraformCommand model
				terraformCommandModel := new(schematicsv1.TerraformCommand)
				terraformCommandModel.Command = core.StringPtr("testString")
				terraformCommandModel.CommandParams = core.StringPtr("testString")
				terraformCommandModel.CommandName = core.StringPtr("testString")
				terraformCommandModel.CommandDesc = core.StringPtr("testString")
				terraformCommandModel.CommandOnError = core.StringPtr("testString")
				terraformCommandModel.CommandDependsOn = core.StringPtr("testString")
				terraformCommandModel.CommandStatus = core.StringPtr("testString")

				// Construct an instance of the RunWorkspaceCommandsOptions model
				runWorkspaceCommandsOptionsModel := new(schematicsv1.RunWorkspaceCommandsOptions)
				runWorkspaceCommandsOptionsModel.WID = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.RefreshToken = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Commands = []schematicsv1.TerraformCommand{*terraformCommandModel}
				runWorkspaceCommandsOptionsModel.OperationName = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Description = core.StringPtr("testString")
				runWorkspaceCommandsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RunWorkspaceCommandsOptions model with no property values
				runWorkspaceCommandsOptionsModelNew := new(schematicsv1.RunWorkspaceCommandsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ApplyWorkspaceCommand(applyWorkspaceCommandOptions *ApplyWorkspaceCommandOptions) - Operation response error`, func() {
		applyWorkspaceCommandPath := "/v1/workspaces/testString/apply"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(applyWorkspaceCommandPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ApplyWorkspaceCommand with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}

				// Construct an instance of the ApplyWorkspaceCommandOptions model
				applyWorkspaceCommandOptionsModel := new(schematicsv1.ApplyWorkspaceCommandOptions)
				applyWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				applyWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				applyWorkspaceCommandOptionsModel.ActionOptions = workspaceActivityOptionsTemplateModel
				applyWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ApplyWorkspaceCommand(applyWorkspaceCommandOptions *ApplyWorkspaceCommandOptions)`, func() {
		applyWorkspaceCommandPath := "/v1/workspaces/testString/apply"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(applyWorkspaceCommandPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"activityid": "Activityid"}`)
				}))
			})
			It(`Invoke ApplyWorkspaceCommand successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ApplyWorkspaceCommand(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}

				// Construct an instance of the ApplyWorkspaceCommandOptions model
				applyWorkspaceCommandOptionsModel := new(schematicsv1.ApplyWorkspaceCommandOptions)
				applyWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				applyWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				applyWorkspaceCommandOptionsModel.ActionOptions = workspaceActivityOptionsTemplateModel
				applyWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ApplyWorkspaceCommandWithContext(ctx, applyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ApplyWorkspaceCommandWithContext(ctx, applyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ApplyWorkspaceCommand with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}

				// Construct an instance of the ApplyWorkspaceCommandOptions model
				applyWorkspaceCommandOptionsModel := new(schematicsv1.ApplyWorkspaceCommandOptions)
				applyWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				applyWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				applyWorkspaceCommandOptionsModel.ActionOptions = workspaceActivityOptionsTemplateModel
				applyWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ApplyWorkspaceCommandOptions model with no property values
				applyWorkspaceCommandOptionsModelNew := new(schematicsv1.ApplyWorkspaceCommandOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DestroyWorkspaceCommand(destroyWorkspaceCommandOptions *DestroyWorkspaceCommandOptions) - Operation response error`, func() {
		destroyWorkspaceCommandPath := "/v1/workspaces/testString/destroy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(destroyWorkspaceCommandPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DestroyWorkspaceCommand with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}

				// Construct an instance of the DestroyWorkspaceCommandOptions model
				destroyWorkspaceCommandOptionsModel := new(schematicsv1.DestroyWorkspaceCommandOptions)
				destroyWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				destroyWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				destroyWorkspaceCommandOptionsModel.ActionOptions = workspaceActivityOptionsTemplateModel
				destroyWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DestroyWorkspaceCommand(destroyWorkspaceCommandOptions *DestroyWorkspaceCommandOptions)`, func() {
		destroyWorkspaceCommandPath := "/v1/workspaces/testString/destroy"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(destroyWorkspaceCommandPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"activityid": "Activityid"}`)
				}))
			})
			It(`Invoke DestroyWorkspaceCommand successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.DestroyWorkspaceCommand(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}

				// Construct an instance of the DestroyWorkspaceCommandOptions model
				destroyWorkspaceCommandOptionsModel := new(schematicsv1.DestroyWorkspaceCommandOptions)
				destroyWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				destroyWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				destroyWorkspaceCommandOptionsModel.ActionOptions = workspaceActivityOptionsTemplateModel
				destroyWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DestroyWorkspaceCommandWithContext(ctx, destroyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DestroyWorkspaceCommandWithContext(ctx, destroyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DestroyWorkspaceCommand with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}

				// Construct an instance of the DestroyWorkspaceCommandOptions model
				destroyWorkspaceCommandOptionsModel := new(schematicsv1.DestroyWorkspaceCommandOptions)
				destroyWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				destroyWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				destroyWorkspaceCommandOptionsModel.ActionOptions = workspaceActivityOptionsTemplateModel
				destroyWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DestroyWorkspaceCommandOptions model with no property values
				destroyWorkspaceCommandOptionsModelNew := new(schematicsv1.DestroyWorkspaceCommandOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PlanWorkspaceCommand(planWorkspaceCommandOptions *PlanWorkspaceCommandOptions) - Operation response error`, func() {
		planWorkspaceCommandPath := "/v1/workspaces/testString/plan"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(planWorkspaceCommandPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PlanWorkspaceCommand with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the PlanWorkspaceCommandOptions model
				planWorkspaceCommandOptionsModel := new(schematicsv1.PlanWorkspaceCommandOptions)
				planWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				planWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				planWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`PlanWorkspaceCommand(planWorkspaceCommandOptions *PlanWorkspaceCommandOptions)`, func() {
		planWorkspaceCommandPath := "/v1/workspaces/testString/plan"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(planWorkspaceCommandPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"activityid": "Activityid"}`)
				}))
			})
			It(`Invoke PlanWorkspaceCommand successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.PlanWorkspaceCommand(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PlanWorkspaceCommandOptions model
				planWorkspaceCommandOptionsModel := new(schematicsv1.PlanWorkspaceCommandOptions)
				planWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				planWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				planWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.PlanWorkspaceCommandWithContext(ctx, planWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.PlanWorkspaceCommandWithContext(ctx, planWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke PlanWorkspaceCommand with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the PlanWorkspaceCommandOptions model
				planWorkspaceCommandOptionsModel := new(schematicsv1.PlanWorkspaceCommandOptions)
				planWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				planWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				planWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PlanWorkspaceCommandOptions model with no property values
				planWorkspaceCommandOptionsModelNew := new(schematicsv1.PlanWorkspaceCommandOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RefreshWorkspaceCommand(refreshWorkspaceCommandOptions *RefreshWorkspaceCommandOptions) - Operation response error`, func() {
		refreshWorkspaceCommandPath := "/v1/workspaces/testString/refresh"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(refreshWorkspaceCommandPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RefreshWorkspaceCommand with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the RefreshWorkspaceCommandOptions model
				refreshWorkspaceCommandOptionsModel := new(schematicsv1.RefreshWorkspaceCommandOptions)
				refreshWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				refreshWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				refreshWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RefreshWorkspaceCommand(refreshWorkspaceCommandOptions *RefreshWorkspaceCommandOptions)`, func() {
		refreshWorkspaceCommandPath := "/v1/workspaces/testString/refresh"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(refreshWorkspaceCommandPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"activityid": "Activityid"}`)
				}))
			})
			It(`Invoke RefreshWorkspaceCommand successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.RefreshWorkspaceCommand(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RefreshWorkspaceCommandOptions model
				refreshWorkspaceCommandOptionsModel := new(schematicsv1.RefreshWorkspaceCommandOptions)
				refreshWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				refreshWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				refreshWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.RefreshWorkspaceCommandWithContext(ctx, refreshWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.RefreshWorkspaceCommandWithContext(ctx, refreshWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke RefreshWorkspaceCommand with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the RefreshWorkspaceCommandOptions model
				refreshWorkspaceCommandOptionsModel := new(schematicsv1.RefreshWorkspaceCommandOptions)
				refreshWorkspaceCommandOptionsModel.WID = core.StringPtr("testString")
				refreshWorkspaceCommandOptionsModel.RefreshToken = core.StringPtr("testString")
				refreshWorkspaceCommandOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RefreshWorkspaceCommandOptions model with no property values
				refreshWorkspaceCommandOptionsModelNew := new(schematicsv1.RefreshWorkspaceCommandOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetWorkspaceInputs(getWorkspaceInputsOptions *GetWorkspaceInputsOptions) - Operation response error`, func() {
		getWorkspaceInputsPath := "/v1/workspaces/testString/template_data/testString/values"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceInputsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceInputs with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceInputsOptions model
				getWorkspaceInputsOptionsModel := new(schematicsv1.GetWorkspaceInputsOptions)
				getWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceInputsOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceInputs(getWorkspaceInputsOptions *GetWorkspaceInputsOptions)`, func() {
		getWorkspaceInputsPath := "/v1/workspaces/testString/template_data/testString/values"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceInputsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"values_metadata": [{"anyKey": "anyValue"}]}`)
				}))
			})
			It(`Invoke GetWorkspaceInputs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceInputs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceInputsOptions model
				getWorkspaceInputsOptionsModel := new(schematicsv1.GetWorkspaceInputsOptions)
				getWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceInputsOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceInputsWithContext(ctx, getWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceInputsWithContext(ctx, getWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceInputs with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceInputsOptions model
				getWorkspaceInputsOptionsModel := new(schematicsv1.GetWorkspaceInputsOptions)
				getWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceInputsOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceInputsOptions model with no property values
				getWorkspaceInputsOptionsModelNew := new(schematicsv1.GetWorkspaceInputsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceWorkspaceInputs(replaceWorkspaceInputsOptions *ReplaceWorkspaceInputsOptions) - Operation response error`, func() {
		replaceWorkspaceInputsPath := "/v1/workspaces/testString/template_data/testString/values"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceWorkspaceInputsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceWorkspaceInputs with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceWorkspaceInputsOptions model
				replaceWorkspaceInputsOptionsModel := new(schematicsv1.ReplaceWorkspaceInputsOptions)
				replaceWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.TID = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				replaceWorkspaceInputsOptionsModel.Values = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}
				replaceWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceWorkspaceInputs(replaceWorkspaceInputsOptions *ReplaceWorkspaceInputsOptions)`, func() {
		replaceWorkspaceInputsPath := "/v1/workspaces/testString/template_data/testString/values"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceWorkspaceInputsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"env_values": [{"anyKey": "anyValue"}], "values": "Values", "variablestore": [{"description": "Description", "name": "Name", "secure": true, "type": "Type", "value": "Value"}]}`)
				}))
			})
			It(`Invoke ReplaceWorkspaceInputs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ReplaceWorkspaceInputs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceWorkspaceInputsOptions model
				replaceWorkspaceInputsOptionsModel := new(schematicsv1.ReplaceWorkspaceInputsOptions)
				replaceWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.TID = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				replaceWorkspaceInputsOptionsModel.Values = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}
				replaceWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceWorkspaceInputsWithContext(ctx, replaceWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceWorkspaceInputsWithContext(ctx, replaceWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceWorkspaceInputs with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceWorkspaceInputsOptions model
				replaceWorkspaceInputsOptionsModel := new(schematicsv1.ReplaceWorkspaceInputsOptions)
				replaceWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.TID = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				replaceWorkspaceInputsOptionsModel.Values = core.StringPtr("testString")
				replaceWorkspaceInputsOptionsModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}
				replaceWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceWorkspaceInputsOptions model with no property values
				replaceWorkspaceInputsOptionsModelNew := new(schematicsv1.ReplaceWorkspaceInputsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAllWorkspaceInputs(getAllWorkspaceInputsOptions *GetAllWorkspaceInputsOptions) - Operation response error`, func() {
		getAllWorkspaceInputsPath := "/v1/workspaces/testString/templates/values"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllWorkspaceInputsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAllWorkspaceInputs with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetAllWorkspaceInputsOptions model
				getAllWorkspaceInputsOptionsModel := new(schematicsv1.GetAllWorkspaceInputsOptions)
				getAllWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				getAllWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAllWorkspaceInputs(getAllWorkspaceInputsOptions *GetAllWorkspaceInputsOptions)`, func() {
		getAllWorkspaceInputsPath := "/v1/workspaces/testString/templates/values"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAllWorkspaceInputsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"runtime_data": [{"engine_cmd": "EngineCmd", "engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "log_store_url": "LogStoreURL", "output_values": [{"anyKey": "anyValue"}], "resources": [[{"anyKey": "anyValue"}]], "state_store_url": "StateStoreURL"}], "shared_data": {"cluster_created_on": "ClusterCreatedOn", "cluster_id": "ClusterID", "cluster_name": "ClusterName", "cluster_type": "ClusterType", "entitlement_keys": [{"anyKey": "anyValue"}], "namespace": "Namespace", "region": "Region", "resource_group_id": "ResourceGroupID", "worker_count": 11, "worker_machine_type": "WorkerMachineType"}, "template_data": [{"env_values": [{"hidden": true, "name": "Name", "secure": true, "value": "Value"}], "folder": "Folder", "has_githubtoken": true, "id": "ID", "template_type": "TemplateType", "uninstall_script_name": "UninstallScriptName", "values": "Values", "values_metadata": [{"anyKey": "anyValue"}], "values_url": "ValuesURL", "variablestore": [{"description": "Description", "name": "Name", "secure": true, "type": "Type", "value": "Value"}]}]}`)
				}))
			})
			It(`Invoke GetAllWorkspaceInputs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetAllWorkspaceInputs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAllWorkspaceInputsOptions model
				getAllWorkspaceInputsOptionsModel := new(schematicsv1.GetAllWorkspaceInputsOptions)
				getAllWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				getAllWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetAllWorkspaceInputsWithContext(ctx, getAllWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetAllWorkspaceInputsWithContext(ctx, getAllWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAllWorkspaceInputs with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetAllWorkspaceInputsOptions model
				getAllWorkspaceInputsOptionsModel := new(schematicsv1.GetAllWorkspaceInputsOptions)
				getAllWorkspaceInputsOptionsModel.WID = core.StringPtr("testString")
				getAllWorkspaceInputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAllWorkspaceInputsOptions model with no property values
				getAllWorkspaceInputsOptionsModelNew := new(schematicsv1.GetAllWorkspaceInputsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptions *GetWorkspaceInputMetadataOptions)`, func() {
		getWorkspaceInputMetadataPath := "/v1/workspaces/testString/template_data/testString/values_metadata"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceInputMetadataPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"anyKey": "anyValue"}]`)
				}))
			})
			It(`Invoke GetWorkspaceInputMetadata successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceInputMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceInputMetadataOptions model
				getWorkspaceInputMetadataOptionsModel := new(schematicsv1.GetWorkspaceInputMetadataOptions)
				getWorkspaceInputMetadataOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceInputMetadataOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceInputMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceInputMetadataWithContext(ctx, getWorkspaceInputMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceInputMetadataWithContext(ctx, getWorkspaceInputMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceInputMetadata with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceInputMetadataOptions model
				getWorkspaceInputMetadataOptionsModel := new(schematicsv1.GetWorkspaceInputMetadataOptions)
				getWorkspaceInputMetadataOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceInputMetadataOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceInputMetadataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceInputMetadataOptions model with no property values
				getWorkspaceInputMetadataOptionsModelNew := new(schematicsv1.GetWorkspaceInputMetadataOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetWorkspaceOutputs(getWorkspaceOutputsOptions *GetWorkspaceOutputsOptions) - Operation response error`, func() {
		getWorkspaceOutputsPath := "/v1/workspaces/testString/output_values"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceOutputsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceOutputs with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceOutputsOptions model
				getWorkspaceOutputsOptionsModel := new(schematicsv1.GetWorkspaceOutputsOptions)
				getWorkspaceOutputsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceOutputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceOutputs(getWorkspaceOutputsOptions *GetWorkspaceOutputsOptions)`, func() {
		getWorkspaceOutputsPath := "/v1/workspaces/testString/output_values"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceOutputsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"folder": "Folder", "id": "ID", "output_values": [{"anyKey": "anyValue"}], "value_type": "ValueType"}]`)
				}))
			})
			It(`Invoke GetWorkspaceOutputs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceOutputs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceOutputsOptions model
				getWorkspaceOutputsOptionsModel := new(schematicsv1.GetWorkspaceOutputsOptions)
				getWorkspaceOutputsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceOutputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceOutputsWithContext(ctx, getWorkspaceOutputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceOutputsWithContext(ctx, getWorkspaceOutputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceOutputs with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceOutputsOptions model
				getWorkspaceOutputsOptionsModel := new(schematicsv1.GetWorkspaceOutputsOptions)
				getWorkspaceOutputsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceOutputsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceOutputsOptions model with no property values
				getWorkspaceOutputsOptionsModelNew := new(schematicsv1.GetWorkspaceOutputsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspaceResources(getWorkspaceResourcesOptions *GetWorkspaceResourcesOptions) - Operation response error`, func() {
		getWorkspaceResourcesPath := "/v1/workspaces/testString/resources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceResourcesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceResources with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceResourcesOptions model
				getWorkspaceResourcesOptionsModel := new(schematicsv1.GetWorkspaceResourcesOptions)
				getWorkspaceResourcesOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceResources(getWorkspaceResourcesOptions *GetWorkspaceResourcesOptions)`, func() {
		getWorkspaceResourcesPath := "/v1/workspaces/testString/resources"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"folder": "Folder", "id": "ID", "null_resources": [{"anyKey": "anyValue"}], "related_resources": [{"anyKey": "anyValue"}], "resources": [{"anyKey": "anyValue"}], "resources_count": 14, "template_type": "TemplateType"}]`)
				}))
			})
			It(`Invoke GetWorkspaceResources successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceResourcesOptions model
				getWorkspaceResourcesOptionsModel := new(schematicsv1.GetWorkspaceResourcesOptions)
				getWorkspaceResourcesOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceResourcesWithContext(ctx, getWorkspaceResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceResourcesWithContext(ctx, getWorkspaceResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceResources with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceResourcesOptions model
				getWorkspaceResourcesOptionsModel := new(schematicsv1.GetWorkspaceResourcesOptions)
				getWorkspaceResourcesOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceResourcesOptions model with no property values
				getWorkspaceResourcesOptionsModelNew := new(schematicsv1.GetWorkspaceResourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspaceState(getWorkspaceStateOptions *GetWorkspaceStateOptions) - Operation response error`, func() {
		getWorkspaceStatePath := "/v1/workspaces/testString/state_stores"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceState with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceStateOptions model
				getWorkspaceStateOptionsModel := new(schematicsv1.GetWorkspaceStateOptions)
				getWorkspaceStateOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceState(getWorkspaceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceState(getWorkspaceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceState(getWorkspaceStateOptions *GetWorkspaceStateOptions)`, func() {
		getWorkspaceStatePath := "/v1/workspaces/testString/state_stores"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceStatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"runtime_data": [{"engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "state_store_url": "StateStoreURL"}]}`)
				}))
			})
			It(`Invoke GetWorkspaceState successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceStateOptions model
				getWorkspaceStateOptionsModel := new(schematicsv1.GetWorkspaceStateOptions)
				getWorkspaceStateOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceState(getWorkspaceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceStateWithContext(ctx, getWorkspaceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceState(getWorkspaceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceStateWithContext(ctx, getWorkspaceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceState with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceStateOptions model
				getWorkspaceStateOptionsModel := new(schematicsv1.GetWorkspaceStateOptions)
				getWorkspaceStateOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceState(getWorkspaceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceStateOptions model with no property values
				getWorkspaceStateOptionsModelNew := new(schematicsv1.GetWorkspaceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceState(getWorkspaceStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions *GetWorkspaceTemplateStateOptions) - Operation response error`, func() {
		getWorkspaceTemplateStatePath := "/v1/workspaces/testString/runtime_data/testString/state_store"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceTemplateStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceTemplateState with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceTemplateStateOptions model
				getWorkspaceTemplateStateOptionsModel := new(schematicsv1.GetWorkspaceTemplateStateOptions)
				getWorkspaceTemplateStateOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceTemplateStateOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceTemplateStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions *GetWorkspaceTemplateStateOptions)`, func() {
		getWorkspaceTemplateStatePath := "/v1/workspaces/testString/runtime_data/testString/state_store"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceTemplateStatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"version": 7, "terraform_version": "TerraformVersion", "serial": 6, "lineage": "Lineage", "modules": [{"anyKey": "anyValue"}]}`)
				}))
			})
			It(`Invoke GetWorkspaceTemplateState successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceTemplateState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceTemplateStateOptions model
				getWorkspaceTemplateStateOptionsModel := new(schematicsv1.GetWorkspaceTemplateStateOptions)
				getWorkspaceTemplateStateOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceTemplateStateOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceTemplateStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceTemplateStateWithContext(ctx, getWorkspaceTemplateStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceTemplateStateWithContext(ctx, getWorkspaceTemplateStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceTemplateState with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceTemplateStateOptions model
				getWorkspaceTemplateStateOptionsModel := new(schematicsv1.GetWorkspaceTemplateStateOptions)
				getWorkspaceTemplateStateOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceTemplateStateOptionsModel.TID = core.StringPtr("testString")
				getWorkspaceTemplateStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceTemplateStateOptions model with no property values
				getWorkspaceTemplateStateOptionsModelNew := new(schematicsv1.GetWorkspaceTemplateStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptions *GetWorkspaceActivityLogsOptions) - Operation response error`, func() {
		getWorkspaceActivityLogsPath := "/v1/workspaces/testString/actions/testString/logs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceActivityLogsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceActivityLogs with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceActivityLogsOptions model
				getWorkspaceActivityLogsOptionsModel := new(schematicsv1.GetWorkspaceActivityLogsOptions)
				getWorkspaceActivityLogsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceActivityLogsOptionsModel.ActivityID = core.StringPtr("testString")
				getWorkspaceActivityLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptions *GetWorkspaceActivityLogsOptions)`, func() {
		getWorkspaceActivityLogsPath := "/v1/workspaces/testString/actions/testString/logs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceActivityLogsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"action_id": "ActionID", "name": "Name", "templates": [{"log_url": "LogURL", "template_id": "TemplateID", "template_type": "TemplateType"}]}`)
				}))
			})
			It(`Invoke GetWorkspaceActivityLogs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceActivityLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceActivityLogsOptions model
				getWorkspaceActivityLogsOptionsModel := new(schematicsv1.GetWorkspaceActivityLogsOptions)
				getWorkspaceActivityLogsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceActivityLogsOptionsModel.ActivityID = core.StringPtr("testString")
				getWorkspaceActivityLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceActivityLogsWithContext(ctx, getWorkspaceActivityLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceActivityLogsWithContext(ctx, getWorkspaceActivityLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceActivityLogs with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceActivityLogsOptions model
				getWorkspaceActivityLogsOptionsModel := new(schematicsv1.GetWorkspaceActivityLogsOptions)
				getWorkspaceActivityLogsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceActivityLogsOptionsModel.ActivityID = core.StringPtr("testString")
				getWorkspaceActivityLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceActivityLogsOptions model with no property values
				getWorkspaceActivityLogsOptionsModelNew := new(schematicsv1.GetWorkspaceActivityLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspaceLogUrls(getWorkspaceLogUrlsOptions *GetWorkspaceLogUrlsOptions) - Operation response error`, func() {
		getWorkspaceLogUrlsPath := "/v1/workspaces/testString/log_stores"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceLogUrlsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceLogUrls with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceLogUrlsOptions model
				getWorkspaceLogUrlsOptionsModel := new(schematicsv1.GetWorkspaceLogUrlsOptions)
				getWorkspaceLogUrlsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceLogUrlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceLogUrls(getWorkspaceLogUrlsOptions *GetWorkspaceLogUrlsOptions)`, func() {
		getWorkspaceLogUrlsPath := "/v1/workspaces/testString/log_stores"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceLogUrlsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"runtime_data": [{"engine_name": "EngineName", "engine_version": "EngineVersion", "id": "ID", "log_store_url": "LogStoreURL"}]}`)
				}))
			})
			It(`Invoke GetWorkspaceLogUrls successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceLogUrls(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceLogUrlsOptions model
				getWorkspaceLogUrlsOptionsModel := new(schematicsv1.GetWorkspaceLogUrlsOptions)
				getWorkspaceLogUrlsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceLogUrlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceLogUrlsWithContext(ctx, getWorkspaceLogUrlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceLogUrlsWithContext(ctx, getWorkspaceLogUrlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceLogUrls with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceLogUrlsOptions model
				getWorkspaceLogUrlsOptionsModel := new(schematicsv1.GetWorkspaceLogUrlsOptions)
				getWorkspaceLogUrlsOptionsModel.WID = core.StringPtr("testString")
				getWorkspaceLogUrlsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceLogUrlsOptions model with no property values
				getWorkspaceLogUrlsOptionsModelNew := new(schematicsv1.GetWorkspaceLogUrlsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTemplateLogs(getTemplateLogsOptions *GetTemplateLogsOptions)`, func() {
		getTemplateLogsPath := "/v1/workspaces/testString/runtime_data/testString/log_store"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplateLogsPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for log_tf_cmd query parameter


					// TODO: Add check for log_tf_prefix query parameter


					// TODO: Add check for log_tf_null_resource query parameter


					// TODO: Add check for log_tf_ansible query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke GetTemplateLogs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetTemplateLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTemplateLogsOptions model
				getTemplateLogsOptionsModel := new(schematicsv1.GetTemplateLogsOptions)
				getTemplateLogsOptionsModel.WID = core.StringPtr("testString")
				getTemplateLogsOptionsModel.TID = core.StringPtr("testString")
				getTemplateLogsOptionsModel.LogTfCmd = core.BoolPtr(true)
				getTemplateLogsOptionsModel.LogTfPrefix = core.BoolPtr(true)
				getTemplateLogsOptionsModel.LogTfNullResource = core.BoolPtr(true)
				getTemplateLogsOptionsModel.LogTfAnsible = core.BoolPtr(true)
				getTemplateLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetTemplateLogs(getTemplateLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetTemplateLogsWithContext(ctx, getTemplateLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetTemplateLogs(getTemplateLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetTemplateLogsWithContext(ctx, getTemplateLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetTemplateLogs with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateLogsOptions model
				getTemplateLogsOptionsModel := new(schematicsv1.GetTemplateLogsOptions)
				getTemplateLogsOptionsModel.WID = core.StringPtr("testString")
				getTemplateLogsOptionsModel.TID = core.StringPtr("testString")
				getTemplateLogsOptionsModel.LogTfCmd = core.BoolPtr(true)
				getTemplateLogsOptionsModel.LogTfPrefix = core.BoolPtr(true)
				getTemplateLogsOptionsModel.LogTfNullResource = core.BoolPtr(true)
				getTemplateLogsOptionsModel.LogTfAnsible = core.BoolPtr(true)
				getTemplateLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetTemplateLogs(getTemplateLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTemplateLogsOptions model with no property values
				getTemplateLogsOptionsModelNew := new(schematicsv1.GetTemplateLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetTemplateLogs(getTemplateLogsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetTemplateActivityLog(getTemplateActivityLogOptions *GetTemplateActivityLogOptions)`, func() {
		getTemplateActivityLogPath := "/v1/workspaces/testString/runtime_data/testString/log_store/actions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplateActivityLogPath))
					Expect(req.Method).To(Equal("GET"))


					// TODO: Add check for log_tf_cmd query parameter


					// TODO: Add check for log_tf_prefix query parameter


					// TODO: Add check for log_tf_null_resource query parameter


					// TODO: Add check for log_tf_ansible query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `"OperationResponse"`)
				}))
			})
			It(`Invoke GetTemplateActivityLog successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetTemplateActivityLog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTemplateActivityLogOptions model
				getTemplateActivityLogOptionsModel := new(schematicsv1.GetTemplateActivityLogOptions)
				getTemplateActivityLogOptionsModel.WID = core.StringPtr("testString")
				getTemplateActivityLogOptionsModel.TID = core.StringPtr("testString")
				getTemplateActivityLogOptionsModel.ActivityID = core.StringPtr("testString")
				getTemplateActivityLogOptionsModel.LogTfCmd = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.LogTfPrefix = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.LogTfNullResource = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.LogTfAnsible = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetTemplateActivityLog(getTemplateActivityLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetTemplateActivityLogWithContext(ctx, getTemplateActivityLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetTemplateActivityLog(getTemplateActivityLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetTemplateActivityLogWithContext(ctx, getTemplateActivityLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetTemplateActivityLog with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateActivityLogOptions model
				getTemplateActivityLogOptionsModel := new(schematicsv1.GetTemplateActivityLogOptions)
				getTemplateActivityLogOptionsModel.WID = core.StringPtr("testString")
				getTemplateActivityLogOptionsModel.TID = core.StringPtr("testString")
				getTemplateActivityLogOptionsModel.ActivityID = core.StringPtr("testString")
				getTemplateActivityLogOptionsModel.LogTfCmd = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.LogTfPrefix = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.LogTfNullResource = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.LogTfAnsible = core.BoolPtr(true)
				getTemplateActivityLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetTemplateActivityLog(getTemplateActivityLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTemplateActivityLogOptions model with no property values
				getTemplateActivityLogOptionsModelNew := new(schematicsv1.GetTemplateActivityLogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetTemplateActivityLog(getTemplateActivityLogOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptions *CreateWorkspaceDeletionJobOptions) - Operation response error`, func() {
		createWorkspaceDeletionJobPath := "/v1/workspace_jobs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkspaceDeletionJobPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["destroy_resources"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateWorkspaceDeletionJob with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CreateWorkspaceDeletionJobOptions model
				createWorkspaceDeletionJobOptionsModel := new(schematicsv1.CreateWorkspaceDeletionJobOptions)
				createWorkspaceDeletionJobOptionsModel.RefreshToken = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewDeleteWorkspaces = core.BoolPtr(true)
				createWorkspaceDeletionJobOptionsModel.NewDestroyResources = core.BoolPtr(true)
				createWorkspaceDeletionJobOptionsModel.NewJob = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewVersion = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewWorkspaces = []string{"testString"}
				createWorkspaceDeletionJobOptionsModel.DestroyResources = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptions *CreateWorkspaceDeletionJobOptions)`, func() {
		createWorkspaceDeletionJobPath := "/v1/workspace_jobs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createWorkspaceDeletionJobPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["destroy_resources"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"job": "Job", "job_id": "JobID"}`)
				}))
			})
			It(`Invoke CreateWorkspaceDeletionJob successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.CreateWorkspaceDeletionJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateWorkspaceDeletionJobOptions model
				createWorkspaceDeletionJobOptionsModel := new(schematicsv1.CreateWorkspaceDeletionJobOptions)
				createWorkspaceDeletionJobOptionsModel.RefreshToken = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewDeleteWorkspaces = core.BoolPtr(true)
				createWorkspaceDeletionJobOptionsModel.NewDestroyResources = core.BoolPtr(true)
				createWorkspaceDeletionJobOptionsModel.NewJob = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewVersion = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewWorkspaces = []string{"testString"}
				createWorkspaceDeletionJobOptionsModel.DestroyResources = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateWorkspaceDeletionJobWithContext(ctx, createWorkspaceDeletionJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateWorkspaceDeletionJobWithContext(ctx, createWorkspaceDeletionJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateWorkspaceDeletionJob with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the CreateWorkspaceDeletionJobOptions model
				createWorkspaceDeletionJobOptionsModel := new(schematicsv1.CreateWorkspaceDeletionJobOptions)
				createWorkspaceDeletionJobOptionsModel.RefreshToken = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewDeleteWorkspaces = core.BoolPtr(true)
				createWorkspaceDeletionJobOptionsModel.NewDestroyResources = core.BoolPtr(true)
				createWorkspaceDeletionJobOptionsModel.NewJob = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewVersion = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.NewWorkspaces = []string{"testString"}
				createWorkspaceDeletionJobOptionsModel.DestroyResources = core.StringPtr("testString")
				createWorkspaceDeletionJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateWorkspaceDeletionJobOptions model with no property values
				createWorkspaceDeletionJobOptionsModelNew := new(schematicsv1.CreateWorkspaceDeletionJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.CreateWorkspaceDeletionJob(createWorkspaceDeletionJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptions *GetWorkspaceDeletionJobStatusOptions) - Operation response error`, func() {
		getWorkspaceDeletionJobStatusPath := "/v1/workspace_jobs/testString/status"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceDeletionJobStatusPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWorkspaceDeletionJobStatus with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceDeletionJobStatusOptions model
				getWorkspaceDeletionJobStatusOptionsModel := new(schematicsv1.GetWorkspaceDeletionJobStatusOptions)
				getWorkspaceDeletionJobStatusOptionsModel.WjID = core.StringPtr("testString")
				getWorkspaceDeletionJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptions *GetWorkspaceDeletionJobStatusOptions)`, func() {
		getWorkspaceDeletionJobStatusPath := "/v1/workspace_jobs/testString/status"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWorkspaceDeletionJobStatusPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"job_status": {"failed": ["Failed"], "in_progress": ["InProgress"], "success": ["Success"], "last_updated_on": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetWorkspaceDeletionJobStatus successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetWorkspaceDeletionJobStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWorkspaceDeletionJobStatusOptions model
				getWorkspaceDeletionJobStatusOptionsModel := new(schematicsv1.GetWorkspaceDeletionJobStatusOptions)
				getWorkspaceDeletionJobStatusOptionsModel.WjID = core.StringPtr("testString")
				getWorkspaceDeletionJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceDeletionJobStatusWithContext(ctx, getWorkspaceDeletionJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetWorkspaceDeletionJobStatusWithContext(ctx, getWorkspaceDeletionJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetWorkspaceDeletionJobStatus with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetWorkspaceDeletionJobStatusOptions model
				getWorkspaceDeletionJobStatusOptionsModel := new(schematicsv1.GetWorkspaceDeletionJobStatusOptions)
				getWorkspaceDeletionJobStatusOptionsModel.WjID = core.StringPtr("testString")
				getWorkspaceDeletionJobStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWorkspaceDeletionJobStatusOptions model with no property values
				getWorkspaceDeletionJobStatusOptionsModelNew := new(schematicsv1.GetWorkspaceDeletionJobStatusOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetWorkspaceDeletionJobStatus(getWorkspaceDeletionJobStatusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateAction(createActionOptions *CreateActionOptions) - Operation response error`, func() {
		createActionPath := "/v2/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createActionPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Github-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Github-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAction with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")

				// Construct an instance of the CreateActionOptions model
				createActionOptionsModel := new(schematicsv1.CreateActionOptions)
				createActionOptionsModel.Name = core.StringPtr("Stop Action")
				createActionOptionsModel.Description = core.StringPtr("This Action can be used to Stop the targets")
				createActionOptionsModel.Location = core.StringPtr("us_south")
				createActionOptionsModel.ResourceGroup = core.StringPtr("testString")
				createActionOptionsModel.Tags = []string{"testString"}
				createActionOptionsModel.UserState = userStateModel
				createActionOptionsModel.SourceReadmeURL = core.StringPtr("testString")
				createActionOptionsModel.Source = externalSourceModel
				createActionOptionsModel.SourceType = core.StringPtr("local")
				createActionOptionsModel.CommandParameter = core.StringPtr("testString")
				createActionOptionsModel.Bastion = targetResourcesetModel
				createActionOptionsModel.Targets = []schematicsv1.TargetResourceGroup{*targetResourceGroupModel}
				createActionOptionsModel.Credentials = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.TriggerRecordID = core.StringPtr("testString")
				createActionOptionsModel.State = actionStateModel
				createActionOptionsModel.SysLock = systemLockModel
				createActionOptionsModel.XGithubToken = core.StringPtr("testString")
				createActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.CreateAction(createActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.CreateAction(createActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateAction(createActionOptions *CreateActionOptions)`, func() {
		createActionPath := "/v2/actions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createActionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Github-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Github-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Stop Action", "description": "This Action can be used to Stop the targets", "location": "us_south", "resource_group": "ResourceGroup", "tags": ["Tags"], "user_state": {"state": "draft", "set_by": "SetBy", "set_at": "2019-01-01T12:00:00"}, "source_readme_url": "SourceReadmeURL", "source": {"source_type": "local", "git": {"git_repo_url": "GitRepoURL", "git_token": "GitToken", "git_repo_folder": "GitRepoFolder", "git_release": "GitRelease", "git_branch": "GitBranch"}}, "source_type": "local", "command_parameter": "CommandParameter", "bastion": {"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}, "targets": [{"name": "Name", "description": "Description", "credential_ref": "CredentialRef", "bastion_ref": "BastionRef", "target_resources": [{"resource_id": "ResourceID", "resource_configs": [{"name": "Name", "value": "Value", "description": "Description"}]}]}], "credentials": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "outputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "trigger_record_id": "TriggerRecordID", "id": "ID", "crn": "Crn", "account": "Account", "source_created_at": "2019-01-01T12:00:00", "source_created_by": "SourceCreatedBy", "source_updated_at": "2019-01-01T12:00:00", "source_updated_by": "SourceUpdatedBy", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "namespace": "Namespace", "state": {"status_code": "normal", "status_job_id": "StatusJobID", "status_message": "StatusMessage"}, "playbook_names": ["PlaybookNames"], "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke CreateAction successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.CreateAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")

				// Construct an instance of the CreateActionOptions model
				createActionOptionsModel := new(schematicsv1.CreateActionOptions)
				createActionOptionsModel.Name = core.StringPtr("Stop Action")
				createActionOptionsModel.Description = core.StringPtr("This Action can be used to Stop the targets")
				createActionOptionsModel.Location = core.StringPtr("us_south")
				createActionOptionsModel.ResourceGroup = core.StringPtr("testString")
				createActionOptionsModel.Tags = []string{"testString"}
				createActionOptionsModel.UserState = userStateModel
				createActionOptionsModel.SourceReadmeURL = core.StringPtr("testString")
				createActionOptionsModel.Source = externalSourceModel
				createActionOptionsModel.SourceType = core.StringPtr("local")
				createActionOptionsModel.CommandParameter = core.StringPtr("testString")
				createActionOptionsModel.Bastion = targetResourcesetModel
				createActionOptionsModel.Targets = []schematicsv1.TargetResourceGroup{*targetResourceGroupModel}
				createActionOptionsModel.Credentials = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.TriggerRecordID = core.StringPtr("testString")
				createActionOptionsModel.State = actionStateModel
				createActionOptionsModel.SysLock = systemLockModel
				createActionOptionsModel.XGithubToken = core.StringPtr("testString")
				createActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.CreateAction(createActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateActionWithContext(ctx, createActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.CreateAction(createActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateActionWithContext(ctx, createActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateAction with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")

				// Construct an instance of the CreateActionOptions model
				createActionOptionsModel := new(schematicsv1.CreateActionOptions)
				createActionOptionsModel.Name = core.StringPtr("Stop Action")
				createActionOptionsModel.Description = core.StringPtr("This Action can be used to Stop the targets")
				createActionOptionsModel.Location = core.StringPtr("us_south")
				createActionOptionsModel.ResourceGroup = core.StringPtr("testString")
				createActionOptionsModel.Tags = []string{"testString"}
				createActionOptionsModel.UserState = userStateModel
				createActionOptionsModel.SourceReadmeURL = core.StringPtr("testString")
				createActionOptionsModel.Source = externalSourceModel
				createActionOptionsModel.SourceType = core.StringPtr("local")
				createActionOptionsModel.CommandParameter = core.StringPtr("testString")
				createActionOptionsModel.Bastion = targetResourcesetModel
				createActionOptionsModel.Targets = []schematicsv1.TargetResourceGroup{*targetResourceGroupModel}
				createActionOptionsModel.Credentials = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				createActionOptionsModel.TriggerRecordID = core.StringPtr("testString")
				createActionOptionsModel.State = actionStateModel
				createActionOptionsModel.SysLock = systemLockModel
				createActionOptionsModel.XGithubToken = core.StringPtr("testString")
				createActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.CreateAction(createActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListActions(listActionsOptions *ListActionsOptions) - Operation response error`, func() {
		listActionsPath := "/v2/actions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listActionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["profile"]).To(Equal([]string{"ids"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListActions with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListActionsOptions model
				listActionsOptionsModel := new(schematicsv1.ListActionsOptions)
				listActionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listActionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listActionsOptionsModel.Sort = core.StringPtr("testString")
				listActionsOptionsModel.Profile = core.StringPtr("ids")
				listActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListActions(listActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListActions(listActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListActions(listActionsOptions *ListActionsOptions)`, func() {
		listActionsPath := "/v2/actions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listActionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["profile"]).To(Equal([]string{"ids"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "offset": 6, "actions": [{"name": "Stop Action", "description": "This Action can be used to Stop the targets", "id": "ID", "crn": "Crn", "location": "us_south", "resource_group": "ResourceGroup", "namespace": "Namespace", "tags": ["Tags"], "playbook_name": "PlaybookName", "user_state": {"state": "draft", "set_by": "SetBy", "set_at": "2019-01-01T12:00:00"}, "state": {"status_code": "normal", "status_message": "StatusMessage"}, "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListActions successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListActionsOptions model
				listActionsOptionsModel := new(schematicsv1.ListActionsOptions)
				listActionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listActionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listActionsOptionsModel.Sort = core.StringPtr("testString")
				listActionsOptionsModel.Profile = core.StringPtr("ids")
				listActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListActions(listActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListActionsWithContext(ctx, listActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListActions(listActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListActionsWithContext(ctx, listActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListActions with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListActionsOptions model
				listActionsOptionsModel := new(schematicsv1.ListActionsOptions)
				listActionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listActionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listActionsOptionsModel.Sort = core.StringPtr("testString")
				listActionsOptionsModel.Profile = core.StringPtr("ids")
				listActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListActions(listActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAction(getActionOptions *GetActionOptions) - Operation response error`, func() {
		getActionPath := "/v2/actions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getActionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["profile"]).To(Equal([]string{"summary"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAction with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetActionOptions model
				getActionOptionsModel := new(schematicsv1.GetActionOptions)
				getActionOptionsModel.ActionID = core.StringPtr("testString")
				getActionOptionsModel.Profile = core.StringPtr("summary")
				getActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetAction(getActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetAction(getActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAction(getActionOptions *GetActionOptions)`, func() {
		getActionPath := "/v2/actions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getActionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["profile"]).To(Equal([]string{"summary"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Stop Action", "description": "This Action can be used to Stop the targets", "location": "us_south", "resource_group": "ResourceGroup", "tags": ["Tags"], "user_state": {"state": "draft", "set_by": "SetBy", "set_at": "2019-01-01T12:00:00"}, "source_readme_url": "SourceReadmeURL", "source": {"source_type": "local", "git": {"git_repo_url": "GitRepoURL", "git_token": "GitToken", "git_repo_folder": "GitRepoFolder", "git_release": "GitRelease", "git_branch": "GitBranch"}}, "source_type": "local", "command_parameter": "CommandParameter", "bastion": {"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}, "targets": [{"name": "Name", "description": "Description", "credential_ref": "CredentialRef", "bastion_ref": "BastionRef", "target_resources": [{"resource_id": "ResourceID", "resource_configs": [{"name": "Name", "value": "Value", "description": "Description"}]}]}], "credentials": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "outputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "trigger_record_id": "TriggerRecordID", "id": "ID", "crn": "Crn", "account": "Account", "source_created_at": "2019-01-01T12:00:00", "source_created_by": "SourceCreatedBy", "source_updated_at": "2019-01-01T12:00:00", "source_updated_by": "SourceUpdatedBy", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "namespace": "Namespace", "state": {"status_code": "normal", "status_job_id": "StatusJobID", "status_message": "StatusMessage"}, "playbook_names": ["PlaybookNames"], "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke GetAction successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetActionOptions model
				getActionOptionsModel := new(schematicsv1.GetActionOptions)
				getActionOptionsModel.ActionID = core.StringPtr("testString")
				getActionOptionsModel.Profile = core.StringPtr("summary")
				getActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetAction(getActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetActionWithContext(ctx, getActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetAction(getActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetActionWithContext(ctx, getActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAction with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetActionOptions model
				getActionOptionsModel := new(schematicsv1.GetActionOptions)
				getActionOptionsModel.ActionID = core.StringPtr("testString")
				getActionOptionsModel.Profile = core.StringPtr("summary")
				getActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetAction(getActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetActionOptions model with no property values
				getActionOptionsModelNew := new(schematicsv1.GetActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetAction(getActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAction(deleteActionOptions *DeleteActionOptions)`, func() {
		deleteActionPath := "/v2/actions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteActionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Force"]).ToNot(BeNil())
					Expect(req.Header["Force"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.Header["Propagate"]).ToNot(BeNil())
					Expect(req.Header["Propagate"][0]).To(Equal(fmt.Sprintf("%v", true)))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAction successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := schematicsService.DeleteAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteActionOptions model
				deleteActionOptionsModel := new(schematicsv1.DeleteActionOptions)
				deleteActionOptionsModel.ActionID = core.StringPtr("testString")
				deleteActionOptionsModel.Force = core.BoolPtr(true)
				deleteActionOptionsModel.Propagate = core.BoolPtr(true)
				deleteActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = schematicsService.DeleteAction(deleteActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				schematicsService.DisableRetries()
				response, operationErr = schematicsService.DeleteAction(deleteActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAction with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the DeleteActionOptions model
				deleteActionOptionsModel := new(schematicsv1.DeleteActionOptions)
				deleteActionOptionsModel.ActionID = core.StringPtr("testString")
				deleteActionOptionsModel.Force = core.BoolPtr(true)
				deleteActionOptionsModel.Propagate = core.BoolPtr(true)
				deleteActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := schematicsService.DeleteAction(deleteActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteActionOptions model with no property values
				deleteActionOptionsModelNew := new(schematicsv1.DeleteActionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = schematicsService.DeleteAction(deleteActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAction(updateActionOptions *UpdateActionOptions) - Operation response error`, func() {
		updateActionPath := "/v2/actions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateActionPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Github-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Github-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAction with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")

				// Construct an instance of the UpdateActionOptions model
				updateActionOptionsModel := new(schematicsv1.UpdateActionOptions)
				updateActionOptionsModel.ActionID = core.StringPtr("testString")
				updateActionOptionsModel.Name = core.StringPtr("Stop Action")
				updateActionOptionsModel.Description = core.StringPtr("This Action can be used to Stop the targets")
				updateActionOptionsModel.Location = core.StringPtr("us_south")
				updateActionOptionsModel.ResourceGroup = core.StringPtr("testString")
				updateActionOptionsModel.Tags = []string{"testString"}
				updateActionOptionsModel.UserState = userStateModel
				updateActionOptionsModel.SourceReadmeURL = core.StringPtr("testString")
				updateActionOptionsModel.Source = externalSourceModel
				updateActionOptionsModel.SourceType = core.StringPtr("local")
				updateActionOptionsModel.CommandParameter = core.StringPtr("testString")
				updateActionOptionsModel.Bastion = targetResourcesetModel
				updateActionOptionsModel.Targets = []schematicsv1.TargetResourceGroup{*targetResourceGroupModel}
				updateActionOptionsModel.Credentials = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.TriggerRecordID = core.StringPtr("testString")
				updateActionOptionsModel.State = actionStateModel
				updateActionOptionsModel.SysLock = systemLockModel
				updateActionOptionsModel.XGithubToken = core.StringPtr("testString")
				updateActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.UpdateAction(updateActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.UpdateAction(updateActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAction(updateActionOptions *UpdateActionOptions)`, func() {
		updateActionPath := "/v2/actions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateActionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["X-Github-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Github-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Stop Action", "description": "This Action can be used to Stop the targets", "location": "us_south", "resource_group": "ResourceGroup", "tags": ["Tags"], "user_state": {"state": "draft", "set_by": "SetBy", "set_at": "2019-01-01T12:00:00"}, "source_readme_url": "SourceReadmeURL", "source": {"source_type": "local", "git": {"git_repo_url": "GitRepoURL", "git_token": "GitToken", "git_repo_folder": "GitRepoFolder", "git_release": "GitRelease", "git_branch": "GitBranch"}}, "source_type": "local", "command_parameter": "CommandParameter", "bastion": {"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}, "targets": [{"name": "Name", "description": "Description", "credential_ref": "CredentialRef", "bastion_ref": "BastionRef", "target_resources": [{"resource_id": "ResourceID", "resource_configs": [{"name": "Name", "value": "Value", "description": "Description"}]}]}], "credentials": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "outputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "trigger_record_id": "TriggerRecordID", "id": "ID", "crn": "Crn", "account": "Account", "source_created_at": "2019-01-01T12:00:00", "source_created_by": "SourceCreatedBy", "source_updated_at": "2019-01-01T12:00:00", "source_updated_by": "SourceUpdatedBy", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "namespace": "Namespace", "state": {"status_code": "normal", "status_job_id": "StatusJobID", "status_message": "StatusMessage"}, "playbook_names": ["PlaybookNames"], "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}}`)
				}))
			})
			It(`Invoke UpdateAction successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.UpdateAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")

				// Construct an instance of the UpdateActionOptions model
				updateActionOptionsModel := new(schematicsv1.UpdateActionOptions)
				updateActionOptionsModel.ActionID = core.StringPtr("testString")
				updateActionOptionsModel.Name = core.StringPtr("Stop Action")
				updateActionOptionsModel.Description = core.StringPtr("This Action can be used to Stop the targets")
				updateActionOptionsModel.Location = core.StringPtr("us_south")
				updateActionOptionsModel.ResourceGroup = core.StringPtr("testString")
				updateActionOptionsModel.Tags = []string{"testString"}
				updateActionOptionsModel.UserState = userStateModel
				updateActionOptionsModel.SourceReadmeURL = core.StringPtr("testString")
				updateActionOptionsModel.Source = externalSourceModel
				updateActionOptionsModel.SourceType = core.StringPtr("local")
				updateActionOptionsModel.CommandParameter = core.StringPtr("testString")
				updateActionOptionsModel.Bastion = targetResourcesetModel
				updateActionOptionsModel.Targets = []schematicsv1.TargetResourceGroup{*targetResourceGroupModel}
				updateActionOptionsModel.Credentials = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.TriggerRecordID = core.StringPtr("testString")
				updateActionOptionsModel.State = actionStateModel
				updateActionOptionsModel.SysLock = systemLockModel
				updateActionOptionsModel.XGithubToken = core.StringPtr("testString")
				updateActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.UpdateAction(updateActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.UpdateActionWithContext(ctx, updateActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.UpdateAction(updateActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.UpdateActionWithContext(ctx, updateActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateAction with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")

				// Construct an instance of the UpdateActionOptions model
				updateActionOptionsModel := new(schematicsv1.UpdateActionOptions)
				updateActionOptionsModel.ActionID = core.StringPtr("testString")
				updateActionOptionsModel.Name = core.StringPtr("Stop Action")
				updateActionOptionsModel.Description = core.StringPtr("This Action can be used to Stop the targets")
				updateActionOptionsModel.Location = core.StringPtr("us_south")
				updateActionOptionsModel.ResourceGroup = core.StringPtr("testString")
				updateActionOptionsModel.Tags = []string{"testString"}
				updateActionOptionsModel.UserState = userStateModel
				updateActionOptionsModel.SourceReadmeURL = core.StringPtr("testString")
				updateActionOptionsModel.Source = externalSourceModel
				updateActionOptionsModel.SourceType = core.StringPtr("local")
				updateActionOptionsModel.CommandParameter = core.StringPtr("testString")
				updateActionOptionsModel.Bastion = targetResourcesetModel
				updateActionOptionsModel.Targets = []schematicsv1.TargetResourceGroup{*targetResourceGroupModel}
				updateActionOptionsModel.Credentials = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				updateActionOptionsModel.TriggerRecordID = core.StringPtr("testString")
				updateActionOptionsModel.State = actionStateModel
				updateActionOptionsModel.SysLock = systemLockModel
				updateActionOptionsModel.XGithubToken = core.StringPtr("testString")
				updateActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.UpdateAction(updateActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateActionOptions model with no property values
				updateActionOptionsModelNew := new(schematicsv1.UpdateActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.UpdateAction(updateActionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateJob(createJobOptions *CreateJobOptions) - Operation response error`, func() {
		createJobPath := "/v2/jobs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateJob with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				jobStatusModel.ActionJobStatus = jobStatusActionModel

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(schematicsv1.CreateJobOptions)
				createJobOptionsModel.RefreshToken = core.StringPtr("testString")
				createJobOptionsModel.CommandObject = core.StringPtr("workspace")
				createJobOptionsModel.CommandObjectID = core.StringPtr("testString")
				createJobOptionsModel.CommandName = core.StringPtr("workspace_init_flow")
				createJobOptionsModel.CommandParameter = core.StringPtr("testString")
				createJobOptionsModel.CommandOptions = []string{"testString"}
				createJobOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				createJobOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				createJobOptionsModel.Tags = []string{"testString"}
				createJobOptionsModel.Location = core.StringPtr("us_south")
				createJobOptionsModel.Status = jobStatusModel
				createJobOptionsModel.Data = jobDataModel
				createJobOptionsModel.Bastion = targetResourcesetModel
				createJobOptionsModel.LogSummary = jobLogSummaryModel
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateJob(createJobOptions *CreateJobOptions)`, func() {
		createJobPath := "/v2/jobs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createJobPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"command_object": "workspace", "command_object_id": "CommandObjectID", "command_name": "workspace_init_flow", "command_parameter": "CommandParameter", "command_options": ["CommandOptions"], "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "tags": ["Tags"], "id": "ID", "name": "Name", "description": "Description", "location": "us_south", "resource_group": "ResourceGroup", "submitted_at": "2019-01-01T12:00:00", "submitted_by": "SubmittedBy", "start_at": "2019-01-01T12:00:00", "end_at": "2019-01-01T12:00:00", "duration": "Duration", "status": {"action_job_status": {"action_name": "ActionName", "status_code": "job_pending", "status_message": "StatusMessage", "bastion_status_code": "none", "bastion_status_message": "BastionStatusMessage", "targets_status_code": "none", "targets_status_message": "TargetsStatusMessage", "updated_at": "2019-01-01T12:00:00"}}, "data": {"job_type": "repo_download_job", "action_job_data": {"action_name": "ActionName", "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "outputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "updated_at": "2019-01-01T12:00:00"}}, "targets": [{"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}], "bastion": {"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}, "log_summary": {"job_id": "JobID", "job_type": "repo_download_job", "log_start_at": "2019-01-01T12:00:00", "log_analyzed_till": "2019-01-01T12:00:00", "elapsed_time": 11, "log_errors": [{"error_code": "ErrorCode", "error_msg": "ErrorMsg", "error_count": 10}], "repo_download_job": {"scanned_file_count": 16, "quarantined_file_count": 20, "detected_filetype": "DetectedFiletype", "inputs_count": "InputsCount", "outputs_count": "OutputsCount"}, "action_job": {"target_count": 11, "task_count": 9, "play_count": 9, "recap": {"target": ["Target"], "ok": 2, "changed": 7, "failed": 6, "skipped": 7, "unreachable": 11}}}, "log_store_url": "LogStoreURL", "state_store_url": "StateStoreURL", "results_url": "ResultsURL", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateJob successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.CreateJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				jobStatusModel.ActionJobStatus = jobStatusActionModel

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(schematicsv1.CreateJobOptions)
				createJobOptionsModel.RefreshToken = core.StringPtr("testString")
				createJobOptionsModel.CommandObject = core.StringPtr("workspace")
				createJobOptionsModel.CommandObjectID = core.StringPtr("testString")
				createJobOptionsModel.CommandName = core.StringPtr("workspace_init_flow")
				createJobOptionsModel.CommandParameter = core.StringPtr("testString")
				createJobOptionsModel.CommandOptions = []string{"testString"}
				createJobOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				createJobOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				createJobOptionsModel.Tags = []string{"testString"}
				createJobOptionsModel.Location = core.StringPtr("us_south")
				createJobOptionsModel.Status = jobStatusModel
				createJobOptionsModel.Data = jobDataModel
				createJobOptionsModel.Bastion = targetResourcesetModel
				createJobOptionsModel.LogSummary = jobLogSummaryModel
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.CreateJob(createJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateJobWithContext(ctx, createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.CreateJob(createJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateJobWithContext(ctx, createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateJob with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				jobStatusModel.ActionJobStatus = jobStatusActionModel

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel

				// Construct an instance of the CreateJobOptions model
				createJobOptionsModel := new(schematicsv1.CreateJobOptions)
				createJobOptionsModel.RefreshToken = core.StringPtr("testString")
				createJobOptionsModel.CommandObject = core.StringPtr("workspace")
				createJobOptionsModel.CommandObjectID = core.StringPtr("testString")
				createJobOptionsModel.CommandName = core.StringPtr("workspace_init_flow")
				createJobOptionsModel.CommandParameter = core.StringPtr("testString")
				createJobOptionsModel.CommandOptions = []string{"testString"}
				createJobOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				createJobOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				createJobOptionsModel.Tags = []string{"testString"}
				createJobOptionsModel.Location = core.StringPtr("us_south")
				createJobOptionsModel.Status = jobStatusModel
				createJobOptionsModel.Data = jobDataModel
				createJobOptionsModel.Bastion = targetResourcesetModel
				createJobOptionsModel.LogSummary = jobLogSummaryModel
				createJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.CreateJob(createJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateJobOptions model with no property values
				createJobOptionsModelNew := new(schematicsv1.CreateJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.CreateJob(createJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListJobs(listJobsOptions *ListJobsOptions) - Operation response error`, func() {
		listJobsPath := "/v2/jobs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["profile"]).To(Equal([]string{"ids"}))

					Expect(req.URL.Query()["resource"]).To(Equal([]string{"workspaces"}))

					Expect(req.URL.Query()["action_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["list"]).To(Equal([]string{"all"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListJobs with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(schematicsv1.ListJobsOptions)
				listJobsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listJobsOptionsModel.Sort = core.StringPtr("testString")
				listJobsOptionsModel.Profile = core.StringPtr("ids")
				listJobsOptionsModel.Resource = core.StringPtr("workspaces")
				listJobsOptionsModel.ActionID = core.StringPtr("testString")
				listJobsOptionsModel.List = core.StringPtr("all")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListJobs(listJobsOptions *ListJobsOptions)`, func() {
		listJobsPath := "/v2/jobs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(0))}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["profile"]).To(Equal([]string{"ids"}))

					Expect(req.URL.Query()["resource"]).To(Equal([]string{"workspaces"}))

					Expect(req.URL.Query()["action_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["list"]).To(Equal([]string{"all"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "offset": 6, "jobs": [{"id": "ID", "name": "Name", "description": "Description", "command_object": "workspace", "command_object_id": "CommandObjectID", "command_name": "workspace_init_flow", "tags": ["Tags"], "location": "us_south", "resource_group": "ResourceGroup", "targets": [{"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}], "submitted_at": "2019-01-01T12:00:00", "submitted_by": "SubmittedBy", "duration": "Duration", "start_at": "2019-01-01T12:00:00", "end_at": "2019-01-01T12:00:00", "status": {"action_job_status": {"action_name": "ActionName", "status_code": "job_pending", "status_message": "StatusMessage", "bastion_status_code": "none", "bastion_status_message": "BastionStatusMessage", "targets_status_code": "none", "targets_status_message": "TargetsStatusMessage", "updated_at": "2019-01-01T12:00:00"}}, "log_summary": {"job_id": "JobID", "job_type": "repo_download_job", "log_start_at": "2019-01-01T12:00:00", "log_analyzed_till": "2019-01-01T12:00:00", "elapsed_time": 11, "log_errors": [{"error_code": "ErrorCode", "error_msg": "ErrorMsg", "error_count": 10}], "repo_download_job": {"scanned_file_count": 16, "quarantined_file_count": 20, "detected_filetype": "DetectedFiletype", "inputs_count": "InputsCount", "outputs_count": "OutputsCount"}, "action_job": {"target_count": 11, "task_count": 9, "play_count": 9, "recap": {"target": ["Target"], "ok": 2, "changed": 7, "failed": 6, "skipped": 7, "unreachable": 11}}}, "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListJobs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListJobs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(schematicsv1.ListJobsOptions)
				listJobsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listJobsOptionsModel.Sort = core.StringPtr("testString")
				listJobsOptionsModel.Profile = core.StringPtr("ids")
				listJobsOptionsModel.Resource = core.StringPtr("workspaces")
				listJobsOptionsModel.ActionID = core.StringPtr("testString")
				listJobsOptionsModel.List = core.StringPtr("all")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListJobsWithContext(ctx, listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListJobsWithContext(ctx, listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListJobs with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := new(schematicsv1.ListJobsOptions)
				listJobsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listJobsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listJobsOptionsModel.Sort = core.StringPtr("testString")
				listJobsOptionsModel.Profile = core.StringPtr("ids")
				listJobsOptionsModel.Resource = core.StringPtr("workspaces")
				listJobsOptionsModel.ActionID = core.StringPtr("testString")
				listJobsOptionsModel.List = core.StringPtr("all")
				listJobsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListJobs(listJobsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceJob(replaceJobOptions *ReplaceJobOptions) - Operation response error`, func() {
		replaceJobPath := "/v2/jobs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceJobPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceJob with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				jobStatusModel.ActionJobStatus = jobStatusActionModel

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel

				// Construct an instance of the ReplaceJobOptions model
				replaceJobOptionsModel := new(schematicsv1.ReplaceJobOptions)
				replaceJobOptionsModel.JobID = core.StringPtr("testString")
				replaceJobOptionsModel.RefreshToken = core.StringPtr("testString")
				replaceJobOptionsModel.CommandObject = core.StringPtr("workspace")
				replaceJobOptionsModel.CommandObjectID = core.StringPtr("testString")
				replaceJobOptionsModel.CommandName = core.StringPtr("workspace_init_flow")
				replaceJobOptionsModel.CommandParameter = core.StringPtr("testString")
				replaceJobOptionsModel.CommandOptions = []string{"testString"}
				replaceJobOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				replaceJobOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				replaceJobOptionsModel.Tags = []string{"testString"}
				replaceJobOptionsModel.Location = core.StringPtr("us_south")
				replaceJobOptionsModel.Status = jobStatusModel
				replaceJobOptionsModel.Data = jobDataModel
				replaceJobOptionsModel.Bastion = targetResourcesetModel
				replaceJobOptionsModel.LogSummary = jobLogSummaryModel
				replaceJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ReplaceJob(replaceJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ReplaceJob(replaceJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceJob(replaceJobOptions *ReplaceJobOptions)`, func() {
		replaceJobPath := "/v2/jobs/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceJobPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"command_object": "workspace", "command_object_id": "CommandObjectID", "command_name": "workspace_init_flow", "command_parameter": "CommandParameter", "command_options": ["CommandOptions"], "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "tags": ["Tags"], "id": "ID", "name": "Name", "description": "Description", "location": "us_south", "resource_group": "ResourceGroup", "submitted_at": "2019-01-01T12:00:00", "submitted_by": "SubmittedBy", "start_at": "2019-01-01T12:00:00", "end_at": "2019-01-01T12:00:00", "duration": "Duration", "status": {"action_job_status": {"action_name": "ActionName", "status_code": "job_pending", "status_message": "StatusMessage", "bastion_status_code": "none", "bastion_status_message": "BastionStatusMessage", "targets_status_code": "none", "targets_status_message": "TargetsStatusMessage", "updated_at": "2019-01-01T12:00:00"}}, "data": {"job_type": "repo_download_job", "action_job_data": {"action_name": "ActionName", "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "outputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "updated_at": "2019-01-01T12:00:00"}}, "targets": [{"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}], "bastion": {"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}, "log_summary": {"job_id": "JobID", "job_type": "repo_download_job", "log_start_at": "2019-01-01T12:00:00", "log_analyzed_till": "2019-01-01T12:00:00", "elapsed_time": 11, "log_errors": [{"error_code": "ErrorCode", "error_msg": "ErrorMsg", "error_count": 10}], "repo_download_job": {"scanned_file_count": 16, "quarantined_file_count": 20, "detected_filetype": "DetectedFiletype", "inputs_count": "InputsCount", "outputs_count": "OutputsCount"}, "action_job": {"target_count": 11, "task_count": 9, "play_count": 9, "recap": {"target": ["Target"], "ok": 2, "changed": 7, "failed": 6, "skipped": 7, "unreachable": 11}}}, "log_store_url": "LogStoreURL", "state_store_url": "StateStoreURL", "results_url": "ResultsURL", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke ReplaceJob successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ReplaceJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				jobStatusModel.ActionJobStatus = jobStatusActionModel

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel

				// Construct an instance of the ReplaceJobOptions model
				replaceJobOptionsModel := new(schematicsv1.ReplaceJobOptions)
				replaceJobOptionsModel.JobID = core.StringPtr("testString")
				replaceJobOptionsModel.RefreshToken = core.StringPtr("testString")
				replaceJobOptionsModel.CommandObject = core.StringPtr("workspace")
				replaceJobOptionsModel.CommandObjectID = core.StringPtr("testString")
				replaceJobOptionsModel.CommandName = core.StringPtr("workspace_init_flow")
				replaceJobOptionsModel.CommandParameter = core.StringPtr("testString")
				replaceJobOptionsModel.CommandOptions = []string{"testString"}
				replaceJobOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				replaceJobOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				replaceJobOptionsModel.Tags = []string{"testString"}
				replaceJobOptionsModel.Location = core.StringPtr("us_south")
				replaceJobOptionsModel.Status = jobStatusModel
				replaceJobOptionsModel.Data = jobDataModel
				replaceJobOptionsModel.Bastion = targetResourcesetModel
				replaceJobOptionsModel.LogSummary = jobLogSummaryModel
				replaceJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ReplaceJob(replaceJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceJobWithContext(ctx, replaceJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ReplaceJob(replaceJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceJobWithContext(ctx, replaceJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceJob with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				jobStatusModel.ActionJobStatus = jobStatusActionModel

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel

				// Construct an instance of the ReplaceJobOptions model
				replaceJobOptionsModel := new(schematicsv1.ReplaceJobOptions)
				replaceJobOptionsModel.JobID = core.StringPtr("testString")
				replaceJobOptionsModel.RefreshToken = core.StringPtr("testString")
				replaceJobOptionsModel.CommandObject = core.StringPtr("workspace")
				replaceJobOptionsModel.CommandObjectID = core.StringPtr("testString")
				replaceJobOptionsModel.CommandName = core.StringPtr("workspace_init_flow")
				replaceJobOptionsModel.CommandParameter = core.StringPtr("testString")
				replaceJobOptionsModel.CommandOptions = []string{"testString"}
				replaceJobOptionsModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				replaceJobOptionsModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				replaceJobOptionsModel.Tags = []string{"testString"}
				replaceJobOptionsModel.Location = core.StringPtr("us_south")
				replaceJobOptionsModel.Status = jobStatusModel
				replaceJobOptionsModel.Data = jobDataModel
				replaceJobOptionsModel.Bastion = targetResourcesetModel
				replaceJobOptionsModel.LogSummary = jobLogSummaryModel
				replaceJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ReplaceJob(replaceJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceJobOptions model with no property values
				replaceJobOptionsModelNew := new(schematicsv1.ReplaceJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ReplaceJob(replaceJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteJob(deleteJobOptions *DeleteJobOptions)`, func() {
		deleteJobPath := "/v2/jobs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteJobPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Refresh_token"]).ToNot(BeNil())
					Expect(req.Header["Refresh_token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Force"]).ToNot(BeNil())
					Expect(req.Header["Force"][0]).To(Equal(fmt.Sprintf("%v", true)))
					Expect(req.Header["Propagate"]).ToNot(BeNil())
					Expect(req.Header["Propagate"][0]).To(Equal(fmt.Sprintf("%v", true)))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteJob successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := schematicsService.DeleteJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteJobOptions model
				deleteJobOptionsModel := new(schematicsv1.DeleteJobOptions)
				deleteJobOptionsModel.JobID = core.StringPtr("testString")
				deleteJobOptionsModel.RefreshToken = core.StringPtr("testString")
				deleteJobOptionsModel.Force = core.BoolPtr(true)
				deleteJobOptionsModel.Propagate = core.BoolPtr(true)
				deleteJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = schematicsService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				schematicsService.DisableRetries()
				response, operationErr = schematicsService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteJob with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the DeleteJobOptions model
				deleteJobOptionsModel := new(schematicsv1.DeleteJobOptions)
				deleteJobOptionsModel.JobID = core.StringPtr("testString")
				deleteJobOptionsModel.RefreshToken = core.StringPtr("testString")
				deleteJobOptionsModel.Force = core.BoolPtr(true)
				deleteJobOptionsModel.Propagate = core.BoolPtr(true)
				deleteJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := schematicsService.DeleteJob(deleteJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteJobOptions model with no property values
				deleteJobOptionsModelNew := new(schematicsv1.DeleteJobOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = schematicsService.DeleteJob(deleteJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetJob(getJobOptions *GetJobOptions) - Operation response error`, func() {
		getJobPath := "/v2/jobs/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["profile"]).To(Equal([]string{"summary"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetJob with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(schematicsv1.GetJobOptions)
				getJobOptionsModel.JobID = core.StringPtr("testString")
				getJobOptionsModel.Profile = core.StringPtr("summary")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetJob(getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetJob(getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetJob(getJobOptions *GetJobOptions)`, func() {
		getJobPath := "/v2/jobs/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getJobPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["profile"]).To(Equal([]string{"summary"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"command_object": "workspace", "command_object_id": "CommandObjectID", "command_name": "workspace_init_flow", "command_parameter": "CommandParameter", "command_options": ["CommandOptions"], "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "tags": ["Tags"], "id": "ID", "name": "Name", "description": "Description", "location": "us_south", "resource_group": "ResourceGroup", "submitted_at": "2019-01-01T12:00:00", "submitted_by": "SubmittedBy", "start_at": "2019-01-01T12:00:00", "end_at": "2019-01-01T12:00:00", "duration": "Duration", "status": {"action_job_status": {"action_name": "ActionName", "status_code": "job_pending", "status_message": "StatusMessage", "bastion_status_code": "none", "bastion_status_message": "BastionStatusMessage", "targets_status_code": "none", "targets_status_message": "TargetsStatusMessage", "updated_at": "2019-01-01T12:00:00"}}, "data": {"job_type": "repo_download_job", "action_job_data": {"action_name": "ActionName", "inputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "outputs": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "settings": [{"name": "Name", "value": "Value", "metadata": {"type": "boolean", "aliases": ["Aliases"], "description": "Description", "default_value": "DefaultValue", "secure": true, "immutable": false, "hidden": true, "options": ["Options"], "min_value": 8, "max_value": 8, "min_length": 9, "max_length": 9, "matches": "Matches", "position": 8, "group_by": "GroupBy", "source": "Source"}, "link": "Link"}], "updated_at": "2019-01-01T12:00:00"}}, "targets": [{"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}], "bastion": {"name": "Name", "type": "Type", "description": "Description", "resource_query": "ResourceQuery", "credential_ref": "CredentialRef", "id": "ID", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "sys_lock": {"sys_locked": false, "sys_locked_by": "SysLockedBy", "sys_locked_at": "2019-01-01T12:00:00"}, "resource_ids": ["ResourceIds"]}, "log_summary": {"job_id": "JobID", "job_type": "repo_download_job", "log_start_at": "2019-01-01T12:00:00", "log_analyzed_till": "2019-01-01T12:00:00", "elapsed_time": 11, "log_errors": [{"error_code": "ErrorCode", "error_msg": "ErrorMsg", "error_count": 10}], "repo_download_job": {"scanned_file_count": 16, "quarantined_file_count": 20, "detected_filetype": "DetectedFiletype", "inputs_count": "InputsCount", "outputs_count": "OutputsCount"}, "action_job": {"target_count": 11, "task_count": 9, "play_count": 9, "recap": {"target": ["Target"], "ok": 2, "changed": 7, "failed": 6, "skipped": 7, "unreachable": 11}}}, "log_store_url": "LogStoreURL", "state_store_url": "StateStoreURL", "results_url": "ResultsURL", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetJob successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(schematicsv1.GetJobOptions)
				getJobOptionsModel.JobID = core.StringPtr("testString")
				getJobOptionsModel.Profile = core.StringPtr("summary")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetJob(getJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetJobWithContext(ctx, getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetJob(getJobOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetJobWithContext(ctx, getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetJob with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetJobOptions model
				getJobOptionsModel := new(schematicsv1.GetJobOptions)
				getJobOptionsModel.JobID = core.StringPtr("testString")
				getJobOptionsModel.Profile = core.StringPtr("summary")
				getJobOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetJob(getJobOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetJobOptions model with no property values
				getJobOptionsModelNew := new(schematicsv1.GetJobOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetJob(getJobOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListJobLogs(listJobLogsOptions *ListJobLogsOptions) - Operation response error`, func() {
		listJobLogsPath := "/v2/jobs/testString/logs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobLogsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListJobLogs with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListJobLogsOptions model
				listJobLogsOptionsModel := new(schematicsv1.ListJobLogsOptions)
				listJobLogsOptionsModel.JobID = core.StringPtr("testString")
				listJobLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListJobLogs(listJobLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListJobLogs(listJobLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListJobLogs(listJobLogsOptions *ListJobLogsOptions)`, func() {
		listJobLogsPath := "/v2/jobs/testString/logs"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobLogsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"job_id": "JobID", "job_name": "JobName", "log_summary": {"job_id": "JobID", "job_type": "repo_download_job", "log_start_at": "2019-01-01T12:00:00", "log_analyzed_till": "2019-01-01T12:00:00", "elapsed_time": 11, "log_errors": [{"error_code": "ErrorCode", "error_msg": "ErrorMsg", "error_count": 10}], "repo_download_job": {"scanned_file_count": 16, "quarantined_file_count": 20, "detected_filetype": "DetectedFiletype", "inputs_count": "InputsCount", "outputs_count": "OutputsCount"}, "action_job": {"target_count": 11, "task_count": 9, "play_count": 9, "recap": {"target": ["Target"], "ok": 2, "changed": 7, "failed": 6, "skipped": 7, "unreachable": 11}}}, "format": "json", "details": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke ListJobLogs successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListJobLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListJobLogsOptions model
				listJobLogsOptionsModel := new(schematicsv1.ListJobLogsOptions)
				listJobLogsOptionsModel.JobID = core.StringPtr("testString")
				listJobLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListJobLogs(listJobLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListJobLogsWithContext(ctx, listJobLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListJobLogs(listJobLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListJobLogsWithContext(ctx, listJobLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListJobLogs with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListJobLogsOptions model
				listJobLogsOptionsModel := new(schematicsv1.ListJobLogsOptions)
				listJobLogsOptionsModel.JobID = core.StringPtr("testString")
				listJobLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListJobLogs(listJobLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListJobLogsOptions model with no property values
				listJobLogsOptionsModelNew := new(schematicsv1.ListJobLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ListJobLogs(listJobLogsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListJobStates(listJobStatesOptions *ListJobStatesOptions) - Operation response error`, func() {
		listJobStatesPath := "/v2/jobs/testString/states"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobStatesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListJobStates with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListJobStatesOptions model
				listJobStatesOptionsModel := new(schematicsv1.ListJobStatesOptions)
				listJobStatesOptionsModel.JobID = core.StringPtr("testString")
				listJobStatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListJobStates(listJobStatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListJobStates(listJobStatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListJobStates(listJobStatesOptions *ListJobStatesOptions)`, func() {
		listJobStatesPath := "/v2/jobs/testString/states"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listJobStatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"job_id": "JobID", "job_name": "JobName", "summary": [{"name": "Name", "type": "number", "value": "Value"}], "format": "Format", "details": "VGhpcyBpcyBhbiBlbmNvZGVkIGJ5dGUgYXJyYXku", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke ListJobStates successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListJobStates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListJobStatesOptions model
				listJobStatesOptionsModel := new(schematicsv1.ListJobStatesOptions)
				listJobStatesOptionsModel.JobID = core.StringPtr("testString")
				listJobStatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListJobStates(listJobStatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListJobStatesWithContext(ctx, listJobStatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListJobStates(listJobStatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListJobStatesWithContext(ctx, listJobStatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListJobStates with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListJobStatesOptions model
				listJobStatesOptionsModel := new(schematicsv1.ListJobStatesOptions)
				listJobStatesOptionsModel.JobID = core.StringPtr("testString")
				listJobStatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListJobStates(listJobStatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListJobStatesOptions model with no property values
				listJobStatesOptionsModelNew := new(schematicsv1.ListJobStatesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ListJobStates(listJobStatesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListSharedDatasets(listSharedDatasetsOptions *ListSharedDatasetsOptions) - Operation response error`, func() {
		listSharedDatasetsPath := "/v2/shared_datasets"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSharedDatasetsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSharedDatasets with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListSharedDatasetsOptions model
				listSharedDatasetsOptionsModel := new(schematicsv1.ListSharedDatasetsOptions)
				listSharedDatasetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ListSharedDatasets(listSharedDatasetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ListSharedDatasets(listSharedDatasetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListSharedDatasets(listSharedDatasetsOptions *ListSharedDatasetsOptions)`, func() {
		listSharedDatasetsPath := "/v2/shared_datasets"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSharedDatasetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"count": 5, "shared_datasets": [{"account": "Account", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "description": "Description", "effected_workspace_ids": ["EffectedWorkspaceIds"], "resource_group": "ResourceGroup", "shared_dataset_data": [{"default_value": "DefaultValue", "description": "Description", "hidden": true, "immutable": false, "matches": "Matches", "max_value": "MaxValue", "max_value_len": "MaxValueLen", "min_value": "MinValue", "min_value_len": "MinValueLen", "options": ["Options"], "override_value": "OverrideValue", "secure": true, "var_aliases": ["VarAliases"], "var_name": "VarName", "var_ref": "VarRef", "var_type": "VarType"}], "shared_dataset_id": "SharedDatasetID", "shared_dataset_name": "SharedDatasetName", "shared_dataset_type": ["SharedDatasetType"], "state": "State", "tags": ["Tags"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "version": "Version"}]}`)
				}))
			})
			It(`Invoke ListSharedDatasets successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ListSharedDatasets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSharedDatasetsOptions model
				listSharedDatasetsOptionsModel := new(schematicsv1.ListSharedDatasetsOptions)
				listSharedDatasetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ListSharedDatasets(listSharedDatasetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListSharedDatasetsWithContext(ctx, listSharedDatasetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ListSharedDatasets(listSharedDatasetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ListSharedDatasetsWithContext(ctx, listSharedDatasetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListSharedDatasets with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the ListSharedDatasetsOptions model
				listSharedDatasetsOptionsModel := new(schematicsv1.ListSharedDatasetsOptions)
				listSharedDatasetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ListSharedDatasets(listSharedDatasetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSharedDataset(createSharedDatasetOptions *CreateSharedDatasetOptions) - Operation response error`, func() {
		createSharedDatasetPath := "/v2/shared_datasets"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSharedDatasetPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSharedDataset with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")

				// Construct an instance of the CreateSharedDatasetOptions model
				createSharedDatasetOptionsModel := new(schematicsv1.CreateSharedDatasetOptions)
				createSharedDatasetOptionsModel.AutoPropagateChange = core.BoolPtr(true)
				createSharedDatasetOptionsModel.Description = core.StringPtr("testString")
				createSharedDatasetOptionsModel.EffectedWorkspaceIds = []string{"testString"}
				createSharedDatasetOptionsModel.ResourceGroup = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetData = []schematicsv1.SharedDatasetData{*sharedDatasetDataModel}
				createSharedDatasetOptionsModel.SharedDatasetName = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetSourceName = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetType = []string{"testString"}
				createSharedDatasetOptionsModel.Tags = []string{"testString"}
				createSharedDatasetOptionsModel.Version = core.StringPtr("testString")
				createSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.CreateSharedDataset(createSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.CreateSharedDataset(createSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateSharedDataset(createSharedDatasetOptions *CreateSharedDatasetOptions)`, func() {
		createSharedDatasetPath := "/v2/shared_datasets"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSharedDatasetPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"account": "Account", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "description": "Description", "effected_workspace_ids": ["EffectedWorkspaceIds"], "resource_group": "ResourceGroup", "shared_dataset_data": [{"default_value": "DefaultValue", "description": "Description", "hidden": true, "immutable": false, "matches": "Matches", "max_value": "MaxValue", "max_value_len": "MaxValueLen", "min_value": "MinValue", "min_value_len": "MinValueLen", "options": ["Options"], "override_value": "OverrideValue", "secure": true, "var_aliases": ["VarAliases"], "var_name": "VarName", "var_ref": "VarRef", "var_type": "VarType"}], "shared_dataset_id": "SharedDatasetID", "shared_dataset_name": "SharedDatasetName", "shared_dataset_type": ["SharedDatasetType"], "state": "State", "tags": ["Tags"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "version": "Version"}`)
				}))
			})
			It(`Invoke CreateSharedDataset successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.CreateSharedDataset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")

				// Construct an instance of the CreateSharedDatasetOptions model
				createSharedDatasetOptionsModel := new(schematicsv1.CreateSharedDatasetOptions)
				createSharedDatasetOptionsModel.AutoPropagateChange = core.BoolPtr(true)
				createSharedDatasetOptionsModel.Description = core.StringPtr("testString")
				createSharedDatasetOptionsModel.EffectedWorkspaceIds = []string{"testString"}
				createSharedDatasetOptionsModel.ResourceGroup = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetData = []schematicsv1.SharedDatasetData{*sharedDatasetDataModel}
				createSharedDatasetOptionsModel.SharedDatasetName = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetSourceName = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetType = []string{"testString"}
				createSharedDatasetOptionsModel.Tags = []string{"testString"}
				createSharedDatasetOptionsModel.Version = core.StringPtr("testString")
				createSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.CreateSharedDataset(createSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateSharedDatasetWithContext(ctx, createSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.CreateSharedDataset(createSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.CreateSharedDatasetWithContext(ctx, createSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateSharedDataset with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")

				// Construct an instance of the CreateSharedDatasetOptions model
				createSharedDatasetOptionsModel := new(schematicsv1.CreateSharedDatasetOptions)
				createSharedDatasetOptionsModel.AutoPropagateChange = core.BoolPtr(true)
				createSharedDatasetOptionsModel.Description = core.StringPtr("testString")
				createSharedDatasetOptionsModel.EffectedWorkspaceIds = []string{"testString"}
				createSharedDatasetOptionsModel.ResourceGroup = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetData = []schematicsv1.SharedDatasetData{*sharedDatasetDataModel}
				createSharedDatasetOptionsModel.SharedDatasetName = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetSourceName = core.StringPtr("testString")
				createSharedDatasetOptionsModel.SharedDatasetType = []string{"testString"}
				createSharedDatasetOptionsModel.Tags = []string{"testString"}
				createSharedDatasetOptionsModel.Version = core.StringPtr("testString")
				createSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.CreateSharedDataset(createSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSharedDataset(getSharedDatasetOptions *GetSharedDatasetOptions) - Operation response error`, func() {
		getSharedDatasetPath := "/v2/shared_datasets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSharedDatasetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSharedDataset with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetSharedDatasetOptions model
				getSharedDatasetOptionsModel := new(schematicsv1.GetSharedDatasetOptions)
				getSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				getSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetSharedDataset(getSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetSharedDataset(getSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSharedDataset(getSharedDatasetOptions *GetSharedDatasetOptions)`, func() {
		getSharedDatasetPath := "/v2/shared_datasets/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSharedDatasetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account": "Account", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "description": "Description", "effected_workspace_ids": ["EffectedWorkspaceIds"], "resource_group": "ResourceGroup", "shared_dataset_data": [{"default_value": "DefaultValue", "description": "Description", "hidden": true, "immutable": false, "matches": "Matches", "max_value": "MaxValue", "max_value_len": "MaxValueLen", "min_value": "MinValue", "min_value_len": "MinValueLen", "options": ["Options"], "override_value": "OverrideValue", "secure": true, "var_aliases": ["VarAliases"], "var_name": "VarName", "var_ref": "VarRef", "var_type": "VarType"}], "shared_dataset_id": "SharedDatasetID", "shared_dataset_name": "SharedDatasetName", "shared_dataset_type": ["SharedDatasetType"], "state": "State", "tags": ["Tags"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "version": "Version"}`)
				}))
			})
			It(`Invoke GetSharedDataset successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetSharedDataset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSharedDatasetOptions model
				getSharedDatasetOptionsModel := new(schematicsv1.GetSharedDatasetOptions)
				getSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				getSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetSharedDataset(getSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetSharedDatasetWithContext(ctx, getSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetSharedDataset(getSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetSharedDatasetWithContext(ctx, getSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetSharedDataset with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetSharedDatasetOptions model
				getSharedDatasetOptionsModel := new(schematicsv1.GetSharedDatasetOptions)
				getSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				getSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetSharedDataset(getSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSharedDatasetOptions model with no property values
				getSharedDatasetOptionsModelNew := new(schematicsv1.GetSharedDatasetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetSharedDataset(getSharedDatasetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceSharedDataset(replaceSharedDatasetOptions *ReplaceSharedDatasetOptions) - Operation response error`, func() {
		replaceSharedDatasetPath := "/v2/shared_datasets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSharedDatasetPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceSharedDataset with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")

				// Construct an instance of the ReplaceSharedDatasetOptions model
				replaceSharedDatasetOptionsModel := new(schematicsv1.ReplaceSharedDatasetOptions)
				replaceSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.AutoPropagateChange = core.BoolPtr(true)
				replaceSharedDatasetOptionsModel.Description = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.EffectedWorkspaceIds = []string{"testString"}
				replaceSharedDatasetOptionsModel.ResourceGroup = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetData = []schematicsv1.SharedDatasetData{*sharedDatasetDataModel}
				replaceSharedDatasetOptionsModel.SharedDatasetName = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetSourceName = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetType = []string{"testString"}
				replaceSharedDatasetOptionsModel.Tags = []string{"testString"}
				replaceSharedDatasetOptionsModel.Version = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ReplaceSharedDataset(replaceSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ReplaceSharedDataset(replaceSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceSharedDataset(replaceSharedDatasetOptions *ReplaceSharedDatasetOptions)`, func() {
		replaceSharedDatasetPath := "/v2/shared_datasets/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceSharedDatasetPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account": "Account", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "description": "Description", "effected_workspace_ids": ["EffectedWorkspaceIds"], "resource_group": "ResourceGroup", "shared_dataset_data": [{"default_value": "DefaultValue", "description": "Description", "hidden": true, "immutable": false, "matches": "Matches", "max_value": "MaxValue", "max_value_len": "MaxValueLen", "min_value": "MinValue", "min_value_len": "MinValueLen", "options": ["Options"], "override_value": "OverrideValue", "secure": true, "var_aliases": ["VarAliases"], "var_name": "VarName", "var_ref": "VarRef", "var_type": "VarType"}], "shared_dataset_id": "SharedDatasetID", "shared_dataset_name": "SharedDatasetName", "shared_dataset_type": ["SharedDatasetType"], "state": "State", "tags": ["Tags"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "version": "Version"}`)
				}))
			})
			It(`Invoke ReplaceSharedDataset successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ReplaceSharedDataset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")

				// Construct an instance of the ReplaceSharedDatasetOptions model
				replaceSharedDatasetOptionsModel := new(schematicsv1.ReplaceSharedDatasetOptions)
				replaceSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.AutoPropagateChange = core.BoolPtr(true)
				replaceSharedDatasetOptionsModel.Description = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.EffectedWorkspaceIds = []string{"testString"}
				replaceSharedDatasetOptionsModel.ResourceGroup = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetData = []schematicsv1.SharedDatasetData{*sharedDatasetDataModel}
				replaceSharedDatasetOptionsModel.SharedDatasetName = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetSourceName = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetType = []string{"testString"}
				replaceSharedDatasetOptionsModel.Tags = []string{"testString"}
				replaceSharedDatasetOptionsModel.Version = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ReplaceSharedDataset(replaceSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceSharedDatasetWithContext(ctx, replaceSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ReplaceSharedDataset(replaceSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceSharedDatasetWithContext(ctx, replaceSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceSharedDataset with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")

				// Construct an instance of the ReplaceSharedDatasetOptions model
				replaceSharedDatasetOptionsModel := new(schematicsv1.ReplaceSharedDatasetOptions)
				replaceSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.AutoPropagateChange = core.BoolPtr(true)
				replaceSharedDatasetOptionsModel.Description = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.EffectedWorkspaceIds = []string{"testString"}
				replaceSharedDatasetOptionsModel.ResourceGroup = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetData = []schematicsv1.SharedDatasetData{*sharedDatasetDataModel}
				replaceSharedDatasetOptionsModel.SharedDatasetName = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetSourceName = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.SharedDatasetType = []string{"testString"}
				replaceSharedDatasetOptionsModel.Tags = []string{"testString"}
				replaceSharedDatasetOptionsModel.Version = core.StringPtr("testString")
				replaceSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ReplaceSharedDataset(replaceSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceSharedDatasetOptions model with no property values
				replaceSharedDatasetOptionsModelNew := new(schematicsv1.ReplaceSharedDatasetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.ReplaceSharedDataset(replaceSharedDatasetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSharedDataset(deleteSharedDatasetOptions *DeleteSharedDatasetOptions) - Operation response error`, func() {
		deleteSharedDatasetPath := "/v2/shared_datasets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSharedDatasetPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteSharedDataset with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the DeleteSharedDatasetOptions model
				deleteSharedDatasetOptionsModel := new(schematicsv1.DeleteSharedDatasetOptions)
				deleteSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				deleteSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.DeleteSharedDataset(deleteSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.DeleteSharedDataset(deleteSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteSharedDataset(deleteSharedDatasetOptions *DeleteSharedDatasetOptions)`, func() {
		deleteSharedDatasetPath := "/v2/shared_datasets/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSharedDatasetPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account": "Account", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "description": "Description", "effected_workspace_ids": ["EffectedWorkspaceIds"], "resource_group": "ResourceGroup", "shared_dataset_data": [{"default_value": "DefaultValue", "description": "Description", "hidden": true, "immutable": false, "matches": "Matches", "max_value": "MaxValue", "max_value_len": "MaxValueLen", "min_value": "MinValue", "min_value_len": "MinValueLen", "options": ["Options"], "override_value": "OverrideValue", "secure": true, "var_aliases": ["VarAliases"], "var_name": "VarName", "var_ref": "VarRef", "var_type": "VarType"}], "shared_dataset_id": "SharedDatasetID", "shared_dataset_name": "SharedDatasetName", "shared_dataset_type": ["SharedDatasetType"], "state": "State", "tags": ["Tags"], "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy", "version": "Version"}`)
				}))
			})
			It(`Invoke DeleteSharedDataset successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.DeleteSharedDataset(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteSharedDatasetOptions model
				deleteSharedDatasetOptionsModel := new(schematicsv1.DeleteSharedDatasetOptions)
				deleteSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				deleteSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.DeleteSharedDataset(deleteSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DeleteSharedDatasetWithContext(ctx, deleteSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.DeleteSharedDataset(deleteSharedDatasetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.DeleteSharedDatasetWithContext(ctx, deleteSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteSharedDataset with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the DeleteSharedDatasetOptions model
				deleteSharedDatasetOptionsModel := new(schematicsv1.DeleteSharedDatasetOptions)
				deleteSharedDatasetOptionsModel.SdID = core.StringPtr("testString")
				deleteSharedDatasetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.DeleteSharedDataset(deleteSharedDatasetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteSharedDatasetOptions model with no property values
				deleteSharedDatasetOptionsModelNew := new(schematicsv1.DeleteSharedDatasetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.DeleteSharedDataset(deleteSharedDatasetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schematicsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL: "https://schematicsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schematicsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
					URL: "https://testService/api",
				})
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				})
				err := schematicsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_URL": "https://schematicsv1/api",
				"SCHEMATICS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMATICS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schematicsService, serviceErr := schematicsv1.NewSchematicsV1UsingExternalConfig(&schematicsv1.SchematicsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schematicsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetKmsSettings(getKmsSettingsOptions *GetKmsSettingsOptions) - Operation response error`, func() {
		getKmsSettingsPath := "/v2/settings/kms"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKmsSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetKmsSettings with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetKmsSettingsOptions model
				getKmsSettingsOptionsModel := new(schematicsv1.GetKmsSettingsOptions)
				getKmsSettingsOptionsModel.Location = core.StringPtr("testString")
				getKmsSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetKmsSettings(getKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetKmsSettings(getKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetKmsSettings(getKmsSettingsOptions *GetKmsSettingsOptions)`, func() {
		getKmsSettingsPath := "/v2/settings/kms"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getKmsSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": "Location", "encryption_scheme": "EncryptionScheme", "resource_group": "ResourceGroup", "primary_crk": {"kms_name": "KmsName", "kms_private_endpoint": "KmsPrivateEndpoint", "key_crn": "KeyCrn"}, "secondary_crk": {"kms_name": "KmsName", "kms_private_endpoint": "KmsPrivateEndpoint", "key_crn": "KeyCrn"}}`)
				}))
			})
			It(`Invoke GetKmsSettings successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetKmsSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetKmsSettingsOptions model
				getKmsSettingsOptionsModel := new(schematicsv1.GetKmsSettingsOptions)
				getKmsSettingsOptionsModel.Location = core.StringPtr("testString")
				getKmsSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetKmsSettings(getKmsSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetKmsSettingsWithContext(ctx, getKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetKmsSettings(getKmsSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetKmsSettingsWithContext(ctx, getKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetKmsSettings with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetKmsSettingsOptions model
				getKmsSettingsOptionsModel := new(schematicsv1.GetKmsSettingsOptions)
				getKmsSettingsOptionsModel.Location = core.StringPtr("testString")
				getKmsSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetKmsSettings(getKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetKmsSettingsOptions model with no property values
				getKmsSettingsOptionsModelNew := new(schematicsv1.GetKmsSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetKmsSettings(getKmsSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceKmsSettings(replaceKmsSettingsOptions *ReplaceKmsSettingsOptions) - Operation response error`, func() {
		replaceKmsSettingsPath := "/v2/settings/kms"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceKmsSettingsPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceKmsSettings with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the KMSSettingsPrimaryCrk model
				kmsSettingsPrimaryCrkModel := new(schematicsv1.KMSSettingsPrimaryCrk)
				kmsSettingsPrimaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KeyCrn = core.StringPtr("testString")

				// Construct an instance of the KMSSettingsSecondaryCrk model
				kmsSettingsSecondaryCrkModel := new(schematicsv1.KMSSettingsSecondaryCrk)
				kmsSettingsSecondaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KeyCrn = core.StringPtr("testString")

				// Construct an instance of the ReplaceKmsSettingsOptions model
				replaceKmsSettingsOptionsModel := new(schematicsv1.ReplaceKmsSettingsOptions)
				replaceKmsSettingsOptionsModel.Location = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.EncryptionScheme = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.ResourceGroup = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.PrimaryCrk = kmsSettingsPrimaryCrkModel
				replaceKmsSettingsOptionsModel.SecondaryCrk = kmsSettingsSecondaryCrkModel
				replaceKmsSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.ReplaceKmsSettings(replaceKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.ReplaceKmsSettings(replaceKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceKmsSettings(replaceKmsSettingsOptions *ReplaceKmsSettingsOptions)`, func() {
		replaceKmsSettingsPath := "/v2/settings/kms"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceKmsSettingsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"location": "Location", "encryption_scheme": "EncryptionScheme", "resource_group": "ResourceGroup", "primary_crk": {"kms_name": "KmsName", "kms_private_endpoint": "KmsPrivateEndpoint", "key_crn": "KeyCrn"}, "secondary_crk": {"kms_name": "KmsName", "kms_private_endpoint": "KmsPrivateEndpoint", "key_crn": "KeyCrn"}}`)
				}))
			})
			It(`Invoke ReplaceKmsSettings successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.ReplaceKmsSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the KMSSettingsPrimaryCrk model
				kmsSettingsPrimaryCrkModel := new(schematicsv1.KMSSettingsPrimaryCrk)
				kmsSettingsPrimaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KeyCrn = core.StringPtr("testString")

				// Construct an instance of the KMSSettingsSecondaryCrk model
				kmsSettingsSecondaryCrkModel := new(schematicsv1.KMSSettingsSecondaryCrk)
				kmsSettingsSecondaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KeyCrn = core.StringPtr("testString")

				// Construct an instance of the ReplaceKmsSettingsOptions model
				replaceKmsSettingsOptionsModel := new(schematicsv1.ReplaceKmsSettingsOptions)
				replaceKmsSettingsOptionsModel.Location = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.EncryptionScheme = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.ResourceGroup = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.PrimaryCrk = kmsSettingsPrimaryCrkModel
				replaceKmsSettingsOptionsModel.SecondaryCrk = kmsSettingsSecondaryCrkModel
				replaceKmsSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.ReplaceKmsSettings(replaceKmsSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceKmsSettingsWithContext(ctx, replaceKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.ReplaceKmsSettings(replaceKmsSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.ReplaceKmsSettingsWithContext(ctx, replaceKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceKmsSettings with error: Operation request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the KMSSettingsPrimaryCrk model
				kmsSettingsPrimaryCrkModel := new(schematicsv1.KMSSettingsPrimaryCrk)
				kmsSettingsPrimaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KeyCrn = core.StringPtr("testString")

				// Construct an instance of the KMSSettingsSecondaryCrk model
				kmsSettingsSecondaryCrkModel := new(schematicsv1.KMSSettingsSecondaryCrk)
				kmsSettingsSecondaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KeyCrn = core.StringPtr("testString")

				// Construct an instance of the ReplaceKmsSettingsOptions model
				replaceKmsSettingsOptionsModel := new(schematicsv1.ReplaceKmsSettingsOptions)
				replaceKmsSettingsOptionsModel.Location = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.EncryptionScheme = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.ResourceGroup = core.StringPtr("testString")
				replaceKmsSettingsOptionsModel.PrimaryCrk = kmsSettingsPrimaryCrkModel
				replaceKmsSettingsOptionsModel.SecondaryCrk = kmsSettingsSecondaryCrkModel
				replaceKmsSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.ReplaceKmsSettings(replaceKmsSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptions *GetDiscoveredKmsInstancesOptions) - Operation response error`, func() {
		getDiscoveredKmsInstancesPath := "/v2/settings/kms_instances"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDiscoveredKmsInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["encryption_scheme"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDiscoveredKmsInstances with error: Operation response processing error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetDiscoveredKmsInstancesOptions model
				getDiscoveredKmsInstancesOptionsModel := new(schematicsv1.GetDiscoveredKmsInstancesOptions)
				getDiscoveredKmsInstancesOptionsModel.EncryptionScheme = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Location = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.ResourceGroup = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Limit = core.Int64Ptr(int64(1))
				getDiscoveredKmsInstancesOptionsModel.Sort = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schematicsService.GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schematicsService.EnableRetries(0, 0)
				result, response, operationErr = schematicsService.GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptions *GetDiscoveredKmsInstancesOptions)`, func() {
		getDiscoveredKmsInstancesPath := "/v2/settings/kms_instances"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDiscoveredKmsInstancesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["encryption_scheme"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["location"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "limit": 5, "offset": 6, "kms_instances": [{"location": "Location", "encryption_scheme": "EncryptionScheme", "resource_group": "ResourceGroup", "kms_crn": "KmsCrn", "kms_name": "KmsName", "kms_private_endpoint": "KmsPrivateEndpoint", "kms_public_endpoint": "KmsPublicEndpoint", "keys": [{"name": "Name", "crn": "Crn", "error": "Error"}]}]}`)
				}))
			})
			It(`Invoke GetDiscoveredKmsInstances successfully`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())
				schematicsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schematicsService.GetDiscoveredKmsInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDiscoveredKmsInstancesOptions model
				getDiscoveredKmsInstancesOptionsModel := new(schematicsv1.GetDiscoveredKmsInstancesOptions)
				getDiscoveredKmsInstancesOptionsModel.EncryptionScheme = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Location = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.ResourceGroup = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Limit = core.Int64Ptr(int64(1))
				getDiscoveredKmsInstancesOptionsModel.Sort = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schematicsService.GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetDiscoveredKmsInstancesWithContext(ctx, getDiscoveredKmsInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				schematicsService.DisableRetries()
				result, response, operationErr = schematicsService.GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = schematicsService.GetDiscoveredKmsInstancesWithContext(ctx, getDiscoveredKmsInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetDiscoveredKmsInstances with error: Operation validation and request error`, func() {
				schematicsService, serviceErr := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schematicsService).ToNot(BeNil())

				// Construct an instance of the GetDiscoveredKmsInstancesOptions model
				getDiscoveredKmsInstancesOptionsModel := new(schematicsv1.GetDiscoveredKmsInstancesOptions)
				getDiscoveredKmsInstancesOptionsModel.EncryptionScheme = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Location = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.ResourceGroup = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Limit = core.Int64Ptr(int64(1))
				getDiscoveredKmsInstancesOptionsModel.Sort = core.StringPtr("testString")
				getDiscoveredKmsInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schematicsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schematicsService.GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDiscoveredKmsInstancesOptions model with no property values
				getDiscoveredKmsInstancesOptionsModelNew := new(schematicsv1.GetDiscoveredKmsInstancesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schematicsService.GetDiscoveredKmsInstances(getDiscoveredKmsInstancesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			schematicsService, _ := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
				URL:           "http://schematicsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewApplyWorkspaceCommandOptions successfully`, func() {
				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				Expect(workspaceActivityOptionsTemplateModel).ToNot(BeNil())
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}
				Expect(workspaceActivityOptionsTemplateModel.Target).To(Equal([]string{"testString"}))
				Expect(workspaceActivityOptionsTemplateModel.TfVars).To(Equal([]string{"testString"}))

				// Construct an instance of the ApplyWorkspaceCommandOptions model
				wID := "testString"
				refreshToken := "testString"
				applyWorkspaceCommandOptionsModel := schematicsService.NewApplyWorkspaceCommandOptions(wID, refreshToken)
				applyWorkspaceCommandOptionsModel.SetWID("testString")
				applyWorkspaceCommandOptionsModel.SetRefreshToken("testString")
				applyWorkspaceCommandOptionsModel.SetActionOptions(workspaceActivityOptionsTemplateModel)
				applyWorkspaceCommandOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(applyWorkspaceCommandOptionsModel).ToNot(BeNil())
				Expect(applyWorkspaceCommandOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(applyWorkspaceCommandOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(applyWorkspaceCommandOptionsModel.ActionOptions).To(Equal(workspaceActivityOptionsTemplateModel))
				Expect(applyWorkspaceCommandOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateActionOptions successfully`, func() {
				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				Expect(userStateModel).ToNot(BeNil())
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()
				Expect(userStateModel.State).To(Equal(core.StringPtr("draft")))
				Expect(userStateModel.SetBy).To(Equal(core.StringPtr("testString")))
				Expect(userStateModel.SetAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				Expect(externalSourceGitModel).ToNot(BeNil())
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")
				Expect(externalSourceGitModel.GitRepoURL).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitToken).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitRepoFolder).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitRelease).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitBranch).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				Expect(externalSourceModel).ToNot(BeNil())
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel
				Expect(externalSourceModel.SourceType).To(Equal(core.StringPtr("local")))
				Expect(externalSourceModel.Git).To(Equal(externalSourceGitModel))

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				Expect(systemLockModel).ToNot(BeNil())
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()
				Expect(systemLockModel.SysLocked).To(Equal(core.BoolPtr(true)))
				Expect(systemLockModel.SysLockedBy).To(Equal(core.StringPtr("testString")))
				Expect(systemLockModel.SysLockedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				Expect(targetResourcesetModel).ToNot(BeNil())
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel
				Expect(targetResourcesetModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.ResourceQuery).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.CredentialRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.SysLock).To(Equal(systemLockModel))

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				Expect(targetResourceConfigModel).ToNot(BeNil())
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")
				Expect(targetResourceConfigModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceConfigModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceConfigModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				Expect(targetResourceModel).ToNot(BeNil())
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}
				Expect(targetResourceModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceModel.ResourceConfigs).To(Equal([]schematicsv1.TargetResourceConfig{*targetResourceConfigModel}))

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				Expect(targetResourceGroupModel).ToNot(BeNil())
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}
				Expect(targetResourceGroupModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.CredentialRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.BastionRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.TargetResources).To(Equal([]schematicsv1.TargetResource{*targetResourceModel}))

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				Expect(variableMetadataModel).ToNot(BeNil())
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")
				Expect(variableMetadataModel.Type).To(Equal(core.StringPtr("boolean")))
				Expect(variableMetadataModel.Aliases).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Options).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.MinValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MinLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.Matches).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Position).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.GroupBy).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Source).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				Expect(variableDataModel).ToNot(BeNil())
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel
				Expect(variableDataModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Metadata).To(Equal(variableMetadataModel))

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				Expect(actionStateModel).ToNot(BeNil())
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")
				Expect(actionStateModel.StatusCode).To(Equal(core.StringPtr("normal")))
				Expect(actionStateModel.StatusJobID).To(Equal(core.StringPtr("testString")))
				Expect(actionStateModel.StatusMessage).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateActionOptions model
				createActionOptionsModel := schematicsService.NewCreateActionOptions()
				createActionOptionsModel.SetName("Stop Action")
				createActionOptionsModel.SetDescription("This Action can be used to Stop the targets")
				createActionOptionsModel.SetLocation("us_south")
				createActionOptionsModel.SetResourceGroup("testString")
				createActionOptionsModel.SetTags([]string{"testString"})
				createActionOptionsModel.SetUserState(userStateModel)
				createActionOptionsModel.SetSourceReadmeURL("testString")
				createActionOptionsModel.SetSource(externalSourceModel)
				createActionOptionsModel.SetSourceType("local")
				createActionOptionsModel.SetCommandParameter("testString")
				createActionOptionsModel.SetBastion(targetResourcesetModel)
				createActionOptionsModel.SetTargets([]schematicsv1.TargetResourceGroup{*targetResourceGroupModel})
				createActionOptionsModel.SetCredentials([]schematicsv1.VariableData{*variableDataModel})
				createActionOptionsModel.SetInputs([]schematicsv1.VariableData{*variableDataModel})
				createActionOptionsModel.SetOutputs([]schematicsv1.VariableData{*variableDataModel})
				createActionOptionsModel.SetSettings([]schematicsv1.VariableData{*variableDataModel})
				createActionOptionsModel.SetTriggerRecordID("testString")
				createActionOptionsModel.SetState(actionStateModel)
				createActionOptionsModel.SetSysLock(systemLockModel)
				createActionOptionsModel.SetXGithubToken("testString")
				createActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createActionOptionsModel).ToNot(BeNil())
				Expect(createActionOptionsModel.Name).To(Equal(core.StringPtr("Stop Action")))
				Expect(createActionOptionsModel.Description).To(Equal(core.StringPtr("This Action can be used to Stop the targets")))
				Expect(createActionOptionsModel.Location).To(Equal(core.StringPtr("us_south")))
				Expect(createActionOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(createActionOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createActionOptionsModel.UserState).To(Equal(userStateModel))
				Expect(createActionOptionsModel.SourceReadmeURL).To(Equal(core.StringPtr("testString")))
				Expect(createActionOptionsModel.Source).To(Equal(externalSourceModel))
				Expect(createActionOptionsModel.SourceType).To(Equal(core.StringPtr("local")))
				Expect(createActionOptionsModel.CommandParameter).To(Equal(core.StringPtr("testString")))
				Expect(createActionOptionsModel.Bastion).To(Equal(targetResourcesetModel))
				Expect(createActionOptionsModel.Targets).To(Equal([]schematicsv1.TargetResourceGroup{*targetResourceGroupModel}))
				Expect(createActionOptionsModel.Credentials).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(createActionOptionsModel.Inputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(createActionOptionsModel.Outputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(createActionOptionsModel.Settings).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(createActionOptionsModel.TriggerRecordID).To(Equal(core.StringPtr("testString")))
				Expect(createActionOptionsModel.State).To(Equal(actionStateModel))
				Expect(createActionOptionsModel.SysLock).To(Equal(systemLockModel))
				Expect(createActionOptionsModel.XGithubToken).To(Equal(core.StringPtr("testString")))
				Expect(createActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateJobOptions successfully`, func() {
				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				Expect(variableMetadataModel).ToNot(BeNil())
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")
				Expect(variableMetadataModel.Type).To(Equal(core.StringPtr("boolean")))
				Expect(variableMetadataModel.Aliases).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Options).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.MinValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MinLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.Matches).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Position).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.GroupBy).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Source).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				Expect(variableDataModel).ToNot(BeNil())
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel
				Expect(variableDataModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Metadata).To(Equal(variableMetadataModel))

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				Expect(jobStatusActionModel).ToNot(BeNil())
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()
				Expect(jobStatusActionModel.ActionName).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.StatusCode).To(Equal(core.StringPtr("job_pending")))
				Expect(jobStatusActionModel.StatusMessage).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.BastionStatusCode).To(Equal(core.StringPtr("none")))
				Expect(jobStatusActionModel.BastionStatusMessage).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.TargetsStatusCode).To(Equal(core.StringPtr("none")))
				Expect(jobStatusActionModel.TargetsStatusMessage).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.UpdatedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				Expect(jobStatusModel).ToNot(BeNil())
				jobStatusModel.ActionJobStatus = jobStatusActionModel
				Expect(jobStatusModel.ActionJobStatus).To(Equal(jobStatusActionModel))

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				Expect(jobDataActionModel).ToNot(BeNil())
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()
				Expect(jobDataActionModel.ActionName).To(Equal(core.StringPtr("testString")))
				Expect(jobDataActionModel.Inputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(jobDataActionModel.Outputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(jobDataActionModel.Settings).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(jobDataActionModel.UpdatedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				Expect(jobDataModel).ToNot(BeNil())
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel
				Expect(jobDataModel.JobType).To(Equal(core.StringPtr("repo_download_job")))
				Expect(jobDataModel.ActionJobData).To(Equal(jobDataActionModel))

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				Expect(systemLockModel).ToNot(BeNil())
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()
				Expect(systemLockModel.SysLocked).To(Equal(core.BoolPtr(true)))
				Expect(systemLockModel.SysLockedBy).To(Equal(core.StringPtr("testString")))
				Expect(systemLockModel.SysLockedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				Expect(targetResourcesetModel).ToNot(BeNil())
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel
				Expect(targetResourcesetModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.ResourceQuery).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.CredentialRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.SysLock).To(Equal(systemLockModel))

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)
				Expect(jobLogSummaryRepoDownloadJobModel).ToNot(BeNil())

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				Expect(jobLogSummaryActionJobRecapModel).ToNot(BeNil())
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))
				Expect(jobLogSummaryActionJobRecapModel.Target).To(Equal([]string{"testString"}))
				Expect(jobLogSummaryActionJobRecapModel.Ok).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Changed).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Failed).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Skipped).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Unreachable).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				Expect(jobLogSummaryActionJobModel).ToNot(BeNil())
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel
				Expect(jobLogSummaryActionJobModel.Recap).To(Equal(jobLogSummaryActionJobRecapModel))

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				Expect(jobLogSummaryModel).ToNot(BeNil())
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel
				Expect(jobLogSummaryModel.JobType).To(Equal(core.StringPtr("repo_download_job")))
				Expect(jobLogSummaryModel.RepoDownloadJob).To(Equal(jobLogSummaryRepoDownloadJobModel))
				Expect(jobLogSummaryModel.ActionJob).To(Equal(jobLogSummaryActionJobModel))

				// Construct an instance of the CreateJobOptions model
				refreshToken := "testString"
				createJobOptionsModel := schematicsService.NewCreateJobOptions(refreshToken)
				createJobOptionsModel.SetRefreshToken("testString")
				createJobOptionsModel.SetCommandObject("workspace")
				createJobOptionsModel.SetCommandObjectID("testString")
				createJobOptionsModel.SetCommandName("workspace_init_flow")
				createJobOptionsModel.SetCommandParameter("testString")
				createJobOptionsModel.SetCommandOptions([]string{"testString"})
				createJobOptionsModel.SetInputs([]schematicsv1.VariableData{*variableDataModel})
				createJobOptionsModel.SetSettings([]schematicsv1.VariableData{*variableDataModel})
				createJobOptionsModel.SetTags([]string{"testString"})
				createJobOptionsModel.SetLocation("us_south")
				createJobOptionsModel.SetStatus(jobStatusModel)
				createJobOptionsModel.SetData(jobDataModel)
				createJobOptionsModel.SetBastion(targetResourcesetModel)
				createJobOptionsModel.SetLogSummary(jobLogSummaryModel)
				createJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createJobOptionsModel).ToNot(BeNil())
				Expect(createJobOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.CommandObject).To(Equal(core.StringPtr("workspace")))
				Expect(createJobOptionsModel.CommandObjectID).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.CommandName).To(Equal(core.StringPtr("workspace_init_flow")))
				Expect(createJobOptionsModel.CommandParameter).To(Equal(core.StringPtr("testString")))
				Expect(createJobOptionsModel.CommandOptions).To(Equal([]string{"testString"}))
				Expect(createJobOptionsModel.Inputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(createJobOptionsModel.Settings).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(createJobOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createJobOptionsModel.Location).To(Equal(core.StringPtr("us_south")))
				Expect(createJobOptionsModel.Status).To(Equal(jobStatusModel))
				Expect(createJobOptionsModel.Data).To(Equal(jobDataModel))
				Expect(createJobOptionsModel.Bastion).To(Equal(targetResourcesetModel))
				Expect(createJobOptionsModel.LogSummary).To(Equal(jobLogSummaryModel))
				Expect(createJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSharedDatasetOptions successfully`, func() {
				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				Expect(sharedDatasetDataModel).ToNot(BeNil())
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")
				Expect(sharedDatasetDataModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(sharedDatasetDataModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(sharedDatasetDataModel.Matches).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MaxValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MaxValueLen).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MinValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MinValueLen).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Options).To(Equal([]string{"testString"}))
				Expect(sharedDatasetDataModel.OverrideValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(sharedDatasetDataModel.VarAliases).To(Equal([]string{"testString"}))
				Expect(sharedDatasetDataModel.VarName).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.VarRef).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.VarType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateSharedDatasetOptions model
				createSharedDatasetOptionsModel := schematicsService.NewCreateSharedDatasetOptions()
				createSharedDatasetOptionsModel.SetAutoPropagateChange(true)
				createSharedDatasetOptionsModel.SetDescription("testString")
				createSharedDatasetOptionsModel.SetEffectedWorkspaceIds([]string{"testString"})
				createSharedDatasetOptionsModel.SetResourceGroup("testString")
				createSharedDatasetOptionsModel.SetSharedDatasetData([]schematicsv1.SharedDatasetData{*sharedDatasetDataModel})
				createSharedDatasetOptionsModel.SetSharedDatasetName("testString")
				createSharedDatasetOptionsModel.SetSharedDatasetSourceName("testString")
				createSharedDatasetOptionsModel.SetSharedDatasetType([]string{"testString"})
				createSharedDatasetOptionsModel.SetTags([]string{"testString"})
				createSharedDatasetOptionsModel.SetVersion("testString")
				createSharedDatasetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSharedDatasetOptionsModel).ToNot(BeNil())
				Expect(createSharedDatasetOptionsModel.AutoPropagateChange).To(Equal(core.BoolPtr(true)))
				Expect(createSharedDatasetOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createSharedDatasetOptionsModel.EffectedWorkspaceIds).To(Equal([]string{"testString"}))
				Expect(createSharedDatasetOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(createSharedDatasetOptionsModel.SharedDatasetData).To(Equal([]schematicsv1.SharedDatasetData{*sharedDatasetDataModel}))
				Expect(createSharedDatasetOptionsModel.SharedDatasetName).To(Equal(core.StringPtr("testString")))
				Expect(createSharedDatasetOptionsModel.SharedDatasetSourceName).To(Equal(core.StringPtr("testString")))
				Expect(createSharedDatasetOptionsModel.SharedDatasetType).To(Equal([]string{"testString"}))
				Expect(createSharedDatasetOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createSharedDatasetOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(createSharedDatasetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateWorkspaceDeletionJobOptions successfully`, func() {
				// Construct an instance of the CreateWorkspaceDeletionJobOptions model
				refreshToken := "testString"
				createWorkspaceDeletionJobOptionsModel := schematicsService.NewCreateWorkspaceDeletionJobOptions(refreshToken)
				createWorkspaceDeletionJobOptionsModel.SetRefreshToken("testString")
				createWorkspaceDeletionJobOptionsModel.SetNewDeleteWorkspaces(true)
				createWorkspaceDeletionJobOptionsModel.SetNewDestroyResources(true)
				createWorkspaceDeletionJobOptionsModel.SetNewJob("testString")
				createWorkspaceDeletionJobOptionsModel.SetNewVersion("testString")
				createWorkspaceDeletionJobOptionsModel.SetNewWorkspaces([]string{"testString"})
				createWorkspaceDeletionJobOptionsModel.SetDestroyResources("testString")
				createWorkspaceDeletionJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createWorkspaceDeletionJobOptionsModel).ToNot(BeNil())
				Expect(createWorkspaceDeletionJobOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceDeletionJobOptionsModel.NewDeleteWorkspaces).To(Equal(core.BoolPtr(true)))
				Expect(createWorkspaceDeletionJobOptionsModel.NewDestroyResources).To(Equal(core.BoolPtr(true)))
				Expect(createWorkspaceDeletionJobOptionsModel.NewJob).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceDeletionJobOptionsModel.NewVersion).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceDeletionJobOptionsModel.NewWorkspaces).To(Equal([]string{"testString"}))
				Expect(createWorkspaceDeletionJobOptionsModel.DestroyResources).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceDeletionJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateWorkspaceOptions successfully`, func() {
				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				Expect(catalogRefModel).ToNot(BeNil())
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")
				Expect(catalogRefModel.DryRun).To(Equal(core.BoolPtr(true)))
				Expect(catalogRefModel.ItemIconURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemID).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemName).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemReadmeURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.LaunchURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.OfferingVersion).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				Expect(sharedTargetDataModel).ToNot(BeNil())
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")
				Expect(sharedTargetDataModel.ClusterCreatedOn).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterName).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterType).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.EntitlementKeys).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(sharedTargetDataModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.WorkerCount).To(Equal(core.Int64Ptr(int64(26))))
				Expect(sharedTargetDataModel.WorkerMachineType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				Expect(workspaceVariableRequestModel).ToNot(BeNil())
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")
				Expect(workspaceVariableRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.UseDefault).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				Expect(templateSourceDataRequestModel).ToNot(BeNil())
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}
				Expect(templateSourceDataRequestModel.EnvValues).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(templateSourceDataRequestModel.Folder).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.InitStateFile).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.UninstallScriptName).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.Values).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.ValuesMetadata).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(templateSourceDataRequestModel.Variablestore).To(Equal([]schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}))

				// Construct an instance of the TemplateRepoRequest model
				templateRepoRequestModel := new(schematicsv1.TemplateRepoRequest)
				Expect(templateRepoRequestModel).ToNot(BeNil())
				templateRepoRequestModel.Branch = core.StringPtr("testString")
				templateRepoRequestModel.Release = core.StringPtr("testString")
				templateRepoRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoRequestModel.URL = core.StringPtr("testString")
				Expect(templateRepoRequestModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoRequestModel.Release).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoRequestModel.RepoShaValue).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoRequestModel.RepoURL).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoRequestModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceStatusRequest model
				workspaceStatusRequestModel := new(schematicsv1.WorkspaceStatusRequest)
				Expect(workspaceStatusRequestModel).ToNot(BeNil())
				workspaceStatusRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusRequestModel.LockedTime = CreateMockDateTime()
				Expect(workspaceStatusRequestModel.Frozen).To(Equal(core.BoolPtr(true)))
				Expect(workspaceStatusRequestModel.FrozenAt).To(Equal(CreateMockDateTime()))
				Expect(workspaceStatusRequestModel.FrozenBy).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusRequestModel.Locked).To(Equal(core.BoolPtr(true)))
				Expect(workspaceStatusRequestModel.LockedBy).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusRequestModel.LockedTime).To(Equal(CreateMockDateTime()))

				// Construct an instance of the CreateWorkspaceOptions model
				createWorkspaceOptionsModel := schematicsService.NewCreateWorkspaceOptions()
				createWorkspaceOptionsModel.SetAppliedShareddataIds([]string{"testString"})
				createWorkspaceOptionsModel.SetCatalogRef(catalogRefModel)
				createWorkspaceOptionsModel.SetDescription("testString")
				createWorkspaceOptionsModel.SetLocation("testString")
				createWorkspaceOptionsModel.SetName("testString")
				createWorkspaceOptionsModel.SetResourceGroup("testString")
				createWorkspaceOptionsModel.SetSharedData(sharedTargetDataModel)
				createWorkspaceOptionsModel.SetTags([]string{"testString"})
				createWorkspaceOptionsModel.SetTemplateData([]schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel})
				createWorkspaceOptionsModel.SetTemplateRef("testString")
				createWorkspaceOptionsModel.SetTemplateRepo(templateRepoRequestModel)
				createWorkspaceOptionsModel.SetType([]string{"testString"})
				createWorkspaceOptionsModel.SetWorkspaceStatus(workspaceStatusRequestModel)
				createWorkspaceOptionsModel.SetXGithubToken("testString")
				createWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createWorkspaceOptionsModel).ToNot(BeNil())
				Expect(createWorkspaceOptionsModel.AppliedShareddataIds).To(Equal([]string{"testString"}))
				Expect(createWorkspaceOptionsModel.CatalogRef).To(Equal(catalogRefModel))
				Expect(createWorkspaceOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.SharedData).To(Equal(sharedTargetDataModel))
				Expect(createWorkspaceOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createWorkspaceOptionsModel.TemplateData).To(Equal([]schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}))
				Expect(createWorkspaceOptionsModel.TemplateRef).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.TemplateRepo).To(Equal(templateRepoRequestModel))
				Expect(createWorkspaceOptionsModel.Type).To(Equal([]string{"testString"}))
				Expect(createWorkspaceOptionsModel.WorkspaceStatus).To(Equal(workspaceStatusRequestModel))
				Expect(createWorkspaceOptionsModel.XGithubToken).To(Equal(core.StringPtr("testString")))
				Expect(createWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteActionOptions successfully`, func() {
				// Construct an instance of the DeleteActionOptions model
				actionID := "testString"
				deleteActionOptionsModel := schematicsService.NewDeleteActionOptions(actionID)
				deleteActionOptionsModel.SetActionID("testString")
				deleteActionOptionsModel.SetForce(true)
				deleteActionOptionsModel.SetPropagate(true)
				deleteActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteActionOptionsModel).ToNot(BeNil())
				Expect(deleteActionOptionsModel.ActionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteActionOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(deleteActionOptionsModel.Propagate).To(Equal(core.BoolPtr(true)))
				Expect(deleteActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteJobOptions successfully`, func() {
				// Construct an instance of the DeleteJobOptions model
				jobID := "testString"
				refreshToken := "testString"
				deleteJobOptionsModel := schematicsService.NewDeleteJobOptions(jobID, refreshToken)
				deleteJobOptionsModel.SetJobID("testString")
				deleteJobOptionsModel.SetRefreshToken("testString")
				deleteJobOptionsModel.SetForce(true)
				deleteJobOptionsModel.SetPropagate(true)
				deleteJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteJobOptionsModel).ToNot(BeNil())
				Expect(deleteJobOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(deleteJobOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteJobOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(deleteJobOptionsModel.Propagate).To(Equal(core.BoolPtr(true)))
				Expect(deleteJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSharedDatasetOptions successfully`, func() {
				// Construct an instance of the DeleteSharedDatasetOptions model
				sdID := "testString"
				deleteSharedDatasetOptionsModel := schematicsService.NewDeleteSharedDatasetOptions(sdID)
				deleteSharedDatasetOptionsModel.SetSdID("testString")
				deleteSharedDatasetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSharedDatasetOptionsModel).ToNot(BeNil())
				Expect(deleteSharedDatasetOptionsModel.SdID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSharedDatasetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWorkspaceActivityOptions successfully`, func() {
				// Construct an instance of the DeleteWorkspaceActivityOptions model
				wID := "testString"
				activityID := "testString"
				deleteWorkspaceActivityOptionsModel := schematicsService.NewDeleteWorkspaceActivityOptions(wID, activityID)
				deleteWorkspaceActivityOptionsModel.SetWID("testString")
				deleteWorkspaceActivityOptionsModel.SetActivityID("testString")
				deleteWorkspaceActivityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWorkspaceActivityOptionsModel).ToNot(BeNil())
				Expect(deleteWorkspaceActivityOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkspaceActivityOptionsModel.ActivityID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkspaceActivityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWorkspaceOptions successfully`, func() {
				// Construct an instance of the DeleteWorkspaceOptions model
				wID := "testString"
				refreshToken := "testString"
				deleteWorkspaceOptionsModel := schematicsService.NewDeleteWorkspaceOptions(wID, refreshToken)
				deleteWorkspaceOptionsModel.SetWID("testString")
				deleteWorkspaceOptionsModel.SetRefreshToken("testString")
				deleteWorkspaceOptionsModel.SetDestroyResources("testString")
				deleteWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWorkspaceOptionsModel).ToNot(BeNil())
				Expect(deleteWorkspaceOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkspaceOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkspaceOptionsModel.DestroyResources).To(Equal(core.StringPtr("testString")))
				Expect(deleteWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDestroyWorkspaceCommandOptions successfully`, func() {
				// Construct an instance of the WorkspaceActivityOptionsTemplate model
				workspaceActivityOptionsTemplateModel := new(schematicsv1.WorkspaceActivityOptionsTemplate)
				Expect(workspaceActivityOptionsTemplateModel).ToNot(BeNil())
				workspaceActivityOptionsTemplateModel.Target = []string{"testString"}
				workspaceActivityOptionsTemplateModel.TfVars = []string{"testString"}
				Expect(workspaceActivityOptionsTemplateModel.Target).To(Equal([]string{"testString"}))
				Expect(workspaceActivityOptionsTemplateModel.TfVars).To(Equal([]string{"testString"}))

				// Construct an instance of the DestroyWorkspaceCommandOptions model
				wID := "testString"
				refreshToken := "testString"
				destroyWorkspaceCommandOptionsModel := schematicsService.NewDestroyWorkspaceCommandOptions(wID, refreshToken)
				destroyWorkspaceCommandOptionsModel.SetWID("testString")
				destroyWorkspaceCommandOptionsModel.SetRefreshToken("testString")
				destroyWorkspaceCommandOptionsModel.SetActionOptions(workspaceActivityOptionsTemplateModel)
				destroyWorkspaceCommandOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(destroyWorkspaceCommandOptionsModel).ToNot(BeNil())
				Expect(destroyWorkspaceCommandOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(destroyWorkspaceCommandOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(destroyWorkspaceCommandOptionsModel.ActionOptions).To(Equal(workspaceActivityOptionsTemplateModel))
				Expect(destroyWorkspaceCommandOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewExternalSource successfully`, func() {
				sourceType := "local"
				model, err := schematicsService.NewExternalSource(sourceType)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetActionOptions successfully`, func() {
				// Construct an instance of the GetActionOptions model
				actionID := "testString"
				getActionOptionsModel := schematicsService.NewGetActionOptions(actionID)
				getActionOptionsModel.SetActionID("testString")
				getActionOptionsModel.SetProfile("summary")
				getActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getActionOptionsModel).ToNot(BeNil())
				Expect(getActionOptionsModel.ActionID).To(Equal(core.StringPtr("testString")))
				Expect(getActionOptionsModel.Profile).To(Equal(core.StringPtr("summary")))
				Expect(getActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAllWorkspaceInputsOptions successfully`, func() {
				// Construct an instance of the GetAllWorkspaceInputsOptions model
				wID := "testString"
				getAllWorkspaceInputsOptionsModel := schematicsService.NewGetAllWorkspaceInputsOptions(wID)
				getAllWorkspaceInputsOptionsModel.SetWID("testString")
				getAllWorkspaceInputsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAllWorkspaceInputsOptionsModel).ToNot(BeNil())
				Expect(getAllWorkspaceInputsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getAllWorkspaceInputsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDiscoveredKmsInstancesOptions successfully`, func() {
				// Construct an instance of the GetDiscoveredKmsInstancesOptions model
				encryptionScheme := "testString"
				location := "testString"
				getDiscoveredKmsInstancesOptionsModel := schematicsService.NewGetDiscoveredKmsInstancesOptions(encryptionScheme, location)
				getDiscoveredKmsInstancesOptionsModel.SetEncryptionScheme("testString")
				getDiscoveredKmsInstancesOptionsModel.SetLocation("testString")
				getDiscoveredKmsInstancesOptionsModel.SetResourceGroup("testString")
				getDiscoveredKmsInstancesOptionsModel.SetLimit(int64(1))
				getDiscoveredKmsInstancesOptionsModel.SetSort("testString")
				getDiscoveredKmsInstancesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDiscoveredKmsInstancesOptionsModel).ToNot(BeNil())
				Expect(getDiscoveredKmsInstancesOptionsModel.EncryptionScheme).To(Equal(core.StringPtr("testString")))
				Expect(getDiscoveredKmsInstancesOptionsModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(getDiscoveredKmsInstancesOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(getDiscoveredKmsInstancesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(getDiscoveredKmsInstancesOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(getDiscoveredKmsInstancesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetJobOptions successfully`, func() {
				// Construct an instance of the GetJobOptions model
				jobID := "testString"
				getJobOptionsModel := schematicsService.NewGetJobOptions(jobID)
				getJobOptionsModel.SetJobID("testString")
				getJobOptionsModel.SetProfile("summary")
				getJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getJobOptionsModel).ToNot(BeNil())
				Expect(getJobOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(getJobOptionsModel.Profile).To(Equal(core.StringPtr("summary")))
				Expect(getJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetKmsSettingsOptions successfully`, func() {
				// Construct an instance of the GetKmsSettingsOptions model
				location := "testString"
				getKmsSettingsOptionsModel := schematicsService.NewGetKmsSettingsOptions(location)
				getKmsSettingsOptionsModel.SetLocation("testString")
				getKmsSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getKmsSettingsOptionsModel).ToNot(BeNil())
				Expect(getKmsSettingsOptionsModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(getKmsSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchematicsVersionOptions successfully`, func() {
				// Construct an instance of the GetSchematicsVersionOptions model
				getSchematicsVersionOptionsModel := schematicsService.NewGetSchematicsVersionOptions()
				getSchematicsVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchematicsVersionOptionsModel).ToNot(BeNil())
				Expect(getSchematicsVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSharedDatasetOptions successfully`, func() {
				// Construct an instance of the GetSharedDatasetOptions model
				sdID := "testString"
				getSharedDatasetOptionsModel := schematicsService.NewGetSharedDatasetOptions(sdID)
				getSharedDatasetOptionsModel.SetSdID("testString")
				getSharedDatasetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSharedDatasetOptionsModel).ToNot(BeNil())
				Expect(getSharedDatasetOptionsModel.SdID).To(Equal(core.StringPtr("testString")))
				Expect(getSharedDatasetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTemplateActivityLogOptions successfully`, func() {
				// Construct an instance of the GetTemplateActivityLogOptions model
				wID := "testString"
				tID := "testString"
				activityID := "testString"
				getTemplateActivityLogOptionsModel := schematicsService.NewGetTemplateActivityLogOptions(wID, tID, activityID)
				getTemplateActivityLogOptionsModel.SetWID("testString")
				getTemplateActivityLogOptionsModel.SetTID("testString")
				getTemplateActivityLogOptionsModel.SetActivityID("testString")
				getTemplateActivityLogOptionsModel.SetLogTfCmd(true)
				getTemplateActivityLogOptionsModel.SetLogTfPrefix(true)
				getTemplateActivityLogOptionsModel.SetLogTfNullResource(true)
				getTemplateActivityLogOptionsModel.SetLogTfAnsible(true)
				getTemplateActivityLogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTemplateActivityLogOptionsModel).ToNot(BeNil())
				Expect(getTemplateActivityLogOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateActivityLogOptionsModel.TID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateActivityLogOptionsModel.ActivityID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateActivityLogOptionsModel.LogTfCmd).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateActivityLogOptionsModel.LogTfPrefix).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateActivityLogOptionsModel.LogTfNullResource).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateActivityLogOptionsModel.LogTfAnsible).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateActivityLogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTemplateLogsOptions successfully`, func() {
				// Construct an instance of the GetTemplateLogsOptions model
				wID := "testString"
				tID := "testString"
				getTemplateLogsOptionsModel := schematicsService.NewGetTemplateLogsOptions(wID, tID)
				getTemplateLogsOptionsModel.SetWID("testString")
				getTemplateLogsOptionsModel.SetTID("testString")
				getTemplateLogsOptionsModel.SetLogTfCmd(true)
				getTemplateLogsOptionsModel.SetLogTfPrefix(true)
				getTemplateLogsOptionsModel.SetLogTfNullResource(true)
				getTemplateLogsOptionsModel.SetLogTfAnsible(true)
				getTemplateLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTemplateLogsOptionsModel).ToNot(BeNil())
				Expect(getTemplateLogsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateLogsOptionsModel.TID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateLogsOptionsModel.LogTfCmd).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateLogsOptionsModel.LogTfPrefix).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateLogsOptionsModel.LogTfNullResource).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateLogsOptionsModel.LogTfAnsible).To(Equal(core.BoolPtr(true)))
				Expect(getTemplateLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceActivityLogsOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceActivityLogsOptions model
				wID := "testString"
				activityID := "testString"
				getWorkspaceActivityLogsOptionsModel := schematicsService.NewGetWorkspaceActivityLogsOptions(wID, activityID)
				getWorkspaceActivityLogsOptionsModel.SetWID("testString")
				getWorkspaceActivityLogsOptionsModel.SetActivityID("testString")
				getWorkspaceActivityLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceActivityLogsOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceActivityLogsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceActivityLogsOptionsModel.ActivityID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceActivityLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceActivityOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceActivityOptions model
				wID := "testString"
				activityID := "testString"
				getWorkspaceActivityOptionsModel := schematicsService.NewGetWorkspaceActivityOptions(wID, activityID)
				getWorkspaceActivityOptionsModel.SetWID("testString")
				getWorkspaceActivityOptionsModel.SetActivityID("testString")
				getWorkspaceActivityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceActivityOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceActivityOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceActivityOptionsModel.ActivityID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceActivityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceDeletionJobStatusOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceDeletionJobStatusOptions model
				wjID := "testString"
				getWorkspaceDeletionJobStatusOptionsModel := schematicsService.NewGetWorkspaceDeletionJobStatusOptions(wjID)
				getWorkspaceDeletionJobStatusOptionsModel.SetWjID("testString")
				getWorkspaceDeletionJobStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceDeletionJobStatusOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceDeletionJobStatusOptionsModel.WjID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceDeletionJobStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceInputMetadataOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceInputMetadataOptions model
				wID := "testString"
				tID := "testString"
				getWorkspaceInputMetadataOptionsModel := schematicsService.NewGetWorkspaceInputMetadataOptions(wID, tID)
				getWorkspaceInputMetadataOptionsModel.SetWID("testString")
				getWorkspaceInputMetadataOptionsModel.SetTID("testString")
				getWorkspaceInputMetadataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceInputMetadataOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceInputMetadataOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceInputMetadataOptionsModel.TID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceInputMetadataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceInputsOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceInputsOptions model
				wID := "testString"
				tID := "testString"
				getWorkspaceInputsOptionsModel := schematicsService.NewGetWorkspaceInputsOptions(wID, tID)
				getWorkspaceInputsOptionsModel.SetWID("testString")
				getWorkspaceInputsOptionsModel.SetTID("testString")
				getWorkspaceInputsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceInputsOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceInputsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceInputsOptionsModel.TID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceInputsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceLogUrlsOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceLogUrlsOptions model
				wID := "testString"
				getWorkspaceLogUrlsOptionsModel := schematicsService.NewGetWorkspaceLogUrlsOptions(wID)
				getWorkspaceLogUrlsOptionsModel.SetWID("testString")
				getWorkspaceLogUrlsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceLogUrlsOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceLogUrlsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceLogUrlsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceOptions model
				wID := "testString"
				getWorkspaceOptionsModel := schematicsService.NewGetWorkspaceOptions(wID)
				getWorkspaceOptionsModel.SetWID("testString")
				getWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceOutputsOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceOutputsOptions model
				wID := "testString"
				getWorkspaceOutputsOptionsModel := schematicsService.NewGetWorkspaceOutputsOptions(wID)
				getWorkspaceOutputsOptionsModel.SetWID("testString")
				getWorkspaceOutputsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceOutputsOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceOutputsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceOutputsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceReadmeOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceReadmeOptions model
				wID := "testString"
				getWorkspaceReadmeOptionsModel := schematicsService.NewGetWorkspaceReadmeOptions(wID)
				getWorkspaceReadmeOptionsModel.SetWID("testString")
				getWorkspaceReadmeOptionsModel.SetRef("testString")
				getWorkspaceReadmeOptionsModel.SetFormatted("markdown")
				getWorkspaceReadmeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceReadmeOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceReadmeOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceReadmeOptionsModel.Ref).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceReadmeOptionsModel.Formatted).To(Equal(core.StringPtr("markdown")))
				Expect(getWorkspaceReadmeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceResourcesOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceResourcesOptions model
				wID := "testString"
				getWorkspaceResourcesOptionsModel := schematicsService.NewGetWorkspaceResourcesOptions(wID)
				getWorkspaceResourcesOptionsModel.SetWID("testString")
				getWorkspaceResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceResourcesOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceResourcesOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceStateOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceStateOptions model
				wID := "testString"
				getWorkspaceStateOptionsModel := schematicsService.NewGetWorkspaceStateOptions(wID)
				getWorkspaceStateOptionsModel.SetWID("testString")
				getWorkspaceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceStateOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceStateOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWorkspaceTemplateStateOptions successfully`, func() {
				// Construct an instance of the GetWorkspaceTemplateStateOptions model
				wID := "testString"
				tID := "testString"
				getWorkspaceTemplateStateOptionsModel := schematicsService.NewGetWorkspaceTemplateStateOptions(wID, tID)
				getWorkspaceTemplateStateOptionsModel.SetWID("testString")
				getWorkspaceTemplateStateOptionsModel.SetTID("testString")
				getWorkspaceTemplateStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWorkspaceTemplateStateOptionsModel).ToNot(BeNil())
				Expect(getWorkspaceTemplateStateOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceTemplateStateOptionsModel.TID).To(Equal(core.StringPtr("testString")))
				Expect(getWorkspaceTemplateStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewJobData successfully`, func() {
				jobType := "repo_download_job"
				model, err := schematicsService.NewJobData(jobType)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListActionsOptions successfully`, func() {
				// Construct an instance of the ListActionsOptions model
				listActionsOptionsModel := schematicsService.NewListActionsOptions()
				listActionsOptionsModel.SetOffset(int64(0))
				listActionsOptionsModel.SetLimit(int64(1))
				listActionsOptionsModel.SetSort("testString")
				listActionsOptionsModel.SetProfile("ids")
				listActionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listActionsOptionsModel).ToNot(BeNil())
				Expect(listActionsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listActionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listActionsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listActionsOptionsModel.Profile).To(Equal(core.StringPtr("ids")))
				Expect(listActionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListJobLogsOptions successfully`, func() {
				// Construct an instance of the ListJobLogsOptions model
				jobID := "testString"
				listJobLogsOptionsModel := schematicsService.NewListJobLogsOptions(jobID)
				listJobLogsOptionsModel.SetJobID("testString")
				listJobLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listJobLogsOptionsModel).ToNot(BeNil())
				Expect(listJobLogsOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(listJobLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListJobStatesOptions successfully`, func() {
				// Construct an instance of the ListJobStatesOptions model
				jobID := "testString"
				listJobStatesOptionsModel := schematicsService.NewListJobStatesOptions(jobID)
				listJobStatesOptionsModel.SetJobID("testString")
				listJobStatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listJobStatesOptionsModel).ToNot(BeNil())
				Expect(listJobStatesOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(listJobStatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListJobsOptions successfully`, func() {
				// Construct an instance of the ListJobsOptions model
				listJobsOptionsModel := schematicsService.NewListJobsOptions()
				listJobsOptionsModel.SetOffset(int64(0))
				listJobsOptionsModel.SetLimit(int64(1))
				listJobsOptionsModel.SetSort("testString")
				listJobsOptionsModel.SetProfile("ids")
				listJobsOptionsModel.SetResource("workspaces")
				listJobsOptionsModel.SetActionID("testString")
				listJobsOptionsModel.SetList("all")
				listJobsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listJobsOptionsModel).ToNot(BeNil())
				Expect(listJobsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listJobsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listJobsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listJobsOptionsModel.Profile).To(Equal(core.StringPtr("ids")))
				Expect(listJobsOptionsModel.Resource).To(Equal(core.StringPtr("workspaces")))
				Expect(listJobsOptionsModel.ActionID).To(Equal(core.StringPtr("testString")))
				Expect(listJobsOptionsModel.List).To(Equal(core.StringPtr("all")))
				Expect(listJobsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceGroupOptions successfully`, func() {
				// Construct an instance of the ListResourceGroupOptions model
				listResourceGroupOptionsModel := schematicsService.NewListResourceGroupOptions()
				listResourceGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceGroupOptionsModel).ToNot(BeNil())
				Expect(listResourceGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSchematicsLocationOptions successfully`, func() {
				// Construct an instance of the ListSchematicsLocationOptions model
				listSchematicsLocationOptionsModel := schematicsService.NewListSchematicsLocationOptions()
				listSchematicsLocationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSchematicsLocationOptionsModel).ToNot(BeNil())
				Expect(listSchematicsLocationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSharedDatasetsOptions successfully`, func() {
				// Construct an instance of the ListSharedDatasetsOptions model
				listSharedDatasetsOptionsModel := schematicsService.NewListSharedDatasetsOptions()
				listSharedDatasetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSharedDatasetsOptionsModel).ToNot(BeNil())
				Expect(listSharedDatasetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWorkspaceActivitiesOptions successfully`, func() {
				// Construct an instance of the ListWorkspaceActivitiesOptions model
				wID := "testString"
				listWorkspaceActivitiesOptionsModel := schematicsService.NewListWorkspaceActivitiesOptions(wID)
				listWorkspaceActivitiesOptionsModel.SetWID("testString")
				listWorkspaceActivitiesOptionsModel.SetOffset(int64(0))
				listWorkspaceActivitiesOptionsModel.SetLimit(int64(1))
				listWorkspaceActivitiesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWorkspaceActivitiesOptionsModel).ToNot(BeNil())
				Expect(listWorkspaceActivitiesOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(listWorkspaceActivitiesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listWorkspaceActivitiesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listWorkspaceActivitiesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWorkspacesOptions successfully`, func() {
				// Construct an instance of the ListWorkspacesOptions model
				listWorkspacesOptionsModel := schematicsService.NewListWorkspacesOptions()
				listWorkspacesOptionsModel.SetOffset(int64(0))
				listWorkspacesOptionsModel.SetLimit(int64(1))
				listWorkspacesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWorkspacesOptionsModel).ToNot(BeNil())
				Expect(listWorkspacesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listWorkspacesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listWorkspacesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPlanWorkspaceCommandOptions successfully`, func() {
				// Construct an instance of the PlanWorkspaceCommandOptions model
				wID := "testString"
				refreshToken := "testString"
				planWorkspaceCommandOptionsModel := schematicsService.NewPlanWorkspaceCommandOptions(wID, refreshToken)
				planWorkspaceCommandOptionsModel.SetWID("testString")
				planWorkspaceCommandOptionsModel.SetRefreshToken("testString")
				planWorkspaceCommandOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(planWorkspaceCommandOptionsModel).ToNot(BeNil())
				Expect(planWorkspaceCommandOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(planWorkspaceCommandOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(planWorkspaceCommandOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRefreshWorkspaceCommandOptions successfully`, func() {
				// Construct an instance of the RefreshWorkspaceCommandOptions model
				wID := "testString"
				refreshToken := "testString"
				refreshWorkspaceCommandOptionsModel := schematicsService.NewRefreshWorkspaceCommandOptions(wID, refreshToken)
				refreshWorkspaceCommandOptionsModel.SetWID("testString")
				refreshWorkspaceCommandOptionsModel.SetRefreshToken("testString")
				refreshWorkspaceCommandOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(refreshWorkspaceCommandOptionsModel).ToNot(BeNil())
				Expect(refreshWorkspaceCommandOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(refreshWorkspaceCommandOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(refreshWorkspaceCommandOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceJobOptions successfully`, func() {
				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				Expect(variableMetadataModel).ToNot(BeNil())
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")
				Expect(variableMetadataModel.Type).To(Equal(core.StringPtr("boolean")))
				Expect(variableMetadataModel.Aliases).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Options).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.MinValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MinLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.Matches).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Position).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.GroupBy).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Source).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				Expect(variableDataModel).ToNot(BeNil())
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel
				Expect(variableDataModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Metadata).To(Equal(variableMetadataModel))

				// Construct an instance of the JobStatusAction model
				jobStatusActionModel := new(schematicsv1.JobStatusAction)
				Expect(jobStatusActionModel).ToNot(BeNil())
				jobStatusActionModel.ActionName = core.StringPtr("testString")
				jobStatusActionModel.StatusCode = core.StringPtr("job_pending")
				jobStatusActionModel.StatusMessage = core.StringPtr("testString")
				jobStatusActionModel.BastionStatusCode = core.StringPtr("none")
				jobStatusActionModel.BastionStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.TargetsStatusCode = core.StringPtr("none")
				jobStatusActionModel.TargetsStatusMessage = core.StringPtr("testString")
				jobStatusActionModel.UpdatedAt = CreateMockDateTime()
				Expect(jobStatusActionModel.ActionName).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.StatusCode).To(Equal(core.StringPtr("job_pending")))
				Expect(jobStatusActionModel.StatusMessage).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.BastionStatusCode).To(Equal(core.StringPtr("none")))
				Expect(jobStatusActionModel.BastionStatusMessage).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.TargetsStatusCode).To(Equal(core.StringPtr("none")))
				Expect(jobStatusActionModel.TargetsStatusMessage).To(Equal(core.StringPtr("testString")))
				Expect(jobStatusActionModel.UpdatedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the JobStatus model
				jobStatusModel := new(schematicsv1.JobStatus)
				Expect(jobStatusModel).ToNot(BeNil())
				jobStatusModel.ActionJobStatus = jobStatusActionModel
				Expect(jobStatusModel.ActionJobStatus).To(Equal(jobStatusActionModel))

				// Construct an instance of the JobDataAction model
				jobDataActionModel := new(schematicsv1.JobDataAction)
				Expect(jobDataActionModel).ToNot(BeNil())
				jobDataActionModel.ActionName = core.StringPtr("testString")
				jobDataActionModel.Inputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Outputs = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.Settings = []schematicsv1.VariableData{*variableDataModel}
				jobDataActionModel.UpdatedAt = CreateMockDateTime()
				Expect(jobDataActionModel.ActionName).To(Equal(core.StringPtr("testString")))
				Expect(jobDataActionModel.Inputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(jobDataActionModel.Outputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(jobDataActionModel.Settings).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(jobDataActionModel.UpdatedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the JobData model
				jobDataModel := new(schematicsv1.JobData)
				Expect(jobDataModel).ToNot(BeNil())
				jobDataModel.JobType = core.StringPtr("repo_download_job")
				jobDataModel.ActionJobData = jobDataActionModel
				Expect(jobDataModel.JobType).To(Equal(core.StringPtr("repo_download_job")))
				Expect(jobDataModel.ActionJobData).To(Equal(jobDataActionModel))

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				Expect(systemLockModel).ToNot(BeNil())
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()
				Expect(systemLockModel.SysLocked).To(Equal(core.BoolPtr(true)))
				Expect(systemLockModel.SysLockedBy).To(Equal(core.StringPtr("testString")))
				Expect(systemLockModel.SysLockedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				Expect(targetResourcesetModel).ToNot(BeNil())
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel
				Expect(targetResourcesetModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.ResourceQuery).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.CredentialRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.SysLock).To(Equal(systemLockModel))

				// Construct an instance of the JobLogSummaryRepoDownloadJob model
				jobLogSummaryRepoDownloadJobModel := new(schematicsv1.JobLogSummaryRepoDownloadJob)
				Expect(jobLogSummaryRepoDownloadJobModel).ToNot(BeNil())

				// Construct an instance of the JobLogSummaryActionJobRecap model
				jobLogSummaryActionJobRecapModel := new(schematicsv1.JobLogSummaryActionJobRecap)
				Expect(jobLogSummaryActionJobRecapModel).ToNot(BeNil())
				jobLogSummaryActionJobRecapModel.Target = []string{"testString"}
				jobLogSummaryActionJobRecapModel.Ok = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Changed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Failed = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Skipped = core.Float64Ptr(float64(72.5))
				jobLogSummaryActionJobRecapModel.Unreachable = core.Float64Ptr(float64(72.5))
				Expect(jobLogSummaryActionJobRecapModel.Target).To(Equal([]string{"testString"}))
				Expect(jobLogSummaryActionJobRecapModel.Ok).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Changed).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Failed).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Skipped).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(jobLogSummaryActionJobRecapModel.Unreachable).To(Equal(core.Float64Ptr(float64(72.5))))

				// Construct an instance of the JobLogSummaryActionJob model
				jobLogSummaryActionJobModel := new(schematicsv1.JobLogSummaryActionJob)
				Expect(jobLogSummaryActionJobModel).ToNot(BeNil())
				jobLogSummaryActionJobModel.Recap = jobLogSummaryActionJobRecapModel
				Expect(jobLogSummaryActionJobModel.Recap).To(Equal(jobLogSummaryActionJobRecapModel))

				// Construct an instance of the JobLogSummary model
				jobLogSummaryModel := new(schematicsv1.JobLogSummary)
				Expect(jobLogSummaryModel).ToNot(BeNil())
				jobLogSummaryModel.JobType = core.StringPtr("repo_download_job")
				jobLogSummaryModel.RepoDownloadJob = jobLogSummaryRepoDownloadJobModel
				jobLogSummaryModel.ActionJob = jobLogSummaryActionJobModel
				Expect(jobLogSummaryModel.JobType).To(Equal(core.StringPtr("repo_download_job")))
				Expect(jobLogSummaryModel.RepoDownloadJob).To(Equal(jobLogSummaryRepoDownloadJobModel))
				Expect(jobLogSummaryModel.ActionJob).To(Equal(jobLogSummaryActionJobModel))

				// Construct an instance of the ReplaceJobOptions model
				jobID := "testString"
				refreshToken := "testString"
				replaceJobOptionsModel := schematicsService.NewReplaceJobOptions(jobID, refreshToken)
				replaceJobOptionsModel.SetJobID("testString")
				replaceJobOptionsModel.SetRefreshToken("testString")
				replaceJobOptionsModel.SetCommandObject("workspace")
				replaceJobOptionsModel.SetCommandObjectID("testString")
				replaceJobOptionsModel.SetCommandName("workspace_init_flow")
				replaceJobOptionsModel.SetCommandParameter("testString")
				replaceJobOptionsModel.SetCommandOptions([]string{"testString"})
				replaceJobOptionsModel.SetInputs([]schematicsv1.VariableData{*variableDataModel})
				replaceJobOptionsModel.SetSettings([]schematicsv1.VariableData{*variableDataModel})
				replaceJobOptionsModel.SetTags([]string{"testString"})
				replaceJobOptionsModel.SetLocation("us_south")
				replaceJobOptionsModel.SetStatus(jobStatusModel)
				replaceJobOptionsModel.SetData(jobDataModel)
				replaceJobOptionsModel.SetBastion(targetResourcesetModel)
				replaceJobOptionsModel.SetLogSummary(jobLogSummaryModel)
				replaceJobOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceJobOptionsModel).ToNot(BeNil())
				Expect(replaceJobOptionsModel.JobID).To(Equal(core.StringPtr("testString")))
				Expect(replaceJobOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(replaceJobOptionsModel.CommandObject).To(Equal(core.StringPtr("workspace")))
				Expect(replaceJobOptionsModel.CommandObjectID).To(Equal(core.StringPtr("testString")))
				Expect(replaceJobOptionsModel.CommandName).To(Equal(core.StringPtr("workspace_init_flow")))
				Expect(replaceJobOptionsModel.CommandParameter).To(Equal(core.StringPtr("testString")))
				Expect(replaceJobOptionsModel.CommandOptions).To(Equal([]string{"testString"}))
				Expect(replaceJobOptionsModel.Inputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(replaceJobOptionsModel.Settings).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(replaceJobOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(replaceJobOptionsModel.Location).To(Equal(core.StringPtr("us_south")))
				Expect(replaceJobOptionsModel.Status).To(Equal(jobStatusModel))
				Expect(replaceJobOptionsModel.Data).To(Equal(jobDataModel))
				Expect(replaceJobOptionsModel.Bastion).To(Equal(targetResourcesetModel))
				Expect(replaceJobOptionsModel.LogSummary).To(Equal(jobLogSummaryModel))
				Expect(replaceJobOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceKmsSettingsOptions successfully`, func() {
				// Construct an instance of the KMSSettingsPrimaryCrk model
				kmsSettingsPrimaryCrkModel := new(schematicsv1.KMSSettingsPrimaryCrk)
				Expect(kmsSettingsPrimaryCrkModel).ToNot(BeNil())
				kmsSettingsPrimaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsPrimaryCrkModel.KeyCrn = core.StringPtr("testString")
				Expect(kmsSettingsPrimaryCrkModel.KmsName).To(Equal(core.StringPtr("testString")))
				Expect(kmsSettingsPrimaryCrkModel.KmsPrivateEndpoint).To(Equal(core.StringPtr("testString")))
				Expect(kmsSettingsPrimaryCrkModel.KeyCrn).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the KMSSettingsSecondaryCrk model
				kmsSettingsSecondaryCrkModel := new(schematicsv1.KMSSettingsSecondaryCrk)
				Expect(kmsSettingsSecondaryCrkModel).ToNot(BeNil())
				kmsSettingsSecondaryCrkModel.KmsName = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KmsPrivateEndpoint = core.StringPtr("testString")
				kmsSettingsSecondaryCrkModel.KeyCrn = core.StringPtr("testString")
				Expect(kmsSettingsSecondaryCrkModel.KmsName).To(Equal(core.StringPtr("testString")))
				Expect(kmsSettingsSecondaryCrkModel.KmsPrivateEndpoint).To(Equal(core.StringPtr("testString")))
				Expect(kmsSettingsSecondaryCrkModel.KeyCrn).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceKmsSettingsOptions model
				replaceKmsSettingsOptionsModel := schematicsService.NewReplaceKmsSettingsOptions()
				replaceKmsSettingsOptionsModel.SetLocation("testString")
				replaceKmsSettingsOptionsModel.SetEncryptionScheme("testString")
				replaceKmsSettingsOptionsModel.SetResourceGroup("testString")
				replaceKmsSettingsOptionsModel.SetPrimaryCrk(kmsSettingsPrimaryCrkModel)
				replaceKmsSettingsOptionsModel.SetSecondaryCrk(kmsSettingsSecondaryCrkModel)
				replaceKmsSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceKmsSettingsOptionsModel).ToNot(BeNil())
				Expect(replaceKmsSettingsOptionsModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(replaceKmsSettingsOptionsModel.EncryptionScheme).To(Equal(core.StringPtr("testString")))
				Expect(replaceKmsSettingsOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(replaceKmsSettingsOptionsModel.PrimaryCrk).To(Equal(kmsSettingsPrimaryCrkModel))
				Expect(replaceKmsSettingsOptionsModel.SecondaryCrk).To(Equal(kmsSettingsSecondaryCrkModel))
				Expect(replaceKmsSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceSharedDatasetOptions successfully`, func() {
				// Construct an instance of the SharedDatasetData model
				sharedDatasetDataModel := new(schematicsv1.SharedDatasetData)
				Expect(sharedDatasetDataModel).ToNot(BeNil())
				sharedDatasetDataModel.DefaultValue = core.StringPtr("testString")
				sharedDatasetDataModel.Description = core.StringPtr("testString")
				sharedDatasetDataModel.Hidden = core.BoolPtr(true)
				sharedDatasetDataModel.Immutable = core.BoolPtr(true)
				sharedDatasetDataModel.Matches = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValue = core.StringPtr("testString")
				sharedDatasetDataModel.MaxValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.MinValue = core.StringPtr("testString")
				sharedDatasetDataModel.MinValueLen = core.StringPtr("testString")
				sharedDatasetDataModel.Options = []string{"testString"}
				sharedDatasetDataModel.OverrideValue = core.StringPtr("testString")
				sharedDatasetDataModel.Secure = core.BoolPtr(true)
				sharedDatasetDataModel.VarAliases = []string{"testString"}
				sharedDatasetDataModel.VarName = core.StringPtr("testString")
				sharedDatasetDataModel.VarRef = core.StringPtr("testString")
				sharedDatasetDataModel.VarType = core.StringPtr("testString")
				Expect(sharedDatasetDataModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(sharedDatasetDataModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(sharedDatasetDataModel.Matches).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MaxValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MaxValueLen).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MinValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.MinValueLen).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Options).To(Equal([]string{"testString"}))
				Expect(sharedDatasetDataModel.OverrideValue).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(sharedDatasetDataModel.VarAliases).To(Equal([]string{"testString"}))
				Expect(sharedDatasetDataModel.VarName).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.VarRef).To(Equal(core.StringPtr("testString")))
				Expect(sharedDatasetDataModel.VarType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceSharedDatasetOptions model
				sdID := "testString"
				replaceSharedDatasetOptionsModel := schematicsService.NewReplaceSharedDatasetOptions(sdID)
				replaceSharedDatasetOptionsModel.SetSdID("testString")
				replaceSharedDatasetOptionsModel.SetAutoPropagateChange(true)
				replaceSharedDatasetOptionsModel.SetDescription("testString")
				replaceSharedDatasetOptionsModel.SetEffectedWorkspaceIds([]string{"testString"})
				replaceSharedDatasetOptionsModel.SetResourceGroup("testString")
				replaceSharedDatasetOptionsModel.SetSharedDatasetData([]schematicsv1.SharedDatasetData{*sharedDatasetDataModel})
				replaceSharedDatasetOptionsModel.SetSharedDatasetName("testString")
				replaceSharedDatasetOptionsModel.SetSharedDatasetSourceName("testString")
				replaceSharedDatasetOptionsModel.SetSharedDatasetType([]string{"testString"})
				replaceSharedDatasetOptionsModel.SetTags([]string{"testString"})
				replaceSharedDatasetOptionsModel.SetVersion("testString")
				replaceSharedDatasetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceSharedDatasetOptionsModel).ToNot(BeNil())
				Expect(replaceSharedDatasetOptionsModel.SdID).To(Equal(core.StringPtr("testString")))
				Expect(replaceSharedDatasetOptionsModel.AutoPropagateChange).To(Equal(core.BoolPtr(true)))
				Expect(replaceSharedDatasetOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replaceSharedDatasetOptionsModel.EffectedWorkspaceIds).To(Equal([]string{"testString"}))
				Expect(replaceSharedDatasetOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(replaceSharedDatasetOptionsModel.SharedDatasetData).To(Equal([]schematicsv1.SharedDatasetData{*sharedDatasetDataModel}))
				Expect(replaceSharedDatasetOptionsModel.SharedDatasetName).To(Equal(core.StringPtr("testString")))
				Expect(replaceSharedDatasetOptionsModel.SharedDatasetSourceName).To(Equal(core.StringPtr("testString")))
				Expect(replaceSharedDatasetOptionsModel.SharedDatasetType).To(Equal([]string{"testString"}))
				Expect(replaceSharedDatasetOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(replaceSharedDatasetOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(replaceSharedDatasetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceWorkspaceInputsOptions successfully`, func() {
				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				Expect(workspaceVariableRequestModel).ToNot(BeNil())
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")
				Expect(workspaceVariableRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.UseDefault).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceWorkspaceInputsOptions model
				wID := "testString"
				tID := "testString"
				replaceWorkspaceInputsOptionsModel := schematicsService.NewReplaceWorkspaceInputsOptions(wID, tID)
				replaceWorkspaceInputsOptionsModel.SetWID("testString")
				replaceWorkspaceInputsOptionsModel.SetTID("testString")
				replaceWorkspaceInputsOptionsModel.SetEnvValues([]interface{}{map[string]interface{}{"anyKey": "anyValue"}})
				replaceWorkspaceInputsOptionsModel.SetValues("testString")
				replaceWorkspaceInputsOptionsModel.SetVariablestore([]schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel})
				replaceWorkspaceInputsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceWorkspaceInputsOptionsModel).ToNot(BeNil())
				Expect(replaceWorkspaceInputsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(replaceWorkspaceInputsOptionsModel.TID).To(Equal(core.StringPtr("testString")))
				Expect(replaceWorkspaceInputsOptionsModel.EnvValues).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(replaceWorkspaceInputsOptionsModel.Values).To(Equal(core.StringPtr("testString")))
				Expect(replaceWorkspaceInputsOptionsModel.Variablestore).To(Equal([]schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}))
				Expect(replaceWorkspaceInputsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceWorkspaceOptions successfully`, func() {
				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				Expect(catalogRefModel).ToNot(BeNil())
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")
				Expect(catalogRefModel.DryRun).To(Equal(core.BoolPtr(true)))
				Expect(catalogRefModel.ItemIconURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemID).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemName).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemReadmeURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.LaunchURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.OfferingVersion).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				Expect(sharedTargetDataModel).ToNot(BeNil())
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")
				Expect(sharedTargetDataModel.ClusterCreatedOn).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterName).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterType).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.EntitlementKeys).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(sharedTargetDataModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.WorkerCount).To(Equal(core.Int64Ptr(int64(26))))
				Expect(sharedTargetDataModel.WorkerMachineType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				Expect(workspaceVariableRequestModel).ToNot(BeNil())
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")
				Expect(workspaceVariableRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.UseDefault).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				Expect(templateSourceDataRequestModel).ToNot(BeNil())
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}
				Expect(templateSourceDataRequestModel.EnvValues).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(templateSourceDataRequestModel.Folder).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.InitStateFile).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.UninstallScriptName).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.Values).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.ValuesMetadata).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(templateSourceDataRequestModel.Variablestore).To(Equal([]schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}))

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				Expect(templateRepoUpdateRequestModel).ToNot(BeNil())
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")
				Expect(templateRepoUpdateRequestModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.Release).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.RepoShaValue).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.RepoURL).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				Expect(workspaceStatusUpdateRequestModel).ToNot(BeNil())
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()
				Expect(workspaceStatusUpdateRequestModel.Frozen).To(Equal(core.BoolPtr(true)))
				Expect(workspaceStatusUpdateRequestModel.FrozenAt).To(Equal(CreateMockDateTime()))
				Expect(workspaceStatusUpdateRequestModel.FrozenBy).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusUpdateRequestModel.Locked).To(Equal(core.BoolPtr(true)))
				Expect(workspaceStatusUpdateRequestModel.LockedBy).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusUpdateRequestModel.LockedTime).To(Equal(CreateMockDateTime()))

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				Expect(workspaceStatusMessageModel).ToNot(BeNil())
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")
				Expect(workspaceStatusMessageModel.StatusCode).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusMessageModel.StatusMsg).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceWorkspaceOptions model
				wID := "testString"
				replaceWorkspaceOptionsModel := schematicsService.NewReplaceWorkspaceOptions(wID)
				replaceWorkspaceOptionsModel.SetWID("testString")
				replaceWorkspaceOptionsModel.SetCatalogRef(catalogRefModel)
				replaceWorkspaceOptionsModel.SetDescription("testString")
				replaceWorkspaceOptionsModel.SetName("testString")
				replaceWorkspaceOptionsModel.SetSharedData(sharedTargetDataModel)
				replaceWorkspaceOptionsModel.SetTags([]string{"testString"})
				replaceWorkspaceOptionsModel.SetTemplateData([]schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel})
				replaceWorkspaceOptionsModel.SetTemplateRepo(templateRepoUpdateRequestModel)
				replaceWorkspaceOptionsModel.SetType([]string{"testString"})
				replaceWorkspaceOptionsModel.SetWorkspaceStatus(workspaceStatusUpdateRequestModel)
				replaceWorkspaceOptionsModel.SetWorkspaceStatusMsg(workspaceStatusMessageModel)
				replaceWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceWorkspaceOptionsModel).ToNot(BeNil())
				Expect(replaceWorkspaceOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(replaceWorkspaceOptionsModel.CatalogRef).To(Equal(catalogRefModel))
				Expect(replaceWorkspaceOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replaceWorkspaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(replaceWorkspaceOptionsModel.SharedData).To(Equal(sharedTargetDataModel))
				Expect(replaceWorkspaceOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(replaceWorkspaceOptionsModel.TemplateData).To(Equal([]schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}))
				Expect(replaceWorkspaceOptionsModel.TemplateRepo).To(Equal(templateRepoUpdateRequestModel))
				Expect(replaceWorkspaceOptionsModel.Type).To(Equal([]string{"testString"}))
				Expect(replaceWorkspaceOptionsModel.WorkspaceStatus).To(Equal(workspaceStatusUpdateRequestModel))
				Expect(replaceWorkspaceOptionsModel.WorkspaceStatusMsg).To(Equal(workspaceStatusMessageModel))
				Expect(replaceWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunWorkspaceCommandsOptions successfully`, func() {
				// Construct an instance of the TerraformCommand model
				terraformCommandModel := new(schematicsv1.TerraformCommand)
				Expect(terraformCommandModel).ToNot(BeNil())
				terraformCommandModel.Command = core.StringPtr("testString")
				terraformCommandModel.CommandParams = core.StringPtr("testString")
				terraformCommandModel.CommandName = core.StringPtr("testString")
				terraformCommandModel.CommandDesc = core.StringPtr("testString")
				terraformCommandModel.CommandOnError = core.StringPtr("testString")
				terraformCommandModel.CommandDependsOn = core.StringPtr("testString")
				terraformCommandModel.CommandStatus = core.StringPtr("testString")
				Expect(terraformCommandModel.Command).To(Equal(core.StringPtr("testString")))
				Expect(terraformCommandModel.CommandParams).To(Equal(core.StringPtr("testString")))
				Expect(terraformCommandModel.CommandName).To(Equal(core.StringPtr("testString")))
				Expect(terraformCommandModel.CommandDesc).To(Equal(core.StringPtr("testString")))
				Expect(terraformCommandModel.CommandOnError).To(Equal(core.StringPtr("testString")))
				Expect(terraformCommandModel.CommandDependsOn).To(Equal(core.StringPtr("testString")))
				Expect(terraformCommandModel.CommandStatus).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the RunWorkspaceCommandsOptions model
				wID := "testString"
				refreshToken := "testString"
				runWorkspaceCommandsOptionsModel := schematicsService.NewRunWorkspaceCommandsOptions(wID, refreshToken)
				runWorkspaceCommandsOptionsModel.SetWID("testString")
				runWorkspaceCommandsOptionsModel.SetRefreshToken("testString")
				runWorkspaceCommandsOptionsModel.SetCommands([]schematicsv1.TerraformCommand{*terraformCommandModel})
				runWorkspaceCommandsOptionsModel.SetOperationName("testString")
				runWorkspaceCommandsOptionsModel.SetDescription("testString")
				runWorkspaceCommandsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runWorkspaceCommandsOptionsModel).ToNot(BeNil())
				Expect(runWorkspaceCommandsOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(runWorkspaceCommandsOptionsModel.RefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(runWorkspaceCommandsOptionsModel.Commands).To(Equal([]schematicsv1.TerraformCommand{*terraformCommandModel}))
				Expect(runWorkspaceCommandsOptionsModel.OperationName).To(Equal(core.StringPtr("testString")))
				Expect(runWorkspaceCommandsOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(runWorkspaceCommandsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateActionOptions successfully`, func() {
				// Construct an instance of the UserState model
				userStateModel := new(schematicsv1.UserState)
				Expect(userStateModel).ToNot(BeNil())
				userStateModel.State = core.StringPtr("draft")
				userStateModel.SetBy = core.StringPtr("testString")
				userStateModel.SetAt = CreateMockDateTime()
				Expect(userStateModel.State).To(Equal(core.StringPtr("draft")))
				Expect(userStateModel.SetBy).To(Equal(core.StringPtr("testString")))
				Expect(userStateModel.SetAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the ExternalSourceGit model
				externalSourceGitModel := new(schematicsv1.ExternalSourceGit)
				Expect(externalSourceGitModel).ToNot(BeNil())
				externalSourceGitModel.GitRepoURL = core.StringPtr("testString")
				externalSourceGitModel.GitToken = core.StringPtr("testString")
				externalSourceGitModel.GitRepoFolder = core.StringPtr("testString")
				externalSourceGitModel.GitRelease = core.StringPtr("testString")
				externalSourceGitModel.GitBranch = core.StringPtr("testString")
				Expect(externalSourceGitModel.GitRepoURL).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitToken).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitRepoFolder).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitRelease).To(Equal(core.StringPtr("testString")))
				Expect(externalSourceGitModel.GitBranch).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ExternalSource model
				externalSourceModel := new(schematicsv1.ExternalSource)
				Expect(externalSourceModel).ToNot(BeNil())
				externalSourceModel.SourceType = core.StringPtr("local")
				externalSourceModel.Git = externalSourceGitModel
				Expect(externalSourceModel.SourceType).To(Equal(core.StringPtr("local")))
				Expect(externalSourceModel.Git).To(Equal(externalSourceGitModel))

				// Construct an instance of the SystemLock model
				systemLockModel := new(schematicsv1.SystemLock)
				Expect(systemLockModel).ToNot(BeNil())
				systemLockModel.SysLocked = core.BoolPtr(true)
				systemLockModel.SysLockedBy = core.StringPtr("testString")
				systemLockModel.SysLockedAt = CreateMockDateTime()
				Expect(systemLockModel.SysLocked).To(Equal(core.BoolPtr(true)))
				Expect(systemLockModel.SysLockedBy).To(Equal(core.StringPtr("testString")))
				Expect(systemLockModel.SysLockedAt).To(Equal(CreateMockDateTime()))

				// Construct an instance of the TargetResourceset model
				targetResourcesetModel := new(schematicsv1.TargetResourceset)
				Expect(targetResourcesetModel).ToNot(BeNil())
				targetResourcesetModel.Name = core.StringPtr("testString")
				targetResourcesetModel.Type = core.StringPtr("testString")
				targetResourcesetModel.Description = core.StringPtr("testString")
				targetResourcesetModel.ResourceQuery = core.StringPtr("testString")
				targetResourcesetModel.CredentialRef = core.StringPtr("testString")
				targetResourcesetModel.SysLock = systemLockModel
				Expect(targetResourcesetModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.ResourceQuery).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.CredentialRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourcesetModel.SysLock).To(Equal(systemLockModel))

				// Construct an instance of the TargetResourceConfig model
				targetResourceConfigModel := new(schematicsv1.TargetResourceConfig)
				Expect(targetResourceConfigModel).ToNot(BeNil())
				targetResourceConfigModel.Name = core.StringPtr("testString")
				targetResourceConfigModel.Value = core.StringPtr("testString")
				targetResourceConfigModel.Description = core.StringPtr("testString")
				Expect(targetResourceConfigModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceConfigModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceConfigModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TargetResource model
				targetResourceModel := new(schematicsv1.TargetResource)
				Expect(targetResourceModel).ToNot(BeNil())
				targetResourceModel.ResourceID = core.StringPtr("testString")
				targetResourceModel.ResourceConfigs = []schematicsv1.TargetResourceConfig{*targetResourceConfigModel}
				Expect(targetResourceModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceModel.ResourceConfigs).To(Equal([]schematicsv1.TargetResourceConfig{*targetResourceConfigModel}))

				// Construct an instance of the TargetResourceGroup model
				targetResourceGroupModel := new(schematicsv1.TargetResourceGroup)
				Expect(targetResourceGroupModel).ToNot(BeNil())
				targetResourceGroupModel.Name = core.StringPtr("testString")
				targetResourceGroupModel.Description = core.StringPtr("testString")
				targetResourceGroupModel.CredentialRef = core.StringPtr("testString")
				targetResourceGroupModel.BastionRef = core.StringPtr("testString")
				targetResourceGroupModel.TargetResources = []schematicsv1.TargetResource{*targetResourceModel}
				Expect(targetResourceGroupModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.CredentialRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.BastionRef).To(Equal(core.StringPtr("testString")))
				Expect(targetResourceGroupModel.TargetResources).To(Equal([]schematicsv1.TargetResource{*targetResourceModel}))

				// Construct an instance of the VariableMetadata model
				variableMetadataModel := new(schematicsv1.VariableMetadata)
				Expect(variableMetadataModel).ToNot(BeNil())
				variableMetadataModel.Type = core.StringPtr("boolean")
				variableMetadataModel.Aliases = []string{"testString"}
				variableMetadataModel.Description = core.StringPtr("testString")
				variableMetadataModel.DefaultValue = core.StringPtr("testString")
				variableMetadataModel.Secure = core.BoolPtr(true)
				variableMetadataModel.Immutable = core.BoolPtr(true)
				variableMetadataModel.Hidden = core.BoolPtr(true)
				variableMetadataModel.Options = []string{"testString"}
				variableMetadataModel.MinValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxValue = core.Int64Ptr(int64(38))
				variableMetadataModel.MinLength = core.Int64Ptr(int64(38))
				variableMetadataModel.MaxLength = core.Int64Ptr(int64(38))
				variableMetadataModel.Matches = core.StringPtr("testString")
				variableMetadataModel.Position = core.Int64Ptr(int64(38))
				variableMetadataModel.GroupBy = core.StringPtr("testString")
				variableMetadataModel.Source = core.StringPtr("testString")
				Expect(variableMetadataModel.Type).To(Equal(core.StringPtr("boolean")))
				Expect(variableMetadataModel.Aliases).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Immutable).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(variableMetadataModel.Options).To(Equal([]string{"testString"}))
				Expect(variableMetadataModel.MinValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxValue).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MinLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.MaxLength).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.Matches).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Position).To(Equal(core.Int64Ptr(int64(38))))
				Expect(variableMetadataModel.GroupBy).To(Equal(core.StringPtr("testString")))
				Expect(variableMetadataModel.Source).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VariableData model
				variableDataModel := new(schematicsv1.VariableData)
				Expect(variableDataModel).ToNot(BeNil())
				variableDataModel.Name = core.StringPtr("testString")
				variableDataModel.Value = core.StringPtr("testString")
				variableDataModel.Metadata = variableMetadataModel
				Expect(variableDataModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(variableDataModel.Metadata).To(Equal(variableMetadataModel))

				// Construct an instance of the ActionState model
				actionStateModel := new(schematicsv1.ActionState)
				Expect(actionStateModel).ToNot(BeNil())
				actionStateModel.StatusCode = core.StringPtr("normal")
				actionStateModel.StatusJobID = core.StringPtr("testString")
				actionStateModel.StatusMessage = core.StringPtr("testString")
				Expect(actionStateModel.StatusCode).To(Equal(core.StringPtr("normal")))
				Expect(actionStateModel.StatusJobID).To(Equal(core.StringPtr("testString")))
				Expect(actionStateModel.StatusMessage).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateActionOptions model
				actionID := "testString"
				updateActionOptionsModel := schematicsService.NewUpdateActionOptions(actionID)
				updateActionOptionsModel.SetActionID("testString")
				updateActionOptionsModel.SetName("Stop Action")
				updateActionOptionsModel.SetDescription("This Action can be used to Stop the targets")
				updateActionOptionsModel.SetLocation("us_south")
				updateActionOptionsModel.SetResourceGroup("testString")
				updateActionOptionsModel.SetTags([]string{"testString"})
				updateActionOptionsModel.SetUserState(userStateModel)
				updateActionOptionsModel.SetSourceReadmeURL("testString")
				updateActionOptionsModel.SetSource(externalSourceModel)
				updateActionOptionsModel.SetSourceType("local")
				updateActionOptionsModel.SetCommandParameter("testString")
				updateActionOptionsModel.SetBastion(targetResourcesetModel)
				updateActionOptionsModel.SetTargets([]schematicsv1.TargetResourceGroup{*targetResourceGroupModel})
				updateActionOptionsModel.SetCredentials([]schematicsv1.VariableData{*variableDataModel})
				updateActionOptionsModel.SetInputs([]schematicsv1.VariableData{*variableDataModel})
				updateActionOptionsModel.SetOutputs([]schematicsv1.VariableData{*variableDataModel})
				updateActionOptionsModel.SetSettings([]schematicsv1.VariableData{*variableDataModel})
				updateActionOptionsModel.SetTriggerRecordID("testString")
				updateActionOptionsModel.SetState(actionStateModel)
				updateActionOptionsModel.SetSysLock(systemLockModel)
				updateActionOptionsModel.SetXGithubToken("testString")
				updateActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateActionOptionsModel).ToNot(BeNil())
				Expect(updateActionOptionsModel.ActionID).To(Equal(core.StringPtr("testString")))
				Expect(updateActionOptionsModel.Name).To(Equal(core.StringPtr("Stop Action")))
				Expect(updateActionOptionsModel.Description).To(Equal(core.StringPtr("This Action can be used to Stop the targets")))
				Expect(updateActionOptionsModel.Location).To(Equal(core.StringPtr("us_south")))
				Expect(updateActionOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(updateActionOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(updateActionOptionsModel.UserState).To(Equal(userStateModel))
				Expect(updateActionOptionsModel.SourceReadmeURL).To(Equal(core.StringPtr("testString")))
				Expect(updateActionOptionsModel.Source).To(Equal(externalSourceModel))
				Expect(updateActionOptionsModel.SourceType).To(Equal(core.StringPtr("local")))
				Expect(updateActionOptionsModel.CommandParameter).To(Equal(core.StringPtr("testString")))
				Expect(updateActionOptionsModel.Bastion).To(Equal(targetResourcesetModel))
				Expect(updateActionOptionsModel.Targets).To(Equal([]schematicsv1.TargetResourceGroup{*targetResourceGroupModel}))
				Expect(updateActionOptionsModel.Credentials).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(updateActionOptionsModel.Inputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(updateActionOptionsModel.Outputs).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(updateActionOptionsModel.Settings).To(Equal([]schematicsv1.VariableData{*variableDataModel}))
				Expect(updateActionOptionsModel.TriggerRecordID).To(Equal(core.StringPtr("testString")))
				Expect(updateActionOptionsModel.State).To(Equal(actionStateModel))
				Expect(updateActionOptionsModel.SysLock).To(Equal(systemLockModel))
				Expect(updateActionOptionsModel.XGithubToken).To(Equal(core.StringPtr("testString")))
				Expect(updateActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateWorkspaceOptions successfully`, func() {
				// Construct an instance of the CatalogRef model
				catalogRefModel := new(schematicsv1.CatalogRef)
				Expect(catalogRefModel).ToNot(BeNil())
				catalogRefModel.DryRun = core.BoolPtr(true)
				catalogRefModel.ItemIconURL = core.StringPtr("testString")
				catalogRefModel.ItemID = core.StringPtr("testString")
				catalogRefModel.ItemName = core.StringPtr("testString")
				catalogRefModel.ItemReadmeURL = core.StringPtr("testString")
				catalogRefModel.ItemURL = core.StringPtr("testString")
				catalogRefModel.LaunchURL = core.StringPtr("testString")
				catalogRefModel.OfferingVersion = core.StringPtr("testString")
				Expect(catalogRefModel.DryRun).To(Equal(core.BoolPtr(true)))
				Expect(catalogRefModel.ItemIconURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemID).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemName).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemReadmeURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.ItemURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.LaunchURL).To(Equal(core.StringPtr("testString")))
				Expect(catalogRefModel.OfferingVersion).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SharedTargetData model
				sharedTargetDataModel := new(schematicsv1.SharedTargetData)
				Expect(sharedTargetDataModel).ToNot(BeNil())
				sharedTargetDataModel.ClusterCreatedOn = core.StringPtr("testString")
				sharedTargetDataModel.ClusterID = core.StringPtr("testString")
				sharedTargetDataModel.ClusterName = core.StringPtr("testString")
				sharedTargetDataModel.ClusterType = core.StringPtr("testString")
				sharedTargetDataModel.EntitlementKeys = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				sharedTargetDataModel.Namespace = core.StringPtr("testString")
				sharedTargetDataModel.Region = core.StringPtr("testString")
				sharedTargetDataModel.ResourceGroupID = core.StringPtr("testString")
				sharedTargetDataModel.WorkerCount = core.Int64Ptr(int64(26))
				sharedTargetDataModel.WorkerMachineType = core.StringPtr("testString")
				Expect(sharedTargetDataModel.ClusterCreatedOn).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterName).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ClusterType).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.EntitlementKeys).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(sharedTargetDataModel.Namespace).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(sharedTargetDataModel.WorkerCount).To(Equal(core.Int64Ptr(int64(26))))
				Expect(sharedTargetDataModel.WorkerMachineType).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceVariableRequest model
				workspaceVariableRequestModel := new(schematicsv1.WorkspaceVariableRequest)
				Expect(workspaceVariableRequestModel).ToNot(BeNil())
				workspaceVariableRequestModel.Description = core.StringPtr("testString")
				workspaceVariableRequestModel.Name = core.StringPtr("testString")
				workspaceVariableRequestModel.Secure = core.BoolPtr(true)
				workspaceVariableRequestModel.Type = core.StringPtr("testString")
				workspaceVariableRequestModel.UseDefault = core.BoolPtr(true)
				workspaceVariableRequestModel.Value = core.StringPtr("testString")
				Expect(workspaceVariableRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.Secure).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(workspaceVariableRequestModel.UseDefault).To(Equal(core.BoolPtr(true)))
				Expect(workspaceVariableRequestModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the TemplateSourceDataRequest model
				templateSourceDataRequestModel := new(schematicsv1.TemplateSourceDataRequest)
				Expect(templateSourceDataRequestModel).ToNot(BeNil())
				templateSourceDataRequestModel.EnvValues = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Folder = core.StringPtr("testString")
				templateSourceDataRequestModel.InitStateFile = core.StringPtr("testString")
				templateSourceDataRequestModel.Type = core.StringPtr("testString")
				templateSourceDataRequestModel.UninstallScriptName = core.StringPtr("testString")
				templateSourceDataRequestModel.Values = core.StringPtr("testString")
				templateSourceDataRequestModel.ValuesMetadata = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				templateSourceDataRequestModel.Variablestore = []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}
				Expect(templateSourceDataRequestModel.EnvValues).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(templateSourceDataRequestModel.Folder).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.InitStateFile).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.UninstallScriptName).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.Values).To(Equal(core.StringPtr("testString")))
				Expect(templateSourceDataRequestModel.ValuesMetadata).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(templateSourceDataRequestModel.Variablestore).To(Equal([]schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel}))

				// Construct an instance of the TemplateRepoUpdateRequest model
				templateRepoUpdateRequestModel := new(schematicsv1.TemplateRepoUpdateRequest)
				Expect(templateRepoUpdateRequestModel).ToNot(BeNil())
				templateRepoUpdateRequestModel.Branch = core.StringPtr("testString")
				templateRepoUpdateRequestModel.Release = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoShaValue = core.StringPtr("testString")
				templateRepoUpdateRequestModel.RepoURL = core.StringPtr("testString")
				templateRepoUpdateRequestModel.URL = core.StringPtr("testString")
				Expect(templateRepoUpdateRequestModel.Branch).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.Release).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.RepoShaValue).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.RepoURL).To(Equal(core.StringPtr("testString")))
				Expect(templateRepoUpdateRequestModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the WorkspaceStatusUpdateRequest model
				workspaceStatusUpdateRequestModel := new(schematicsv1.WorkspaceStatusUpdateRequest)
				Expect(workspaceStatusUpdateRequestModel).ToNot(BeNil())
				workspaceStatusUpdateRequestModel.Frozen = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.FrozenAt = CreateMockDateTime()
				workspaceStatusUpdateRequestModel.FrozenBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.Locked = core.BoolPtr(true)
				workspaceStatusUpdateRequestModel.LockedBy = core.StringPtr("testString")
				workspaceStatusUpdateRequestModel.LockedTime = CreateMockDateTime()
				Expect(workspaceStatusUpdateRequestModel.Frozen).To(Equal(core.BoolPtr(true)))
				Expect(workspaceStatusUpdateRequestModel.FrozenAt).To(Equal(CreateMockDateTime()))
				Expect(workspaceStatusUpdateRequestModel.FrozenBy).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusUpdateRequestModel.Locked).To(Equal(core.BoolPtr(true)))
				Expect(workspaceStatusUpdateRequestModel.LockedBy).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusUpdateRequestModel.LockedTime).To(Equal(CreateMockDateTime()))

				// Construct an instance of the WorkspaceStatusMessage model
				workspaceStatusMessageModel := new(schematicsv1.WorkspaceStatusMessage)
				Expect(workspaceStatusMessageModel).ToNot(BeNil())
				workspaceStatusMessageModel.StatusCode = core.StringPtr("testString")
				workspaceStatusMessageModel.StatusMsg = core.StringPtr("testString")
				Expect(workspaceStatusMessageModel.StatusCode).To(Equal(core.StringPtr("testString")))
				Expect(workspaceStatusMessageModel.StatusMsg).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateWorkspaceOptions model
				wID := "testString"
				updateWorkspaceOptionsModel := schematicsService.NewUpdateWorkspaceOptions(wID)
				updateWorkspaceOptionsModel.SetWID("testString")
				updateWorkspaceOptionsModel.SetCatalogRef(catalogRefModel)
				updateWorkspaceOptionsModel.SetDescription("testString")
				updateWorkspaceOptionsModel.SetName("testString")
				updateWorkspaceOptionsModel.SetSharedData(sharedTargetDataModel)
				updateWorkspaceOptionsModel.SetTags([]string{"testString"})
				updateWorkspaceOptionsModel.SetTemplateData([]schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel})
				updateWorkspaceOptionsModel.SetTemplateRepo(templateRepoUpdateRequestModel)
				updateWorkspaceOptionsModel.SetType([]string{"testString"})
				updateWorkspaceOptionsModel.SetWorkspaceStatus(workspaceStatusUpdateRequestModel)
				updateWorkspaceOptionsModel.SetWorkspaceStatusMsg(workspaceStatusMessageModel)
				updateWorkspaceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateWorkspaceOptionsModel).ToNot(BeNil())
				Expect(updateWorkspaceOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkspaceOptionsModel.CatalogRef).To(Equal(catalogRefModel))
				Expect(updateWorkspaceOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkspaceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateWorkspaceOptionsModel.SharedData).To(Equal(sharedTargetDataModel))
				Expect(updateWorkspaceOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(updateWorkspaceOptionsModel.TemplateData).To(Equal([]schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel}))
				Expect(updateWorkspaceOptionsModel.TemplateRepo).To(Equal(templateRepoUpdateRequestModel))
				Expect(updateWorkspaceOptionsModel.Type).To(Equal([]string{"testString"}))
				Expect(updateWorkspaceOptionsModel.WorkspaceStatus).To(Equal(workspaceStatusUpdateRequestModel))
				Expect(updateWorkspaceOptionsModel.WorkspaceStatusMsg).To(Equal(workspaceStatusMessageModel))
				Expect(updateWorkspaceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadTemplateTarOptions successfully`, func() {
				// Construct an instance of the UploadTemplateTarOptions model
				wID := "testString"
				tID := "testString"
				uploadTemplateTarOptionsModel := schematicsService.NewUploadTemplateTarOptions(wID, tID)
				uploadTemplateTarOptionsModel.SetWID("testString")
				uploadTemplateTarOptionsModel.SetTID("testString")
				uploadTemplateTarOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				uploadTemplateTarOptionsModel.SetFileContentType("testString")
				uploadTemplateTarOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadTemplateTarOptionsModel).ToNot(BeNil())
				Expect(uploadTemplateTarOptionsModel.WID).To(Equal(core.StringPtr("testString")))
				Expect(uploadTemplateTarOptionsModel.TID).To(Equal(core.StringPtr("testString")))
				Expect(uploadTemplateTarOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(uploadTemplateTarOptionsModel.FileContentType).To(Equal(core.StringPtr("testString")))
				Expect(uploadTemplateTarOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
