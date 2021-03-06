/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CategoriesObservation struct {
}

type CategoriesParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ClusterReferencesObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ClusterReferencesParameters struct {

	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type CurrentClusterReferenceListObservation struct {
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type CurrentClusterReferenceListParameters struct {
}

type ImageObservation struct {
	APIVersion *string `json:"apiVersion,omitempty" tf:"api_version,omitempty"`

	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	ClusterUUID *string `json:"clusterUuid,omitempty" tf:"cluster_uuid,omitempty"`

	CurrentClusterReferenceList []CurrentClusterReferenceListObservation `json:"currentClusterReferenceList,omitempty" tf:"current_cluster_reference_list,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	RetrievalURIList []*string `json:"retrievalUriList,omitempty" tf:"retrieval_uri_list,omitempty"`

	SizeBytes *float64 `json:"sizeBytes,omitempty" tf:"size_bytes,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ImageParameters struct {

	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZoneReference map[string]*string `json:"availabilityZoneReference,omitempty" tf:"availability_zone_reference,omitempty"`

	// +kubebuilder:validation:Optional
	Categories []CategoriesParameters `json:"categories,omitempty" tf:"categories,omitempty"`

	// +kubebuilder:validation:Optional
	Checksum map[string]*string `json:"checksum,omitempty" tf:"checksum,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterReferences []ClusterReferencesParameters `json:"clusterReferences,omitempty" tf:"cluster_references,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OwnerReference map[string]*string `json:"ownerReference,omitempty" tf:"owner_reference,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectReference map[string]*string `json:"projectReference,omitempty" tf:"project_reference,omitempty"`

	// +kubebuilder:validation:Optional
	SourcePath *string `json:"sourcePath,omitempty" tf:"source_path,omitempty"`

	// +kubebuilder:validation:Optional
	SourceURI *string `json:"sourceUri,omitempty" tf:"source_uri,omitempty"`

	// +kubebuilder:validation:Optional
	Version map[string]*string `json:"version,omitempty" tf:"version,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nutanixjet}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec"`
	Status            ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
