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
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/googleapi"
)

func resourceAccessContextManagerServicePerimeterResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerServicePerimeterResourceCreate,
		Read:   resourceAccessContextManagerServicePerimeterResourceRead,
		Delete: resourceAccessContextManagerServicePerimeterResourceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerServicePerimeterResourceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"perimeter_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the Service Perimeter to add this resource to.`,
			},
			"resource": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A GCP resource that is inside of the service perimeter.
Currently only projects are allowed.
Format: projects/{project_number}`,
			},
		},
	}
}

func resourceAccessContextManagerServicePerimeterResourceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	resourceProp, err := expandAccessContextManagerServicePerimeterResourceResource(d.Get("resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource"); !isEmptyValue(reflect.ValueOf(resourceProp)) && (ok || !reflect.DeepEqual(v, resourceProp)) {
		obj["resource"] = resourceProp
	}

	lockName, err := replaceVars(d, config, "{{perimeter_name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServicePerimeterResource: %#v", obj)

	obj, err = resourceAccessContextManagerServicePerimeterResourcePatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	url, err = addQueryParams(url, map[string]string{"updateMask": "status.resources"})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServicePerimeterResource: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{perimeter_name}}/{{resource}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = accessContextManagerOperationWaitTime(
		config, res, "Creating ServicePerimeterResource",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServicePerimeterResource: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServicePerimeterResource %q: %#v", d.Id(), res)

	return resourceAccessContextManagerServicePerimeterResourceRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterResourceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerServicePerimeterResource %q", d.Id()))
	}

	res, err = flattenNestedAccessContextManagerServicePerimeterResource(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing AccessContextManagerServicePerimeterResource because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("resource", flattenAccessContextManagerServicePerimeterResourceResource(res["resource"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeterResource: %s", err)
	}

	return nil
}

func resourceAccessContextManagerServicePerimeterResourceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lockName, err := replaceVars(d, config, "{{perimeter_name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceAccessContextManagerServicePerimeterResourcePatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return handleNotFoundError(err, d, "ServicePerimeterResource")
	}
	url, err = addQueryParams(url, map[string]string{"updateMask": "status.resources"})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Deleting ServicePerimeterResource %q", d.Id())

	res, err := sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ServicePerimeterResource")
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Deleting ServicePerimeterResource",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServicePerimeterResource %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerServicePerimeterResourceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	parts, err := getImportIdQualifiers([]string{"accessPolicies/(?P<accessPolicy>[^/]+)/servicePerimeters/(?P<perimeter>[^/]+)/(?P<resource>.+)"}, d, config, d.Id())
	if err != nil {
		return nil, err
	}

	d.Set("perimeter_name", fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", parts["accessPolicy"], parts["perimeter"]))
	d.Set("resource", parts["resource"])
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerServicePerimeterResourceResource(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAccessContextManagerServicePerimeterResourceResource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func flattenNestedAccessContextManagerServicePerimeterResource(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["status"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})

	v, ok = res["resources"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value status.resources. Actual value: %v", v)
	}

	_, item, err := resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedResource, err := expandAccessContextManagerServicePerimeterResourceResource(d.Get("resource"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		// List response only contains the ID - construct a response object.
		item := map[string]interface{}{
			"resource": itemRaw,
		}

		itemResource := flattenAccessContextManagerServicePerimeterResourceResource(item["resource"], d, meta.(*Config))
		if !reflect.DeepEqual(itemResource, expectedResource) {
			log.Printf("[DEBUG] Skipping item with resource= %#v, looking for %#v)", itemResource, expectedResource)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceAccessContextManagerServicePerimeterResourcePatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerServicePerimeterResourceListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create ServicePerimeterResource, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"resources": append(currItems, obj["resource"]),
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// PatchUpdateEncoder handles creating request data to PATCH parent resource
// with list including updated object.
func resourceAccessContextManagerServicePerimeterResourcePatchUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	items, err := resourceAccessContextManagerServicePerimeterResourceListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d, meta, items)
	if err != nil {
		return nil, err
	}

	// Return error if item to update does not exist.
	if item == nil {
		return nil, fmt.Errorf("Unable to update ServicePerimeterResource %q - not found in list", d.Id())
	}

	// Merge new object into old.
	for k, v := range obj {
		item[k] = v
	}
	items[idx] = item

	// Return list with new item added
	res := map[string]interface{}{
		"resources": items,
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceAccessContextManagerServicePerimeterResourcePatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerServicePerimeterResourceListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, &googleapi.Error{
			Code:    404,
			Message: "ServicePerimeterResource not found in list",
		}
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"resources": updatedItems,
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceAccessContextManagerServicePerimeterResourceListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*Config)
	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return nil, err
	}
	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool
	if v, ok = res["status"]; ok && v != nil {
		res = v.(map[string]interface{})
	} else {
		return nil, nil
	}

	v, ok = res["resources"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "resources"`)
		}
		return ls, nil
	}
	return nil, nil
}