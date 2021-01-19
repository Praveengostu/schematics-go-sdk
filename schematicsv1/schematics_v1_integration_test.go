// +build integration

/**
 * (C) Copyright IBM Corp. 2020.
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
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the schematicsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

func createSampleWorkspaceWithoutRepoURL() *schematicsv1.WorkspaceResponse {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, err := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	templateSourceDataRequestModel := &schematicsv1.TemplateSourceDataRequest{
		Folder: core.StringPtr("."),
		Type:   core.StringPtr("terraform_v0.11.14"),
		Variablestore: []schematicsv1.WorkspaceVariableRequest{
			{
				Name:  core.StringPtr("variable_name1"),
				Value: core.StringPtr("variable_value1"),
			},
			{
				Name:  core.StringPtr("variable_name2"),
				Value: core.StringPtr("variable_value2"),
			},
		},
	}

	createWorkspaceOptions := &schematicsv1.CreateWorkspaceOptions{
		Description:  core.StringPtr(""),
		Name:         core.StringPtr("myworkspace"),
		TemplateData: []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel},
		Type:         []string{"terraform_v0.11.14"},
		Tags:         []string{},
	}

	workspaceResponse, detailedResponse, err := schematicsService.CreateWorkspace(createWorkspaceOptions)

	if err != nil {
		fmt.Printf("Failed to create the workspaces : %v and the response is %s", err, detailedResponse)
		panic(err)
	}

	waitForWorkspaceStatus(workspaceResponse.ID, "DRAFT")

	return workspaceResponse

}

func getWorkspaceByID(wid *string) *schematicsv1.WorkspaceResponse {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, err := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	workspaceResponse, detailedResponse, err := schematicsService.GetWorkspace(&schematicsv1.GetWorkspaceOptions{
		WID: wid,
	})

	if err != nil {
		fmt.Printf("Failed to get the workspace : %v and the response is %s", err, detailedResponse)
	}

	return workspaceResponse
}

func getWorkspaceActivityByID(wid *string, activityid *string) *schematicsv1.WorkspaceActivity {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, err := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	activityResponse, detailedResponse, err := schematicsService.GetWorkspaceActivity(&schematicsv1.GetWorkspaceActivityOptions{
		WID:        wid,
		ActivityID: activityid,
	})

	if err != nil {
		fmt.Printf("Failed to get the workspace activity : %v and the response is %s", err, detailedResponse)
	}

	return activityResponse
}

func deleteWorkspaceByID(wid *string) {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, err := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	_, detailedResponse, err := schematicsService.DeleteWorkspace(&schematicsv1.DeleteWorkspaceOptions{
		WID:          wid,
		RefreshToken: core.StringPtr("Refresh Token"),
	})

	if err != nil {
		fmt.Printf("Failed to delete the workspace : %v and the response is %s", err, detailedResponse)
	}

}

func waitForWorkspaceStatus(wid *string, status string) {
	var workspaceStatus string
	for strings.Compare(workspaceStatus, status) != 0 {
		workspaceStatus = *getWorkspaceByID(wid).Status
		//fmt.Println(workspaceStatus)
		time.Sleep(2 * time.Second)
	}
}

func waitForWorkspaceActivityStatus(wid *string, activityid *string, status string) {
	var activitystatus string
	for strings.Compare(activitystatus, status) != 0 {
		activitystatus = *getWorkspaceActivityByID(wid, activityid).Status
		//fmt.Println(activitystatus)
		time.Sleep(2 * time.Second)
	}
}

func uploadTarFile(ws *schematicsv1.WorkspaceResponse) *schematicsv1.TemplateRepoTarUploadResponse {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, _ := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	fileDir, _ := os.Getwd()
	fileName := "tf_cloudless_sleepy_git_archive.tar"
	filePath := path.Join(fileDir, "tarfiles", fileName)
	fileReader, _ := os.Open(filePath)
	fileReaderWrapper := ioutil.NopCloser(fileReader)

	uploadTarOptions := &schematicsv1.UploadTemplateTarOptions{
		WID:             ws.ID,
		TID:             ws.TemplateData[0].ID,
		File:            fileReaderWrapper,
		FileContentType: core.StringPtr("multipart/form-data"),
	}

	uploadResponse, detailedResponse, err := schematicsService.UploadTemplateTar(uploadTarOptions)

	if err != nil {
		fmt.Printf("Failed to delete the workspace : %v and the response is %s", err, detailedResponse)
	}

	waitForWorkspaceStatus(ws.ID, "INACTIVE")

	return uploadResponse

}

func refreshWorkspaceActionByID(wid *string) *string {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, _ := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	refreshWorkspaceCommandOptions := &schematicsv1.RefreshWorkspaceCommandOptions{
		WID:          wid,
		RefreshToken: core.StringPtr("token"),
	}

	refreshResult, detailedResponse, err := schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptions)

	if err != nil {
		fmt.Printf("Failed to destroy action for the workspace : %v and the response is %s", err, detailedResponse)
	}

	waitForWorkspaceActivityStatus(wid, refreshResult.Activityid, "COMPLETED")

	return refreshResult.Activityid

}

func planWorkspaceByID(wid *string) *string {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, _ := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	planWorkspaceCommandOptions := &schematicsv1.PlanWorkspaceCommandOptions{
		WID:          wid,
		RefreshToken: core.StringPtr("refresh_token"),
	}

	planResult, detailedResponse, err := schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptions)

	if err != nil {
		fmt.Printf("Failed to generate plan : %v and the response is %s", err, detailedResponse)
	}

	waitForWorkspaceActivityStatus(wid, planResult.Activityid, "COMPLETED")

	return planResult.Activityid

}

func planWorkspaceByIDWithoutWait(wid *string) *string {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, _ := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	planWorkspaceCommandOptions := &schematicsv1.PlanWorkspaceCommandOptions{
		WID:          wid,
		RefreshToken: core.StringPtr("refresh_token"),
	}

	planResult, detailedResponse, err := schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptions)

	if err != nil {
		fmt.Printf("Failed to generate plan : %v and the response is %s", err, detailedResponse)
	}

	return planResult.Activityid

}

func applyWorkspaceByID(wid *string) *string {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, _ := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	applyWorkspaceCommandOptions := &schematicsv1.ApplyWorkspaceCommandOptions{
		WID:          wid,
		RefreshToken: core.StringPtr("token"),
	}

	applyResult, detailedResponse, err := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptions)

	if err != nil {
		fmt.Printf("Failed to generate plan : %v and the response is %s", err, detailedResponse)
	}

	waitForWorkspaceActivityStatus(wid, applyResult.Activityid, "COMPLETED")

	return applyResult.Activityid

}

func applyWorkspaceByIDWithoutWait(wid *string) *string {

	schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

	schematicsService, _ := schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

	applyWorkspaceCommandOptions := &schematicsv1.ApplyWorkspaceCommandOptions{
		WID:          wid,
		RefreshToken: core.StringPtr("token"),
	}

	applyResult, detailedResponse, err := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptions)

	if err != nil {
		fmt.Printf("Failed to generate plan : %v and the response is %s", err, detailedResponse)
	}

	return applyResult.Activityid

}

var _ = Describe(`SchematicsV1 Integration Tests`, func() {

	const externalConfigFile = "../schematics_v1.env"

	var (
		err               error
		schematicsService *schematicsv1.SchematicsV1
		serviceURL        string
		config            map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(schematicsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			schematicsServiceOptions := &schematicsv1.SchematicsV1Options{}

			schematicsService, err = schematicsv1.NewSchematicsV1UsingExternalConfig(schematicsServiceOptions)

			Expect(err).To(BeNil())
			Expect(schematicsService).ToNot(BeNil())
			Expect(schematicsService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ListSchematicsLocation - List supported schematics locations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSchematicsLocation(listSchematicsLocationOptions *ListSchematicsLocationOptions)`, func() {

			listSchematicsLocationOptions := &schematicsv1.ListSchematicsLocationOptions{}

			schematicsLocations, response, err := schematicsService.ListSchematicsLocation(listSchematicsLocationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsLocations).ToNot(BeNil())

		})
	})

	Describe(`ListResourceGroup - List of resource groups in the Account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListResourceGroup(listResourceGroupOptions *ListResourceGroupOptions)`, func() {

			listResourceGroupOptions := &schematicsv1.ListResourceGroupOptions{}

			resourceGroupResponse, response, err := schematicsService.ListResourceGroup(listResourceGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupResponse).ToNot(BeNil())

		})
	})

	Describe(`GetSchematicsVersion - Get schematics version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchematicsVersion(getSchematicsVersionOptions *GetSchematicsVersionOptions)`, func() {

			getSchematicsVersionOptions := &schematicsv1.GetSchematicsVersionOptions{}

			versionResponse, response, err := schematicsService.GetSchematicsVersion(getSchematicsVersionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(versionResponse).ToNot(BeNil())

		})
	})

	Describe(`ListWorkspaces - List all workspace definitions`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
		})

		It(`ListWorkspaces(listWorkspacesOptions *ListWorkspacesOptions)`, func() {

			listWorkspacesOptions := &schematicsv1.ListWorkspacesOptions{
				Offset: core.Int64Ptr(int64(0)),
				Limit:  core.Int64Ptr(int64(1)),
			}

			workspaceResponseList, response, err := schematicsService.ListWorkspaces(listWorkspacesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponseList).ToNot(BeNil())

		})

		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`CreateWorkspace - Create workspace definition`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateWorkspace(createWorkspaceOptions *CreateWorkspaceOptions)`, func() {

			templateSourceDataRequestModel := &schematicsv1.TemplateSourceDataRequest{
				Folder: core.StringPtr("."),
				Type:   core.StringPtr("terraform_v0.11.14"),
				Variablestore: []schematicsv1.WorkspaceVariableRequest{
					{
						Name:  core.StringPtr("variable_name1"),
						Value: core.StringPtr("variable_value1"),
					},
					{
						Name:  core.StringPtr("variable_name2"),
						Value: core.StringPtr("variable_value2"),
					},
				},
			}

			createWorkspaceOptions := &schematicsv1.CreateWorkspaceOptions{
				Description:  core.StringPtr("Sample Workspace"),
				Name:         core.StringPtr("myworkspace"),
				TemplateData: []schematicsv1.TemplateSourceDataRequest{*templateSourceDataRequestModel},
				TemplateRef:  core.StringPtr("testString"),
				Type:         []string{"terraform_v0.11.14"},
			}

			workspaceResponse, response, err := schematicsService.CreateWorkspace(createWorkspaceOptions)

			ws = workspaceResponse

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(workspaceResponse).ToNot(BeNil())

		})

		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspace - Get workspace definition`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
		})

		It(`GetWorkspace(getWorkspaceOptions *GetWorkspaceOptions)`, func() {

			getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{
				WID: ws.ID,
			}

			workspaceResponse, response, err := schematicsService.GetWorkspace(getWorkspaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())

		})

		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})

	})

	Describe(`ReplaceWorkspace - Replace the workspace definition`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
		})

		It(`ReplaceWorkspace(replaceWorkspaceOptions *ReplaceWorkspaceOptions)`, func() {

			templateRepoRequestUpdateModel := &schematicsv1.TemplateRepoUpdateRequest{
				URL:     core.StringPtr("https://github.com/ptaube/tf_cloudless_sleepy"),
				RepoURL: core.StringPtr("literal string"),
			}

			updateWorkspaceOptions := &schematicsv1.ReplaceWorkspaceOptions{
				WID:          ws.ID,
				Description:  core.StringPtr(""),
				Name:         core.StringPtr("myworkspace"),
				Type:         []string{"terraform_v0.12.20"},
				TemplateRepo: templateRepoRequestUpdateModel,
				CatalogRef: &schematicsv1.CatalogRef{
					ItemID:          core.StringPtr("literal string"),
					ItemName:        core.StringPtr("literal string"),
					ItemURL:         core.StringPtr("literal string"),
					ItemReadmeURL:   core.StringPtr("literal string"),
					ItemIconURL:     core.StringPtr("literal string"),
					OfferingVersion: core.StringPtr("literal string"),
					LaunchURL:       core.StringPtr("literal string"),
				},
			}

			workspaceResponse, detailedResponse, err := schematicsService.ReplaceWorkspace(updateWorkspaceOptions)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())

			waitForWorkspaceStatus(ws.ID, "INACTIVE")

		})

		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`UpdateWorkspace - Update the workspace definition`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
		})

		It(`UpdateWorkspace(updateWorkspaceOptions *UpdateWorkspaceOptions)`, func() {

			updateWorkspaceOptions := &schematicsv1.UpdateWorkspaceOptions{
				WID:         ws.ID,
				Description: core.StringPtr("Updated"),
				Name:        core.StringPtr("myworkspace"),
				Type:        []string{"terraform_v0.12.20"},
				WorkspaceStatus: &schematicsv1.WorkspaceStatusUpdateRequest{
					Frozen:   core.BoolPtr(false),
					FrozenBy: core.StringPtr("<frozen_by>"),
				},
			}

			workspaceResponse, detailedResponse, err := schematicsService.UpdateWorkspace(updateWorkspaceOptions)

			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(workspaceResponse).ToNot(BeNil())

		})

		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`UploadTemplateTar - Upload template tar file for the workspace`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
		})

		It(`UploadTemplateTar(uploadTemplateTarOptions *UploadTemplateTarOptions)`, func() {

			fileDir, _ := os.Getwd()
			fileName := "tf_cloudless_sleepy_git_archive.tar"
			filePath := path.Join(fileDir, "tarfiles", fileName)
			fileReader, _ := os.Open(filePath)
			fileReaderWrapper := ioutil.NopCloser(fileReader) // no-op Close method wrapping the provided Reader

			uploadTarOptions := &schematicsv1.UploadTemplateTarOptions{
				WID:             ws.ID,
				TID:             ws.TemplateData[0].ID,
				File:            fileReaderWrapper,
				FileContentType: core.StringPtr("multipart/form-data"),
			}

			resultUpload, httpResponse, err := schematicsService.UploadTemplateTar(uploadTarOptions)
			Expect(err).To(BeNil())
			Expect(httpResponse.StatusCode).To(Equal(200))
			Expect(resultUpload).ToNot(BeNil())

			waitForWorkspaceStatus(ws.ID, "INACTIVE")

		})

		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceReadme - Get the workspace readme`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`GetWorkspaceReadme(getWorkspaceReadmeOptions *GetWorkspaceReadmeOptions)`, func() {

			getWorkspaceReadmeOptions := &schematicsv1.GetWorkspaceReadmeOptions{
				WID: ws.ID,
			}

			templateReadme, response, err := schematicsService.GetWorkspaceReadme(getWorkspaceReadmeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateReadme).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`ListWorkspaceActivities - List all workspace activities`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			refreshWorkspaceActionByID(ws.ID)
		})
		It(`ListWorkspaceActivities(listWorkspaceActivitiesOptions *ListWorkspaceActivitiesOptions)`, func() {

			listWorkspaceActivitiesOptions := &schematicsv1.ListWorkspaceActivitiesOptions{
				WID: ws.ID,
			}

			workspaceActivities, response, err := schematicsService.ListWorkspaceActivities(listWorkspaceActivitiesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivities).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceActivity - Get workspace activity details`, func() {

		var ws *schematicsv1.WorkspaceResponse
		var activityID *string

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			activityID = refreshWorkspaceActionByID(ws.ID)
		})
		It(`GetWorkspaceActivity(getWorkspaceActivityOptions *GetWorkspaceActivityOptions)`, func() {

			getWorkspaceActivityOptions := &schematicsv1.GetWorkspaceActivityOptions{
				WID:        ws.ID,
				ActivityID: activityID,
			}

			workspaceActivity, response, err := schematicsService.GetWorkspaceActivity(getWorkspaceActivityOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivity).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`RunWorkspaceCommands - Run terraform Commands`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})

		It(`RunWorkspaceCommands(runWorkspaceCommandsOptions *RunWorkspaceCommandsOptions)`, func() {

			applyWorkspaceCommandOptions := &schematicsv1.ApplyWorkspaceCommandOptions{
				WID:          ws.ID,
				RefreshToken: core.StringPtr("refresh_token"),
			}

			workspaceActivityApplyResult, response, err := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityApplyResult).ToNot(BeNil())

			waitForWorkspaceActivityStatus(ws.ID, workspaceActivityApplyResult.Activityid, "COMPLETED")

			terraformCommandModel := &schematicsv1.TerraformCommand{
				Command:       core.StringPtr("state show"),
				CommandParams: core.StringPtr("data.template_file.test"),
				CommandName:   core.StringPtr("State Show"),
			}

			runWorkspaceCommandsOptions := &schematicsv1.RunWorkspaceCommandsOptions{
				WID:           ws.ID,
				RefreshToken:  core.StringPtr("refresh_token"),
				Commands:      []schematicsv1.TerraformCommand{*terraformCommandModel},
				OperationName: core.StringPtr("State_Show"),
			}

			workspaceActivityCommandResult, response, err := schematicsService.RunWorkspaceCommands(runWorkspaceCommandsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityCommandResult).ToNot(BeNil())

			waitForWorkspaceActivityStatus(ws.ID, workspaceActivityCommandResult.Activityid, "COMPLETED")

		})

		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`ApplyWorkspaceCommand - Run schematics workspace 'apply' activity`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`ApplyWorkspaceCommand(applyWorkspaceCommandOptions *ApplyWorkspaceCommandOptions)`, func() {

			applyWorkspaceCommandOptions := &schematicsv1.ApplyWorkspaceCommandOptions{
				WID:          ws.ID,
				RefreshToken: core.StringPtr("refresh_token"),
			}

			workspaceActivityApplyResult, response, err := schematicsService.ApplyWorkspaceCommand(applyWorkspaceCommandOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityApplyResult).ToNot(BeNil())

			waitForWorkspaceActivityStatus(ws.ID, workspaceActivityApplyResult.Activityid, "COMPLETED")

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`DestroyWorkspaceCommand - Run workspace 'destroy' activity`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
		})

		It(`DestroyWorkspaceCommand(destroyWorkspaceCommandOptions *DestroyWorkspaceCommandOptions)`, func() {

			destroyWorkspaceCommandOptions := &schematicsv1.DestroyWorkspaceCommandOptions{
				WID:          ws.ID,
				RefreshToken: core.StringPtr("refresh_token"),
			}

			workspaceActivityDestroyResult, response, err := schematicsService.DestroyWorkspaceCommand(destroyWorkspaceCommandOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityDestroyResult).ToNot(BeNil())

		})
	})

	Describe(`PlanWorkspaceCommand - Run workspace 'plan' activity,`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`PlanWorkspaceCommand(planWorkspaceCommandOptions *PlanWorkspaceCommandOptions)`, func() {

			planWorkspaceCommandOptions := &schematicsv1.PlanWorkspaceCommandOptions{
				WID:          ws.ID,
				RefreshToken: core.StringPtr("refresh_token"),
			}

			workspaceActivityPlanResult, response, err := schematicsService.PlanWorkspaceCommand(planWorkspaceCommandOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityPlanResult).ToNot(BeNil())

			waitForWorkspaceActivityStatus(ws.ID, workspaceActivityPlanResult.Activityid, "COMPLETED")

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`RefreshWorkspaceCommand - Run workspace 'refresh' activity`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`RefreshWorkspaceCommand(refreshWorkspaceCommandOptions *RefreshWorkspaceCommandOptions)`, func() {

			refreshWorkspaceCommandOptions := &schematicsv1.RefreshWorkspaceCommandOptions{
				WID:          ws.ID,
				RefreshToken: core.StringPtr("refresh_token"),
			}

			workspaceActivityRefreshResult, response, err := schematicsService.RefreshWorkspaceCommand(refreshWorkspaceCommandOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(workspaceActivityRefreshResult).ToNot(BeNil())

			waitForWorkspaceActivityStatus(ws.ID, workspaceActivityRefreshResult.Activityid, "COMPLETED")

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceInputs - Get the input values of the workspace`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`GetWorkspaceInputs(getWorkspaceInputsOptions *GetWorkspaceInputsOptions)`, func() {

			getWorkspaceInputsOptions := &schematicsv1.GetWorkspaceInputsOptions{
				WID: ws.ID,
				TID: ws.TemplateData[0].ID,
			}

			templateValues, response, err := schematicsService.GetWorkspaceInputs(getWorkspaceInputsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateValues).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`ReplaceWorkspaceInputs - Replace the input values for the workspace`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`ReplaceWorkspaceInputs(replaceWorkspaceInputsOptions *ReplaceWorkspaceInputsOptions)`, func() {

			workspaceVariableRequestModel := &schematicsv1.WorkspaceVariableRequest{
				Description: core.StringPtr("Sample Variable"),
				Name:        core.StringPtr("sample_var"),
				Secure:      core.BoolPtr(true),
				Type:        core.StringPtr("string"),
				UseDefault:  core.BoolPtr(true),
				Value:       core.StringPtr("var_value"),
			}

			envVariables := []interface{}{
				map[string]interface{}{
					"KEY1": "VALUE1",
					"KEY2": "VALUE2",
				}}

			replaceWorkspaceInputsOptions := &schematicsv1.ReplaceWorkspaceInputsOptions{
				WID:           ws.ID,
				TID:           ws.TemplateData[0].ID,
				EnvValues:     envVariables,
				Values:        core.StringPtr("testString"),
				Variablestore: []schematicsv1.WorkspaceVariableRequest{*workspaceVariableRequestModel},
			}

			userValues, response, err := schematicsService.ReplaceWorkspaceInputs(replaceWorkspaceInputsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userValues).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetAllWorkspaceInputs - Get all the input values of the workspace`, func() {

		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`GetAllWorkspaceInputs(getAllWorkspaceInputsOptions *GetAllWorkspaceInputsOptions)`, func() {

			getAllWorkspaceInputsOptions := &schematicsv1.GetAllWorkspaceInputsOptions{
				WID: ws.ID,
			}

			workspaceTemplateValuesResponse, response, err := schematicsService.GetAllWorkspaceInputs(getAllWorkspaceInputsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceTemplateValuesResponse).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceInputMetadata - Get the input metadata of the workspace`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
		})
		It(`GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptions *GetWorkspaceInputMetadataOptions)`, func() {

			getWorkspaceInputMetadataOptions := &schematicsv1.GetWorkspaceInputMetadataOptions{
				WID: ws.ID,
				TID: ws.TemplateData[0].ID,
			}

			result, response, err := schematicsService.GetWorkspaceInputMetadata(getWorkspaceInputMetadataOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceOutputs - Get all the output values of the workspace`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			applyWorkspaceByID(ws.ID)
		})
		It(`GetWorkspaceOutputs(getWorkspaceOutputsOptions *GetWorkspaceOutputsOptions)`, func() {

			getWorkspaceOutputsOptions := &schematicsv1.GetWorkspaceOutputsOptions{
				WID: ws.ID,
			}

			outputValuesItem, response, err := schematicsService.GetWorkspaceOutputs(getWorkspaceOutputsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outputValuesItem).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceResources - Get all the resources created by the workspace`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			applyWorkspaceByID(ws.ID)
		})
		It(`GetWorkspaceResources(getWorkspaceResourcesOptions *GetWorkspaceResourcesOptions)`, func() {

			getWorkspaceResourcesOptions := &schematicsv1.GetWorkspaceResourcesOptions{
				WID: ws.ID,
			}

			templateResources, response, err := schematicsService.GetWorkspaceResources(getWorkspaceResourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResources).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceState - Get the workspace state`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			applyWorkspaceByID(ws.ID)
		})
		It(`GetWorkspaceState(getWorkspaceStateOptions *GetWorkspaceStateOptions)`, func() {

			getWorkspaceStateOptions := &schematicsv1.GetWorkspaceStateOptions{
				WID: ws.ID,
			}

			stateStoreResponseList, response, err := schematicsService.GetWorkspaceState(getWorkspaceStateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stateStoreResponseList).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceTemplateState - Get the template state`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			applyWorkspaceByID(ws.ID)
		})
		It(`GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions *GetWorkspaceTemplateStateOptions)`, func() {

			getWorkspaceTemplateStateOptions := &schematicsv1.GetWorkspaceTemplateStateOptions{
				WID: ws.ID,
				TID: ws.TemplateData[0].ID,
			}

			templateStateStore, response, err := schematicsService.GetWorkspaceTemplateState(getWorkspaceTemplateStateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateStateStore).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceActivityLogs - Get the workspace activity log urls`, func() {
		var ws *schematicsv1.WorkspaceResponse
		var activityID *string

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			activityID = applyWorkspaceByID(ws.ID)
		})
		It(`GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptions *GetWorkspaceActivityLogsOptions)`, func() {

			getWorkspaceActivityLogsOptions := &schematicsv1.GetWorkspaceActivityLogsOptions{
				WID:        ws.ID,
				ActivityID: activityID,
			}

			workspaceActivityLogs, response, err := schematicsService.GetWorkspaceActivityLogs(getWorkspaceActivityLogsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceActivityLogs).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetWorkspaceLogUrls - Get all workspace log urls`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			applyWorkspaceByID(ws.ID)
		})

		It(`GetWorkspaceLogUrls(getWorkspaceLogUrlsOptions *GetWorkspaceLogUrlsOptions)`, func() {

			getWorkspaceLogUrlsOptions := &schematicsv1.GetWorkspaceLogUrlsOptions{
				WID: ws.ID,
			}

			logStoreResponseList, response, err := schematicsService.GetWorkspaceLogUrls(getWorkspaceLogUrlsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logStoreResponseList).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetTemplateLogs - Get all template logs`, func() {
		var ws *schematicsv1.WorkspaceResponse

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			applyWorkspaceByID(ws.ID)
		})
		It(`GetTemplateLogs(getTemplateLogsOptions *GetTemplateLogsOptions)`, func() {

			getTemplateLogsOptions := &schematicsv1.GetTemplateLogsOptions{
				WID:               ws.ID,
				TID:               ws.TemplateData[0].ID,
				LogTfCmd:          core.BoolPtr(true),
				LogTfPrefix:       core.BoolPtr(true),
				LogTfNullResource: core.BoolPtr(true),
				LogTfAnsible:      core.BoolPtr(true),
			}

			result, response, err := schematicsService.GetTemplateLogs(getTemplateLogsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})

	Describe(`GetTemplateActivityLog - Get the template activity logs`, func() {
		var ws *schematicsv1.WorkspaceResponse
		var activityID *string

		BeforeEach(func() {
			shouldSkipTest()
			ws = createSampleWorkspaceWithoutRepoURL()
			uploadTarFile(ws)
			activityID = applyWorkspaceByID(ws.ID)
		})
		It(`GetTemplateActivityLog(getTemplateActivityLogOptions *GetTemplateActivityLogOptions)`, func() {

			getTemplateActivityLogOptions := &schematicsv1.GetTemplateActivityLogOptions{
				WID:               ws.ID,
				TID:               ws.TemplateData[0].ID,
				ActivityID:        activityID,
				LogTfCmd:          core.BoolPtr(true),
				LogTfPrefix:       core.BoolPtr(true),
				LogTfNullResource: core.BoolPtr(true),
				LogTfAnsible:      core.BoolPtr(true),
			}

			result, response, err := schematicsService.GetTemplateActivityLog(getTemplateActivityLogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		AfterEach(func() {
			deleteWorkspaceByID(ws.ID)
		})
	})
})

//
// Utility functions are declared in the unit test file
//
