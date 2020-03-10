// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var SecretManagerSecretIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"secret_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type SecretManagerSecretIamUpdater struct {
	project  string
	secretId string
	d        *schema.ResourceData
	Config   *Config
}

func SecretManagerSecretIamUpdaterProducer(d *schema.ResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}
	values["project"] = project
	if v, ok := d.GetOk("secret_id"); ok {
		values["secret_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/secrets/(?P<secret_id>[^/]+)", "(?P<project>[^/]+)/(?P<secret_id>[^/]+)", "(?P<secret_id>[^/]+)"}, d, config, d.Get("secret_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SecretManagerSecretIamUpdater{
		project:  values["project"],
		secretId: values["secret_id"],
		d:        d,
		Config:   config,
	}

	d.Set("project", u.project)
	d.Set("secret_id", u.GetResourceId())

	return u, nil
}

func SecretManagerSecretIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/secrets/(?P<secret_id>[^/]+)", "(?P<project>[^/]+)/(?P<secret_id>[^/]+)", "(?P<secret_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &SecretManagerSecretIamUpdater{
		project:  values["project"],
		secretId: values["secret_id"],
		d:        d,
		Config:   config,
	}
	d.Set("secret_id", u.GetResourceId())
	d.SetId(u.GetResourceId())
	return nil
}

func (u *SecretManagerSecretIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifySecretUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	policy, err := sendRequest(u.Config, "GET", project, url, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *SecretManagerSecretIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifySecretUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *SecretManagerSecretIamUpdater) qualifySecretUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{SecretManagerBasePath}}%s:%s", fmt.Sprintf("projects/%s/secrets/%s", u.project, u.secretId), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *SecretManagerSecretIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/secrets/%s", u.project, u.secretId)
}

func (u *SecretManagerSecretIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-secretmanager-secret-%s", u.GetResourceId())
}

func (u *SecretManagerSecretIamUpdater) DescribeResource() string {
	return fmt.Sprintf("secretmanager secret %q", u.GetResourceId())
}
