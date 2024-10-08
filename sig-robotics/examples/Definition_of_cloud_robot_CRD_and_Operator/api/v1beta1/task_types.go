/*
Copyright 2022 The KubeEdge Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TaskSpec defines the desired state of Task
type TaskSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	OrderID              *uint                  `json:"orderID,omitempty"`
	TaskID               *uint                  `json:"taskID,omitempty"`
	RequiredSensors      []Sensor               `json:"requiredSensors,omitempty"`	// List of sensors required to complete this task
	Allocated            int                   `json:"allocated,omitempty,default=-1"` // -1 if not assigned, otherwise the assigned robotID
	TaskDestinationID    uint                   `json:"taskDestinationID,omitempty"`
	PointStateSequence   []PointStateSequence   `json:"pointStateSequence,omitempty"`	// The point sequence contained in the current task
	SegmentStateSequence []SegmentStateSequence `json:"segmentStateSequence,omitempty"`
}

// TaskStatus defines the observed state of Task
type TaskStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Task is the Schema for the tasks API
type Task struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TaskSpec   `json:"spec,omitempty"`
	Status TaskStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TaskList contains a list of Task
type TaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Task `json:"items"`
}

type PointStateSequence struct {
	SequenceID *uint    `json:"sequenceID,omitempty"`
	PointID    *uint    `json:"pointID,omitempty"`
	Position_x *float32 `json:"position_x,omitempty"`
	Position_y *float32 `json:"position_y,omitempty"`
	Allocated  uint16   `json:"allocated,omitempty"` // 0x00:allocated, 0x01: preallocate
}

type SegmentStateSequence struct {
	SequenceID uint   `json:"sequenceID,omitempty"`
	SegmentID  uint   `json:"segmentID,omitempty"`
	Allocated  uint16 `json:"allocated,omitempty"` // 0x00:allocated, 0x01: preallocate
}

func init() {
	SchemeBuilder.Register(&Task{}, &TaskList{})
}
