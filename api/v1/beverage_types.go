/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BeverageSpec defines the desired state of Beverage
type BeverageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Brew is the beer type.
	// +kubebuilder:validation:Enum=Lager;IPA;IIPA;XPA;PaleAle;Stout
	Brew string `json:"brew,omitempty"`
}

// BeverageStatus defines the observed state of Beverage
type BeverageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Tasty bool `json:"tasty"`
}

// +kubebuilder:object:root=true

// Beverage is the Schema for the beverages API
type Beverage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BeverageSpec   `json:"spec,omitempty"`
	Status BeverageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BeverageList contains a list of Beverage
type BeverageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Beverage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Beverage{}, &BeverageList{})
}
