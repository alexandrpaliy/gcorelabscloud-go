package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/alexandrpaliy/gcorelabscloud-go/gcore/faas/v1/faas"
	"github.com/alexandrpaliy/gcorelabscloud-go/pagination"
	fake "github.com/alexandrpaliy/gcorelabscloud-go/testhelper/client"
	"github.com/stretchr/testify/require"

	log "github.com/sirupsen/logrus"

	th "github.com/alexandrpaliy/gcorelabscloud-go/testhelper"
)

func prepareNamespaceTestURLParams(projectID int, regionID int, nsName string) string { // nolint
	return fmt.Sprintf("/v1/faas/namespaces/%d/%d/%s", projectID, regionID, nsName)
}

func prepareNamespaceTestURL() string { // nolint
	return fmt.Sprintf("/v1/faas/namespaces/%d/%d", fake.ProjectID, fake.RegionID)
}

func prepareGetNamespaceTestURL(nsName string) string {
	return prepareNamespaceTestURLParams(fake.ProjectID, fake.RegionID, nsName)
}

func prepareDeleteNamespaceTestURL(nsName string) string {
	return prepareNamespaceTestURLParams(fake.ProjectID, fake.RegionID, nsName)
}

func prepareUpdateNamespaceTestURL(nsName string) string {
	return prepareNamespaceTestURLParams(fake.ProjectID, fake.RegionID, nsName)
}

func prepareFunctionTestURLParams(projectID int, regionID int, nsName, fName string) string { // nolint
	return fmt.Sprintf("/v1/faas/namespaces/%d/%d/%s/functions/%s", projectID, regionID, nsName, fName)
}

func prepareFunctionTestURL(nsName string) string { // nolint
	return fmt.Sprintf("/v1/faas/namespaces/%d/%d/%s/functions", fake.ProjectID, fake.RegionID, nsName)
}

func prepareGetFunctionTestURL(nsName, fName string) string {
	return prepareFunctionTestURLParams(fake.ProjectID, fake.RegionID, nsName, fName)
}

func prepareDeleteFunctionTestURL(nsName, fName string) string {
	return prepareFunctionTestURLParams(fake.ProjectID, fake.RegionID, nsName, fName)
}

func prepareUpdateFunctionTestURL(nsName, fName string) string {
	return prepareFunctionTestURLParams(fake.ProjectID, fake.RegionID, nsName, fName)
}

func prepareKeyTestURL() string {
	return fmt.Sprintf("/v1/faas/keys/%d/%d", fake.ProjectID, fake.RegionID)
}

func prepareGetKeyTestURL(kName string) string {
	return fmt.Sprintf("/v1/faas/keys/%d/%d/%s", fake.ProjectID, fake.RegionID, kName)
}

func prepareDeleteKeyTestURL(kName string) string {
	return fmt.Sprintf("/v1/faas/keys/%d/%d/%s", fake.ProjectID, fake.RegionID, kName)
}

func prepareUpdateKeyTestURL(kName string) string {
	return fmt.Sprintf("/v1/faas/keys/%d/%d/%s", fake.ProjectID, fake.RegionID, kName)
}

func TestGetNamespace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareGetNamespaceTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, getNamespaceResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	ns, err := faas.GetNamespace(client, nsName).Extract()
	require.NoError(t, err)
	require.Equal(t, expectedNs, *ns)
}

func TestListNamespaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareNamespaceTestURL(), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, listNamespaceResponse)
		if err != nil {
			log.Error(err)
		}
	})

	var count int
	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	err := faas.ListNamespace(client, nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := faas.ExtractNamespaces(page)
		require.NoError(t, err)
		ct := actual[0]
		require.Equal(t, expectedNs, ct)
		require.Equal(t, expectedNsSlice, actual)
		return true, nil
	})
	th.AssertNoErr(t, err)

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListAllNamespaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareNamespaceTestURL(), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, listNamespaceResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	actual, err := faas.ListNamespaceALL(client, nil)
	require.NoError(t, err)

	ct := actual[0]
	require.Equal(t, expectedNs, ct)
	require.Equal(t, expectedNsSlice, actual)

}

func TestCreateNamespace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareNamespaceTestURL(), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createNamespaceRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, taskResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	opts := faas.CreateNamespaceOpts{
		Name:        "string",
		Description: "long string",
		Envs:        map[string]string{"ENV_VAR": "value 1"},
	}
	task, err := faas.CreateNamespace(client, opts).Extract()
	require.NoError(t, err)
	require.Equal(t, tasks1, *task)
}

func TestDeleteNamespace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareDeleteNamespaceTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, taskResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	task, err := faas.DeleteNamespace(client, nsName).Extract()
	require.NoError(t, err)
	require.Equal(t, tasks1, *task)
}

func TestUpdateNamespace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareUpdateNamespaceTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateNamespaceRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, taskResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	opts := faas.UpdateNamespaceOpts{
		Description: "long string",
		Envs:        map[string]string{"ENV_VAR": "value 1"},
	}
	task, err := faas.UpdateNamespace(client, nsName, opts).Extract()
	require.NoError(t, err)
	require.Equal(t, tasks1, *task)
}

