// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGenerateDataSet struct {
}

func (*validateOpGenerateDataSet) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGenerateDataSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GenerateDataSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGenerateDataSetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartSupportDataExport struct {
}

func (*validateOpStartSupportDataExport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartSupportDataExport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartSupportDataExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartSupportDataExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGenerateDataSetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGenerateDataSet{}, middleware.After)
}

func addOpStartSupportDataExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartSupportDataExport{}, middleware.After)
}

func validateOpGenerateDataSetInput(v *GenerateDataSetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GenerateDataSetInput"}
	if len(v.DataSetType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DataSetType"))
	}
	if v.DataSetPublicationDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSetPublicationDate"))
	}
	if v.RoleNameArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleNameArn"))
	}
	if v.DestinationS3BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationS3BucketName"))
	}
	if v.SnsTopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnsTopicArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartSupportDataExportInput(v *StartSupportDataExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartSupportDataExportInput"}
	if len(v.DataSetType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DataSetType"))
	}
	if v.FromDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FromDate"))
	}
	if v.RoleNameArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleNameArn"))
	}
	if v.DestinationS3BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationS3BucketName"))
	}
	if v.SnsTopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnsTopicArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
