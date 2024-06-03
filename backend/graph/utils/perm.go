package utils

import "github.com/hawa130/computility-cloud/graph/model"

func StringArrToPolicy(arr []string) *model.CPolicy {
	return &model.CPolicy{
		Sub: arr[0],
		Obj: arr[1],
		Act: arr[2],
	}
}

func StringArrToGroup(arr []string) *model.CGroup {
	return &model.CGroup{
		Sub: arr[0],
		Obj: arr[1],
	}
}

func PolicyToAnyArr(policy *model.CRequestInput) []interface{} {
	return []interface{}{policy.Sub, policy.Obj, policy.Act}
}

func PolicyToStringArr(policy *model.CRequestInput) []string {
	return []string{policy.Sub, policy.Obj, policy.Act}
}

func GroupToAnyArr(group *model.CGroupInput) []interface{} {
	return []interface{}{group.Sub, group.Obj}
}

func GroupToStringArr(group *model.CGroupInput) []string {
	return []string{group.Sub, group.Obj}
}

func ToResult(input *model.CRequestInput, res bool) *model.CPolicyResult {
	return &model.CPolicyResult{
		Sub: input.Sub,
		Obj: input.Obj,
		Act: input.Act,
		Ok:  res,
	}
}

func ToGroupResult(input *model.CGroupInput, res bool) *model.CGroupResult {
	return &model.CGroupResult{
		Sub: input.Sub,
		Obj: input.Obj,
		Ok:  res,
	}
}

func ToUpdateResult(new, old *model.CRequestInput, res bool) *model.UpdateCPolicy {
	return &model.UpdateCPolicy{
		New: &model.CPolicy{
			Sub: new.Sub,
			Obj: new.Obj,
			Act: new.Act,
		},
		Old: &model.CPolicy{
			Sub: old.Sub,
			Obj: old.Obj,
			Act: old.Act,
		},
		Ok: res,
	}
}

func ToGroupUpdateResult(new, old *model.CGroupInput, res bool) *model.UpdateCGroup {
	return &model.UpdateCGroup{
		New: &model.CGroup{
			Sub: new.Sub,
			Obj: new.Obj,
		},
		Old: &model.CGroup{
			Sub: old.Sub,
			Obj: old.Obj,
		},
		Ok: res,
	}
}

func PoliciesToStringArr(input []*model.CRequestInput) [][]string {
	result := make([][]string, 0)
	for _, policy := range input {
		result = append(result, []string{policy.Sub, policy.Obj, policy.Act})
	}
	return result
}

func GroupsToStringArr(input []*model.CGroupInput) [][]string {
	result := make([][]string, 0)
	for _, group := range input {
		result = append(result, []string{group.Sub, group.Obj})
	}
	return result
}

func ToBatchResult(input []*model.CRequestInput, ok bool) *model.BatchCPolicy {
	result := make([]*model.CPolicy, 0)
	for _, policy := range input {
		result = append(result, &model.CPolicy{
			Sub: policy.Sub,
			Obj: policy.Obj,
			Act: policy.Act,
		})
	}
	return &model.BatchCPolicy{
		Data: result,
		Ok:   ok,
	}
}

func ToBatchGroupResult(input []*model.CGroupInput, ok bool) *model.BatchCGroup {
	result := make([]*model.CGroup, 0)
	for _, group := range input {
		result = append(result, &model.CGroup{
			Sub: group.Sub,
			Obj: group.Obj,
		})
	}
	return &model.BatchCGroup{
		Data: result,
		Ok:   ok,
	}
}
