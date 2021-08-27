package cmdutils

import "github.com/wizedkyle/sumologic-go-sdk/service/cip/types"

func GenerateFieldsMap(fieldNames []string, fieldValues []string) map[string]string {
	if len(fieldNames) > 0 && len(fieldValues) > 0 {
		fieldsMap := make(map[string]string)
		for i, _ := range fieldNames {
			fieldsMap[fieldNames[i]] = fieldValues[i]
			i++
		}
		return fieldsMap
	} else {
		return nil
	}
}

func GenerateLookupTableFields(fieldNames []string, fieldTypes []string) []types.LookupTableField {
	if len(fieldNames) > 0 && len(fieldTypes) > 0 {
		fields := []types.LookupTableField{}
		for i, _ := range fieldNames {
			field := types.LookupTableField{
				FieldName: fieldNames[i],
				FieldType: fieldTypes[i],
			}
			fields = append(fields, field)
			i++
		}
		return fields
	} else {
		return nil
	}
}

func GenerateLookupTableColumns(columnNames []string, columnValues []string) []types.TableRow {
	if len(columnNames) > 0 && len(columnValues) > 0 {
		fields := []types.TableRow{}
		for i, _ := range columnNames {
			field := types.TableRow{
				ColumnName:  columnNames[i],
				ColumnValue: columnValues[i],
			}
			fields = append(fields, field)
			i++
		}
		return fields
	} else {
		return nil
	}
}

func GenerateCidrList(ipAddresses []string, descriptions []string) types.CidrList {
	addressList := types.CidrList{}
	if len(ipAddresses) > 0 {
		for i, _ := range ipAddresses {
			address := types.Cidr{
				Cidr:        ipAddresses[i],
				Description: descriptions[i],
			}
			addressList.Data = append(addressList.Data, address)
			i++
		}
		return addressList
	} else {
		return addressList
	}
}
