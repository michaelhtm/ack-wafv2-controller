	// LoggingConfiguration is applied via a separate PutLoggingConfiguration call.
	// Defer it to the update path: mark the resource unsynced so the runtime
	// requeues and the delta-driven update applies the logging configuration.
	if ko.Spec.LoggingConfiguration != nil {
		msg := "Logging configuration update pending; resource will be requeued to sync logging configuration"
		ackcondition.SetSynced(&resource{ko}, corev1.ConditionFalse, &msg, nil)
	}
