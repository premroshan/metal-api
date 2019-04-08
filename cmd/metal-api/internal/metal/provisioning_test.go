package metal

import (
	"testing"

	"git.f-i-ts.de/cloud-native/metallib/zapup"
)

var (
	SuccessfulEventCycle = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventBootingNewKernel,
		},
		ProvisioningEvent{
			Event: ProvisioningEventInstalling,
		},
		ProvisioningEvent{
			Event: ProvisioningEventWaiting,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
	}
	CrashEventCycle = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
	}
	CycleWithPlannedReboot = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventWaiting,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPlannedReboot,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
	}
	CycleWithPlannedRebootAndError = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventInstalling,
		},
		ProvisioningEvent{
			Event: ProvisioningEventWaiting,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPlannedReboot,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
	}
	CycleWithPlannedRebootAndImmediateError = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventInstalling,
		},
		ProvisioningEvent{
			Event: ProvisioningEventWaiting,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPlannedReboot,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
	}
	CycleWithACrash = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventCrashed,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
	}
	CycleWithReset = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventResetFailCount,
		},
		ProvisioningEvent{
			Event: ProvisioningEventWaiting,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventCrashed,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
	}
	SuccessfulEventCycleWithBadHistory = ProvisioningEvents{
		ProvisioningEvent{
			Event: ProvisioningEventBootingNewKernel,
		},
		ProvisioningEvent{
			Event: ProvisioningEventInstalling,
		},
		ProvisioningEvent{
			Event: ProvisioningEventWaiting,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
		ProvisioningEvent{
			Event: ProvisioningEventPreparing,
		},
		ProvisioningEvent{
			Event: ProvisioningEventRegistering,
		},
	}
)

func TestProvisioning_IncompleteCycles(t *testing.T) {

	tests := []struct {
		name           string
		eventContainer ProvisioningEventContainer
		want           string
	}{
		{
			name: "TestProvisioning_IncompleteCycles Test 1",
			eventContainer: ProvisioningEventContainer{
				Events: SuccessfulEventCycle,
			},
			want: "0",
		},
		{
			name: "TestProvisioning_IncompleteCycles Test 2",
			eventContainer: ProvisioningEventContainer{
				Events: CrashEventCycle,
			},
			want: "2",
		},
		{
			name: "TestProvisioning_IncompleteCycles Test 3",
			eventContainer: ProvisioningEventContainer{
				Events: CycleWithPlannedReboot,
			},
			want: "0",
		},
		{
			name: "TestProvisioning_IncompleteCycles Test 4",
			eventContainer: ProvisioningEventContainer{
				Events: CycleWithPlannedRebootAndError,
			},
			want: "1",
		},
		{
			name: "TestProvisioning_IncompleteCycles Test 5",
			eventContainer: ProvisioningEventContainer{
				Events: CycleWithPlannedRebootAndImmediateError,
			},
			want: "1",
		},
		{
			name: "TestProvisioning_IncompleteCycles Test 6",
			eventContainer: ProvisioningEventContainer{
				Events: CycleWithACrash,
			},
			want: "1",
		},
		{
			name: "TestProvisioning_IncompleteCycles Test 7",
			eventContainer: ProvisioningEventContainer{
				Events: CycleWithReset,
			},
			want: "0",
		},
		{
			name: "TestProvisioning_IncompleteCycles Test 8",
			eventContainer: ProvisioningEventContainer{
				Events: SuccessfulEventCycleWithBadHistory,
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.eventContainer.CalculateIncompleteCycles(zapup.MustRootLogger().Sugar()); got != tt.want {
				t.Errorf("CalculateIncompleteCycles() = %v, want %v", got, tt.want)
			}
		})
	}
}