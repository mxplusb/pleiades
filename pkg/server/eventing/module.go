/*
 * Copyright (c) 2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

package eventing

import (
	"github.com/mxplusb/pleiades/pkg/messaging"
	"go.uber.org/fx"
)

var EventingModule = fx.Module("eventing",
	fx.Provide(messaging.NewEmbeddedMessagingWithDefaults),
	fx.Provide(NewEventServer),
	fx.Invoke(NewLifecycleManager),
	fx.Provide(NewStreamClient),
	fx.Provide(NewPubSubClient))
