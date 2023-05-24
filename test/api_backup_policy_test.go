/*
Taikun - WebApi

Testing BackupPolicyApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/kibernetik542/baburgo"
)

func Test_openapi_BackupPolicyApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BackupPolicyApiService BackupByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupByName(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupClearProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupClearProject(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDeleteBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupDeleteBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDeleteBackupLocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupDeleteBackupLocation(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDeleteRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupDeleteRestore(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDeleteSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupDeleteSchedule(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDescribeBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupDescribeBackup(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDescribeRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupDescribeRestore(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDescribeSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var name string

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupDescribeSchedule(context.Background(), projectId, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupDisableBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupDisableBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupEnableBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupEnableBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupImportBackupStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupImportBackupStorage(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupListAllBackupStorages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupListAllBackupStorages(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupListAllBackups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupListAllBackups(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupListAllDeleteBackupRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupListAllDeleteBackupRequests(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupListAllRestores", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupListAllRestores(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupListAllSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.BackupPolicyApi.BackupListAllSchedules(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BackupPolicyApiService BackupRestoreBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BackupPolicyApi.BackupRestoreBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
