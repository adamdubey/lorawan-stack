// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	time "time"
)

func (dst *ApplicationPackage) SetFields(src *ApplicationPackage, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "default_f_port":
			if len(subs) > 0 {
				return fmt.Errorf("'default_f_port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DefaultFPort = src.DefaultFPort
			} else {
				var zero uint32
				dst.DefaultFPort = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPackages) SetFields(src *ApplicationPackages, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "packages":
			if len(subs) > 0 {
				return fmt.Errorf("'packages' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Packages = src.Packages
			} else {
				dst.Packages = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPackageAssociationIdentifiers) SetFields(src *ApplicationPackageAssociationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "f_port":
			if len(subs) > 0 {
				return fmt.Errorf("'f_port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FPort = src.FPort
			} else {
				var zero uint32
				dst.FPort = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPackageAssociation) SetFields(src *ApplicationPackageAssociation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPackageAssociationIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &ApplicationPackageAssociationIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				dst.CreatedAt = nil
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				dst.UpdatedAt = nil
			}
		case "package_name":
			if len(subs) > 0 {
				return fmt.Errorf("'package_name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PackageName = src.PackageName
			} else {
				var zero string
				dst.PackageName = zero
			}
		case "data":
			if len(subs) > 0 {
				return fmt.Errorf("'data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Data = src.Data
			} else {
				dst.Data = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPackageAssociations) SetFields(src *ApplicationPackageAssociations, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "associations":
			if len(subs) > 0 {
				return fmt.Errorf("'associations' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Associations = src.Associations
			} else {
				dst.Associations = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetApplicationPackageAssociationRequest) SetFields(src *GetApplicationPackageAssociationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPackageAssociationIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &ApplicationPackageAssociationIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListApplicationPackageAssociationRequest) SetFields(src *ListApplicationPackageAssociationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.Ids == nil) && dst.Ids == nil {
					continue
				}
				if src != nil {
					newSrc = src.Ids
				}
				if dst.Ids != nil {
					newDst = dst.Ids
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.Ids = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					dst.Ids = nil
				}
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetApplicationPackageAssociationRequest) SetFields(src *SetApplicationPackageAssociationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "association":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPackageAssociation
				if (src == nil || src.Association == nil) && dst.Association == nil {
					continue
				}
				if src != nil {
					newSrc = src.Association
				}
				if dst.Association != nil {
					newDst = dst.Association
				} else {
					newDst = &ApplicationPackageAssociation{}
					dst.Association = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Association = src.Association
				} else {
					dst.Association = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPackageDefaultAssociationIdentifiers) SetFields(src *ApplicationPackageDefaultAssociationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "f_port":
			if len(subs) > 0 {
				return fmt.Errorf("'f_port' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FPort = src.FPort
			} else {
				var zero uint32
				dst.FPort = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPackageDefaultAssociation) SetFields(src *ApplicationPackageDefaultAssociation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPackageDefaultAssociationIdentifiers
				if src != nil {
					newSrc = &src.Ids
				}
				newDst = &dst.Ids
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					var zero ApplicationPackageDefaultAssociationIdentifiers
					dst.Ids = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				var zero time.Time
				dst.UpdatedAt = zero
			}
		case "package_name":
			if len(subs) > 0 {
				return fmt.Errorf("'package_name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PackageName = src.PackageName
			} else {
				var zero string
				dst.PackageName = zero
			}
		case "data":
			if len(subs) > 0 {
				return fmt.Errorf("'data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Data = src.Data
			} else {
				dst.Data = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ApplicationPackageDefaultAssociations) SetFields(src *ApplicationPackageDefaultAssociations, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "defaults":
			if len(subs) > 0 {
				return fmt.Errorf("'defaults' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Defaults = src.Defaults
			} else {
				dst.Defaults = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetApplicationPackageDefaultAssociationRequest) SetFields(src *GetApplicationPackageDefaultAssociationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPackageDefaultAssociationIdentifiers
				if src != nil {
					newSrc = &src.Ids
				}
				newDst = &dst.Ids
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					var zero ApplicationPackageDefaultAssociationIdentifiers
					dst.Ids = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListApplicationPackageDefaultAssociationRequest) SetFields(src *ListApplicationPackageDefaultAssociationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.Ids
				}
				newDst = &dst.Ids
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Ids = src.Ids
				} else {
					var zero ApplicationIdentifiers
					dst.Ids = zero
				}
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetApplicationPackageDefaultAssociationRequest) SetFields(src *SetApplicationPackageDefaultAssociationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "default":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationPackageDefaultAssociation
				if src != nil {
					newSrc = &src.Default
				}
				newDst = &dst.Default
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Default = src.Default
				} else {
					var zero ApplicationPackageDefaultAssociation
					dst.Default = zero
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
