// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type Profile -header-file ../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package profile

import (
	"github.com/siderolabs/talos/pkg/machinery/meta"
)

// DeepCopy generates a deep copy of Profile.
func (o Profile) DeepCopy() Profile {
	var cp Profile = o
	if o.SecureBoot != nil {
		cp.SecureBoot = new(bool)
		*cp.SecureBoot = *o.SecureBoot
	}
	if o.Customization.ExtraKernelArgs != nil {
		cp.Customization.ExtraKernelArgs = make([]string, len(o.Customization.ExtraKernelArgs))
		copy(cp.Customization.ExtraKernelArgs, o.Customization.ExtraKernelArgs)
	}
	if o.Customization.MetaContents != nil {
		cp.Customization.MetaContents = make([]meta.Value, len(o.Customization.MetaContents))
		copy(cp.Customization.MetaContents, o.Customization.MetaContents)
	}
	if o.Input.SecureBoot != nil {
		cp.Input.SecureBoot = new(SecureBootAssets)
		*cp.Input.SecureBoot = *o.Input.SecureBoot
	}
	if o.Input.SystemExtensions != nil {
		cp.Input.SystemExtensions = make([]ContainerAsset, len(o.Input.SystemExtensions))
		copy(cp.Input.SystemExtensions, o.Input.SystemExtensions)
	}
	if o.Overlay != nil {
		cp.Overlay = new(OverlayOptions)
		*cp.Overlay = *o.Overlay
		if o.Overlay.Options != nil {
			cp.Overlay.Options = make(map[string]any, len(o.Overlay.Options))
			for k4, v4 := range o.Overlay.Options {
				cp.Overlay.Options[k4] = v4
			}
		}
	}
	if o.Output.ImageOptions != nil {
		cp.Output.ImageOptions = new(ImageOptions)
		*cp.Output.ImageOptions = *o.Output.ImageOptions
	}
	if o.Output.ISOOptions != nil {
		cp.Output.ISOOptions = new(ISOOptions)
		*cp.Output.ISOOptions = *o.Output.ISOOptions
	}
	return cp
}
