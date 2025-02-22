// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BucketACL.
func (mg *BucketACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketObject.
func (mg *BucketObject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketObjectACL.
func (mg *BucketObjectACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Key),
		Extract:      resource.ExtractParamPath("key", true),
		Reference:    mg.Spec.ForProvider.KeyRef,
		Selector:     mg.Spec.ForProvider.KeySelector,
		To: reference.To{
			List:    &BucketObjectList{},
			Managed: &BucketObject{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Key")
	}
	mg.Spec.ForProvider.Key = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Key),
		Extract:      resource.ExtractParamPath("key", true),
		Reference:    mg.Spec.InitProvider.KeyRef,
		Selector:     mg.Spec.InitProvider.KeySelector,
		To: reference.To{
			List:    &BucketObjectList{},
			Managed: &BucketObject{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Key")
	}
	mg.Spec.InitProvider.Key = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketPolicy.
func (mg *BucketPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BucketReplication.
func (mg *BucketReplication) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DestinationBucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.ForProvider.DestinationBucketRef,
		Selector:     mg.Spec.ForProvider.DestinationBucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DestinationBucket")
	}
	mg.Spec.ForProvider.DestinationBucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DestinationBucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DestinationBucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.InitProvider.DestinationBucketRef,
		Selector:     mg.Spec.InitProvider.DestinationBucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DestinationBucket")
	}
	mg.Spec.InitProvider.DestinationBucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DestinationBucketRef = rsp.ResolvedReference

	return nil
}
