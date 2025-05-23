package conditions

import (
	"github.com/kyma-project/btp-manager/api/v1alpha1"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Reason string

// gophers_reasons_section_start
const (
	ReconcileSucceeded                                Reason = "ReconcileSucceeded"
	ReconcileFailed                                   Reason = "ReconcileFailed"
	Initialized                                       Reason = "Initialized"
	Processing                                        Reason = "Processing"
	WrongNamespaceOrName                              Reason = "WrongNamespaceOrName"
	ChartInstallFailed                                Reason = "ChartInstallFailed"
	ConsistencyCheckFailed                            Reason = "ConsistencyCheckFailed"
	MissingSecret                                     Reason = "MissingSecret"
	InvalidSecret                                     Reason = "InvalidSecret"
	HardDeleting                                      Reason = "HardDeleting"
	ResourceRemovalFailed                             Reason = "ResourceRemovalFailed"
	SoftDeleting                                      Reason = "SoftDeleting"
	Updated                                           Reason = "Updated"
	UpdateCheck                                       Reason = "UpdateCheck"
	UpdateCheckSucceeded                              Reason = "UpdateCheckSucceeded"
	InconsistentChart                                 Reason = "InconsistentChart"
	UpdateDone                                        Reason = "UpdateDone"
	PreparingInstallInfoFailed                        Reason = "PreparingInstallInfoFailed"
	ChartPathEmpty                                    Reason = "ChartPathEmpty"
	DeletionOfOrphanedResourcesFailed                 Reason = "DeletionOfOrphanedResourcesFailed"
	ServiceInstancesAndBindingsNotCleaned             Reason = "ServiceInstancesAndBindingsNotCleaned"
	StoringChartDetailsFailed                         Reason = "StoringChartDetailsFailed"
	GettingConfigMapFailed                            Reason = "GettingConfigMapFailed"
	ProvisioningFailed                                Reason = "ProvisioningFailed"
	GettingDefaultCredentialsSecretFailed             Reason = "GettingDefaultCredentialsSecretFailed"
	GettingSapBtpServiceOperatorConfigMapFailed       Reason = "GettingSapBtpServiceOperatorConfigMapFailed"
	CredentialsNamespaceChanged                       Reason = "CredentialsNamespaceChanged"
	ClusterIdChanged                                  Reason = "ClusterIdChanged"
	AnnotatingSecretFailed                            Reason = "AnnotatingSecretFailed"
	GettingSapBtpServiceOperatorClusterIdSecretFailed Reason = "GettingSapBtpServiceOperatorClusterIdSecretFailed"
)

// gophers_reasons_section_end

const (
	ReadyType = "Ready"
)

type Metadata struct {
	Status metav1.ConditionStatus
	State  v1alpha1.State
}

// gophers_metadata_section_start
var Reasons = map[Reason]Metadata{
	ReconcileSucceeded:                    {Status: metav1.ConditionTrue, State: v1alpha1.StateReady},       //Ready;Reconciled successfully
	UpdateDone:                            {Status: metav1.ConditionTrue, State: v1alpha1.StateReady},       //Ready;Update done
	UpdateCheckSucceeded:                  {Status: metav1.ConditionTrue, State: v1alpha1.StateReady},       //Ready;Update not required
	ReconcileFailed:                       {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Reconciliation failed
	Updated:                               {Status: metav1.ConditionFalse, State: v1alpha1.StateProcessing}, //Processing;Resource has been updated
	Initialized:                           {Status: metav1.ConditionFalse, State: v1alpha1.StateProcessing}, //Processing;Initial processing or chart is inconsistent
	ChartInstallFailed:                    {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Failure during chart installation
	ConsistencyCheckFailed:                {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Failure during consistency check
	Processing:                            {Status: metav1.ConditionFalse, State: v1alpha1.StateProcessing}, //Processing;Final State after deprovisioning
	WrongNamespaceOrName:                  {Status: metav1.ConditionFalse, State: v1alpha1.StateWarning},    //Warning;Wrong namespace or name
	MissingSecret:                         {Status: metav1.ConditionFalse, State: v1alpha1.StateWarning},    //Warning;`sap-btp-manager` Secret was not found - create proper Secret
	InvalidSecret:                         {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;`sap-btp-manager` Secret does not contain required data - create proper Secret
	HardDeleting:                          {Status: metav1.ConditionFalse, State: v1alpha1.StateDeleting},   //Deleting;Trying to hard delete
	ResourceRemovalFailed:                 {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Some resources can still be present due to errors while deprovisioning
	SoftDeleting:                          {Status: metav1.ConditionFalse, State: v1alpha1.StateDeleting},   //Deleting;Trying to soft-delete after hard-delete failed
	UpdateCheck:                           {Status: metav1.ConditionFalse, State: v1alpha1.StateProcessing}, //Processing;Checking for updates
	InconsistentChart:                     {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Chart is inconsistent, reconciliation initialized
	PreparingInstallInfoFailed:            {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Error while preparing installation information
	ChartPathEmpty:                        {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;No chart path available for processing
	DeletionOfOrphanedResourcesFailed:     {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Deletion of orphaned resources failed
	StoringChartDetailsFailed:             {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Failure of storing chart details
	GettingConfigMapFailed:                {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Getting ConfigMap failed
	ProvisioningFailed:                    {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Provisioning failed
	ServiceInstancesAndBindingsNotCleaned: {Status: metav1.ConditionFalse, State: v1alpha1.StateWarning},    //Warning;Deprovisioning blocked because of ServiceInstances and/or ServiceBindings existence
	GettingDefaultCredentialsSecretFailed: {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Getting default credentials Secret failed
	AnnotatingSecretFailed:                {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Annotating the required Secret failed
	GettingSapBtpServiceOperatorConfigMapFailed:       {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Getting SAP BTP service operator ConfigMap failed
	CredentialsNamespaceChanged:                       {Status: metav1.ConditionFalse, State: v1alpha1.StateProcessing}, //Processing;Credentials namespace changed
	ClusterIdChanged:                                  {Status: metav1.ConditionFalse, State: v1alpha1.StateProcessing}, //Processing;Cluster ID changed
	GettingSapBtpServiceOperatorClusterIdSecretFailed: {Status: metav1.ConditionFalse, State: v1alpha1.StateError},      //Error;Getting SAP BTP service operator Cluster ID Secret failed
}

// gophers_metadata_section_end

func ConditionFromExistingReason(reason Reason, message string) *metav1.Condition {
	metadata, found := Reasons[reason]
	if found {
		return &metav1.Condition{
			Status:             metadata.Status,
			Reason:             string(reason),
			Message:            message,
			Type:               ReadyType,
			ObservedGeneration: 0,
		}
	}
	return nil
}

// This is required because of difference between Conditions declarations
// In BtpOperator we have Status.Conditions []*Condition instead of Status.Conditions []Condition
func SetStatusCondition(conditions *[]*metav1.Condition, newCondition metav1.Condition) {
	conditionsCnt := len(*conditions)
	var conditionsArray = make([]metav1.Condition, conditionsCnt, conditionsCnt)
	for i := 0; i < conditionsCnt; i++ {
		conditionsArray[i] = *(*conditions)[i]
	}

	apimeta.SetStatusCondition(&conditionsArray, newCondition)

	for i := 0; i < conditionsCnt; i++ {
		(*conditions)[i] = &conditionsArray[i]
	}
	if len(conditionsArray) > conditionsCnt {
		*conditions = append(*conditions, &metav1.Condition{})
		(*conditions)[conditionsCnt] = &conditionsArray[conditionsCnt]
	}
}
