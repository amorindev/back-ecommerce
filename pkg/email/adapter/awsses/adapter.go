package awsses

import "com.fernando/pkg/email/port"

var _ port.EmailAdapter = &SesAdapter{}

type SesAdapter struct{}