func TestGetFunction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareGetFunctionTestURL(nsName, fName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, getFunctionResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	f, err := faas.GetFunction(client, nsName, fName).Extract()
	require.NoError(t, err)
	require.Equal(t, expectedF, *f)
}

func TestListFunctions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareFunctionTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, listFunctionResponse)
		if err != nil {
			log.Error(err)
		}
	})

	var count int
	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	err := faas.ListFunctions(client, nsName, nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := faas.ExtractFunctions(page)
		require.NoError(t, err)
		ct := actual[0]
		require.Equal(t, expectedF, ct)
		require.Equal(t, expectedFSlice, actual)
		return true, nil
	})
	th.AssertNoErr(t, err)

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListAllFunctions(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareFunctionTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, listFunctionResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	actual, err := faas.ListFunctionsALL(client, nsName, nil)
	require.NoError(t, err)
	ct := actual[0]
	require.Equal(t, expectedF, ct)
	require.Equal(t, expectedFSlice, actual)

}

func TestCreateFunction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareFunctionTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createFunctionRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, taskResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	opts := faas.CreateFunctionOpts{
		Name:        fName,
		Description: "Function description",
		Envs: map[string]string{
			"ENV_VAR": "value 1",
		},
		Runtime: "python3.7.12",
		Timeout: 5,
		Flavor:  "64mCPU-64MB",
		Autoscaling: faas.FunctionAutoscaling{
			MinInstances: pointer.To(1),
			MaxInstances: pointer.To(2),
		},
		EnableApiKey: pointer.To(true),
		Keys:         []string{"key-one"},
		CodeText:     "def main(): print('It works!')",
		MainMethod:   "main",
	}
	task, err := faas.CreateFunction(client, nsName, opts).Extract()
	require.NoError(t, err)
	require.Equal(t, tasks1, *task)
}

func TestDeleteFunction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareDeleteFunctionTestURL(nsName, fName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, taskResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	task, err := faas.DeleteFunction(client, nsName, fName).Extract()
	require.NoError(t, err)
	require.Equal(t, tasks1, *task)
}

func TestUpdateFunction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareUpdateFunctionTestURL(nsName, fName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateFunctionRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, taskResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/namespaces", "v1")
	opts := faas.UpdateFunctionOpts{
		Description: "string",
		Envs:        map[string]string{"property1": "string"},
		Timeout:     180,
		Flavor:      "string",
		Autoscaling: &faas.FunctionAutoscaling{
			MinInstances: pointer.To(1),
			MaxInstances: pointer.To(2),
		},
		CodeText:   "string",
		MainMethod: "string",
	}
	task, err := faas.UpdateFunction(client, nsName, fName, opts).Extract()
	require.NoError(t, err)
	require.Equal(t, tasks1, *task)
}

func TestGetKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareGetKeyTestURL(kName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, getKeyResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/keys", "v1")
	key, err := faas.GetKey(client, kName).Extract()
	require.NoError(t, err)
	require.Equal(t, expectedKey, *key)
}

func TestListKeys(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareKeyTestURL(), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, listKeysResponse)
		if err != nil {
			log.Error(err)
		}
	})

	var count int
	client := fake.ServiceTokenClient("faas/keys", "v1")
	err := faas.ListKeys(client, nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := faas.ExtractKeys(page)
		require.NoError(t, err)
		ct := actual[0]
		require.Equal(t, expectedKey, ct)
		require.Equal(t, expectedKeysSlice, actual)
		return true, nil
	})
	th.AssertNoErr(t, err)

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestListAllKeys(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareKeyTestURL(), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, listKeysResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/keys", "v1")
	actual, err := faas.ListKeysAll(client, nil)
	require.NoError(t, err)

	ct := actual[0]
	require.Equal(t, expectedKey, ct)
	require.Equal(t, expectedKeysSlice, actual)

}

func TestCreateKeys(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareKeyTestURL(), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createKeyRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, createKeyResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/keys", "v1")
	opts := faas.CreateKeyOpts{
		Name:        "test-key",
		Description: "description",
		Functions: []faas.KeysFunction{
			{
				Name:      "function",
				Namespace: "namespace",
			},
		},
	}
	key, err := faas.CreateKey(client, &opts)
	require.NoError(t, err)
	require.Equal(t, expectedKey, key)
}

func TestDeleteKeys(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareDeleteKeyTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Accept", "application/json")

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	client := fake.ServiceTokenClient("faas/keys", "v1")
	err := faas.DeleteKey(client, nsName)
	require.NoError(t, err)
}

func TestUpdateKeys(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(prepareUpdateKeyTestURL(nsName), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "Authorization", fmt.Sprintf("Bearer %s", fake.AccessToken))
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateKeyRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := fmt.Fprint(w, updateKeyResponse)
		if err != nil {
			log.Error(err)
		}
	})

	client := fake.ServiceTokenClient("faas/keys", "v1")
	opts := faas.UpdateKeyOpts{
		Description: "long string",
		Functions: []faas.KeysFunction{
			{
				Name:      "function1",
				Namespace: "namespace1",
			},
			{
				Name:      "function2",
				Namespace: "namespace1",
			},
		},
	}
	key, err := faas.UpdateKey(client, nsName, opts)
	require.NoError(t, err)
	require.Equal(t, expectedUpdatedKey, key)
}
