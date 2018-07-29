package verification

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/util/runtime"
)

func UpdateOnVerificationPodLaunch(message string, unsignedImage string, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionTrue, v1alpha2.ImageExecutionConditionInitialization)

	imageVerificationRequestCopy.Status.StartTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseRunning)
}

func UpdateOnImageVerificationInitializationFailure(message string, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionFalse, v1alpha2.ImageExecutionConditionInitialization)

	imageVerificationRequestCopy.Status.StartTime = condition.LastTransitionTime
	imageVerificationRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseFailed)
}

func updateImageVerificationRequest(imageVerificationRequest *v1alpha2.ImageVerificationRequest, condition v1alpha2.ImageExecutionCondition, phase v1alpha2.ImageExecutionPhase) error {

	imageVerificationRequest.Status.Conditions = append(imageVerificationRequest.Status.Conditions, condition)
	imageVerificationRequest.Status.Phase = phase

	err := sdk.Update(imageVerificationRequest)

	return err
}

func UpdateOnImageVerificationCompletionError(message string, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionFalse, v1alpha2.ImageExecutionConditionFinished)

	imageVerificationRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseFailed)
}

func UpdateOnImageVerificationCompletionSuccess(message string, totalRules int, passedRules int, failedRules int, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionTrue, v1alpha2.ImageExecutionConditionFinished)

	scanResult := &v1alpha2.ScanResult{
		FailedRules: failedRules,
		PassedRules: passedRules,
		TotalRules:  totalRules,
	}

	imageVerificationRequestCopy.Status.ScanResult = *scanResult
	imageVerificationRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseCompleted)
}

func LaunchVerificationPod(config config.Config, image string, ownerID string, ownerReference string) (string, error) {

	pod, err := createVerificationPod(config.SignScanImage, config.TargetProject, image, ownerID, ownerReference, config.TargetServiceAccount)

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

func createVerificationPod(signScanImage string, targetProject string, image string, ownerID string, ownerReference string, serviceAccount string) (*corev1.Pod, error) {

	priv := true

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        ownerID,
			Namespace:   targetProject,
			Labels:      map[string]string{"type": "image-verification"},
			Annotations: map[string]string{"cop.redhat.com/owner": ownerReference, "cop.redhat.com/type": "image-verification"},
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

func DeleteVerificationPod(name string, namespace string) error {

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

