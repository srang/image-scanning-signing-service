package scanning

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/redhat-cop/image-scanning-signing-service/pkg/apis/cop/v1alpha2"
	"github.com/redhat-cop/image-scanning-signing-service/pkg/config"
	"github.com/redhat-cop/image-scanning-signing-service/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"

	"github.com/sirupsen/logrus"

	corev1 "k8s.io/api/core/v1"
)

func UpdateOnScanningPodLaunch(message string, imageScanningRequest v1alpha2.ImageScanningRequest) error {
	imageScanningRequestCopy := imageScanningRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionTrue, v1alpha2.ImageExecutionConditionInitialization)

	imageScanningRequestCopy.Status.StartTime = condition.LastTransitionTime

	return updateImageScanningRequest(imageScanningRequestCopy, condition, v1alpha2.PhaseRunning)
}

func UpdateOnImageScanningInitializationFailure(message string, imageScanningRequest v1alpha2.ImageScanningRequest) error {
	imageScanningRequestCopy := imageScanningRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionFalse, v1alpha2.ImageExecutionConditionInitialization)

	imageScanningRequestCopy.Status.StartTime = condition.LastTransitionTime
	imageScanningRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageScanningRequest(imageScanningRequestCopy, condition, v1alpha2.PhaseFailed)
}

func updateImageScanningRequest(imageScanningRequest *v1alpha2.ImageScanningRequest, condition v1alpha2.ImageExecutionCondition, phase v1alpha2.ImageExecutionPhase) error {

	imageScanningRequest.Status.Conditions = append(imageScanningRequest.Status.Conditions, condition)
	imageScanningRequest.Status.Phase = phase

	err := sdk.Update(imageScanningRequest)

	return err
}

func UpdateOnImageScanningCompletionError(message string, imageScanningRequest v1alpha2.ImageScanningRequest) error {
	imageScanningRequestCopy := imageScanningRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionFalse, v1alpha2.ImageExecutionConditionFinished)

	imageScanningRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageScanningRequest(imageScanningRequestCopy, condition, v1alpha2.PhaseFailed)
}

func UpdateOnImageScanningCompletionSuccess(message string, totalRules int, passedRules int, failedRules int, imageScanningRequest v1alpha2.ImageScanningRequest) error {
	imageScanningRequestCopy := imageScanningRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionTrue, v1alpha2.ImageExecutionConditionFinished)

	scanResult := &v1alpha2.ScanResult{
		FailedRules: failedRules,
		PassedRules: passedRules,
		TotalRules:  totalRules,
	}

	imageScanningRequestCopy.Status.ScanResult = *scanResult
	imageScanningRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageScanningRequest(imageScanningRequestCopy, condition, v1alpha2.PhaseCompleted)
}

func LaunchScanningPod(config config.Config, image string, ownerID string, ownerReference string) (string, error) {

	pod, err := createScanningPod(config.SignScanImage, config.TargetProject, image, ownerID, ownerReference, config.TargetServiceAccount)

	if err != nil {
		logrus.Errorf("Error Generating Pod: %v'", err)
		return "", err
	}

	err = sdk.Create(pod)

	if err != nil {
		logrus.Errorf("Error Creating Pod: %v'", err)
		return "", err
	}

	var key string
	if key, err = cache.MetaNamespaceKeyFunc(pod); err != nil {
		runtime.HandleError(err)
		return "", err
	}

	return key, nil
}

func createScanningPod(signScanImage string, targetProject string, image string, ownerID string, ownerReference string, serviceAccount string) (*corev1.Pod, error) {

	priv := true

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        ownerID,
			Namespace:   targetProject,
			Labels:      map[string]string{"type": "image-scanning"},
			Annotations: map[string]string{"cop.redhat.com/owner": ownerReference, "cop.redhat.com/type": "image-scanning"},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				Name:            "image-scanner",
				Image:           signScanImage,
				ImagePullPolicy: corev1.PullAlways,
				Command:         []string{"/bin/bash", "-c", "/usr/local/bin/scan-image"},
				Env: []corev1.EnvVar{
					{
						Name:  "IMAGE",
						Value: image,
					},
				},
				SecurityContext: &corev1.SecurityContext{
					Privileged: &priv,
				},
				Ports: []corev1.ContainerPort{
					{
						Name:          "webdav",
						ContainerPort: 8080,
						Protocol:      "TCP",
					},
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "docker-socket",
						MountPath: "/var/run/docker.sock",
					},
				},
			}},
			RestartPolicy:      corev1.RestartPolicyNever,
			ServiceAccountName: serviceAccount,
			Volumes: []corev1.Volume{
				{
					Name: "docker-socket",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/var/run/docker.sock",
						},
					},
				},
			},
		},
	}

	return pod, nil
}

func DeleteScanningPod(name string, namespace string) error {

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}

	return sdk.Delete(pod)

}
