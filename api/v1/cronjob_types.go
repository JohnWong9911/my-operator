/*
Copyright 2024.

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
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CronJob. Edit cronjob_types.go to remove/update
	Foo string `json:"foo,omitempty"`
	// Schedule is the cron schedule
	Schedule string `json:"schedule"`

	// JobTemplate is the template for the job
	JobTemplate batchv1.JobSpec `json:"jobTemplate"`
}

// JobTemplateSpec describes the data a Job should have when created from a template
type JobTemplateSpec struct {
	Spec JobSpec `json:"spec"`
}

// JobSpec describes how the job execution will look like
type JobSpec struct {
	Template PodTemplateSpec `json:"template"`
}

// PodTemplateSpec describes the data a pod should have when created from a template
type PodTemplateSpec struct {
	Spec PodSpec `json:"spec"`
}

// PodSpec is a description of a pod
type PodSpec struct {
	Containers    []Container `json:"containers"`
	RestartPolicy string      `json:"restartPolicy"`
}

// Container represents a single container that is expected to be run on the host
type Container struct {
	Name  string   `json:"name"`
	Image string   `json:"image"`
	Args  []string `json:"args"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
